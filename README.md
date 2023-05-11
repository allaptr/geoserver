# Geo Server

Given GPS coordinates of a location,the geoserver tells us in which state, if any, the point is located. 
Some simplified geometries are included in https://github.com/allaptr/geoserver/blob/master/cmd/geoserver/data/states.json (so greatly simplified,that some of the smaller ones disappear). 

## TODO
- Update & improve the states border data, especially Florida, because Miami is not in Florida, according to the current data. 
- Add Alaska, Hawaii, and other states absent from the data.

### Build
```
$ go install ./...
```
### Run
```
$ ./geoserver &
```

## Send POST Request / get Response
```
  $ curl  -d "longitude=-122.335167&latitude=47.608013" http://localhost:9000/
  ["Washington"]
  $ curl -d "longitude=-70.290136&latitude=43.697047" http://localhost:9000/
  ["Maine"]
  $ curl -d "longitude=-81.387652&latitude=28.5450541" http://localhost:9000/
  ["Florida"]
```


