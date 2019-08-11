package main

import "fmt"

func main() {
	var m float32
	var p float32
	fmt.Println("Сколько рублей?")
	fmt.Scanln(&m)
	fmt.Println("Kaкой процент?")
	fmt.Scanln(&p)
	a := (p / 100 * m) + m*5
	fmt.Println("За 5 лет накопление составит", a)

}
