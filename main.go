package main

func main(){
  // main_dial()
}

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
