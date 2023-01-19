package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main(){
	exibeNomes()
	//exibeIntroducao()	

	for{
//exibeMenu()
		comando := leComando()
		
		switch comando {
			case 1:			
				iniciarMonitoramento()
			case 2:
				fmt.Println("Exibindo logs...")	
			case 0:
				fmt.Println("Saindo do programa")
				os.Exit(0)		
			default:
				fmt.Println("Não conheço este comando")
				os.Exit(-1)
		}
	}

	
	
}

func exibeIntroducao(){
	nome := "Jaque"
	versao := 1.1
	fmt.Println("Olá, sra", nome)
	fmt.Println("Este programa esta na versão", versao)
}

func exibeMenu(){
	fmt.Println("-----------------------------------")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando()int {	
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento(){
	fmt.Println("Monitorando...")	
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"




	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	}else{
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
	}

}


func exibeNomes(){
	nomes := []string{"Douglas", "Daniel", "Bernardo"}	
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu do slice tem", len(nomes), "itens")
	fmt.Println("O meu do slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu do slice tem", len(nomes), "itens")
	fmt.Println("O meu do slice tem capacidade para", cap(nomes), "itens")

}
