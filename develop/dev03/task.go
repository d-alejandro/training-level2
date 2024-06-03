package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

/*
=== Утилита sort ===

Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные
параметры): на входе подается файл из несортированных строк, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель —
пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/*
./task3 -r input.txt
*/
func main() {
	descendingOrderFlag := flag.Bool("r", false, "descending order")
	sortByNumberFlag := flag.Bool("n", false, "sort by number")

	flag.Parse()

	arguments := flag.Args()

	if len(arguments) == 0 {
		fmt.Println("Please set the input file name.")
		os.Exit(1)
	}

	strings := readFile(arguments[0])

	writeFile(strings)

	fmt.Println(*descendingOrderFlag, *sortByNumberFlag, strings)
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

func writeFile(strings []string) {
	const OutputFileName = "output.txt"
	file, err := os.Create(getCurrentPath() + "/" + OutputFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, line := range strings {
		if _, errorWrite := fmt.Fprintln(file, line); errorWrite != nil {
			fmt.Println(errorWrite)
			os.Exit(1)
		}
	}

	closeFile(file)
}

func getCurrentPath() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return path
}
