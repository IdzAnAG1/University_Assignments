package Third_Semester

import (
	"fmt"
	"math"
	"time"
)

// Функция реализующая первую лабораторную работу
func (l1 *Lab_1) impl() (start, result [10]int) {

	var StartingArray [10]int

	for i := 0; i < 10; i++ {
		fmt.Printf("Введите %d-ый элемент : ", i)
		fmt.Scan(&StartingArray[i])
	}
	start = StartingArray

	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			if StartingArray[i] > StartingArray[j] {
				StartingArray[i], StartingArray[j] = StartingArray[j], StartingArray[i]
			}
		}
	}

	result = StartingArray
	return
}

// Функция для теста первой лабораторной работы
func (l1 *Lab_1) implTest(input [10]int) (start, result [10]int) {
	start = input
	result = input
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return
}

// Функция реализующая вторую лабораторную работу
func (l2 *Lab_2) impl() (start [10]int, min, index int) {
	var (
		StartArr [10]int
	)

	min = math.MaxInt
	for i := 0; i < 10; i++ {

		fmt.Printf("Введите %d-ый элемент : ", i)
		fmt.Scan(&StartArr[i])
	}
	start = StartArr

	for i := 0; i < len(StartArr); i++ {
		if min > StartArr[i] {
			min = StartArr[i]
			index = i
		}
	}

	return
}

// Функция для теста второй лабораторной работы
func (l2 *Lab_2) implTest(input [10]int) (start [10]int, min, index int) {
	start = input

	for i := 0; i < len(start); i++ {
		if min > start[i] {
			min = start[i]
			index = i
		}
	}

	return
}
