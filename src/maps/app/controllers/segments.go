package controllers

import (
  "github.com/revel/revel"
  "log"
  "net/http"
)

type Segments struct {
	*revel.Controller
}

func (c Segments) Explore(term string) revel.Result {
  var err error
  var res *http.Response

  res, err = http.Get("https://www.strava.com/api/v3/segments/explore?bounds=37.821362,-122.505373,37.842038,-122.465977")
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  log.Println(res.Body)

  return c.RenderJson([]byte(""))
}
