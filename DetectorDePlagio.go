package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
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
	PrintNumPalvras(lines)
}

func ContadorDePalavras(str []string) map[string]int { //Abrir arquivo e verificar a incidencia de cada palavra
	counts := make(map[string]int)
	f := func(c rune) bool {
		return unicode.IsSpace(c) || c == '.' || c == ',' || c == '(' || c == ')' || c == '{' || c == '}' || c == ';' ||
			c == '=' || c == '+' || c == '-' || c == '<' || c == '>' || c == '|' || c == '\\' || c == '"' || c == '*' ||
			c == '&' || c == '/' || c == '#' || c == '%'
	}
	for _, line := range str {
		wordList := strings.FieldsFunc(line, f)
		for _, word := range wordList {
			_, ok := counts[word]
			if ok {
				counts[word] += 1
			} else {
				counts[word] = 1
			}
		}
	}
	return counts
}

func PrintNumPalvras(lines []string) {
	fmt.Println("=================================")
	for index, element := range ContadorDePalavras(lines) {
		fmt.Println(index, "=>", element)
	}
}

func GeradorDeToken() { //Troca palvras chave por tokens predenidos
}

func MaiorSubstring() { //Dado dois arquivos o algoritmo retorna a maior substring detectada
}