package stringformating

// Schreibe eine Funktion, die Name und Alter zu einem formatierten String zusammensetzt.
import "fmt"

func FormatPerson(name string, age int) string {
	// TODO: formatierten String zurückgeben
	return fmt.Sprintf("Name: %s, Alter: %d", name, age)
}
