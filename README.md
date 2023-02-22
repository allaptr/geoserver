# State Server!

Given GPS coordinates of a location,the state server tells us in which state, if any, the point is located. 
Some simplified geometries are included in https://github.com/allaptr/stateserver/blob/master/cmd/state-server/data/states.json (so greatly simplified,that some of the smaller ones disappear). I am planning to update & improve the states border data.
## Behavior
```
  $ ./state-server &
  [1] 21507
  $ curl  -d "longitude=-77.036133&latitude=40.513799" http://localhost:8080/
  ["Pennsylvania"]
  $
  ```


