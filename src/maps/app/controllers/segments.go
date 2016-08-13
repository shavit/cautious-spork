package controllers

import (
  "github.com/revel/revel"
  "os"
  "log"
  "net/http"
)

type Segments struct {
	*revel.Controller
}

func (c Segments) Explore(term string) revel.Result {
  var err error
  var client *http.Client = &http.Client{}
  var res *http.Response
  var req *http.Request
  var accessToken string = os.Getenv("STRAVA_ACCESS_TOKEN")

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

  log.Println(res.Body)

  return c.RenderJson([]byte(""))
}
