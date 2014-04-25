package connect

import (
  "fmt"
  "time"
  "os"
  "net/http"
)

type Ping struct {
  time time.Duration
  status int
  bstatus bool
}

type Connection struct {
  requests, threats, connections int
  url string
}

func pingURL (url string) *Ping {
  t0 := time.Now()
  response, err := http.Get(url)
  if err != nil {
    fmt.Printf("NO SUCH URL?\n")
    os.Exit(1)
  }
  t1 := time.Now()
  return &Ping{
    time:     t1.Sub(t0),
    status:   response.StatusCode,
    bstatus:  true,
  }
}

func BlowPing (url string, ch chan<- *Ping) {
  p := pingURL(url)
  if !p.bstatus {
    fmt.Println("...wow, that last ping was bad!")
  } else {
    ch <- p
  }
}

func Blowpipe (c *Connection) {
  blowpipeChan := make(chan *Ping)
  for i := 0; i < c.requests; i++ {
    go BlowPing(c.url, blowpipeChan)
  }
  p := &Ping{time: time.Second, status: 200, bstatus: true}
  timeA := time.Now()
  timeB := time.Now()
  for i := 0; i < c.requests; i++ {
    p = <-blowpipeChan
    timeA = timeA.Add(p.time)
  }
  fmt.Printf("The average call took %f ms .\n", float64(timeA.Sub(timeB))/float64(c.requests)/float64(1000000))
}

func Establish (requests int, threats int, connections int, url string) {

  c := &Connection{
    requests:     requests,
    threats:      threats,
    connections:  connections,
    url:          url,
  }
  fmt.Printf("Establish Connection to: %s\n", c.url)
  fmt.Printf("\t- ping: %s...", c.url)

  ping := pingURL(c.url)
  if !ping.bstatus {
    fmt.Printf(" - connection refused!")
    os.Exit(1)
  } else {
    fmt.Printf(" - ping: %v status: %d\n", ping.time, ping.status)
  }

  Blowpipe(c)
}
