package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
	"sync"
	"unicode"
	"runtime"
)

var wg sync.WaitGroup
var ww sync.WaitGroup

func main() {
	tam, _ := fileCount("Arquivos para verificação")
	fmt.Println("Numero de aquivos: ", tam)

	result := make([][]float64, tam)
	headers := make([]string, tam)
	for i := range result {
		result[i] = make([]float64, tam)
	}
	similaridade_cos(result, headers)
	mostra_mat(result, headers)
}

func mostra_mat(result [][]float64, headers []string) {
	aux := 0
	for k := range headers {
		if len(headers[k]) > aux {
			aux = len(headers[k])
		}
	}
	for i := range headers {
		for len(headers[i]) < aux {
			headers[i] += " "
		}
	}
	var l = 0
	for l < aux {
		fmt.Print(" ")
		l++
	}
	for k := range headers {
		fmt.Printf("| Arquivo %d \t|", k+1)
	}
	fmt.Println("")
	for i := range result {
		fmt.Print(headers[i], "")
		for j := range result {
			fmt.Printf("|  %6.2f %% \t|", (result[i][j] * 100))
		}
		fmt.Println()
	}
}
func similaridade_cos(result [][]float64, headers []string) {
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
		fmt.Println("Abrindo arquivo: " + path)
		headers[idx] = f.Name()
		idx += 1
		fullLine = myreadFile(path)
		conteudoArqs = append(conteudoArqs, fullLine)
	}

	for i := range conteudoArqs {
		for j := range conteudoArqs {
			ww.Add(1)
			go similaridade(conteudoArqs[i], conteudoArqs[j], result, i, j)
		}
		fmt.Println(runtime.NumGoroutine())
		ww.Wait()
	}
}

func Cosine(a []float64, b []float64, result [][]float64, i int, j int) {
	count := 0
	length_a := len(a)
	length_b := len(b)
	if length_a > length_b {
		count = length_a
	} else {
		count = length_b
	}
	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= length_a {
			s2 += math.Pow(b[k], 2)
			continue
		}
		if k >= length_b {
			s1 += math.Pow(a[k], 2)
			continue
		}
		sumA += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	// if s1 == 0 || s2 == 0 {
	// 	return 0.0, errors.New("Vectors should not be null (all zeros)")
	// }
	result[i][j] = sumA / (math.Sqrt(s1) * math.Sqrt(s2))
}
func similaridade(arq_0 string, arq_1 string, result [][]float64, x int, y int) {
	arr_i := Tokenização(arq_0, arq_1)
	f := func(c rune) bool {
		return unicode.IsSpace(c) || c == '.' || c == ',' || c == '(' || c == ')' || c == '{' || c == '}' || c == ';' ||
			c == '=' || c == '+' || c == '-' || c == '<' || c == '>' || c == '|' || c == '\\' || c == '"' || c == '*' ||
			c == '&' || c == '/' || c == '#' || c == '%'
	}
	wordList0 := strings.FieldsFunc(arq_0, f)
	wordList1 := strings.FieldsFunc(arq_1, f)
	a := make([]float64, len(arr_i))
	b := make([]float64, len(arr_i))

	for i := 0; i < len(wordList0); i++ {
		aux := Find(arr_i, wordList0[i])
		a[aux] += 1
	}
	for i := 0; i < len(wordList1); i++ {
		aux := Find(arr_i, wordList1[i])
		b[aux] += 1
	}
	Cosine(a, b, result, x, y)
	ww.Done()
}
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
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
func Tokenização(str string, str2 string) []string { //Abrir arquivo e verificar a incidencia de cada palavra
	f := func(c rune) bool {
		return unicode.IsSpace(c) || c == '.' || c == ',' || c == '(' || c == ')' || c == '{' || c == '}' || c == ';' ||
			c == '=' || c == '+' || c == '-' || c == '<' || c == '>' || c == '|' || c == '\\' || c == '"' || c == '*' ||
			c == '&' || c == '/' || c == '#' || c == '%'
	}
	var array_str []string
	wordList := strings.FieldsFunc(str, f)
	wordList2 := strings.FieldsFunc(str2, f)

	for i := 0; i < len(wordList); i++ {
		if contains(array_str, wordList[i]) == false {
			array_str = append(array_str, wordList[i])
		}
	}
	for i := 0; i < len(wordList2); i++ {
		if contains(array_str, wordList2[i]) == false {
			array_str = append(array_str, wordList2[i])
		}
	}
	return array_str
}
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
