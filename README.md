# Detector de plágio em código fonte 

Este programa foi desenvolvido para a matéria de Linguagem de programação da Universidade de Brasília. O intuito do software é aplicar algumas técnicas conhecidas na análise de plágio em código fonte juntamente com o paralelismo fornecido pela linguagem Go.

A ideia é auxiliar professores na correção de atividades submetidas por alunos tendo em vista que a verificação manual se torna inviável em grandes turmas. A combinação dos pares necessários pode ser expressa por:

*n(n-1) /2*

Onde *n* é o número de submissões.

## Algoritmos Adotados

### Algoritmo da maior substring comum
Dado duas strings, este algoritmo é capaz de encontrar a maior substring comum entre as entradas. Com este resultado é possível ver graficamente onde o plagio está presente.

### Similaridade do Cosseno
Dado duas entradas o algoritmo gera um vetor com todas as palavras existentes, depois dessa etapa dois novos vetores são gerados contendo as incidências de cada palavra em cada entrada, com estes dois vetores em mãos é possível efetuar o cálculo do cosseno, dessa forma é possível descobrir o angulo entre os dois vetores e determinar o quão semelhante são os dois arquivos.

## Sobre o Paralelismo
Em nossa abordagem utilizamos o paralelismo para executar a verificação de um arquivo com os demais, por exemplo ao mesmo tempo que existe a verificação do arquivo A com todos os demais, existe a verificação do arquivo B com todos os demais.

No terminal, clone o projeto:  

```  
git clone https://github.com/DanielrCardoso/DetectorDePlagio.git 
``` 

Verifique a existência dos diretórios "Arquivos para verificação" e "depuracao" 

No diretório "Arquivos para verificação" coloque todos os arquivos para verificação 

Execute o código: 

```  
go run main.go 
``` 
Ao final da execução o diretório "depuracao" conterá os resultados referentes a cada par de arquivos comparados e no terminal será exibido uma tabela referente a execução do algoritmo de similaridade do cosseno. 
