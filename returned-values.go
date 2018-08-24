package main

import "fmt"

func E(inA float64, inB float64, inC float64) (outA []float64, outB []float64, outC []float64) {
	return []float64{inA}, []float64{inB}, []float64{inC}
}

func vals() []int {
	return []int{3}
}

func main() {

	a := vals()
	fmt.Println(a)
	_, c, _ := E(float64(0.0), float64(0.0), float64(0.0))
	fmt.Println(c)
}
