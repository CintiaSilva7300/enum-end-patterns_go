package main

import (
	"bufio"
	"enum-exemplos/taxa"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Lista de nomes dos arquivos a serem lidos
	arquivos := []string{"scanner.txt", "scanner1.txt", "scanner2.txt", "scanner3.txt"}

	// Loop através dos nomes dos arquivos
	for _, nomeArquivo := range arquivos {
		// Abra o arquivo de texto
		file, err := os.Open(nomeArquivo)
		if err != nil {
			fmt.Println("Erro ao abrir o arquivo:", err)
			continue // Continue para o próximo arquivo se houver erro ao abrir este
		}
		defer file.Close()

		scanner := bufio.NewScanner(file) // Crie um scanner para ler o arquivo linha por linha do arquivo

		// Processar cada linha do arquivo
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Fields(line) // Separe os campos da linha

			// Verifique se há pelo menos 3 campos
			if len(fields) < 3 {
				fmt.Println("Linha do arquivo incompleta:", line)
				continue
			}

			flag := fields[0]                       // Primeiro campo é a bandeira
			parcela, err := strconv.Atoi(fields[1]) // Segundo campo é o número de parcelas
			if err != nil {
				fmt.Println("Erro ao converter número de parcelas:", err)
				continue
			}

			tipoTransacao := fields[2] // Terceiro campo é o tipo de transação

			// Mapeamento manual das strings para os valores de Flag
			var flagValue taxa.Flag
			switch flag {
			case "MASTERCARD":
				flagValue = taxa.MASTERCARD
			case "VISA":
				flagValue = taxa.VISA
			case "AMEX":
				flagValue = taxa.AMEX
			case "ELO":
				flagValue = taxa.ELO
			default:
				fmt.Println("Bandeira não reconhecida:", flag)
				continue
			}

			// Crie um cartão com base nos dados do arquivo
			card := taxa.Card{
				Flag: flagValue,
				Type: taxa.Type{Nome: flag, Type: tipoTransacao},
			}

			// Obtenha as taxas com base na bandeira do cartão
			taxas := card.Flag.TaxInstallments()

			// Encontre a taxa correta para a parcela especificada
			var taxa float64
			for _, t := range taxas {
				if t.Parcela == parcela {
					taxa = t.Taxa
					break
				}
			}

			// Imprima as informações do cartão e da transação
			fmt.Printf("Arquivo: %s, Parcela: %d, Taxa: %.2f, Bandeira: %s, Tipo da transacao: %s\n", nomeArquivo, parcela, taxa, flag, tipoTransacao)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Erro ao ler o arquivo:", err)
		}
	}
}

//exemplo do arquivo a ser lido
//BANDEIRA, PARCELAS, credito ou debito

// ex:
//MASTERCARD 2 credito
