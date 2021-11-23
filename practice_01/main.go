package main

import "fmt"

func main() {

/* 	var studentName string = "John Doe"
	var grade int = 77
	var isPassed bool = true */

/* 	studentName := "John Doe"
	grade := 77
	isPassed := true */

/* 	var (
		studentName string = "John Doe"
		grade int = 77
		isPassed = true
	)  */

	// var studentName, grade, isPassed = "John Doe", 77, true

	studentName, grade, isPassed := "John Doe", 77, true
	fmt.Println(studentName, grade, isPassed)

	// Declaration Kavramı : "var studentName string" bu ifade ile declaration yapmış olduk yani studentName isminde string tipli bir değişken kullanacağımı belirttim.
	// Assing Kavramı : "var grade int = 77" bu değeri hafızada kapladığı yere yerleştir anlamına geliyor. 
	// Initialization Kavramı : Bir değişkeni oluşturup başlatma anlamına gelir.
	// Initial Value : İlk değer anlamına gelir.
	//Dynamic Data Type and Static Data Type : 
	// Python Dinamik veri tiplidir, yani veri tipini int verdiysek sonradan stringe dönüştürebiliriz. Golang static veri tiplidir, bir değişken numeric tipliyse stringe dönüştürülemez.

/* 	var x int = 77
	fmt.Println(x)
	x := "beyza" 
	fmt.Println(x) */ //HATA ALIRIZ

	// := hem declare eder hem de assign eder, = ise sadece assign eder.
/* 	studentName := "john Doe"
	studentName := "Beyza kucukdilli" */ // HATA ALIRIZ

/* 	isim := "John Doe"
	isim = "Beyza Kucukdilli"
	fmt.Println(isim) */ //DÜZGÜN ÇALIŞIR
	
}