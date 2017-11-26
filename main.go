package main

// func main(){
//   // main_dial()
//   main_tunnel()
// }

func main_listen(){
  // listen in another goroutine please
  go listen()
  print("this is a stupid http server. it is listening carefully...")
  sleep(11.5)
  print("it ended.")
}

func main_dial(){
  dialonce()
}

func main_tunnel(){
  TunnelSetup()
  WaitForKey()
}

func WaitForKey(){
  print("Press any key to exit")
  scan()
}

func QSmain(){
  //0. obtain destination server from cli args
  var dest = GetConnectionDestinationFromCli()
  if dest=="not specified" {
    // user did not specify the connect flag
    // start in server mode
    print("server code not implemented yet...")
  }else{
    // start in client mode
    QuietsocksClientInit(dest)
  }

  WaitForKey()
}

func main(){
  QSmain()
}
