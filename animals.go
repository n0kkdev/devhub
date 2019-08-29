package animals

import "fmt"

//Dog struct
type Dog struct {
	Paws int
}

//Ant struct
type Ant struct {
	Tail bool
}

//GetPaws ...
func (c Dog) GetPaws() int {
	return c.Paws
}

//SetPaws ...
func (c *Dog) SetPaws(paws int) {
	c.Paws = paws
}

//Mammals ...
func (c Dog) Mammals() string {
	if c.Paws == 4 {
		a := "Mammals"
		fmt.Println(a)
	} else {
		a := "Not Mammals"
		fmt.Println(a)
	}
	return ""
}

//GetTail  ...
func (b Ant) GetTail() bool {
	return b.Tail
}

//SetTail  ...
func (b *Ant) SetTail(tail bool) {
	b.Tail = tail
}

//Mammals ...
func (b Ant) Mammals() string {
	if b.Tail == true {
		a := "Mammals"
		fmt.Println(a)
	} else {
		a := "Not Mammals"
		fmt.Println(a)
	}
	return ""
}

//ComAnimal ...
func ComAnimal(mammals ...Animal) string {
	res := ""
	for _, mammal := range mammals {
		res = mammal.Mammals()
	}
	return res
}

//Animal ...
type Animal interface {
	Mammals() string
}
