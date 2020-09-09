package main

import (
	"fmt"
)

func main() {

	arr := []string{"-214748364801", "-214748364802"}
	max := ""
	secondMax := ""
	if len(arr) < 0 {
		fmt.Println(-1)
	} else {
		for _, number := range arr {
			if number > max {
				secondMax = max
				max = number
			}
			if number > secondMax && number < max {
				secondMax = number
			}
		}
		if secondMax == "" {
			fmt.Println(-1)
		} else {
			fmt.Println(secondMax)
		}
	}
}
