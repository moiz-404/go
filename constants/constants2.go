package main

import (
	"fmt"
)

const slogan string = "lets just do golang for"

func main() {
	const day = 15
	const e int = 5
	const d = "days"
	fmt.Println(slogan + " 15 " + d)
	fmt.Println(slogan, day, d)
	fmt.Println("Now", e, "left out of", day)

}
