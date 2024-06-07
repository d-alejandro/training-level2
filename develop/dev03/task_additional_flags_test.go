package main

import (
	"d-alejandro/training-level2/develop/dev03/sorting"
	"reflect"
	"testing"
)

func TestArraySortingSortWithAdditionalFlags(t *testing.T) {
	inputArray := []string{
		"1) Jul    20K * 1P",
		"2) Dec 2M * 3T",
		"3) Mar   11G * 11G",
		"4) Mar 1P * 2M",
		"5) May 3T * 20K",
		"6)   Jan 30B * 2K",
		"7) Jan 2K * 30B",
	}

	tests := []struct {
		name     string
		flags    *sorting.FlagDTO
		expected []string
	}{
		{
			name: "test flags: -k 2 -M -b",
			flags: &sorting.FlagDTO{
				ColumnNumberFlag: 2,
				MonthNameFlag:    true,
				EndingSpaceFlag:  true,
			},
			expected: []string{
				"6) Jan 30B * 2K",
				"7) Jan 2K * 30B",
				"3) Mar 11G * 11G",
				"4) Mar 1P * 2M",
				"5) May 3T * 20K",
				"1) Jul 20K * 1P",
				"2) Dec 2M * 3T",
			},
		},
		{
			name: "test flags: -k 3 -h",
			flags: &sorting.FlagDTO{
				ColumnNumberFlag: 3,
				HumanNumericFlag: true,
			},
			expected: []string{
				"6)   Jan 30B * 2K",
				"7) Jan 2K * 30B",
				"1) Jul    20K * 1P",
				"2) Dec 2M * 3T",
				"3) Mar   11G * 11G",
				"5) May 3T * 20K",
				"4) Mar 1P * 2M",
			},
		},
		{
			name: "test flags: -k 5 -h -r -c",
			flags: &sorting.FlagDTO{
				ColumnNumberFlag:    5,
				HumanNumericFlag:    true,
				DescendingOrderFlag: true,
				SortCheckFlag:       true,
			},
			expected: []string{
				"true",
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
