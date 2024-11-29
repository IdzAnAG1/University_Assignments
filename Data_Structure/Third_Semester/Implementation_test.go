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

func Test_Lab_2_Impl(t *testing.T) {
	tests := []struct {
		name        string
		input       [10]int
		expectedArr [10]int
		expectedMin int
	}{
		{
			name:        "Случайный тест",
			input:       [10]int{4, 8, 9, 1, 0, 5, 7, 2, 0, 1},
			expectedArr: [10]int{4, 8, 9, 1, 0, 5, 7, 2, 0, 1},
			expectedMin: 0,
		},
		{
			name:        "Тест с большими числами",
			input:       [10]int{12346, 4561, 5343, 1, 1095, 467, 7, 7692, 6543, 1246},
			expectedArr: [10]int{12346, 4561, 5343, 1, 1095, 467, 7, 7692, 6543, 1246},
			expectedMin: 1,
		},
		{
			name:        "Тест с отрицательными числами",
			input:       [10]int{-49, -8, -9, -67, -2, -95, -7, -2, -100, -11},
			expectedArr: [10]int{-49, -8, -9, -67, -2, -95, -7, -2, -100, -11},
			expectedMin: -100,
		},
		{
			name:        "Смешанный тест",
			input:       [10]int{-49, 8956, -9, 67, -25, -95, -7754653, 2, 0, -1},
			expectedArr: [10]int{-49, 8956, -9, 67, -25, -95, -7754653, 2, 0, -1},
			expectedMin: -7754653,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l2 := &Lab_2{}

			start, minimum, _ := l2.implTest(test.input)

			if start != test.input {
				t.Errorf("Не сходится входной массив ожидалось %v получили %v", test.input, start)
			}

			if minimum != test.expectedMin {
				t.Errorf("Не сходится, ожидалось %v получили %v", test.expectedMin, minimum)
			}

		})
	}
}
