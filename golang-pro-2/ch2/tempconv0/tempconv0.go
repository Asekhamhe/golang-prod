// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

// Celsius type declaration
type Celsius float64

// Fahrenheit type declaration
type Fahrenheit float64

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

const (
	// AbsoluteZeroC constant
	AbsoluteZeroC Celsius = -273.15
	// FreezingC constant
	FreezingC Celsius = 0
	// BoilingC constant
	BoilingC Celsius = 100
)

// CToF conversion
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC conversion
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
