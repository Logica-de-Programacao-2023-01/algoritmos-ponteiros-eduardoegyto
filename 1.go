package main

import "fmt"

func soma(ptr *int, n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	*ptr = sum
}

func main() {
	var num int
	soma(&num, 5)
	fmt.Println(num) 
}
