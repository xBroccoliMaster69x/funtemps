package conv

import ("math")


/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
  celsius:= (Farhenheit - 32) * (5/9)
	return 
}
func FarhenheitToKelvin(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
  kelvin:= (Farhenheit - 32) * (5/9) + 273.15
	return 
}
func CelsiusToFahrenheit(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
  farhenheit:= celsius * (5/9) + 32
	return 
}
func CelsiusToKelvin(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
  kelvin:= celsius + 273.15
	return 
}
func KelvinToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
  celsius:= kelvin - 273.15
	return 
}
func KelvinToFarhenheit(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
  farhenheit:= (kelvin - 273.15) * (9/5) + 32
	return 
}

// De andre konverteringsfunksjonene implementere her
// ...