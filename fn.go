package main

import   "fmt"

func main() {

	var firstnum int32
	var secondnum int32
    var result int64
	fmt.Println("enter firstnum")
	fmt.Scan(&firstnum)

	fmt.Println("enter secondnum")
	fmt.Scan(&secondnum)
	result=int64(firstnum)+int64(secondnum)
     
	fmt.Println("result =", result)
}
