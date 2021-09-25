package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	readFile, err := os.Open("Arquivos para verificação/teste.c")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	for _, line := range lines {
		fmt.Println(line)
	}
}

func ContadorDePalavras() { //Abrir arquivo e verificar a incidencia de cada palavra

}

func GeradorDeToken() { //Troca palvras chave por tokens predenidos

}

func MaiorSubstring() { //Dado dois arquivos o algoritmo retorna a maior substring detectada

}
