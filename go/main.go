package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 4
const delay = 5


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
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento(){
	fmt.Println("Monitorando...")	

    sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites{
			fmt.Println("Testando o site", i, ":", site )
			testaSite(site)
		}	
		time.Sleep(delay * time.Second)	
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string){
	
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	}else{
		fmt.Println("Site:", site, "esta com problemas. Status code:", resp.StatusCode)
	}
}

func lerSitesDoArquivo()[]string{

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil{
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for{

		linha, err := leitor.ReadString('\n')

		linha = strings.TrimSpace(linha)
		
		sites = append(sites, linha)

		if err == io.EOF{
			break
		}
	
	}

	arquivo.Close()

	return sites
}