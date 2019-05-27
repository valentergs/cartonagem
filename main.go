package main

import (
	"fmt"
	"os"
	"time"
)

var largura float64
var comprimento float64
var altura float64
var espessura float64
var quantidade float64
var arquivo string

func main() {
	fmt.Print("Largura: ")
	fmt.Scanln(&largura)
	fmt.Print("Comprimento: ")
	fmt.Scanln(&comprimento)
	fmt.Print("Altura: ")
	fmt.Scanln(&altura)
	fmt.Print("Espessura: ")
	fmt.Scanln(&espessura)
	fmt.Print("Quantidade: ")
	fmt.Scanln(&quantidade)
	fmt.Print("Nome do arquivo: ")
	fmt.Scanln(&arquivo)

	f, err := os.Create(arquivo + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	hoje := time.Now().Format(time.RFC822)

	f.WriteString(fmt.Sprintf("Arquivo: %v.txt - %s\n\n", arquivo, hoje))
	f.WriteString("==================\n\n")
	f.WriteString(fmt.Sprintf("Largura: %vcm\n\n", largura))
	f.WriteString(fmt.Sprintf("Comprimento: %vcm\n\n", comprimento))
	f.WriteString(fmt.Sprintf("Altura: %vcm\n\n", altura))
	f.WriteString(fmt.Sprintf("Espessura: %vcm\n\n", espessura))
	f.WriteString(fmt.Sprintf("Quantidade: %v un\n\n", quantidade))
	f.WriteString("==================\n\n")
	f.WriteString(fmt.Sprintf("Base: %vcm x %vcm - %v un\n\n", largura, comprimento, quantidade))
	f.WriteString(fmt.Sprintf("Lateral D/E: %vcm x %vcm - %v un\n\n", comprimento, altura, (quantidade * 2)))
	f.WriteString(fmt.Sprintf("Lateral F/F: %vcm x %vcm - %v un\n\n", comprimento, altura, (quantidade * 2)))
	f.WriteString("CAPA =============\n\n")
	f.WriteString(fmt.Sprintf("Tampa: %vcm x %vcm - %v un\n\n", (largura + 1.5), (comprimento + 1), quantidade))
	f.WriteString(fmt.Sprintf("Fundo: %vcm x %vcm - %v un\n\n", (largura + 1.5), (comprimento + 1), quantidade))
	f.WriteString(fmt.Sprintf("Lombada: %vcm x %vcm - %v un\n\n", (largura + 1.5), altura, quantidade))
	f.WriteString("ACABAMENTO =======\n\n")
	f.WriteString(fmt.Sprintf("Interno Tampa: %vcm x %vcm - %v un\n\n", (largura + 1.5 - 0.5), (comprimento + 1 - 0.5), quantidade))
	f.WriteString(fmt.Sprintf("Externo Lateral D/E: %vcm x %vcm - %v un\n\n", (comprimento + espessura), (altura - 0.1), (quantidade * 2)))
	f.WriteString(fmt.Sprintf("Externo Lateral F: %.1fcm x %vcm - %v un\n\n", ((comprimento + espessura) - 0.1), (altura - 0.1), quantidade))
}
