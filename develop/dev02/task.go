package main

import (
	"bufio"
	"d-alejandro/training-level2/develop/dev02/parsing"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println("Введите строку:")

	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	inputString = strings.TrimSuffix(inputString, "\n")
	outputMessage := unpackString(inputString)

	fmt.Println("Output: ", outputMessage)
}

func unpackString(input string) string {
	arrayRune := []rune(input)

	var outputMessage string

	arrayLength := len(arrayRune)

	switch {
	case arrayLength == 0:
		outputMessage = ""
	case unicode.IsDigit(arrayRune[0]):
		outputMessage = "некорректная строка"
	case arrayLength == 1:
		outputMessage = string(arrayRune)
	default:
		parser := parsing.NewParser(arrayRune)
		outputMessage = parser.Parse()
	}
	return outputMessage
}
