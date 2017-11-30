package main

// func main(){
//   // main_dial()
//   main_tunnel()
// }
import (
  // "strings"
  // "bufio"
  // "os"
)

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
  // WaitForString("e")
}

// var stdin *bufio.Reader
// func InitReader() {
//   stdin = bufio.NewReader(os.Stdin)
// }
//
// func ReadLine()(string, error){
//   s,err := stdin.ReadString('\n')
//   if err!=nil{
//     return "",err
//   }else{
//     return s[:len(s)-1],nil
//   }
// }

// func WaitForString(str string){
//   for{
//     print("Enter \""+str+"\" to break")
//     s, err := ReadLine()
//     if err!=nil{
//       print(err)
//     }
//     if s == str {
//       break
//     }
//   }
// }

func WaitForever(){
  for{
    sleep(1.)
  }
}

func QSmain(){
  // InitReader()
  //0. obtain destination server from cli args
  var dest = GetConnectionDestinationFromCli()
  if dest=="not specified" {
    // user did not specify the connect flag
    // start in server mode
    // print("server code not implemented yet...")
    QuietsocksServerInit()
  }else{
    // start in client mode
    QuietsocksClientInit(dest)
  }

  WaitForever()
  // WaitForString("e")
}

func main(){
  QSmain()
}
