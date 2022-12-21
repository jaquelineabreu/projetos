package main

import (
	"fmt"
	"reflect"
)

func main(){
	// var nome string = "Jaque"
	// var idade int = 24
	// var nome = "Jaque"
	// var idade  = 24
	// var versao = 1.1
	nome := "Jaque"
	idade  := 24
	versao := 1.1
	fmt.Println("Olá, sra", nome, "sua idade é", idade)
	fmt.Println("Este programa esta na versão", versao)

	fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))
}