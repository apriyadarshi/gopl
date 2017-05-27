package tempconv

import (
	"flag"
	"fmt"
	"gopl/ch2/tempconv0"
)

//*celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ tempconv0.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv0.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv0.FToC(tempconv0.Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = tempconv0.KToC(tempconv0.Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

/*CelsiusFlag defines a Celsius flag with the specified name,
  default value, and usage, and returns the address of the
  flag variable.
  The flag argument must have a quantity and a unit eg 100C
*/
func CelsiusFlag(name string, value tempconv0.Celsius, usage string) *tempconv0.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
