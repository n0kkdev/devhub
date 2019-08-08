package main

import "fmt"

func main() {
	const course float32 = 65.05
	var r float32
	fmt.Println("Сколько рублей?")
	fmt.Scanln(&r)
	var d float32
	d = r / course
	fmt.Printf("%.2f", d)
	fmt.Println(" Долларов")
}