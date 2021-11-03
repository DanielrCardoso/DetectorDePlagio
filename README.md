# Detector de plágio em código fonte 

Este programa foi desenvolvido para a matéria de Linguagem de programação da Universidade de Brasília. O intuito do software é aplicar algumas técnicas conhecidas na análise de plágio em código fonte juntamente com o paralelismo fornecido pela linguagem Go. 

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
