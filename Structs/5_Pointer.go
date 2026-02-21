package structs

// ConvertToFahrenheit rechnet eine Temperatur in Grad Celsius
// in Fahrenheit um und verändert die ursprüngliche Variable.
// Formel: F = C * 1.8 + 32
func ConvertToFahrenheit(temp *float64) {
	//TODO
	var F float64
	F = *temp*1.8 + 32
	*temp = F
}
