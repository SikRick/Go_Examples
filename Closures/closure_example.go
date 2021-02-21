package main

import "fmt"

//simpleFunction is a normal names function
func simpleFunction(arg string) {
	fmt.Println(arg)
}

func generateAnonymousFunction() func(string) {
	return func(arg string) {
		fmt.Println(arg)
	}
}

func main() {
	simpleFunction("Hello World!")

	//An anonymous function inside the main function
	func(arg string) {
		fmt.Printf("Im inside the main and i say %s \n", arg)
	}("Hey There!")

	newFunction := generateAnonymousFunction()
	newFunction("Hello from the new function!")
}
