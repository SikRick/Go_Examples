package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	dict := make(map[string]string)

	dict["Name"] = "SikRick"
	dict["Place"] = "Whiterun"
	dict["Occupation"] = "Thane of Whiterun, Arch Mage of the college of winterhold"

	fmt.Println(dict)

	jsonString, err := json.Marshal(dict)
	if err != nil {
		fmt.Println("Couldnt compile to a JSON string")
		return
	}
	fmt.Println(jsonString)
}
