package Third_Semester

import "fmt"

type Lab_1 struct{}
type Lab_2 struct{}
type Lab_3 struct{}

// Для Lab_3

type Stack struct {
	items    [10]int
	top      int
	capacity int
}

func NewStack() *Stack {
	return &Stack{
		top:      -1,
		capacity: 10,
	}
}

func (l1 *Lab_1) Run() {
	start, result := l1.impl()
	fmt.Println("Исходный массив", start)
	fmt.Println("Результативный массив", result)
}

func (l2 *Lab_2) Run() {
	start, res, i := l2.impl()
	fmt.Println("Исходный массив", start)
	fmt.Printf("Минимальное число в массиве %d под номером %d", res, i)
}

func (l3 *Lab_3) Run() {
	var Temp Lab_3
	Temp.impl()
}
