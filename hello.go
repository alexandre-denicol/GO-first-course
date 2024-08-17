package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Alexandre"
	var versao float32 = 1.1 //não precisa declarar o tipo da variável (padrão float64) e nem precisa da palavra "var"
	idade := 47
	fmt.Println("Hello World!!", nome)
	fmt.Println("Sua idade é", idade)
	fmt.Println("A versão desse software é", versao)
	fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
}
