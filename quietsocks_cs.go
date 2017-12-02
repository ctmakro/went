package main

import (
  "net"
  "golang.org/x/net/context"
  "github.com/armon/go-socks5"
)

var QsClientPort = 8118
var QsServerPort = 8338
// changed to 80, let's see if the performance gets any better.
var QsSocksPort = 8228 // must not be occupied on server
var LocalHost = "127.0.0.1"

func initSocks5Server(){
  // Create a SOCKS5 server

  // dial function modified from src of github/armon/go-socks5
  var dialfunc = func(ctx context.Context, net_, addr string) (net.Conn, error) {
    print("[socks5] dialing", addr, "...")
		return net.Dial(net_, addr)
	}

  conf := &socks5.Config{}
  conf.Dial = dialfunc

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
func CreateTransparentTunnel(fromPort int, toAddr string, toPort int, asServer bool){
  // define handler
  var handler = func(incoming net.Conn){
    var idstr = "[" + str(idCounter) + "]"
    idCounter++

    var remoteAddr = incoming.RemoteAddr().String()
    // print(idstr, "Accepted connection from", remoteAddr)

    // 2. for each incoming connection
    // dial the destination:
    var outgoing net.Conn
    var err any
    if asServer {
      outgoing, err = Dial(toAddr,toPort)
    } else {
      // outgoing, err = DialWithHttpHeader(toAddr,toPort)
      outgoing, err = Dial(toAddr,toPort)
    }

    if err!=nil {
      print(idstr, "outgoing dial err:",err)
      incoming.Close()
      return // do not execute further
    }

    print(idstr, "accept from", remoteAddr, "to", outgoing.RemoteAddr().String())
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
  }

  //1. listen to a port
  // Listen(fromPort, func(incoming net.Conn){
  if asServer {
    // skip http header in incoming connection
    // ListenSkipHTTPHeader(fromPort, handler)
    Listen(fromPort, handler)
  }else{
    // listen as is
    Listen(fromPort, handler)
  }


}

func QuietsocksServerInit(){
  go initSocks5Server()
  print("Socks5 listening on port",QsSocksPort)

  CreateTransparentTunnel(QsServerPort,"127.0.0.1",QsSocksPort, true)

  // listen successful
  print("Please connect with quietsocks client to this-machine:" + str(QsServerPort))
}

func QuietsocksClientInit(dest string){ // dest: destination server addr
  // run as a client of the quietsocks protocol.
  print("destination specified:", dest)

  CreateTransparentTunnel(QsClientPort,dest,QsServerPort, false)

  // listen successful
  print("Please set the SOCKS5 proxy of your browser to 127.0.0.1:"+str(QsClientPort))

}
