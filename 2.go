package main

import "fmt"

func numeros(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
}

func main() {
	num := 7
	numeros(&num)
	fmt.Println(num) 
}
