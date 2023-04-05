package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://user1:123456@developer.com:80/hello?name=airell&age=23"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	if u.Hostname() == "" {
		panic("ivalid URL")
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name, age)

	// dapat dimanfaatkan untuk authentifikasi
	fmt.Println("Authentication Username :", u.User.Username())
	pass, isErr := u.User.Password()
	if !isErr {
		panic(isErr)
	}
	fmt.Println("Authentication Username :", pass)
}
