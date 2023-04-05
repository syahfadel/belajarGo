package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	// terdapat tag yang disesuaikan dengan field pada data json
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

// struct untuk encoding
type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	// json yang akan di decoding
	var jsonString = `
	{
		"full_name" : "Syah Fadel Putra Dwingga",
		"email" : "syahfadel@email.com",
		"age" : 23
	}
	`

	var result Employee

	// unmarshal digunakan untuk mendecode JSON kepada struct
	// argumen pertama merupakan slice of byte daru JSON. Argumen kedua meletakan pointer dari variabel yang akan disimpan
	fmt.Println("===================DECODING TO STUCT====================")
	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name: ", result.FullName)
	fmt.Println("email: ", result.Email)
	fmt.Println("age: ", result.Age)

	fmt.Println("===================DECODING TO MAP====================")
	// decoding json to map
	var resultMap map[string]interface{}
	var err2 = json.Unmarshal([]byte(jsonString), &resultMap)
	if err2 != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name: ", resultMap["full_name"])
	fmt.Println("email: ", resultMap["email"])
	fmt.Println("age: ", resultMap["age"])

	fmt.Println("===================DECODING TO EMPTY INTERFACE====================")
	// decoding json to empty interface
	var temp interface{}
	var err3 = json.Unmarshal([]byte(jsonString), &temp)
	if err3 != nil {
		fmt.Println(err.Error())
		return
	}

	// terdapat type assertion
	var resultEmpty = temp.(map[string]interface{})

	fmt.Println("full_name: ", resultEmpty["full_name"])
	fmt.Println("email: ", resultEmpty["email"])
	fmt.Println("age: ", resultEmpty["age"])

	fmt.Println("===================DECODING Array of JSON to SLICE OF STRUCT====================")
	// decoding json to empty interface
	var jsonArrayString = `[
		{
			"full_name" : "Syah Fadel Putra Dwingga",
			"email" : "syahfadel@email.com",
			"age" : 23
		},
		{
			"full_name" : "Putra Dwingga",
		"email" : "putra@email.com",
		"age" : 22
	}
	]`
	var res []Employee
	var err4 = json.Unmarshal([]byte(jsonArrayString), &res)
	if err4 != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range res {
		fmt.Printf("index %d: %+v\n", i+1, v)
	}
	fmt.Println("===================Encoding objek ke JSON====================")
	var objek = []User{{"Syah Fadel", 23}, {"Putra", 18}}

	var jsonData, err5 = json.Marshal(objek)

	if err5 != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonStringEncode = string(jsonData)
	fmt.Println(jsonStringEncode)
}
