package main

import "fmt"

//simpleFunction is a normal names function
func simpleFunction(arg string) {
	fmt.Println(arg)
}

func main() {
	simpleFunction("Hello World!")

	//An anonymous function inside the main function
	func(arg string) {
		fmt.Printf("Im inside the main and i say %s \n", arg)
	}("Hey There!")
}
