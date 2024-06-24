package main

import (
	"d-alejandro/training-level2/develop/dev09/getter"
	"flag"
	"fmt"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/*
./dev09 -l 2 https://www.acdc.com/
*/
func main() {
	levelMaxFlag := flag.Int("l", 1, "maximum nesting depth of pages")

	flag.Parse()

	if *levelMaxFlag < 1 {
		fmt.Println("flag value is less than one.")
		os.Exit(1)
	}

	arguments := flag.Args()

	if len(arguments) != 1 {
		fmt.Println("url not found.")
		os.Exit(1)
	}

	url := arguments[0]
	webGetter := getter.NewWebGetter(*levelMaxFlag)

	if err := webGetter.Execute(url); err != nil {
		fmt.Printf("download failed.\n%s", err.Error())
		os.Exit(1)
	}

	fmt.Println("download complete.")
}
