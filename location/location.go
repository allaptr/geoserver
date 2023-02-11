package location

type GeoLocation interface {
	// Given the geo location of a point, return the name of the state of the location
	// Return 'None' if not found
	LocationState([]float64) string
}

func LocationState(location []float64) string {
	for _, plgn := range polygons {
		// check the point is within the state bounds
		if !plgn.Bounds.within(location) {
			continue
		}
		// Use more accurate algorithm to see if it's inside the state polygon
		if plgn.Path.contains(location) {
			return plgn.StateName
		}
	}
	return "None"
}
