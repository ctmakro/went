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

func dialonce(){
  conn, err := net.Dial("tcp", "baidu.com:80")
  if err != nil {print(err);return}

  conn.Write([]byte(`GET / HTTP/1.0
    Accept: text/html

    `))

  

}

func main(){
  dialonce()
}
