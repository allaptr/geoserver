package location

import "state-server/maps"

type GeoLocation interface {
	// Given the geo location of a point, return the name of the state of the location
	// Return 'None' if not found
	LocationState([]float64) string
}

func LocationState(location []float64) string {
	for _, plgn := range maps.GetStatePolygons() {
		// check the point is within the state bounds
		if !plgn.Bounds.Within(location) {
			continue
		}
		// Use more accurate algorithm to see if it's inside the state polygon
		if plgn.Path.Contains(location) {
			return plgn.StateName
		}
	}
	return "None"
}
