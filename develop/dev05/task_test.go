package main

import (
	"d-alejandro/training-level2/develop/dev05/search"
	"reflect"
	"testing"
)

func TestNewTextSearch(t *testing.T) {
	type args struct {
		flags      *search.TextSearchFlagDTO
		pattern    string
		inputArray []string
	}
	tests := []struct {
		name     string
		args     args
		expected []string
	}{
		{
			name: "no flags",
			args: args{
				flags:   &search.TextSearchFlagDTO{},
				pattern: "^May",
				inputArray: []string{
					"computer 7 Jan 20K",
					"May 3 Jul  2M",
					"LAPTOP 7 May 11G",
					"laptop|May 5 Jan 30B 3",
				},
			},
			expected: []string{
				"\u001B[31mMay\u001B[0m 3 Jul  2M",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			textSearch := search.NewTextSearch(test.args.flags)
			actualResponse := textSearch.Search(test.args.pattern, test.args.inputArray)

			if !reflect.DeepEqual(actualResponse, test.expected) {
				t.Errorf("textSearch.Search() = %v, expected: %v", actualResponse, test.expected)
			}
		})
	}
}
