package main

import (
  "flag"
)

func GetConnectionDestinationFromCli() string {
  // try parse "--connect [addr]" from cli
  var dest = flag.String("connect", "not specified", "The server address you want to connect to as a client" )

  flag.Parse()

  return *dest
}
