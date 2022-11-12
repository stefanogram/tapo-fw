package main
 
import (
   "fmt"
   "net"
)
 
func main() {
   foo, err := net.InterfaceAddrs()
 
   if err == nil {
      for _, v := range foo {
         fmt.Println(v)
 
      }
   }
}
