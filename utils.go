package main

import (
  "net"
  "reflect"
  "fmt"
  "time"
  "io"
)

// let's get pythonic
var print = fmt.Println
var scan = fmt.Scanln
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

func ReadTillCrLf(c net.Conn) ([]byte, any){
  var bigbuf []byte // bigbuf is unbounded
  rbuf := make([]byte, 512) // rbuf is bounded
  for {
    n,err := c.Read(rbuf)
    if err != nil {
      return nil,err
    }

    // if no error met, meaning we read something from the socket.
    print("read",n,"bytes from socket...")

    bigbuf = append(bigbuf,rbuf[:n]...) //append rbuf[0:n] to buf

    // if last four byte is /r/n/r/n, we reached the end of http request.
    if reflect.DeepEqual(bigbuf[len(bigbuf)-4:], []byte("\r\n\r\n")){
      print("request ended")
      // print(string(bigbuf))
      // break
      return bigbuf, nil
    }

  }
}

func ReadTillEOF(c net.Conn) ([]byte, any){
  var bigbuf []byte // bigbuf is unbounded
  rbuf := make([]byte, 512) // rbuf is bounded
  for {
    n,err := c.Read(rbuf)
    if err != nil {
      if err == io.EOF{
        print("EOF met")
        return bigbuf, nil // finally we reached the end.
      }
      return nil,err
    }

    // if no error met, meaning we read something from the socket.
    print("read",n,"bytes from socket...")

    bigbuf = append(bigbuf,rbuf[:n]...) //append rbuf[0:n] to buf
  }
}
