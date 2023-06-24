package formatter

import "fmt"

type lookupItem struct {
	Value  float64
	Symbol string
}

func ShortNumber(num float64, digits int) string {
	lookup := []lookupItem{
		{Value: 1, Symbol: ""},
		{Value: 1e3, Symbol: "k"},
		{Value: 1e6, Symbol: "M"},
		{Value: 1e9, Symbol: "G"},
		{Value: 1e12, Symbol: "T"},
		{Value: 1e15, Symbol: "P"},
		{Value: 1e18, Symbol: "E"},
	}

	for i := len(lookup) - 1; i >= 0; i-- {
		item := lookup[i]
		if num >= item.Value {
			result := num / item.Value
			formatted := fmt.Sprintf("%.*f", digits, result)

			// Remove trailing zeros
			for formatted[len(formatted)-1] == '0' {
				formatted = formatted[:len(formatted)-1]
			}

			// Remove the decimal point if there are no decimal digits
			if formatted[len(formatted)-1] == '.' {
				formatted = formatted[:len(formatted)-1]
			}

			return formatted + item.Symbol
		}
	}

	return "0"
}
