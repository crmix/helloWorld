package main

import "fmt"

func main() {

	var first int
	var second int
	var resultofaddition int
	var resultofsubtruction int
	fmt.Println("enter first number :")
	fmt.Scan(&first)
	fmt.Println("enter second number :")
	fmt.Scan(&second)

	resultofaddition = first + second
	fmt.Println("result of addition: ", resultofaddition)
	resultofsubtruction = first - second
	fmt.Println("result of subtruction: ", resultofsubtruction)
}
