package main

import "fmt"

func enRango(min float64, max float64, num ...float64) []float64 {
	var result []float64
	for _, numero := range num {
		if numero >= min && numero <= max {
			result = append(result, numero)
		}
	}
	return result
}

func main() {
	fmt.Println(enRango(-10, 10, 4.1, 12, -12, -5.2))
}
