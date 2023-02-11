package location

import (
	"math"
	"state-server/data"

	"github.com/pkg/errors"
)

// Each state border creates a closed path, that forms a polygon shape
type Polygon struct {
	StateName string
	Bounds    bounds
	Path      closedPath
}

var polygons []Polygon

func CreateMap(dataFilePath string) error {
	d, err := data.Load(dataFilePath)
	if err != nil {
		return errors.Wrap(err, "load")
	}
	plgns, err := createPolygons(d)
	if err != nil {
		return errors.Wrap(err, "createPolygons")
	}
	polygons = plgns
	return nil
}

// Populates the collection of polygons from the serialized states border data
func createPolygons(data []data.StateData) ([]Polygon, error) {
	plgns := make([]Polygon, len(data))
	for ind, d := range data {
		path, bounds := newClosedPath(d.Border)
		plgns[ind].Path = path
		plgns[ind].Bounds = bounds
		plgns[ind].StateName = d.State
	}
	return plgns, nil
}

// polygon bounds
type bounds struct {
	minY float64
	minX float64
	maxY float64
	maxX float64
}

func (b *bounds) adjustBounds(borderPont []float64) {
	if borderPont[0] < b.minX {
		b.minX = borderPont[0]
	}
	if borderPont[0] > b.maxX {
		b.maxX = borderPont[0]
	}
	if borderPont[1] < b.minY {
		b.minY = borderPont[1]
	}
	if borderPont[1] > b.maxY {
		b.maxY = borderPont[1]
	}
}

// check if the location point is within the polygon (state) bounds
func (b *bounds) within(point []float64) bool {
	if point[0] >= b.minX && point[0] <= b.maxX && point[1] >= b.minY && point[1] <= b.maxY {
		return true
	}
	return false
}

// ClosedPath is composed of connected line segments that form a polygon
type closedPath struct {
	segments []segment
}

func newClosedPath(border [][]float64) (closedPath, bounds) {
	bnd := bounds{
		maxX: -1000,
	}
	bnd.adjustBounds(border[0])
	// closing the path by connecting the last point of the border to the first.
	segments := []segment{newSegment(border[len(border)-1], border[0])}
	for i := 1; i < len(border); i++ {
		bnd.adjustBounds(border[i])
		segments = append(segments, newSegment(border[i-1], border[i]))
	}
	path := closedPath{
		segments: segments,
	}
	return path, bnd
}

// The method 'contains' returns whether the point is inside the closed path (polygon)
// It implements the algorithm described at https://www.geeksforgeeks.org/how-to-check-if-a-given-point-lies-inside-a-polygon/
// The code is not handling the `point ‘g’` special case described in the algorithm above.
func (p closedPath) contains(point []float64) bool {
	count := 0
	for _, seg := range p.segments {
		if seg.intersects(point) {
			count++
		}
	}
	return count%2 == 1
}

// Segment is a line of a limited length that starts at the point 'begin' and ends at the point 'end'
type segment struct {
	begin []float64
	end   []float64
}

func newSegment(begin, end []float64) segment {
	return segment{
		begin: begin,
		end:   end,
	}
}

// The method 'intersects' determines whether the horizontal line started at the point 'p' that is extending to the right
// into the infinity would intersect with the line segment 's'
func (s segment) intersects(p []float64) bool {
	coeff := math.Abs(s.begin[0]-s.end[0]) / math.Abs(s.begin[1]-s.end[1])
	// xSect* are values on the x coordinate where the horizontal line and the segment would intersect.
	// This determines the range of x & y coordinates for the point location to intesect with the given segment
	xSect1 := s.end[0] + math.Abs(p[1]-s.end[1])*coeff
	xSect2 := s.end[0] - math.Abs(s.end[1]-p[1])*coeff
	xSect3 := s.end[0] - math.Abs(s.end[1]-p[1])*coeff
	xSect4 := s.end[0] + math.Abs(s.end[1]-p[1])*coeff
	// Below are the cases of the spacial relationship between the segment and the point location
	if s.begin[1] > s.end[1] && s.begin[0] > s.end[0] {
		if p[1] <= s.begin[1] && p[1] >= s.end[1] && p[0] < xSect1 {
			return true
		}
		return false
	} else if s.begin[1] > s.end[1] && s.begin[0] < s.end[0] {
		if p[1] <= s.begin[1] && p[1] >= s.end[1] && p[0] <= xSect2 {
			return true
		}
		return false
	} else if s.begin[1] < s.end[1] && s.begin[0] < s.end[0] {
		if p[1] <= s.begin[1] && p[1] >= s.end[1] && p[0] <= xSect3 {
			return true
		}
		return false
	} else if s.begin[1] < s.end[1] && s.begin[0] > s.end[0] {
		if p[1] <= s.end[1] && p[1] >= s.begin[1] && p[0] <= xSect4 {
			return true
		}
		return false
	} else if s.begin[0] == s.end[0] { //case: the segment is vertical
		if s.begin[1] < s.end[1] && p[1] >= s.begin[1] && p[1] <= s.end[1] {
			return true
		} else if s.begin[1] > s.end[1] && p[1] <= s.begin[1] && p[1] >= s.end[1] {
			return true
		}
		return false
	} else if s.begin[1] == s.end[1] && p[1] == s.begin[1] { //case: the segment is horizontal
		if s.begin[0] <= s.end[0] && p[0] >= s.begin[0] && p[0] <= s.end[0] {
			return true
		} else if s.begin[0] >= s.end[0] && p[0] >= s.end[0] && p[0] <= s.begin[0] {
			return true
		}
		return false
	}
	return false
}
