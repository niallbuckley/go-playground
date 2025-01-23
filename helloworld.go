package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Message string `json:"message"`
}

func main() {
	var variableString = "world"

	// Create a Data struct instance
	data := Data{
		Message: "hello " + variableString,
	}

	// Convert the struct to JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))
}

