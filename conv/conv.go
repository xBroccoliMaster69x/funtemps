package conv

import "math"

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FahrenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFahrenheit
    ...
*/

func roundFloat(value float64) float64 {
    return math.Round(value*100)/100
}

// Konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(Fahrenheit float64) float64 {
    celsius:= (Fahrenheit - 32.00) * (5.00/9.00)
    return roundFloat(celsius)
}

//konverterer Fahrenheit til Kelvin
func FahrenheitToKelvin(Fahrenheit float64) float64 {
    kelvin:= (Fahrenheit - 32.00) * (5.00/9.00) + 273.15
    return roundFloat(kelvin)
}

//konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(celsius float64) float64 {
    fahrenheit:= celsius * (9.00/5.00) + 32.00
    return roundFloat(fahrenheit)
}

//konverterer Celsius til Kelvin
func CelsiusToKelvin(celsius float64) float64 {
    kelvin:= celsius + 273.15
    return roundFloat(kelvin)
}

//konverterer kelvin til celsius
func KelvinToCelsius(kelvin float64) float64 {
    celsius:= kelvin - 273.15
    return roundFloat(celsius)
}

//konverterer kelvin til fahrenheit
func KelvinToFahrenheit(value float64) float64 {
    fahrenheit:= (value - 273.15) * (9.00/5.00) + 32.00
    return roundFloat(fahrenheit)
}

// De andre konverteringsfunksjonene implementere her
// ...