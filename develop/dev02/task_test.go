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
			argument: "QQ3f2d12y",
			expected: "QQQQffddddddddddddy",
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
