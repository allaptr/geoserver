package main

import (
	"flag"
	"fmt"
	"geoserver/location"
	"geoserver/maps"
	"net/http"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	port := flag.Int("port", 9000, "Port the service is listening on")
	flag.Parse()

	err := maps.CreateMap("cmd/geoserver/data/states.json")
	if err != nil {
		log.Fatal(err)
	}

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

	log.Infof("Starting the geoserver on port %d ... \n", *port)
	listenAddress := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
