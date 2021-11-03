package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
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
	headers := make([]string, a)

	// err_dir := os.Mkdir("depuracao", 0755)
	// if err_dir != nil {
	// 	log.Fatal(err_dir)
	// }

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
		headers[idx] = f.Name() // repetido (atenção)
		idx++
		fullLine = myreadFile(path)
		conteudoArqs = append(conteudoArqs, fullLine)
	}

	for content := range conteudoArqs {
		wg.Add(1)
		go RunPlag(a, conteudoArqs[content], conteudoArqs, headers, content)
	}
	fmt.Println("goroutines_MSC2: ", runtime.NumGoroutine()) // debbug
	wg.Wait()

	result := make([][]float64, a)
	for i := range result {
		result[i] = make([]float64, a)
	}
	similaridade_cos(result)
	mostra_mat(result, headers)
}

func gera_dados(arq string, pos []int, tam []int, headers_0 string, headers_1 string) {
	var arq_aux []string
	var arq_result string
	var pos_aux []int
	var tamanho_plag = 3

	for i := 0; i < len(pos); i++ {
		if tam[i] > tamanho_plag {
			arq_aux = append(arq_aux, "<mark>"+arq[pos[i]-1:pos[i]+tam[i]-1]+" --> "+headers_1+"</mark>")
			pos_aux = append(pos_aux, (pos[i])-1)
			pos_aux = append(pos_aux, (tam[i]))

		}
		if pos[i] == 0 {
			i = len(pos)
		}
	}

	aux_idx := 0
	aux_idx2 := 0
	for i := 0; i < len(arq); i++ {
		if aux_idx < len(pos_aux) && i == pos_aux[aux_idx] {
			arq_result += arq_aux[aux_idx2]
			i += pos_aux[aux_idx+1]
			aux_idx2 += 1
			aux_idx += 2
		} else {
			arq_result += arq[i : i+1]
		}
	}

	f, err := os.Create(filepath.Join("depuracao", filepath.Base(headers_0+"_"+headers_1+".html")))
	if err != nil {
		log.Fatalln("My program broke", err.Error())
	}

	defer f.Close()

	str := "<!DOCTYPE html>\n<html>\n<head>\n<style>\nmark { \nbackground-color: red;\ncolor: black;\n}\n</style>\n</head>\n<body>\n<p>Arquivo:</p>\n<div container>" + arq_result + "</div>\n<p>Fim do arquivo.\n</p>\n</body>\n<style></style></html>"

	bs := []byte(str)

	_, err = f.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file", err.Error())
	}

}

func RunPlag(tamanho int, archive string, data []string, headers []string, arq_idx int) {
	i := 0
	for i = 0; i < tamanho; i++ {
		if archive != data[i] {
			ArrayResult, ArrayResult2 := make([]int, tamanho*len(archive)*len(data[i])), make([]int, tamanho*len(archive)*len(data[i]))
			MaiorSubstringComumdo(archive, data[i], ArrayResult, ArrayResult2)
			gera_dados(archive, ArrayResult, ArrayResult2, headers[arq_idx], headers[i])
		}
	}
	wg.Done()
}

func myreadFile(path string) string { // repetido (y)
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
		fullLine += (line + "\n")
	}
	return fullLine
}

func fileCount(caminho string) (int, error) { // repetido (y)
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
func similaridade_cos(result [][]float64) {
	var fullLine = ""
	var path = ""
	files, err := ioutil.ReadDir("Arquivos para verificação")
	if err != nil {
		log.Fatal(err)
	}
	var conteudoArqs []string
	for _, f := range files { // repetido (atenção)
		path = "Arquivos para verificação/" + f.Name()
		fullLine = myreadFile(path)
		conteudoArqs = append(conteudoArqs, fullLine)
	}
	for i := range conteudoArqs {
		for j := range conteudoArqs {
			ww.Add(1)
			go similaridade(conteudoArqs[i], conteudoArqs[j], result, i, j)
		}
		fmt.Println("goroutines_cos: ", runtime.NumGoroutine()) // debbug
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
func Teste() {
	fmt.Println("Hello World!")
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
