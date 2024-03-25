package calculo

import "enum-exemplos/taxa"

func MasterCard(parcela int) float64 {
	return taxa.Flag(0).TaxInstallments()[parcela-1].Taxa
}
