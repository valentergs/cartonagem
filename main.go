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
	f.WriteString("==================\n\n")
	f.WriteString(fmt.Sprintf("Base: %vcm x %vcm\n\n", largura, comprimento))
	f.WriteString(fmt.Sprintf("Lateral Esquerda: %vcm x %vcm\n\n", comprimento, altura))
	f.WriteString(fmt.Sprintf("Lateral Direita: %vcm x %vcm\n\n", comprimento, altura))
	f.WriteString(fmt.Sprintf("Lateral Frente: %vcm x %vcm\n\n", (largura + espessura), altura))
	f.WriteString(fmt.Sprintf("Lateral Fundo: %vcm x %vcm\n\n", (largura + espessura), altura))
	f.WriteString("CAPA =============\n\n")
	f.WriteString(fmt.Sprintf("Tampa: %vcm x %vcm\n\n", (largura + 1.5), (comprimento + 1)))
	f.WriteString(fmt.Sprintf("Fundo: %vcm x %vcm\n\n", (largura + 1.5), (comprimento + 1)))
	f.WriteString(fmt.Sprintf("Lombada: %vcm x %vcm\n\n", (largura + 1.5), altura))
	f.WriteString("ACABAMENTO =======\n\n")
	f.WriteString(fmt.Sprintf("Interno Tampa: %vcm x %vcm\n\n", (largura + 1.5 - 0.5), (comprimento + 1 - 0.5)))
	f.WriteString(fmt.Sprintf("Externo Lateral Direita: %vcm x %vcm\n\n", (comprimento + espessura), (altura - 0.1)))
	f.WriteString(fmt.Sprintf("Externo Lateral Esquerda: %vcm x %vcm\n\n", (comprimento + espessura), (altura - 0.1)))
	f.WriteString(fmt.Sprintf("Externo Lateral Frente: %.1fcm x %vcm\n\n", ((comprimento + espessura) - 0.1), (altura - 0.1)))
	f.WriteString(fmt.Sprintf("Externo Lateral Fundo: %.1fcm x %vcm\n\n", ((comprimento + espessura) - 0.1), (altura - 0.1)))
}
