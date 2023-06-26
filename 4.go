package main

import "fmt"

func soma(ptr *int) {
	num := *ptr % 100
	digit1 := num / 10
	digit2 := num % 10
	sum := digit1 + digit2
	*ptr = sum
}

func main() {
	num := 1234
	soma(&num)
	fmt.Println(num) 
}
