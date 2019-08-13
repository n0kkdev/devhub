package main

import "fmt"

//проверка на целое
func evenNumb() {
	var num int
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}

//проверка деления на 3
func divThree() {
	var num int
	fmt.Scanln(&num)
	if num%3 == 0 {
		fmt.Println("Число делится на 3")
	} else {
		fmt.Println("Число не делится на 3")
	}
}

//вывод чисел Фибоначчи
func numFib() {
	var f1, f2, f3 int64
	f1 = 1
	f2 = 1
	for i := 3; i <= 100; i++ {
		f3 = f2 + f1
		f1 = f2
		f2 = f3
		fmt.Println(f3)
	}
}

func main() {
	fmt.Println("Введите число: ")
	evenNumb()
	fmt.Println("Введите число для деления: ")
	divThree()
	fmt.Println("Числа Фибоначчи: ")
	numFib()
}
