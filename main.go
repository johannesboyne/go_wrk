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
  url := ""

  for _, element := range args {
    if m, _ := regexp.MatchString("^-r.*", element); m {
      i, _ := strconv.Atoi(element[2:len(element)])
      requests = i
    }
  }


  fmt.Printf("Requests: %d \n", requests)

  if m, _ := regexp.MatchString("^http*", args[len(args)-1]); !m {
    fmt.Println("ERROR: Please use a proper URL (e.g.: http(s)://awesomeurl.com %s")
    os.Exit(1)
  } else {
    url = args[len(args)-1]
    connect.Establish(requests, url)
  }

}
