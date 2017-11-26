package main

import (
  "net"
  "github.com/armon/go-socks5"
)

var QsClientPort = 8118
var QsServerPort = 8338
var QsSocksPort = 8228 // must not be occupied on server
var LocalHost = "127.0.0.1"

func initSocks5Server(){
  // Create a SOCKS5 server
  conf := &socks5.Config{}
  server, err := socks5.New(conf)
  if err != nil {
    panic(err)
  }

  // Create SOCKS5 proxy on localhost
  if err := server.ListenAndServe("tcp", "127.0.0.1:"+str(QsSocksPort)); err != nil {
    panic(err)
  }
}

var idCounter = 0
func CreateTransparentTunnel(fromPort int, toAddr string, toPort int){
  //1. listen to a port
  Listen(fromPort, func(incoming net.Conn){
    var idstr = "[" + str(idCounter) + "]"
    idCounter++

    var remoteAddr = incoming.RemoteAddr().String()
    print(idstr, "Accepted connection from", remoteAddr)

    // 2. for each incoming connection
    // dial the destination:
    var outgoing, err = Dial(toAddr,toPort)
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
}

func QuietsocksServerInit(){
  go initSocks5Server()
  print("Socks5 listening on port",QsSocksPort)

  CreateTransparentTunnel(QsServerPort,"127.0.0.1",QsSocksPort)

  // listen successful
  print("Please connect with quietsocks client to this-machine:" + str(QsServerPort))
}

func QuietsocksClientInit(dest string){ // dest: destination server addr
  // run as a client of the quietsocks protocol.
  print("destination specified:", dest)

  CreateTransparentTunnel(QsClientPort,dest,QsServerPort)

  // listen successful
  print("Please set the SOCKS5 proxy of your browser to 127.0.0.1:"+str(QsClientPort))

}