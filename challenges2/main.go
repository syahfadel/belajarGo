package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Printf("nilai i = %d\n", i)
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			unicodes := [7]int{0x421, 0x410, 0x428, 0x410, 0x420, 0x412, 0x41E}
			for k, unicode := range unicodes {
				fmt.Printf("character %#U starts at byte position %d\n", unicode, k*2)

			}
		} else {
			fmt.Printf("Nilai j = %v\n", j)
		}

	}
}
