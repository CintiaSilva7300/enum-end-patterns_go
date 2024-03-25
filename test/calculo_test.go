package test

import (
	"enum-exemplos/taxa"
	calculo "enum-exemplos/type"
	"fmt"

	"testing"
)

// esses teste testam os valores de taxa se estao corretos para diferentes números de parcelas.
func TestMasterCard(t *testing.T) {
	tests := []struct {
		parcela int
		want    float64
	}{
		{1, 1.03},
		{2, 1.04},
		{3, 1.05},
		{4, 1.06},
	}

	for _, tt := range tests {
		if got := calculo.MasterCard(tt.parcela); got != tt.want {
			t.Errorf("MasterCard(%d) = %v, want %v", tt.parcela, got, tt.want)
		}
	}
}

func TestVisa(t *testing.T) {
	tests := []struct {
		parcela int
		want    float64
	}{
		{1, 2.02},
		{2, 2.03},
		{3, 2.04},
		{4, 2.05},
	}

	for _, tt := range tests {
		if got := calculo.Visa(tt.parcela); got != tt.want {
			t.Errorf("Visa(%d) = %v, want %v", tt.parcela, got, tt.want)
		}
	}
}

func TestAmex(t *testing.T) {
	tests := []struct {
		parcela int
		want    float64
	}{
		{1, 3.03},
		{2, 3.04},
		{3, 3.05},
		{4, 3.06},
	}

	for _, tt := range tests {
		if got := calculo.Amex(tt.parcela); got != tt.want {
			t.Errorf("Amex(%d) = %v, want %v", tt.parcela, got, tt.want)
		}
	}
}

func TestElo(t *testing.T) {
	tests := []struct {
		parcela int
		want    float64
	}{
		{1, 4.04},
		{2, 4.05},
		{3, 4.06},
		{4, 4.07},
		{5, 4.08},
		{6, 4.09},
		{7, 4.10},
		{8, 4.11},
		{9, 4.12},
		{10, 4.13},
	}

	for _, tt := range tests {
		if got := calculo.Elo(tt.parcela); got != tt.want {
			t.Errorf("Elo(%d) = %v, want %v", tt.parcela, got, tt.want)
		}
	}
}

//testa a bandeira
func TestIsCreditCard(t *testing.T) {
	cartaoCredito := taxa.Type{Type: taxa.Credito}

	if cartaoCredito.IsCreditCard() {
		fmt.Println("Este é um cartão de crédito")
	} else {
		fmt.Println("Este não é um cartão de crédito")
	}
}

func TestIsDebitCard(t *testing.T) {
	cartaoDebito := taxa.Type{Type: taxa.Debito}

	if cartaoDebito.IsDebitCard() {
		fmt.Println("Este é um cartão de debito")
	} else {
		fmt.Println("Este não é um cartão de debito")
	}
}
