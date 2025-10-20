package main

import (
	"fmt"
	"os"
)

func main() {
	var x int
	fmt.Println("Введите имя: ")
	fmt.Fscan(os.Stdin, &x)
	for x < 12307 {
		if x < 0 {
			x *= -1
		}
		if x%7 == 0 {
			x *= 39
		}
		if x%9 == 0 {
			x *= 13
			x += 1
			fmt.Println(x)
			continue
		} else {
			x += 2
			x *= 3
		}
		fmt.Println(x)
	}
}
