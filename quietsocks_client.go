package main

import (
  "net"
)

var QsClientPort = 8118
var QsServerPort = 8338

func QuietsocksClientInit(dest string){ // dest: destination server addr
  // run as a client of the quietsocks protocol.

  print("destination specified:", dest)

  var counter = 0

  //1. listen to a port
  Listen(QsClientPort, func(incoming net.Conn){
    var idstr = "[" + str(counter) + "]"
    counter++

    var remoteAddr = incoming.RemoteAddr().String()
    print(idstr, "Accepted connection from", remoteAddr)

    // 2. for each incoming connection
    // dial the destination:
    var outgoing, err = Dial(dest,QsServerPort)
    if err!=nil {
      print(idstr, "outgoing dial err:",err)
      incoming.Close()
      return // do not execute further
    }

    var ch1,ch2 = make(chan string), make(chan string) // pass something to em

    // 3. pipe two into each other, reversing the bytes
    go DirectionalPipeWithReversedBytes(incoming, outgoing, ch1)
    go DirectionalPipeWithReversedBytes(outgoing, incoming, ch2)

    // tell them something via the channels
    ch1 <- idstr
    ch2 <- idstr

    // wait for them to end
    <- ch1
    <- ch2

    print(idstr, "both connection should have ended")
  })

  // listen successful
  print("Please set the SOCKS5 proxy of your browser to 127.0.0.1:"+str(QsClientPort))

}
