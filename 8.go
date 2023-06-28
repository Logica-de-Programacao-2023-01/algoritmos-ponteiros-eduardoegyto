package main

import "fmt"

func modificarValor(p *int) {
	*p = 10
}

func main() {
	var valor int
	fmt.Println("Valor antes:", valor)
	modificarValor(&valor)
	fmt.Println("Valor depois:", valor)
}
