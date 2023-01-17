package main

import (
	"fmt"
)

func main(){
	nome := "Jaque"
	versao := 1.1
	fmt.Println("Olá, sra", nome)
	fmt.Println("Este programa esta na versão", versao)
	fmt.Println("-----------------------------------")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")		
	}else if comando == 2 {
		fmt.Println("Exibindo logs...")				
	}else if comando == 0 {
		fmt.Println("Saindo do programa")
	}else{
		fmt.Println("Não conheço este comando")
	}

}