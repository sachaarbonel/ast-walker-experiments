package main

import "fmt"

func E() ([]float64, []float64, []float64) {
	return []float64{0.0, 0.0}, []float64{0.0, 0.0}, []float64{0.0, 0.0}
}

func vals() []int {
	return []int{3}
}

func main() {

	a := vals()
	fmt.Println(a)
	_, c, _ := E()
	fmt.Println(c)
}
