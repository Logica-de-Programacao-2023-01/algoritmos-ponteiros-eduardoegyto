package main

import "fmt"

type Livro struct {
	Título string
	Autor  string
}

func alterarTítulo(l *Livro) {
	if l.Autor == "Anônimo" {
		l.Título = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Título: "Livro X",
		Autor:  "Anônimo",
	}

	fmt.Println("Título antes:", livro.Título)
	alterarTítulo(&livro)
	fmt.Println("Título depois:", livro.Título)
}
