package controllers

import (
  "github.com/revel/revel"
  "golang.org/x/net/websocket"
  "maps/app/socket"
  "log"
)

type Activities struct {
	*revel.Controller
}

// Subscribe to the activity stream
func (c Activities) Stream(user string, ws *websocket.Conn) revel.Result {
  log.Println("---> New subscriber")
  var err error
  var ok bool
  var event socket.Event
  var newMessages chan string = make(chan string)
  var subscription socket.Subscription = socket.Subscribe()

  defer subscription.Cancel()
  socket.Join(user)
  defer socket.Leave(user)

  for _, event = range subscription.Events {
    if websocket.JSON.Send(ws, &event) != nil {
      log.Println("---> User disconnected")
      return nil
    }
  }

  // In order to select between websoscket messages and subscription events,
  //  we need to channel websocket events together.
  var msg string
  go func (){
    // var msg string

    for {
      err = websocket.Message.Receive(ws, &msg)
      if err != nil {
        close(newMessages)
        return
      }

      newMessages <- msg
    }
  }()

  // Listen for new events
  for {
    select {
    case event := <- subscription.New:
      if websocket.JSON.Send(ws, &event) != nil {
        log.Println("---> User disconnected")
        return nil
      }

    case msg, ok = <- newMessages:
      if ok == false {
        // If the channel is closed, the users unsubscribed
        return nil
      }

      // if the channel is available, send a message
      socket.Say(user, msg)
    }
  }

  return nil
	// return c.RenderJson([]byte(""))
}
