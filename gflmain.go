package main

import "fmt"
import "strconv"
import "github.com/garclak/gflgo/gflConst"



func main() {
	l := gflConst.NewStateC()
	fmt.Printf("Type: " + strconv.Itoa(l.Const("admin")) + "\n")
}

/* Console Text Entry
import "os"
import "bufio"

reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString('\n')
fmt.Println(text)
*/
