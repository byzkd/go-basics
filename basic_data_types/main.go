package main

import "fmt"

func main() {

	name := "Beyza"
	var age uint8 = 23
	isMarried := false
	var weight = 53.5

	age = 41
	

	fmt.Println(name, age, isMarried, weight)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n",age)
	fmt.Printf("%T\n", isMarried)
	fmt.Printf("%T", weight)
}
