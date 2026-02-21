package structs

// Struct zur Darstellung einer Bewertung
type Rating struct {
	Name  string // Name der bewertenden Person
	Score int    // Bewertung zwischen z. B. 1 und 5
}

// AverageScore berechnet den Durchschnittswert aller Bewertungen
func AverageScore(ratings []Rating) float64 {
	// TODO: Implementieren
	var total int = 0
	for _, ratiratings := range ratings {
		total += ratiratings.Score
	}
	return float64(total) / float64(len(ratings))
}
