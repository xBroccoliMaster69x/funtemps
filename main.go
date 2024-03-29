package main

import (
	"flag"
	"fmt"
	"github.com/xBroccoliMaster69x/funtemps/conv"
)

func main() {
	var fahr, cels, kelv float64
	var out string

	//Definerer flag variabler, vet dette staar tidligere men torr ikke fjerne dem.
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")

	// Parser flag
	flag.Parse()

	// if statment som sjekker input flag som ble satt og temperaturen den skal convertere fra
	var inputUnit string
	var inputValue float64

	if fahr != 0.0 {
		inputValue = fahr
		inputUnit = "F"
	} else if cels != 0.0 {
		inputValue = cels
		inputUnit = "C"
	} else if kelv != 0.0 {
		inputValue = kelv
		inputUnit = "K"
	} else {
		fmt.Println("Error: No input temperature provided")
		return
	}

	//switch statment som sjekker output flagget som ble satt og temperaturen den skal convertere til
	var outputUnit string
	var outputValue float64

	if inputUnit == "K" {
		switch out {
		case "C":
			outputUnit = "C"
			outputValue = conv.KelvinToCelsius(inputValue)
		case "F":
			outputUnit = "F"
			outputValue = conv.KelvinToFahrenheit(inputValue)
		case "K":
			outputUnit = "K"
			outputValue = inputValue
		default:
			fmt.Println("Error: Invalid output temperature unit")
			return
		}
	} else if inputUnit == "C" {
		switch out {
		case "K":
			outputUnit = "K"
			outputValue = conv.CelsiusToKelvin(inputValue)
		case "F":
			outputUnit = "F"
			outputValue = conv.CelsiusToFahrenheit(inputValue)
		case "C":
			outputUnit = "C"
			outputValue = inputValue
		default:
			fmt.Println("Error: Invalid output temperature unit")
			return
		}
	} else if inputUnit == "F" {
		switch out {
		case "K":
			outputUnit = "K"
			outputValue = conv.FahrenheitToKelvin(inputValue)
		case "C":
			outputUnit = "C"
			outputValue = conv.FahrenheitToCelsius(inputValue)
		case "F":
			outputUnit = "F"
			outputValue = inputValue
		default:
			fmt.Println("Error: Invalid output temperature unit")
			return
		}

	}

	// Output result
	fmt.Printf("%.2f%s is %.2f%s\n", inputValue, inputUnit, outputValue, outputUnit)
}

/* kode som jeg i dag fant ut ikke oppnodde oppgave-kravet :)
//funksjon for valg av tempereaturkonverterings funksjon.
func tmp() string {
    var tmpFunction string
    fmt.Println("Please select a function:")
    fmt.Println("1. FahrenheitToCelsius")
    fmt.Println("2. FahrenheitToKelvin")
    fmt.Println("3. CelsiusToFahrenheit")
    fmt.Println("4. CelsiusToKelvin")
    fmt.Println("5. KelvinToFahrenheit")
    fmt.Println("6. KelvinToCelsius")
    var input int
    fmt.Scanln(&input)
    if input == 1 {
		fmt.Println("please Select a Fahreneit temperature to convert Celsius")
        tmpFunction = conv.FahrenheitToCelsius(Fahreneit)
		return roundFloat(celsius)
    } else if input == 2 {
		fmt.Println("please Select a Fahrenheit temperature to convert Kelvin")
        tmpFunction = conv.FahrenheitToKelvin(Fahreneit)
		return roundFloat(kelvin)
    } else if input == 3 {
		fmt.Println("please Select a Celsius temperature to convert Fahrenheit")
        tmpFunction = conv.CelsiusToFahrenheit(Celsius)
		return roundFloat(fahrenheit)
    } else if input == 4 {
		fmt.Println("please Select a Celsius temperature to convert Kelvin")
        tmpFunction = conv.CelsiusToKelvin(Celsius)
		return roundFloat(kelvin)
    } else if input == 5 {
		fmt.Println("please Select a Kelvin temperature to convert Fahrenheit")
        tmpFunction = conv.KelvinToFahrenheit(Kelvin)
		return roundFloat(celsius)
    } else if input == 6 {
		fmt.Println("please Select a Kelvin temperature to convert Celsius")
        tmpFunction = conv.KelvinToCelsius(Kelvin)
		return roundFloat(fahrenheit)
    } else {
        tmpFunction = "Invalid input"
    }
    return tmpFunction
}
//funksjon som printer om du har valgt funtemps
func funtemps() string {
    return "you have selected funtemps"
}

func main() {
    var tmpOrFunTemps string
    flag.StringVar(&tmpOrFunTemps,"tmpOrFunTemps","","tmp or funtemps?")
    flag.Parse()
	if tmpOrFunTemps == "" {
		fmt.Println("Please choose wether you wish to use temperature converter(tmp) or would like to hear a fun temperature fact(funtemps)")
		fmt.Scanln(&tmpOrFunTemps)
	}

	if tmpOrFunTemps == "tmp"{
		tmpFunction = tmp()
		fmt.Println(tmpFunction)
	} else if tmpOrFunTemps == "funtemps" {
		fmt.Println(funtemps())
	} else {
		fmt.Println(tmpOrFunTemps, "is not a valid funciton name")
	}

}
*/

/**
    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
    pakkene implementeres.
    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
    hvor mange flagg og argumenter er spesifisert på kommandolinje.
        fmt.Println("len(flag.Args())", len(flag.Args()))
		    fmt.Println("flag.NFlag()", flag.NFlag())
    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
    brukes for å utelukke ugyldige kombinasjoner:
    -F, -C, -K kan ikke brukes samtidig
    disse tre kan brukes med -out, men ikke med -funfacts
    -funfacts kan brukes kun med -t
    ...
    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
    implementert i flag-pakken og at den vil skrive ut "Usage" med
    beskrivelsene av flagg-variablene, som angitt i parameter fire til
    funksjonene Float64Var og StringVar
*/
/*
	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println(fahr, out, funfacts)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		fmt.Println("0°F er -17.78°C")
	}

}
*/

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
/*
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
*/
