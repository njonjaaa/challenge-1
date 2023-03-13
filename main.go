package main

import "fmt"

func main() {

	var i = 21

	fmt.Printf("%v \n", i)
	fmt.Printf("%T\n", i)

	var persen = "%"
	fmt.Println(persen)

	var statement bool = true
	fmt.Printf("%t\n", statement)

	russia := 'Я'
	fmt.Printf("%+q\n", russia)

	number := 21
	fmt.Printf("%d\n", number)
	fmt.Printf("%o\n", number)

	huruf := 16
	fmt.Printf("%x\n", huruf)

	var decimal float64 = 123.456
	fmt.Printf("%f\n", decimal)
	fmt.Printf("%e\n", decimal)

}
