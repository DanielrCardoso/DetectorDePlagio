package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"unicode"
	// "runtime"
)

var wg sync.WaitGroup

var ww sync.WaitGroup

func main() {
	fmt.Println("Versao Apresentação")
	a, _ := fileCount("Arquivos para verificação")
	fmt.Println("Numero de aquivos: ", a)
	fmt.Println("Arquivos encontrados: ")
	//showNameFiles("Arquivos para verificação")

	var fullLine = ""
	var path = ""
	files, err := ioutil.ReadDir("Arquivos para verificação")
	if err != nil {
		log.Fatal(err)
	}
	var conteudoArqs []string
	for _, f := range files {
		path = "Arquivos para verificação/" + f.Name()
		fmt.Println("Abrindo arquivo: " + path)
		fullLine = myreadFile(path)
		//fmt.Println(fullLine)
		PrintNumPalvras(fullLine)
		//fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++=")
		conteudoArqs = append(conteudoArqs, fullLine)
	}

	for cacatua := range conteudoArqs {
		wg.Add(1)
		go RunPlag(a, conteudoArqs[cacatua], conteudoArqs)
	}
	wg.Wait()
	// fmt.Println("go:",runtime.NumGoroutine())
}

func RunPlag(tamanho int, archive string, data []string) {
	i := 0
	for i = 0; i < tamanho; i++ {
		ww.Add(1)
		ArrayResult, ArrayResult2 := make([]int, tamanho*10), make([]int, tamanho*10)
		//go MaiorSubstringComumdo(archive, data[i], ArrayResult, ArrayResult2)
		go lcs(archive, data[i], ArrayResult, ArrayResult2)
		//fmt.Println("arquivo: ", archive, " data:", data[i])
		for abacate := range ArrayResult {
			fmt.Println(ArrayResult[abacate], " , ", ArrayResult2[abacate])

		}
	}
	ww.Wait()
	wg.Done()
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

func lcs(str1, str2 string, resultA []int, resultB []int) {
	m, m2 := MaiorSubstringComumdo(str1, str2)
	//printMatriz(m, len(str1), len(str2))
	resultA, resultB = posiPlagio(m2, len(str1), len(str2), resultA, resultB)
	allsimular := strplagio(len(str1), len(str2), m, str1, str2)
	//escreverTexto(allsimular, "Resultados")
	printplagio(allsimular, 2)
	ww.Done()
}

func MaiorSubstringComumdo(str1, str2 string) ([][]int, [][]int) {
	len1, len2 := len(str1), len(str2)
	maior, i, j := 0, 0, 0
	m := make([][]int, len1+1)
	m2 := make([][]int, len1+1)
	for a := range m {
		m[a] = make([]int, len2+1)
		m2[a] = make([]int, len2+1)
	}

	for i = 0; i <= len1; i++ {
		for j = 0; j <= len2; j++ {
			if i == 0 || j == 0 {
				m[i][j] = 0
				m2[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				m[i][j] = m[i-1][j-1] + 1
				m2[i][j] = m2[i-1][j-1] + 1

				if m[i][j] > maior {
					maior = m[i][j]
				}
			} else {
				m[i][j] = 0
				m2[i][j] = 0
			}
		}
	}
	return m, m2
}

func posiPlagio(m [][]int, len1, len2 int, resultA []int, resultB []int) ([]int, []int) {
	//printMatriz(m, len1, len2)
	aux1, aux2, idx, i, j := 0, 0, 0, 0, 0
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
	return resultA, resultB
}

func strplagio(len1, len2 int, m [][]int, str1, str2 string) []string {
	var plagio bytes.Buffer
	todosplagios := make([]string, len1+len2)

	i, j := 0, 0
	var r, s int
	var pos int
	pos = 0
	for HaveNotNull(m, len1, len2) {
		for i = 0; i <= len1; i++ {
			for j = 0; j <= len2; j++ {
				if m[i][j] != 0 {
					for r, s = i, j; r <= len1 && s <= len2; r, s = r+1, s+1 {
						if m[r][s] == 0 {
							break
						} else {
							plagio.WriteByte(str1[r-1])
							m[r][s] = 0
						}
					}
					todosplagios[pos] = plagio.String()
					pos++
					plagio.Reset()
				}
			}
		}
	}
	// for aa := range todosplagios {
	// 	if len(todosplagios[aa]) > 2 {
	// 		fmt.Println(todosplagios[aa])
	// 	}
	// }

	// var hh int
	// hh = 0
	// for hh < len1 {
	// 	fmt.Print(string(str1[hh]))
	// 	hh++
	// }
	return todosplagios
}

func printplagio(similaridades []string, valorbase int) {
	for aa := range similaridades {
		if len(similaridades[aa]) > valorbase {
			fmt.Println(similaridades[aa])
		}
	}
}
func HaveNotNull(table [][]int, linha, coluna int) bool {
	i, j := 0, 0
	for i = 0; i <= linha; i++ {
		for j = 0; j <= coluna; j++ {
			if table[i][j] != 0 {
				return true
			}
		}
	}
	return false
}

func printPosicoes(ArrayResult, ArrayResult2 []int) {

}

// func escreverTexto(linhas []string, caminhoDoArquivo string) error {
// 	arquivo, err := os.Create(caminhoDoArquivo)
// 	if err != nil {
// 		return err
// 	}

// 	defer arquivo.Close()

// 	escritor := bufio.NewWriter(arquivo)
// 	for _, linha := range linhas {
// 		fmt.Fprintln(escritor, linha)
// 	}

// 	return escritor.Flush()
// }
