package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; i < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
