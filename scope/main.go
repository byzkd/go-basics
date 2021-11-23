package main

import "fmt"

/*  var packVar = "Package Scope"
var funcVar ="Func(Package) Scope" */

func main() {

/* 	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}
	 funcVar := "Func Scope"

	fmt.Println(funcVar) */ //aynı isme sahip olsa bile kendi scopundaki degiskeni ekrana bastı
	// fmt.Println(packVar)
	// myFunc()

	var name = "beyza"
	name, surname := "arda" , "kucukdilli"
	fmt.Println(name, surname)

}

/* func myFunc() {
	fmt.Println(funcVar)
} */




//Block scope da kısa göstrim yapılır, paket degiskenin de yapılmaz. Cunku tanıyamaz onun ne oldugunu