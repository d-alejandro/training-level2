package main

import (
	"encoding/json"
	"reflect"
	"testing"
	"unsafe"
)

func Test_searchAnagramFromArray(t *testing.T) {
	tests := []struct {
		name     string
		input    unsafe.Pointer
		expected *map[string]*[]string
	}{
		{
			name:  `"Пятак", "Пятка", "Тяпка", "Листок", "Слиток", "Палка", "Столик"`,
			input: unsafe.Pointer(&[...]string{"Пятак", "Пятка", "Тяпка", "Листок", "Слиток", "Палка", "Столик"}),
			expected: &map[string]*[]string{
				"пятак":  {"пятка", "тяпка"},
				"листок": {"слиток", "столик"},
			},
		},
		{
			name: `"Полка", "Сон", "Кабан", "Банка", "Бакан", "Отк", "Ток", "Кот", "Кит", "ТИК", "Икт"`,
			input: unsafe.Pointer(
				&[...]string{"Полка", "Сон", "Кабан", "Банка", "Бакан", "Отк", "Ток", "Кот", "Кит", "ТИК", "Икт"},
			),
			expected: &map[string]*[]string{
				"кабан": {"бакан", "банка"},
				"отк":   {"кот", "ток"},
				"кит":   {"икт", "тик"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := searchAnagramFromArray(test.input); !reflect.DeepEqual(result, test.expected) {
				resultJSON, _ := json.MarshalIndent(result, "", "  ")
				expectedJSON, _ := json.MarshalIndent(test.expected, "", "  ")

				t.Errorf(
					"convertArrayPointerToSlice() = %s, \nexpected %v",
					string(resultJSON),
					string(expectedJSON),
				)
			}
		})
	}
}
