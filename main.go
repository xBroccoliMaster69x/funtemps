package main

import (
    "flag"
    "fmt"
	"github.com/xBroccoliMaster69x/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cels float64
var kelv float64
var out string
var funfacts string

var tmpOrFunTemps string
var tmpFunction string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

    /*
       Her er eksempler på hvordan man implementerer parsing av flagg.
       For eksempel, kommando
           funtemps -F 0 -out C
       skal returnere output: 0°F er -17.78°C
    */

    // Definerer og initialiserer flagg-variablene
    flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
    // Du må selv definere flag-variablene for "C" og "K"
    flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
    flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")
    flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
    flag.StringVar(&funfacts, "funfacts", "sun", "fun-facts om sun - Solen, luna - Månen og terra - Jorden")

    // Du må selv definere flag-variabelen for -t flagget, som bestemmer
    // hvilken temperaturskala skal brukes når funfacts skal vises

}



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
        tmpFunction = conv.FahrenheitToCelsius()
    } else if input == 2 {
        tmpFunction = conv.FahrenheitToKelvin()
    } else if input == 3 {
        tmpFunction = conv.CelsiusToFahrenheit()
    } else if input == 4 {
        tmpFunction = conv.CelsiusToKelvin()
    } else if input == 5 {
        tmpFunction = conv.KelvinToFahrenheit()
    } else if input == 6 {
        tmpFunction = conv.KelvinToCelsius()
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
