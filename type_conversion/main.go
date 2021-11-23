package main

import "fmt"

func main() {
	/* 	x := 10
	   	y:= 10.5

	   	fmt.Printf("%v %T\n", x,x)
	   	fmt.Printf("%v %T\n", y,y)

	   	//Type conversion syntax  type(value) => int(y)



	   	fmt.Println(float64(x)+y)
	   	fmt.Printf("%v %T\n", x,x) //type conversion ile yeni bir degisken olusturuyoruz ama esk tipini degistirmiyoruz yani x yine int veri tipinde, degismedi */

	/* 	var a int8 =10
	   	var b int16=15

	   	fmt.Println(a+b) //type mismatched hatası */

	/* 	var x int16 = 128
	   	var y int8

	   	y = int8(x) //128 int8 in sınırını astıgı icin int8'in alabilecegi en kucuk deger olan -128i ekrana bastı

	   	fmt.Println(y) */

	num1 := 106
	str1 := string(num1) //ascii karakter kodu olarak donusturur

	fmt.Printf("%v, %T\n", str1,str1)
	fmt.Printf("%v, %T\n", num1,num1)
}

//type conversition ile donusturdugumuz verilerin string ve numerşc olarak birbirleri ile tutumlu olması gerekir. JS gibi degil!!