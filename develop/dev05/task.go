package main

import (
	"bufio"
	"d-alejandro/training-level2/develop/dev05/search"
	"flag"
	"fmt"
	"os"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	rowsAfterFlag := flag.Int("A", 0, "print rows after")
	rowsBeforeFlag := flag.Int("B", 0, "print rows before")
	rowsContextFlag := flag.Int("C", 0, "print rows around")
	countFlag := flag.Bool("c", false, "print row number")
	ignoreCaseFlag := flag.Bool("i", false, "ignore case")
	invertFlag := flag.Bool("v", false, "invert")
	fixedFlag := flag.Bool("F", false, "fixed")
	lineNumFlag := flag.Bool("n", false, "line num")

	flag.Parse()

	arguments := flag.Args()

	if len(arguments) < 2 {
		fmt.Println("Please set the search pattern and input file name.")
		os.Exit(1)
	}

	textSearchFlagDTO := &search.TextSearchFlagDTO{
		RowsAfterFlag:   *rowsAfterFlag,
		RowsBeforeFlag:  *rowsBeforeFlag,
		RowsContextFlag: *rowsContextFlag,
		CountFlag:       *countFlag,
		IgnoreCaseFlag:  *ignoreCaseFlag,
		InvertFlag:      *invertFlag,
		FixedFlag:       *fixedFlag,
		LineNumFlag:     *lineNumFlag,
	}

	textSearch := search.NewTextSearch(textSearchFlagDTO)

	pattern := arguments[0]
	inputRows := readFile(arguments[1])
	response := textSearch.Search(pattern, inputRows)

	for _, outputRow := range response {
		fmt.Println(outputRow)
	}
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	var strings []string

	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	closeFile(file)

	return strings
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
