package main

import "fmt"

type Livro struct {
	Título string
	Autor  string
	Preço  float64
}

func aplicarDesconto(l *Livro) {
	l.Preço = l.Preço * 0.9
}

func main() {
	livro := Livro{
		Título: "Livro X",
		Autor:  "Autor Y",
		Preço:  50.0,
	}

	fmt.Println("Preço antes:", livro.Preço)
	aplicarDesconto(&livro)
	fmt.Println("Preço depois:", livro.Preço)
}
