package models

import (
  "gopkg.in/mgo.v2/bson"
  "time"
)

type Segment struct {
  Id bson.ObjectId `bson:"_id" json:"_id"`
  Uid string `bson:"uid" json:"uid"`
  ResourceState int `bson:"resourceState" json:"resourceState"`
  Name string `bson:"name" json:"name"`
  ActivityType string `bson:"activityType" json:"activityType"`
  Distance float64 `bson:"distance" json:"distance"`
  AverageGrade float64 `bson:"averageGrade" json:"averageGrade"`
  MaximumGrade float64 `bson:"maximumGrade" json:"maximumGrade"`
  ElevationHigh float64 `bson:"elevationHigh" json:"elevationHigh"`
  ElevationLow float64 `bson:"elevationLow" json:"elevationLow"`
  StartPoint  []float64 `bson:"startPoint" json:"startPoint"`
  EndPoint  []float64 `bson:"endPoint" json:"endPoint"`
  ClimbCategory int `bson:"climbCategory" json:"climbCategory"`
  City string `bson:"city" json:"city"`
  State string `bson:"state" json:"state"`
  Country string `bson:"country" json:"country"`
  Private bool `bson:"private" json:"private"`
  Starred bool `bson:"starred" json:"starred"`
  Created time.Time `bson:"created" json:"created"`
  Updated time.Time `bson:"updated" json:"updated"`
  TotalElevationGain float64 `bson:"totalElevationGain" json:"totalElevationGain"`
  Map string `bson:"map" json:"map"`
  EffortCount int `bson:"effortCount" json:"effortCount"`
  AthleteCount int `bson:"athleteCount" json:"athleteCount"`
  Hazardous bool `bson:"hazardous" json:"hazardous"`
  StarCount int `bson:"starCount" json:"starCount"`
}
