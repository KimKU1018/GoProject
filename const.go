package main

import "fmt"

func main() {
	const PI1 float64 = 3.1415926535
	var PI2 float64 = 3.1415926565

	PI2 = 4

	fmt.Println("원주율: %f\n", PI1)
	fmt.Println("원주율: %f\n", PI2)

}
