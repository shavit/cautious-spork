package controllers

import (
  "github.com/revel/revel"
  "log"
  "maps/app"
  "maps/app/models"
)

type Routes struct {
	*revel.Controller
}

func (c Routes) Index() revel.Result {
  var route models.Route
  log.Println(route)
  log.Println(app.DB)

	return c.RenderJson([]byte("{}"))
}
