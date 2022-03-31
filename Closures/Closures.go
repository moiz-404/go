package main

import "fmt"

func return_anony_func() func(string) {
	return func(msg string) {
		fmt.Println(msg)
	}
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	func(msg string) {
		fmt.Println(msg)
	}("hello")

	print_anony1 := func(msg string) {
		fmt.Println(msg)
	}
	print_anony1("hello retrun anonymous function")

	print_anony := return_anony_func()
	print_anony("hello retrun anonymous function")

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
