package main

import (
	"fmt"

	"./animals"
)

func main() {
	var c animals.Dog
	c.SetPaws(4)
	var b animals.Ant
	b.SetTail(false)
	fmt.Println(animals.ComAnimal(c, b))
}
