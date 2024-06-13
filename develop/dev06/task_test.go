package main

import (
	"d-alejandro/training-level2/develop/dev06/cut"
	"reflect"
	"testing"
)

func TestTextCut_Cut(t *testing.T) {
	type args struct {
		flags     *cut.TextCutFlagDTO
		inputRows []string
	}
	tests := []struct {
		name     string
		args     args
		expected []string
	}{
		{
			name: "default flags (-f 1 -d '\t')",
			args: args{
				flags: &cut.TextCutFlagDTO{
					FieldsFlag:    "1",
					DelimiterFlag: "\t",
				},
				inputRows: []string{
					"computer	7	Jan	20K",
					"mouse	3	Jul	2M",
					"laptop_5_Jan_30B",
					"debian	4	Dec	1E",
				},
			},
			expected: []string{
				"computer",
				"mouse",
				"laptop_5_Jan_30B",
				"debian",
			},
		},
		{
			name: "-f 2,3 -d '\t' -s",
			args: args{
				flags: &cut.TextCutFlagDTO{
					FieldsFlag:    "2,3",
					DelimiterFlag: "\t",
					SeparatedFlag: true,
				},
				inputRows: []string{
					"computer	7	Jan	20K",
					"mouse	3	Jul	2M",
					"laptop_5_Jan_30B",
					"debian	4	Dec	1E",
				},
			},
			expected: []string{
				"7	Jan",
				"3	Jul",
				"4	Dec",
			},
		},
		{
			name: "-f 1,4 -d '-'",
			args: args{
				flags: &cut.TextCutFlagDTO{
					FieldsFlag:    "1,4",
					DelimiterFlag: "-",
				},
				inputRows: []string{
					"computer-7-Jan-20K",
					"mouse-3-Jul-2M",
					"laptop-5-Jan-30B",
					"debian-4-Dec-1E",
				},
			},
			expected: []string{
				"computer-20K",
				"mouse-2M",
				"laptop-30B",
				"debian-1E",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			textCut := cut.NewTextCut(test.args.flags)

			if response := textCut.Cut(test.args.inputRows); !reflect.DeepEqual(response, test.expected) {
				t.Errorf("textCut.Cut() = %v, expected: %v", response, test.expected)
			}
		})
	}
}
