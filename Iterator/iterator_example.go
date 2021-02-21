package main

import "fmt"

//using closure to create a sort of iterator/ generator
func arrayIterator(array []int) func() int {
	length := len(array)
	index := -1
	return func() int {
		if index <= length {
			index++
		}
		return array[index]
	}
}

func main() {

	array := []int{1, 2, 3, 4, 5}
	iterrator := arrayIterator(array)
	fmt.Println(iterrator())
	fmt.Println(iterrator())
	fmt.Println(iterrator())

}
