package controllers

import (
  "github.com/revel/revel"
  "os"
  "io/ioutil"
  "log"
  "net/http"
  "encoding/json"
  "maps/app/models"
)

type Segments struct {
	*revel.Controller
}

func (c Segments) Explore(term string) revel.Result {
  var err error
  var accessToken string = os.Getenv("STRAVA_ACCESS_TOKEN")
  var client *http.Client = &http.Client{}
  var res *http.Response
  var req *http.Request

  req, err = http.NewRequest("GET", "https://www.strava.com/api/v3/segments/explore?bounds=37.821362,-122.505373,37.842038,-122.465977", nil)
  if err != nil {
    log.Fatal(err)
  }
  req.Header.Add("Authorization", "Bearer "+accessToken)

  res, err = client.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()

  // Read the API response
  data, err := ioutil.ReadAll(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  // Create a JSON response
  var jsonData struct {
    Segments []models.Segment `json: "segments"`
  }
  if json.Unmarshal(data, &jsonData); err != nil {
    log.Fatal(err)
  }

  // Save the data
  

  return c.RenderJson(jsonData)
}
