package main

import(
  "net"
  "io"
)

func DirectionalPipe(r net.Conn, w net.Conn){

  var rbuf = make([]byte, 2048)
  for {
    n,err := r.Read(rbuf)
    if err!=nil {
      if err==io.EOF{
        // EOF need not error report
      }else{
        print("!read error", err)
      }

      w.Close()
      r.Close()
      break
    }

    n,err = w.Write(rbuf[:n])
    if err!=nil {
      print("!write error", err)
      r.Close()
      w.Close()
      break
    }
  }
}

// reverse every byte of the buffer (operate on the buffer directly)
func reverseBytes(buf *[]byte) {
  // for i:=0; i<len(*buf); i++{
  //   (*buf)[i] = byte(255) - (*buf)[i]
  // }
  for i:= range *buf{
    (*buf)[i] = byte(255) - (*buf)[i]
  }
}

func DirectionalPipeWithReversedBytes(r net.Conn, w net.Conn, c chan string){
  var idstr = <- c
  var rbuf = make([]byte, 8192)
  for {
    n,err := r.Read(rbuf)
    if err!=nil {
      if err == io.EOF{
        // EOF need not error report
      }else{
        print(idstr, "read error:", err)
      }

      w.Close()
      r.Close()
      break
    }

    var reversed = rbuf[:n]
    reverseBytes(&reversed)

    n,err = w.Write(reversed)
    if err!=nil {
      print(idstr, "write error:", err)
      r.Close()
      w.Close()
      break
    }
  }
  c <- "nothing" // signal end of execution
}

func TunnelSetup(){
  //1. listen to a port
  go Listen(1080, func(incoming net.Conn){
    // 2. for each incoming connection
    // dial the destination:

    outgoing, err := Dial("127.0.0.1",8118)
    if err!=nil {
      print("outgoing dial err",err)
      incoming.Close()
    }

    // 3. pipe two into each other
    go DirectionalPipe(incoming, outgoing)
    go DirectionalPipe(outgoing, incoming)

    // 4. dont care
  })
}
