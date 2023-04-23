package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CNY = "CNY"
	RUB = "RUB"
)

// IsSupportesCurrency return true if the currency is supported
func IsSupportesCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CNY, RUB:
		return true
	}

	return false
}
