package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){

	exibeIntroducao()	

	for{
		exibeMenu()
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
	sites := []string{"https://random-status-code.herokuapp.com/","https://www.alura.com.br","https://www.caelum.com.br"}

	for i, site := range sites{
		fmt.Println("Estou passando na posicao", i, "do meu slice e essa posicao tem o site: ", site )
	}

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	}else{
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
	}

}


