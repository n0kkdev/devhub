package stack

var x []string

// Push добавит новый элемент в стек
func Push(str string) {
	x = append(x, str)
}

// Pop вернет последний добавленный в стек элемент
func Pop() string {
	numOfElements := len(x)

	if numOfElements == 0 {
		return "Стек пуст"
	}
	popElem := x[numOfElements-1]
	x = x[:numOfElements-1]
	return popElem
}
