package main

import (
  "fmt"
  "os"
  "regexp"
  "strconv"
  "./connect"
)

func main () {
  args := os.Args

  requests    := 1
  connections := 1   // not used atm.
  url := ""

  for _, element := range args {
    if m, _ := regexp.MatchString("^-r.*", element); m {
      i, _ := strconv.Atoi(element[2:len(element)])
      requests = i
    } else if m, _ := regexp.MatchString("^-c.*", element); m {
      i, _ := strconv.Atoi(element[2:len(element)])
      connections = i
    }
  }


  fmt.Printf("Requests: %d, Connections: %d \n", requests, connections)

  if m, _ := regexp.MatchString("^http*", args[len(args)-1]); !m {
    fmt.Println("ERROR: Please use a proper URL (e.g.: http(s)://awesomeurl.com %s")
    os.Exit(1)
  } else {
    url = args[len(args)-1]
    connect.Establish(requests, connections, url)
  }

}
