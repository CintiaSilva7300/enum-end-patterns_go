package calculo

import "enum-exemplos/taxa"

func Visa(parcela int) float64 {
	return taxa.Flag(1).TaxInstallments()[parcela-1].Taxa
}
