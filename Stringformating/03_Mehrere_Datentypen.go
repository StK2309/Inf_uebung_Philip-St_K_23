package stringformating

import "fmt"

// Formatiere folgende Werte:
// -Produktname
// -Stückzahl
// -Einzelpreis
func FormatOrder(product string, amount int, price float64) string {
	// TODO
	return fmt.Sprintf("Produkt: %s | Menge: %d | Einzelpreis: %.2f €", product, amount, price)
}
