package controllers

import "github.com/revel/revel"

type Routes struct {
	*revel.Controller
}

func (c Routes) Index() revel.Result {
	return c.RenderJson([]byte("{}"))
}
