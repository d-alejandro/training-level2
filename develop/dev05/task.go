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
	rowCountAfter := flag.Int("A", 0, "print row count after")
	rowCountBefore := flag.Int("B", 0, "print row count before")
	rowCountAround := flag.Int("C", 0, "print row count around")
	rowCount := flag.Int("c", 0, "print row count")

	flag.Parse()

	arguments := flag.Args()

	if len(arguments) < 2 {
		fmt.Println("Please set the search pattern and input file name.")
		os.Exit(1)
	}

	textSearchFlagDTO := &search.TextSearchFlagDTO{
		RowCountAfter:  *rowCountAfter,
		RowCountBefore: *rowCountBefore,
		RowCountAround: *rowCountAround,
		RowCount:       *rowCount,
	}

	textSearch := search.NewTextSearch(textSearchFlagDTO)
	rows := readFile(arguments[0])
	response := textSearch.Search(rows)
	fmt.Println(response)
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

func colorToRed(text string) string {
	return "\u001B[31m" + text + "\u001B[0m"
}
