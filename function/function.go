package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func test(my_func func(int) int) {
	fmt.Println(my_func(7))
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	test_func := func(x int) int {
		return x * 7
	}
	test(test_func)
}
