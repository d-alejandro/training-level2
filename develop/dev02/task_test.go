package main

import "testing"

func Test_unpackString(t *testing.T) {
	tests := []struct {
		argument string
		expected string
	}{
		{
			argument: "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			argument: "abcd",
			expected: "abcd",
		},
		{
			argument: "45",
			expected: "некорректная строка",
		},
		{
			argument: "",
			expected: "",
		},
		{
			argument: "v",
			expected: "v",
		},
		{
			argument: "QQ3f2d9y",
			expected: "QQQQffdddddddddy",
		},
	}
	for _, test := range tests {
		t.Run(test.argument, func(t *testing.T) {
			if actualString := unpackString(test.argument); actualString != test.expected {
				t.Errorf(`unpackString("%v") = "%v", expected "%v"`, test.argument, actualString, test.expected)
			}
		})
	}
}

func Test_unpackStringWithEscapeSymbol(t *testing.T) {
	tests := []struct {
		argument string
		expected string
	}{
		{
			argument: `qwe\4\5`,
			expected: "qwe45",
		},
		{
			argument: `qwe\45`,
			expected: "qwe44444",
		},
		{
			argument: `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			argument: `Xv3\1\2R1\\3yy\72z3`,
			expected: `Xvvv12R\\\yy77zzz`,
		},
	}
	for _, test := range tests {
		t.Run(test.argument, func(t *testing.T) {
			if actualString := unpackString(test.argument); actualString != test.expected {
				t.Errorf(`unpackString("%v") = "%v", expected "%v"`, test.argument, actualString, test.expected)
			}
		})
	}
}
