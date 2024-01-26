package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	R1 = flag.Float64("r1", 10000.0, "Resistor R1 value in ohms")
	R2 = flag.Float64("r2", 1000.0, "Resistor R2 value in ohms")
	CuF = flag.Float64("c", 10.0, "Capacitor C value in microfarads (µF)")
)

func main() {
	// Parse the command-line flags
	flag.Parse()

	// Convert the capacitor value from µF to F
	C := *CuF * 1e-6

	// Frequency range: 20Hz to 2000Hz
	for f := 20.0; f <= 2000; f += 10 {
		gain := computeGain(f, C)
		fmt.Printf("Frequency: %4.0f Hz, Gain: %f\n", f, gain)
	}
}

func computeGain(f float64, C float64) float64 {
	omega := 2 * math.Pi * f
	// Calculate the gain using the formula
	gain := 1 + (*R1/(*R2))*(1/(1+1i*omega*(*R2)*C))
	// Return the magnitude of the gain
	return math.Abs(gain)
}

