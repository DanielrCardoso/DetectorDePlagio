package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	// "runtime"
	"strings"
	"sync"
	"unicode"
)

var wg sync.WaitGroup
var ww sync.WaitGroup

func main() {
	a, _ := fileCount("Arquivos para verificação")
	fmt.Println("Numero de aquivos: ", a)
	fmt.Println("Arquivos encontrados: ")
	//showNameFiles("Arquivos para verificação")
	headers := make([]string, a)

	var fullLine = ""
	var path = ""
	files, err := ioutil.ReadDir("Arquivos para verificação")
	if err != nil {
		log.Fatal(err)
	}
	idx := 0
	var conteudoArqs []string
	for _, f := range files {
		path = "Arquivos para verificação/" + f.Name()
		fmt.Print("Abrindo arquivo: " + path)
		headers[idx] = f.Name()
		idx++
		fullLine = myreadFile(path)
		//fmt.Println(fullLine)
		//PrintNumPalvras(fullLine)
		//fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++=")
		conteudoArqs = append(conteudoArqs, fullLine)
	}

	RunPlag(a, conteudoArqs[0], conteudoArqs, headers, 0)
	// RunPlag(a, conteudoArqs[1], conteudoArqs)

	fmt.Println("head!", headers)

	// for cacatua := range conteudoArqs {
	// 	// wg.Add(1)
	// 	RunPlag(a, conteudoArqs[cacatua], conteudoArqs)
	// }
	// fmt.Println(runtime.NumGoroutine())
	// // wg.Wait()
}

func gera_dados(arq string, pos []int, tam []int, headers_0 string, headers_1 string) {
	// var arq_aux []string
	// var arq_result string
	// var tamanho_plag = 3

	f, err := os.Create("Dados.html")
	if err != nil {
		log.Fatalln("My program broke", err.Error())
	}

	defer f.Close()

	// str := "<!DOCTYPE html>\n<html>\n<head>\n<style>\nmark { \nbackground-color: red;\ncolor: black;\n}\n</style>\n</head>\n<body>\n<p>Aqui não tem plagio!</p>\n<mark>" + arq_aux + "</mark>\n<p>aqui tbm não.\n</p>\n</body>\n</html>"

	// bs := []byte(str)

	// _, err = f.Write(bs)
	// if err != nil {
	// 	log.Fatalln("error writing to file", err.Error())
	// }

	// fmt.Print(string(bs))
}

func RunPlag(tamanho int, archive string, data []string, headers []string, arq_idx int) {
	i := 0
	for i = 0; i < tamanho; i++ {
		if archive != data[i] {
			// ww.Add(1)
			ArrayResult, ArrayResult2 := make([]int, tamanho*100), make([]int, tamanho*100)
			MaiorSubstringComumdo(archive, data[i], ArrayResult, ArrayResult2)
			fmt.Print("arquivo: ", archive, " data :", data[i])
			for abacate := range ArrayResult {
				if ArrayResult[abacate] != 0 {
					fmt.Println("teste", ArrayResult[abacate], " , ", ArrayResult2[abacate])
				}
			}
			gera_dados(archive, ArrayResult, ArrayResult2, headers[arq_idx], headers[i])
		}
		// ww.Wait()
	}
	// wg.Done()
}

func myreadFile(path string) string {
	readFile, err := os.Open(path)
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
	var fullLine = " "
	for _, line := range lines {
		//fmt.Println(line)
		fullLine += (line + "\n")
	}
	return fullLine
}

func ContadorDePalavras(str string) map[string]int { //Abrir arquivo e verificar a incidencia de cada palavra
	counts := make(map[string]int)
	f := func(c rune) bool {
		return unicode.IsSpace(c) || c == '.' || c == ',' || c == '(' || c == ')' || c == '{' || c == '}' || c == ';' ||
			c == '=' || c == '+' || c == '-' || c == '<' || c == '>' || c == '|' || c == '\\' || c == '"' || c == '*' ||
			c == '&' || c == '/' || c == '#' || c == '%'
	}

	wordList := strings.FieldsFunc(str, f)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func PrintNumPalvras(str string) {
	fmt.Println("=================================")
	for index, element := range ContadorDePalavras(str) {
		fmt.Println(index, "=>", element)
	}
}

func GeradorDeToken() { //Troca palvras chave por tokens predenidos
}

func fileCount(caminho string) (int, error) {
	i := 0
	arquivos, err := ioutil.ReadDir(caminho)
	if err != nil {
		return 0, err
	}
	for _, file := range arquivos {
		if !file.IsDir() {
			i++
		}
	}
	return i, nil
}

func showNameFiles(caminho string) {
	files, err := ioutil.ReadDir(caminho)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
func printFile() {
}

/////////////////////////////////////////////////////////////////////////////////

func printMatriz(table [][]int, linha, coluna int) {
	i, j := 0, 0
	for i = 0; i <= linha; i++ {
		for j = 0; j <= coluna; j++ {
			fmt.Print(table[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func MaiorSubstringComumdo(str1, str2 string, resultA []int, resultB []int) {
	len1, len2 := len(str1), len(str2)
	maior, i, j := 0, 0, 0

	m := make([][]int, len1+1)
	for a := range m {
		m[a] = make([]int, len2+1)
	}

	for i = 0; i <= len1; i++ {
		for j = 0; j <= len2; j++ {
			if i == 0 || j == 0 {
				m[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				m[i][j] = m[i-1][j-1] + 1

				if m[i][j] > maior {
					maior = m[i][j]
				}
			} else {
				m[i][j] = 0
			}
		}
	}

	printMatriz(m, len1, len2)
	aux1, aux2, idx := 0, 0, 0
	for i = 0; i <= len1; i++ {
		for j = 0; j <= len2; j++ {
			if m[i][j] != 0 {
				aux1 = i
				var r, s, flag int
				flag = 0
				r = i
				s = j
				for flag == 0 {
					//printMatriz(m, len1, len2)
					if r >= len1 || s >= len2 {
						aux2 = m[r][s]
						m[r][s] = 0
						flag = 1
					} else if m[r+1][s+1] == 0 {
						aux2 = m[r][s]
						m[r][s] = 0
						flag = 1
					} else {
						m[r][s] = 0
						s++
						r++
					}
				}

				resultA[idx] = aux1
				resultB[idx] = aux2
				idx++
				m[i][j] = 0
			}
		}
	}
	//println("resultA", resultA[0])
	// ww.Done()
	//printMatriz(m, len1, len2)
}
