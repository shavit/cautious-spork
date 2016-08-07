package controllers

import (
  "github.com/revel/revel"
  "log"
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
