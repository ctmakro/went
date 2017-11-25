// not the best approach. when in goland, do the gophers do.

package main

import (
  "fmt"
  // "net"
)
func print(s...interface{}){fmt.Println(s...)}
type any interface{}
func str(instance any) string{
  return fmt.Sprintf("%v",instance)
}

func div(a float32,b float32) float32{
  if b==0. {
    panic("divide by zero")
  }
  return a/b
}

func try(f func(), catch func(any)){
  defer func(){
    if err := recover(); err!=nil {
      catch(err)
    }
  }()
  f()
}

func main(){
  try(func(){
    print("slash", div(.1,0.))
    print("slash", div(.1,.3))
  },func(e any){
    print("error caught:",e)
  })
}
