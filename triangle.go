package main

import (
	"fmt"
	"math"
)

func main() {
	var kat1 float64 = 8
	var kat2 float64 = 3
	var s float64
	var p float64
	k1 := math.Pow(kat1, 2)
	k2 := math.Pow(kat2, 2)
	g := math.Sqrt(k1 + k2)
	s = (g + kat1 + kat2) * 0.5
	p = g + kat1 + kat2

	fmt.Println(kat1, ",", kat2, "Катеты")
	fmt.Printf("%.3f", g)
	fmt.Println(" Гипотенуза")
	fmt.Printf("%.3f", s)
	fmt.Println(" Площадь")
	fmt.Printf("%.3f", p)
	fmt.Println(" Периметр")

}
