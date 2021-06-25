package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	var me string = `{"FirstName":"Armin", "LastName":"Shoeibi"}`
	fmt.Println(me)

	var p Person
	err := json.Unmarshal([]byte(me), &p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p)

	kai := Person{
		FirstName: "Kai",
		LastName:  "Greene",
	}

	byteSlice, err := json.Marshal(kai)

	fmt.Println(string(byteSlice))
}
