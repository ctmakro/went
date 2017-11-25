package main

import (
  "fmt"
  "net"
  // "os"
  "time"
  "io"
  "reflect"
)

// let's get pythonic
var print = fmt.Println
type any interface{}
func str(instance any) string{
  return fmt.Sprintf("%v",instance)
}
func sleep(s float32){
  time.Sleep(
    time.Duration(s * float32(time.Second)),
  )
}
// imagine without all above

var counter = 0

func ReadTillEnd(c net.Conn){
  var bigbuf []byte
  rbuf := make([]byte, 1024)
  for {
    n,err := c.Read(rbuf)
    if err != nil {
      if err != io.EOF { // if error is not end of file
        print("read error:", err)
      }else{
        // if end of file met:
        print("read ended with EOF.")
        break
      }
    }else{
      // if no error met, meaning we read something from the socket.
      print("read",n,"bytes from socket...")

      bigbuf = append(bigbuf,rbuf[:n]...) //append rbuf[0:n] to buf

      // if last four byte is /r/n/r/n, we reached the end of http request.
      if reflect.DeepEqual(bigbuf[len(bigbuf)-4:], []byte("\r\n\r\n")){
        print("request ended")
        print(string(bigbuf))
        break
      }
    }
  }
}

func handle(c net.Conn){
  // close connection when this function ends, no matter what happens
  defer c.Close()

  counter+=1
  print("["+str(counter)+"] we got connection!!")

  // bigbuf := make([]byte, 0, 4096)
  var bigbuf []byte
  // rbuf := make([]byte, 1024)
  var rbuf []byte
  for {
    n,err := c.Read(rbuf)
    if err != nil {
      if err != io.EOF { // if error is not end of file
        print("read error:", err)
      }else{
        // if end of file met:
        print("read ended with EOF.")
        break
      }
    }else{
      // if no error met, meaning we read something from the socket.
      print("read",n,"bytes from socket...")

      bigbuf = append(bigbuf,rbuf[:n]...) //append rbuf[0:n] to buf

      // if last four byte is /r/n/r/n, we reached the end of http request.
      if reflect.DeepEqual(bigbuf[len(bigbuf)-4:], []byte("\r\n\r\n")){
        print("request ended")
        print(string(bigbuf))
        break
      }
    }
  }

  print("total read:", len(bigbuf))

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

func main(){
  // listen in another goroutine please
  go listen()
  print("this is a stupid http server. it is listening carefully...")
  sleep(11.5)
  print("it ended.")
}
