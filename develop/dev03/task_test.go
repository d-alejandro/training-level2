package main

import (
	"d-alejandro/training-level2/develop/dev03/sorting"
	"reflect"
	"testing"
)

func TestArraySortingSort(t *testing.T) {
	inputArray := []string{
		"a) 21 Test1",
		"b) 1 Test2  ",
		"c)      2 Test3",
		"d) 21 Test4",
		"e) 50 Test5",
		"f) 1 Test6",
	}

	tests := []struct {
		name     string
		flags    *sorting.FlagDTO
		expected []string
	}{
		{
			name: "test flags: -k 2 -n -r -u",
			flags: &sorting.FlagDTO{
				ColumnNumberFlag:    2,
				SortByNumberFlag:    true,
				DescendingOrderFlag: true,
				UniqueFlag:          true,
			},
			expected: []string{
				"e) 50 Test5",
				"a) 21 Test1",
				"c)      2 Test3",
				"b) 1 Test2  ",
			},
		},
		{
			name: "test flags: -r",
			flags: &sorting.FlagDTO{
				DescendingOrderFlag: true,
			},
			expected: []string{
				"f) 1 Test6",
				"e) 50 Test5",
				"d) 21 Test4",
				"c)      2 Test3",
				"b) 1 Test2  ",
				"a) 21 Test1",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arraySorting := sorting.NewArraySorting(test.flags)
			if actualResponse := arraySorting.Sort(inputArray); !reflect.DeepEqual(actualResponse, test.expected) {
				t.Errorf("Sort(\n%#v\n) = \n%#v, \nexpected: \n%#v", inputArray, actualResponse, test.expected)
			}
		})
	}
}
