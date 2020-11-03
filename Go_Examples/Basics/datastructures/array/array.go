package array

import "fmt"

func ArrayExample(){
	arr:= [10]int{1,1,1,1,1,1,1,1,1,1}
	for index, item := range arr{
		fmt.Println(index, item)
	}
}