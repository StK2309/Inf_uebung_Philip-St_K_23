package structs

// Struct zur Darstellung einer Bewertung
type Rating struct {
	Name  string // Name der bewertenden Person
	Score int    // Bewertung zwischen z. B. 1 und 5
}

// AverageScore berechnet den Durchschnittswert aller Bewertungen
func AverageScore(ratings []Rating) float64 {
	// TODO: Implementieren
	return 0.1 // Platzhalter
}
