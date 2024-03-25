package calculo

import "enum-exemplos/taxa"

func Elo(parcela int) float64 {
	return taxa.Flag(3).TaxInstallments()[parcela-1].Taxa
}
