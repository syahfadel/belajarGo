package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "selamat malam"
	letters := splitString(word)
	fmt.Println(countLetter(letters))
}

func splitString(word string) []string {
	var result []string
	result = strings.Split(word, "")
	for _, v := range result {
		fmt.Println(v)
	}
	return result
}

func countLetter(letters []string) map[string]int {
	var result map[string]int
	result = map[string]int{}
	for _, v := range letters {
		_, exist := result[v]
		if exist {
			result[v]++
		} else {
			result[v] = 1
		}
	}
	return result
}
