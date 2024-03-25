package calculo

import "enum-exemplos/taxa"

func Amex(parcela int) float64 {
	return taxa.Flag(2).TaxInstallments()[parcela-1].Taxa
}
