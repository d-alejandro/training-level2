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
					"LAPTOP 7 ^May 11G",
					"laptop|May 5 Jan 30B 3",
				},
			},
			expected: []string{"\u001B[31mMay\u001B[0m 3 Jul  2M"},
		},
		{
			name: "not found",
			args: args{
				flags: &search.TextSearchFlagDTO{
					RowsAfterFlag:  1,
					RowsBeforeFlag: 2,
					FixedFlag:      true,
					LineNumFlag:    true,
				},
				pattern:    "test",
				inputArray: []string{"computer 7 Jan 20K"},
			},
			expected: nil,
		},
		{
			name: "not found with flag -c",
			args: args{
				flags:      &search.TextSearchFlagDTO{CountFlag: true},
				pattern:    "test",
				inputArray: []string{"computer 7 Jan 20K"},
			},
			expected: []string{"0"},
		},
		{
			name: "-A 1 -B 2 -n -F",
			args: args{
				flags: &search.TextSearchFlagDTO{
					RowsAfterFlag:  1,
					RowsBeforeFlag: 2,
					FixedFlag:      true,
					LineNumFlag:    true,
				},
				pattern: "^May",
				inputArray: []string{
					"computer 7 Jan 20K",
					"May 3 Jul  2M",
					"LAPTOP 7 ^May 11G",
					"laptop|May 5 Jan 30B 3",
				},
			},
			expected: []string{
				"\u001B[32m1\u001B[0m\u001B[34m-\u001B[0mcomputer 7 Jan 20K",
				"\u001B[32m2\u001B[0m\u001B[34m-\u001B[0mMay 3 Jul  2M",
				"\u001B[32m3\u001B[0m\u001B[34m:\u001B[0mLAPTOP 7 \u001B[31m^May\u001B[0m 11G",
				"\u001B[32m4\u001B[0m\u001B[34m-\u001B[0mlaptop|May 5 Jan 30B 3",
			},
		},
		{
			name: "-A 1",
			args: args{
				flags:   &search.TextSearchFlagDTO{RowsAfterFlag: 1},
				pattern: "May",
				inputArray: []string{
					"May 3 Jul  2M",
					"computer 7 Jan 20K",
					"LAPTOP 7 Jan 11G",
					"laptop|May 5 Jan 30B 3",
				},
			},
			expected: []string{
				"\u001B[31mMay\u001B[0m 3 Jul  2M",
				"computer 7 Jan 20K",
				"\u001B[34m--\u001B[0m",
				"laptop|\u001B[31mMay\u001B[0m 5 Jan 30B 3",
			},
		},
		{
			name: "-C 1 -n -F",
			args: args{
				flags: &search.TextSearchFlagDTO{
					RowsContextFlag: 1,
					FixedFlag:       true,
					LineNumFlag:     true,
				},
				pattern: "^May",
				inputArray: []string{
					"computer 7 Jan 20K",
					"May 3 Jul  2M",
					"LAPTOP 7 ^May 11G",
					"laptop|May 5 Jan 30B 3",
				},
			},
			expected: []string{
				"\u001B[32m2\u001B[0m\u001B[34m-\u001B[0mMay 3 Jul  2M",
				"\u001B[32m3\u001B[0m\u001B[34m:\u001B[0mLAPTOP 7 \u001B[31m^May\u001B[0m 11G",
				"\u001B[32m4\u001B[0m\u001B[34m-\u001B[0mlaptop|May 5 Jan 30B 3",
			},
		},
		{
			name: "-i",
			args: args{
				flags:   &search.TextSearchFlagDTO{IgnoreCaseFlag: true},
				pattern: "laptop",
				inputArray: []string{
					"LAPTOP 7 ^May 11G",
					"computer 7 Jan 20K",
					"May 3 Jul  2M",
					"laptop|May 5 Jan 30B 3",
				},
			},
			expected: []string{
				"\u001B[31mLAPTOP\u001B[0m 7 ^May 11G",
				"\u001B[31mlaptop\u001B[0m|May 5 Jan 30B 3",
			},
		},
		{
			name: "-v",
			args: args{
				flags:   &search.TextSearchFlagDTO{InvertFlag: true},
				pattern: "Jul",
				inputArray: []string{
					"LAPTOP 7 ^May 11G",
					"computer 7 Jan 20K",
					"May 3 Jul  2M",
					"laptop|May 5 Jan 30B 3",
				},
			},
			expected: []string{
				"LAPTOP 7 ^May 11G",
				"computer 7 Jan 20K",
				"laptop|May 5 Jan 30B 3",
			},
		},
		{
			name: "-c -v",
			args: args{
				flags: &search.TextSearchFlagDTO{
					CountFlag:  true,
					InvertFlag: true,
				},
				pattern: "Jul",
				inputArray: []string{
					"LAPTOP 7 ^May 11G",
					"computer 7 Jan 20K",
					"May 3 Jul  2M",
					"laptop|May 5 Jan 30B 3",
				},
			},
			expected: []string{"3"},
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
