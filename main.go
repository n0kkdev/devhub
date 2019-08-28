package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"./car"
	"./stack"
)

var car2 = car.Car{Mark: "Ваз 2114", Year: 2001, Vol: 654, Trunkvol: 310, Start: false, Openw: false}

func main() {
	var trunk car.Car
	trunk.SetVal("Зил 5301", 2003, 4500, 3200, true, true)
	fmt.Println(trunk.GetVal())
	fmt.Println(car.Сartrunk(car2.Vol), car2.Mark, car2.Year, car2.Vol, car.Procvol(car2.Trunkvol, car2.Vol), "%",
		car.Enginestart(car2.Start), car.Openwin(car2.Openw))

	b, err := json.Marshal(trunk)
	c, err := json.Marshal(car2)
	if err != nil {
		fmt.Println(err)
		return
	}
	stack.Push(string(b))
	stack.Push("Добавили в стек структуру trunk")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	stack.Push(string(c))
	stack.Push("Добавили в стек структуру car2")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	var ab = (string(b))
	var ac = (string(c))
	fmt.Println(strings.EqualFold(ab, ac))
	var as = strings.Join([]string{ab, ac}, "")
	as = (strings.ToUpper(as))

	file, err := os.Create("struct.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(as)


	if len(c) > len(b) {
		fmt.Println("длина car2 больше trunk")
	} else if len(c) < len(b) {
		fmt.Println("длина car2 меньше trunk")
	} else if len(c) == len(b) {
		fmt.Println("длина car2 равна trunk")
	} else {
		fmt.Println("длина не известна")
	}

	valcar := reflect.Indirect(reflect.ValueOf(car2))
	fmt.Println("Structure car2: ")
	for i := 0; i < valcar.NumField(); i++ {
		carfield := (valcar.Type().Field(i).Name)
		cartype := (valcar.Field(i).Type().Name())
		fmt.Println("Поле:", carfield, "Тип:", cartype)
	}

	valtrunk := reflect.Indirect(reflect.ValueOf(trunk))
	fmt.Println("Structure trunk: ")
	for i := 0; i < valtrunk.NumField(); i++ {
		trunkfield := (valtrunk.Type().Field(i).Name)
		trunktype := (valtrunk.Field(i).Type().Name())
		fmt.Println("Поле: ", trunkfield, "Тип: ", trunktype)
	}
}
