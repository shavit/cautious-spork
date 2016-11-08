package models

import (
  "gopkg.in/mgo.v2/bson"
  "time"
  "log"
  "maps/app"
)

type Segment struct {
  Id bson.ObjectId `bson:"_id" json:"_id"`
  Uid int `bson:"uid" json:"id"`
  ResourceState int `bson:"resourceState" json:"resource_state"`
  Name string `bson:"name" json:"name"`
  ActivityType string `bson:"activityType" json:"activityType"`
  Distance float64 `bson:"distance" json:"distance"`
  AverageGrade float64 `bson:"averageGrade" json:"avg_grade"`
  MaximumGrade float64 `bson:"maximumGrade" json:"maximum_grade"`
  ElevationHigh float64 `bson:"elevationHigh" json:"elevation_high"`
  ElevationLow float64 `bson:"elevationLow" json:"elevation_low"`
  StartPoint  []float64 `bson:"startPoint" json:"startLatlng"`
  EndPoint  []float64 `bson:"endPoint" json:"endLatlng"`
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
  Path [][]float64 `bson:"path" json:"path"`
  Points string `bson:"points" json:"points"`
  EffortCount int `bson:"effortCount" json:"effort_count"`
  AthleteCount int `bson:"athleteCount" json:"athlete_count"`
  Hazardous bool `bson:"hazardous" json:"hazardous"`
  StarCount int `bson:"starCount" json:"star_count"`
}


func (s *Segment) All() (segments []Segment, e error){
  var data []Segment
  var err error

  err = app.DB.C("segments").Find(bson.M{}).All(&data)
  if err != nil {
    log.Fatal("Error:", err)
  }

  return data, err
}


func (s *Segment) Save(){
  var err error
  var n int

  // Check if the record exists
  n, err = app.DB.C("segments").Find(bson.M{"uid": s.Uid}).Count()
  if err != nil {
    log.Fatal(err)
  }
  log.Println("---> Segment count", n, s.Uid)

  // Insert if the record does not exists
  if n > 0 {
    log.Printf("---> Segment %v exists\n", s.Id.Hex())
    return
  }

  s.Id = bson.NewObjectId()
  s.DecodePolyline()
  log.Println("---> Points ", s.Path)
  err = app.DB.C("segments").Insert(&s)
  if err != nil {
    log.Fatal(err)
  }

}


func (s *Segment) DecodePolyline() {
	var count, index int
  var f int = 1.0e5
  var pointLonLat []float64
  var points [][]float64

	tempLatLng := [2]int{0, 0}

	for index < len(s.Points) {
		var result int
		var b = 0x20
		var shift uint

		for b >= 0x20 {
			b = int(s.Points[index]) - 63
			index++

			result |= (b & 0x1f) << shift
			shift += 5
		}

		// Sign dection
		if result&1 != 0 {
			result = ^(result >> 1)
		} else {
			result = result >> 1
		}

		if count%2 == 0 {
			result += tempLatLng[0]
			tempLatLng[0] = result
		} else {
			result += tempLatLng[1]
			tempLatLng[1] = result

		}
    pointLonLat = append(pointLonLat, float64(tempLatLng[0]) / float64(f))
    pointLonLat = append(pointLonLat, float64(tempLatLng[1]) / float64(f))
    points = append(points, pointLonLat)
    // Clear the pair
    pointLonLat = nil

		count++
	}

  s.Path = points
}
