//This package does celsius to fahrenheit conversions and vice versa

package tempconv0

import "fmt"

type Celsius float64 //Named type with underlying type Float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

//See a package level function declaration func <name>(<params list>) <return type list>  { <func defn> }
// Whereas a method declaration is --------func (<params list>) <name> <return type list> { <func defn> }
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32) //Fahrenheit(<float>) is an explicit type conversion, not a function call. Value is not changed
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
