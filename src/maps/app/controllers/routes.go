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
  var route models.Route
  log.Println(route)
  log.Println(route.All())

	return c.RenderJson([]byte("{}"))
}
