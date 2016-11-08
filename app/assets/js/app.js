
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


function displayRoute(streamData){
  latings = streamData[0].data
  distances = streamData[1].data
  altitudes = streamData[2].data
}

function loadMap(){
  if (window.L) {
    // createMap()
    $("#mapView").html("")
    clearInterval(waitForLInterval)

    map = L.map("mapView", {
      center: latings[Math.round((latings.length)/2)],
      zoom: 10
      })

    polyline = L
      .polyline(streamData[0].data, {
        color: "blue"
      })
      .addTo(map)

    L.tileLayer("http://{s}.tile.osm.org/{z}/{x}/{y}.png", {
      attribution: "&copy; <a href=\"http://osm.org/copyright\">OpenStreetMap</a> contributors"
    })
    .addTo(map)
  }
}

//
//  Run
//

var map
var latings
var distances
var altitudes
var waitForLInterval = setInterval(loadMap, 300)

var routesURL = "/public/js/routes.json"
$.get(routesURL, function(data) {
  streamData = data.data
  displayRoute(streamData)
})
