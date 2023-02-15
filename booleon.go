package main

import "fmt"

func main() {
	fmt.Println("Welcome to our service")

	var amount uint
	const total int = 100
	var remaining uint = uint(total)

	for {

		fmt.Println("Enter amount of money you would like to get")
		fmt.Scan(&amount)

		if amount > remaining {

			fmt.Printf("soorry now %v amount is available\n", remaining)
			continue
		}

		remaining = remaining - amount
		fmt.Println("now", remaining, "amount is available")

		if remaining == 0 {
			fmt.Println("next time try it on")
			break
		}
	}
}
