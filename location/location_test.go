package location

import (
	"state-server/maps"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LocationState(t *testing.T) {
	err := maps.CreateMap("../cmd/state-server/data/states.json")
	assert.NoError(t, err)

	testData := []struct {
		name     string
		location []float64
		response string
	}{
		{"Seattle", []float64{-122.335167, 47.608013}, "Washington"},
		{"Portland", []float64{-122.707969, 45.683114}, "Oregon"},
		{"Salem", []float64{-123.035768, 44.911756}, "Oregon"},
		{"Missoula", []float64{-113.992933, 46.872561}, "Montana"},
		{"San Diego", []float64{-117.150869, 32.758670}, "California"},
		{"Portland", []float64{-70.290136, 43.697047}, "Maine"},
		{"Morgantown", []float64{-79.970465, 39.648044}, "West Virginia"},
		{"Burlington", []float64{-73.016444, 44.492103}, "Vermont"},
		{"Cranville", []float64{-100.846071, 48.268265}, "North Dakota"},
		{"Orlando", []float64{-81.387652, 28.5450541}, "Florida"},
		{"Malden", []float64{-71.071022, 42.429752}, "Massachusetts"},
	}
	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			res := LocationState(td.location)
			assert.Equal(t, td.response, res)
		})
	}
}
