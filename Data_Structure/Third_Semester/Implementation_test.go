package Third_Semester

import (
	"testing"
)

func Test_Lab_1_Impl(t *testing.T) {
	tests := []struct {
		name     string
		input    [10]int
		Expected [10]int
	}{
		{
			name:     "Обычный случай(Случайный ввод символов)",
			input:    [10]int{4, 8, 9, 1, 0, 5, 7, 2, 0, 10},
			Expected: [10]int{0, 0, 1, 2, 4, 5, 7, 8, 9, 10},
		},
		{
			name:     "Обратно отсортированный массив",
			input:    [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			Expected: [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Отрицательные числа",
			input:    [10]int{-1, -199, -6, -8, -2, -3, -4, -5, -9, -10},
			Expected: [10]int{-199, -10, -9, -8, -6, -5, -4, -3, -2, -1},
		},
		{
			name:     "Отрицательные и положительные числа",
			input:    [10]int{1, -179, 0, -8, 2, 13, 4, -5, 19, -10},
			Expected: [10]int{-179, -10, -8, -5, 0, 1, 2, 4, 13, 19},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l1 := &Lab_1{}
			Start, Result := l1.implTest(test.input)
			if Start != test.input {
				t.Errorf("Начальный массив изменен %v, ожидается %v", Start, test.input)
			}
			if Result != test.Expected {
				t.Errorf("Сортировка не верна : получили %v, одидали : %v", Result, test.Expected)
			}
		})
	}
}
