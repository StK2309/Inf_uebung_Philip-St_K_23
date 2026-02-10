package rekursion

import "unicode"

// Erwartet eine verschachtelte Liste von Strings ([][]string)
// und einen gesuchten Buchstaben (z. B. 'a').
// Gibt die Anzahl zurück, wie oft dieser Buchstabe insgesamt vorkommt –
// unabhängig von Groß- oder Kleinschreibung.
//
// Beispiel:
//
//	input := [][]string{
//	    {"Apfel", "Banane"},
//	    {"Ananas"},
//	    {},
//	    {"Avocado", "Apfel"},
//	}
//
// CountLetterNested(input, 'a') → 7
// Hinweis: nutze unicode.ToLower()
// (Achtung: auch große A zählen!)
func CountLetterNested(nested [][]string, letter rune) int {
	// TODO: Funktion implementieren (rekursiv oder iterativ möglich)
	count := 0
	target := unicode.ToLower(letter)

	for _, inner := range nested {
		for _, word := range inner {
			for _, r := range word {
				if unicode.ToLower(r) == target {
					count++
				}
			}
		}
	}
	return count
}
