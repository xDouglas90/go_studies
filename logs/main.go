package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoringTimes = 3
const monitoringDelay = 10 * time.Second

func main() {

	showIntro()

	for {
		showMenu()

		command := readCommand()

		switch command {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
			logsReader("log.txt")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando não identificado")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Douglas"
	version := 1.2
	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("O comando escolhido foi", commandRead)

	return commandRead
}

func initMonitoring() {
	fmt.Println("Monitorando...")

	sites := readSitesFromFile()

	for i := 0; i < monitoringTimes; i++ {
		for i, site := range sites {
			siteTest(i, site)
		}
		time.Sleep(monitoringDelay)
	}
}

func siteTest(pos int, site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if res.StatusCode == 200 {
		fmt.Println(pos+1, "|", "site:", site, "> foi carregado com sucesso!")
		logRegister(site, true)
	} else {
		fmt.Println(pos+1, "|", "site:", site, "> está com problemas. Status Code:", res.StatusCode)
		logRegister(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err != nil {
			break
		}
	}

	file.Close()

	return sites
}

func logRegister(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Fprintln(file, time.Now().Format("02/01/2006 15:04:05 -"), site, "- online:", status)

	file.Close()

}

func logsReader(logs string) {
	file, err := os.Open(logs)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		fmt.Println(line)

		if err != nil {
			break
		}
	}

	file.Close()
}
