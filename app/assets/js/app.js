//
// Configure the websocket
//
var userName = "guest"
var socket = new WebSocket("ws://"+window.location.host+"/api/v1/activities/stream?user="+userName)

var onMessage = function(e){
  console.log("New event: "+e)
}

socket.onmessage = function(e){
  onMessage(JSON.parse(e.data))
}
