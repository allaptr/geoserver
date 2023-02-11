package maps

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

const (
	// All state borders used as is, except for Washington.
	// The Washington border was modified to closer conform with the geographical border on the map.
	// Hey, this is my state, I wanted it to be well rounded :-)
	wash = `{"state": "Washington", "border": [[-122.402015, 48.225216],[-122.700000, 48.999999], [-117.032049, 48.999931], [-116.919132, 45.995175], [-124.079107, 46.267259], [-124.717175, 48.377557], [-122.92315, 48.047963], [-122.402015, 48.225216]]}`
)

type StateData struct {
	State  string      `json: "state"`
	Border [][]float64 `json: "border"`
}

var statesData []StateData

func load(filename string) ([]StateData, error) {
	strs, err := readData(filename)
	if err != nil {
		return nil, errors.Wrap(err, "readData")
	}
	sd, err := populateStateData(strs)
	if err != nil {
		return nil, errors.Wrap(err, "populateStateData")
	}
	return sd, nil
}

func readData(fname string) ([]string, error) {
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		return nil, errors.Errorf("Cannot open file '%s': '%v'", fname, err)
	}
	rdr := bufio.NewReader(f)
	data := []string{}
	cnt := 0
	for {
		d, err := rdr.ReadString([]byte("\n")[0])
		if err == nil {
			data = append(data, d)
		} else if errors.Cause(err) == io.EOF {
			data = append(data, d) // handle case when file does not have EOF character at the end
			break
		} else {
			fmt.Println(cnt)
			return nil, err
		}
		cnt++
	}
	data[0] = wash // updating data for washington state
	return data, nil
}

func populateStateData(sdata []string) ([]StateData, error) {
	stdata := []StateData{}
	for _, sd := range sdata {
		if len(sd) == 0 {
			continue
		}
		var s StateData
		err := json.Unmarshal([]byte(sd), &s)
		if err != nil {
			return nil, errors.Wrapf(err, "unmarshalling '%s'", sd)
		}
		stdata = append(stdata, s)
	}
	return stdata, nil
}
