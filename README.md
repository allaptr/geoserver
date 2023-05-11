# Geo Server

Given GPS coordinates of a location, the geoserver returns the state, the point is located in or None. 
The included state border coordinates are greatly simplified in https://github.com/allaptr/geoserver/blob/master/cmd/geoserver/data/states.json so simplified that some small states are not on the map. 

## TODO
- Update & improve the states border data. 
- Add Alaska, Hawaii, and other states absent from the data.

### Build
```
$ go install ./...
```
### Run
```
$ ./geoserver
```

## Send POST Request / get Response
in a different terminal
```
  $ curl  -d "longitude=-122.335167&latitude=47.608013" http://localhost:9000/
  ["Washington"]
  $ curl -d "longitude=-70.290136&latitude=43.697047" http://localhost:9000/
  ["Maine"]
  $ curl -d "longitude=-80.225870&latitude=25.783156" http://localhost:9000/
  ["Florida"]
```


