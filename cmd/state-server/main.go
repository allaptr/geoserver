package main

import (
	"fmt"
	"log"
	"net/http"
	"state-server/location"
	"state-server/maps"
	"strconv"
)

func init() {
	err := maps.CreateMap("data/states.json")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			latStr := r.FormValue("latitude")
			longStr := r.FormValue("longitude")
			lat, err1 := strconv.ParseFloat(latStr, 64)
			long, err2 := strconv.ParseFloat(longStr, 64)
			if err1 != nil || err2 != nil {
				http.ResponseWriter.WriteHeader(w, http.StatusBadRequest)
				return
			}
			stateName := location.LocationState([]float64{long, lat})
			http.ResponseWriter.WriteHeader(w, http.StatusOK)
			http.ResponseWriter.Write(w, []byte("[\""+stateName+"\"]\n"))
		default:
			http.ResponseWriter.WriteHeader(w, http.StatusMethodNotAllowed)
			return
		}
	})

	fmt.Printf("Starting the state server ... \n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
