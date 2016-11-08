package controllers

import (
  "github.com/revel/revel"
  "log"
  "os"
  "maps/app/models"
)

type Routes struct {
	*revel.Controller
}

func (c Routes) Index() revel.Result {
  var err error
  var route models.Route
  var routes []models.Route

  routes, err = route.All()
  if err != nil {
    log.Fatal(err)
  }

	return c.RenderJson(routes)
}

func (c Routes) StaticRoutes() revel.Result {
  var err error
  var f *os.File

  f, err = os.Open(revel.BasePath+"/public/js/rouets.json")
  defer f.Close()
  if err != nil {
    log.Fatal(err)
  }

  return c.RenderFile(f, revel.Attachment)
}
