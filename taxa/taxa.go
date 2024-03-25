package taxa

type Flag int

type Type struct { // Tipo de cartão
	Nome string
	Type string // "credito" ou "debito"
}

// Verifica se o tipo do cartão é crédito
func (t Type) IsCreditCard() bool {
	return t.Type == "credito"
}

// Verifica se o tipo do cartão é débito
func (t Type) IsDebitCard() bool {
	return t.Type == "debito"
}

type TaxVariable struct {
	Parcela int
	Taxa    float64
}

const (
	MASTERCARD Flag = iota
	VISA
	AMEX
	ELO
)

// Tipo para cada bandeira
const (
	Credito = "credito"
	Debito  = "debito"
)

type Card struct { //
	Flag Flag // tipo de bandeira
	Type Type // tipo de cartão credito ou debito
}

func (f Flag) String() string {
	return [...]string{"MASTERCARD", "VISA", "AMEX", "ELO"}[f]
}

func (f Flag) TaxFixed() float64 {
	return [...]float64{0.02, 0.01, 0.03, 0.04}[f]
}

func (f Flag) TaxInstallments() []TaxVariable {

	masterCard := []TaxVariable{
		{Parcela: 1, Taxa: 1.03},
		{Parcela: 2, Taxa: 1.04},
		{Parcela: 3, Taxa: 1.05},
		{Parcela: 4, Taxa: 1.06},
	}
	visa := []TaxVariable{
		{Parcela: 1, Taxa: 2.02},
		{Parcela: 2, Taxa: 2.03},
		{Parcela: 3, Taxa: 2.04},
		{Parcela: 4, Taxa: 2.05},
	}

	amex := []TaxVariable{
		{Parcela: 1, Taxa: 3.03},
		{Parcela: 2, Taxa: 3.04},
		{Parcela: 3, Taxa: 3.05},
		{Parcela: 4, Taxa: 3.06},
	}

	elo := []TaxVariable{
		{Parcela: 1, Taxa: 4.04},
		{Parcela: 2, Taxa: 4.05},
		{Parcela: 3, Taxa: 4.06},
		{Parcela: 4, Taxa: 4.07},
		{Parcela: 5, Taxa: 4.08},
		{Parcela: 6, Taxa: 4.09},
		{Parcela: 7, Taxa: 4.10},
		{Parcela: 8, Taxa: 4.11},
		{Parcela: 9, Taxa: 4.12},
		{Parcela: 10, Taxa: 4.13},
	}

	return [...][]TaxVariable{
		masterCard,
		visa,
		amex,
		elo,
	}[f]
}
