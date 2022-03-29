package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for z := 1; z <= 10; z++ {
		res := 2 * z
		fmt.Println("2 x", z, "=", res)
	}

}
