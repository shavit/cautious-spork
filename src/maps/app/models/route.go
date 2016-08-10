package models

import (
  "gopkg.in/mgo.v2/bson"
  "time"
  "log"
  "maps/app"
)

type Route struct {
  Id bson.ObjectId `bson:"_id" json:"_id"`
  Uid string `bson:"uid" json:"uid"`
  ResourceState int `bson:"resourceState" json:"resourceState"`
  Name string `bson:"name" json:"name"`
  Description string `bson:"description" json:"description"`
  Athlete string `bson:"athlete" json:"athlete"`
  Distance float64 `bson:"distance" json:"distance"`
  ElevationGain float64 `bson:"elevationGain" json:"elevationGain"`
  RouteType  int `bson:"routeType" json:"routeType"`
  SubType  int `bson:"subType" json:"subType"`
  Private bool `bson:"private" json:"private"`
  Starred bool `bson:"starred" json:"starred"`
  Timestamp int `bson:"timestamp" json:"timestamp"`
  Created time.Time `bson:"created" json:"created"`
  Segments []Segment `bson:"segments" json:"segments"`
}

func (r *Route) All() (routes []Route, e error){
  var data []Route
  var err error

  err = app.DB.C("routes").Find(bson.M{}).All(&data)
  if err != nil {
    log.Fatal("Error:", err)
  }

  return data, err
}
