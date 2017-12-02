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

const crlf = "\r\n"
const HTTPRequestHeader = "POST /index.php HTTP/1.1" + crlf + "Content-Type: application/octet-stream" + crlf + crlf
const HTTPResponseHeader = "HTTP/1.1 200 OK"+crlf+"Content-Type: application/octet-stream"+crlf+"Connection: Keep-Alive"+crlf+crlf

func DialWithHttpHeader(dest string, port int) (net.Conn, any){
  // print("(reqh)", HTTPRequestHeader)

  conn, err := Dial(dest, port)
  if err!=nil {return nil,err}

  // dial successful. but before we return the connection, why not send a fake http header first?
  n,err1 := conn.Write([]byte(HTTPRequestHeader))
  if err1!=nil{return nil,err1}
  // print("(DWHH)n",n)

  // now the server should respond with a HTTP response header. skip it.
  readed, err2 := ReadTillCrLf(conn)
  if err2!=nil {
    print("error on RTCL():",err2)
    return nil,err
  }
  // print("(DWHH)readed", string(readed))
  _,_ = readed,n
  return conn,nil
}
