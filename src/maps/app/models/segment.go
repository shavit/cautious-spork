package models

import (
  "gopkg.in/mgo.v2/bson"
  "time"
)

type Segment struct {
  Id bson.ObjectId `bson:"_id" json:"_id"`
  Uid string `bson:"uid" json:"uid"`
  ResourceState int `bson:"resourceState" json:"resource_state"`
  Name string `bson:"name" json:"name"`
  ActivityType string `bson:"activityType" json:"activityType"`
  Distance float64 `bson:"distance" json:"distance"`
  AverageGrade float64 `bson:"averageGrade" json:"avg_grade"`
  MaximumGrade float64 `bson:"maximumGrade" json:"maximum_grade"`
  ElevationHigh float64 `bson:"elevationHigh" json:"elevation_high"`
  ElevationLow float64 `bson:"elevationLow" json:"elevation_low"`
  StartPoint  []float64 `bson:"startPoint" json:"start_latlng"`
  EndPoint  []float64 `bson:"endPoint" json:"end_latlng"`
  ClimbCategory int `bson:"climbCategory" json:"climb_category"`
  City string `bson:"city" json:"city"`
  State string `bson:"state" json:"state"`
  Country string `bson:"country" json:"country"`
  Private bool `bson:"private" json:"private"`
  Starred bool `bson:"starred" json:"starred"`
  Created time.Time `bson:"created" json:"created_at"`
  Updated time.Time `bson:"updated" json:"updated_at"`
  TotalElevationGain float64 `bson:"totalElevationGain" json:"totalElevationGain"`
  Map string `bson:"map" json:"map"`
  Points string `bson:"points" json:"points"`
  EffortCount int `bson:"effortCount" json:"effort_count"`
  AthleteCount int `bson:"athleteCount" json:"athlete_count"`
  Hazardous bool `bson:"hazardous" json:"hazardous"`
  StarCount int `bson:"starCount" json:"star_count"`
}
