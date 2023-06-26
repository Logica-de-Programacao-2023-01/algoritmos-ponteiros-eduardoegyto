package main

import (
	"fmt"
	"math"
)

func media(ptr *float64) {
	*ptr = (*ptr + math.Pi) / 2
}

func main() {
	num := 3.0
	media(&num)
	fmt.Println(num) 
}
