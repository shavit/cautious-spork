package socket

import (
  "log"
  List "container/list"
  "time"
)

type Event struct {
  Type string // joined, left, message
  User string
  Timestamp int // Unix timestamp
  Body string // The user message
}

type Subscription struct {
  Events []Event // Send all the events
  New <-chan Event // Incoming events
}

const archiveSize int = 10
var (
  // Get all the events
  subscribe chan (chan<- Subscription) = make(chan (chan<- Subscription), 40)
  // Unsubscribe a channel
  unsubscribe chan (<-chan Event) = make(chan (<-chan Event), 40)
  // Receive and publish events
  publish chan Event = make(chan Event, 40)
)

func init(){
  go pubsub()
}

func pubsub(){
  var archive *List.List = List.New()
  var subscribers *List.List = List.New()

  for {
    select {
    case ch := <- subscribe:
      var events []Event
      for e := archive.Front(); e != nil; e = e.Next() {
        // Append and cast each item in the list to Event
        events = append(events, e.Value.(Event))
      }

      subscriber := make(chan Event, 40)
      subscribers.PushBack(subscriber)
      ch <- Subscription{events, subscriber}

    case event := <- publish:
      for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
        ch.Value.(chan Event) <- event
      }

      if archive.Len() >= archiveSize {
        archive.Remove(archive.Front())
      }

      archive.PushBack(event)

    case unsub := <- unsubscribe:
      for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
        if ch.Value.(chan Event) == unsub {
          subscribers.Remove(ch)
          break
        }
      }
    }
  }
}


//
// Methods
//

func (s Subscription) Cancel() {
  // Unsubscribe the channel
  unsubscribe <- s.New
  // Drain it just in case there was a pending publish.
  Drain(s.New)
}

func NewEvent(t string, user string, body string) (e Event) {
  return Event{t, user, int(time.Now().Unix()), body}
}

func Subscribe() (subscription Subscription) {
  var res (chan Subscription) = make(chan Subscription)
  subscribe <- res

  return <- res
}

func Join(user string) {
  publish <- NewEvent("joined", user, "")
}

func Say(user string, body string){
  publish <- NewEvent("message", user, "")
}

func Leave(user string){
  publish <- NewEvent("left", user, "")
}

func Drain(ch <-chan Event){
  var ok bool

  for {
    select {
    case _, ok = <-ch:
      // if ok {
      //   log.Println("---> Reading values")
      // } else {
      //   log.Println("---> Channel is closed")
      //   return
      // }
      if ok == false {
        log.Println("---> Channel is closed")
        return
      }
    default:
      return
    }
  }
}
