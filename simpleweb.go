package main

import (
  // "fmt"
  "net"
  // "os"
  // "time"
  // "io"
  // "reflect"
)

var counter = 0

func handle(c net.Conn){
  // close connection when this function ends, no matter what happens
  defer c.Close()

  counter+=1
  print("["+str(counter)+"] we got connection!!")

  readed, err := ReadTillCrLf(c)
  if err!=nil {
    print("error on RTCL():",err)
  }

  print("total read:", len(readed))
  print("content:", string(readed))

  resp := "HTTP/1.0 200 OK\r\nContent-Type: text/html\r\n\r\n<h1>Hello world!</h1>\r\n"
  c.Write([]byte(resp))
}

func listen(){
  // the listener emits Accept() whenever a new connection arrives
  listener, err := net.Listen("tcp",":8888")
  if err!=nil{
    print("Listen() error:",err)
    return
  }
  for { // poll repeatedly in a blocking fashion
    connection, err := listener.Accept()
    if err!=nil{
      print("Accept() error:",err)
      return
    }

    // start a separate goroutine to handle this specific connection.
    go handle(connection)
  }
}

func Listen(port int, handler func(net.Conn)){
  // the listener emits Accept() whenever a new connection arrives
  var err any
  listener, err := net.Listen("tcp",":"+str(port))
  if err!=nil{
    print("Listen() error:",err)
    panic(err)
  }

  print("Listen() success on", port)

  go func(){
    for { // poll repeatedly in a blocking fashion
      connection, err := listener.Accept()
      if err!=nil{
        print("Accept() error:",err)
        panic(err)
      }

      // start a separate goroutine to handle this specific connection.
      go handler(connection)
    }
  }()
}
