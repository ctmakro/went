package main

import (
  // "fmt"
  "net"
  // "os"
  // "time"
  // "io"
  // "reflect"
)

func dialonce(){
  var err any // err of type any (interface{})

  connection, err := net.Dial("tcp", "www.baidu.com:80")
  if err != nil {print(err);return}

  defer connection.Close()

  n,err := connection.Write([]byte("GET / HTTP/1.0\r\nAccept: text/html\r\n\r\n"))
  if err != nil{print(err);return}

  print("sent",n,"bytes")

  // response sent, now try to read from...

  readed, err := ReadTillEOF(connection)
  if err!=nil {print(err);return}

  print(string(readed))

}

// dial the destination and return connection on success.
func Dial(dest string, port int) (net.Conn, any){
  var err any

  connection, err := net.Dial("tcp", dest+":"+str(port))
  if err != nil {
    print(err)
    return nil, err
  }

  return connection, nil
}

// func main(){
//   dialonce()
// }
