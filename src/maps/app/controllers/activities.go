package controllers

import (
  "github.com/revel/revel"
  "golang.org/x/net/websocket"
  "log"
)

type Activities struct {
	*revel.Controller
}

// Subscribe to the activity stream
func (c Activities) Stream(user string, ws *websocket.Conn) revel.Result {
  log.Println("---> New subscriber")

  // Create a subscription

  // Check if the user subscribed and proceed


	return c.RenderJson([]byte(""))
}
