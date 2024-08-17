package main

import (
	"fmt"
)

func main() {
	versao := 1.1
	var nome string
	fmt.Println("Hello World!!", nome)
	fmt.Println("A versão desse software é", versao)
	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nome)
	fmt.Println("Bem-vindo,", nome)

}
