package main

import "fmt"

func main() {
	var i int = 21
	j := true
	l := 0xf
	u := 0xF13
	var k float64 = 123.456

	fmt.Printf("%v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n")
	fmt.Printf("%t\n", j)
	fmt.Printf("%t\n", j)
	fmt.Printf("%b\n", i)
	fmt.Printf("\u042F\n")
	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%x\n", l)
	fmt.Printf("%X\n", u)
	fmt.Printf("%U\n", 'Я')
	fmt.Printf("%f\n", k)
	fmt.Printf("%E\n", k)
}
