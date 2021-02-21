package closures

import "fmt"

//simpleFunction is a normal names function
func simpleFunction(arg string) {
	fmt.Println(arg)
}

func main() {
	simpleFunction("Hello World!")
}
