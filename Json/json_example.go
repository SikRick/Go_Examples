package main

import (
	"encoding/json"
	"fmt"
)

type profile struct {
	Name       string
	Place      string
	Occupation string
}

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

	//decoding a json string
	var unMarshalled profile
	err = json.Unmarshal(jsonString, &unMarshalled)
	if err != nil {
		fmt.Println("error decoding json")
		return
	}
	fmt.Println(unMarshalled)
}
