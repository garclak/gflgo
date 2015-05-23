package main

import "fmt"
import "strconv"
import "github.com/garclak/gflgo/gflconst"



func main() {
	l := gflconst.NewLogonType()
	a := l.Type("normal")
	//fmt.Println(a)
	fmt.Printf("Type: " + strconv.Itoa(a) + "\n")
}
