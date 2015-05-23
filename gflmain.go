package main

import "fmt"
import "github.com/garclak/gflgo/gflconst"



func main() {
  l := gflconst.LogonType

	fmt.Printf("Hello, world!\n" + l.ConstVal("admin"))
}
