package stringformating

import "fmt"

// Formatiere Name und Punktzahl spaltenbündig.
func FormatScore(name string, score int) string {
	// TODO
	return fmt.Sprintf("%-10s %d", name, score)
}
