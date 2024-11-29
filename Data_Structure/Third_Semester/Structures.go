package Third_Semester

import "fmt"

type Lab_1 struct{}
type Lab_2 struct{}
type Lab_3 struct{}

func (l1 *Lab_1) Run() {
	start, result := l1.impl()
	fmt.Println("Исходный массив", start)
	fmt.Println("Результативный массив", result)
}

func (l2 *Lab_2) Run() {

}

func (l3 *Lab_3) Run() {

}
