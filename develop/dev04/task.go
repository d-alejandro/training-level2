package main

import (
	"fmt"
	"slices"
	"strings"
	"unicode"
	"unsafe"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	wordArray := [...]string{"Пятак", "Пятка", "Тяпка", "Листок", "Слиток", "Палка", "Столик"}

	response := *searchAnagramFromArray(unsafe.Pointer(&wordArray))

	for key, anagramSlice := range response {
		fmt.Println(key, "=>", *anagramSlice)
	}
}

func searchAnagramFromArray(wordArrayPointer unsafe.Pointer) *map[string]*[]string {
	wordSlice := convertArrayPointerToSlice(wordArrayPointer)

	structureSlice := make([]struct {
		word       string
		sortedWord string
	}, len(wordSlice))

	for key, word := range wordSlice {
		lowerCaseWord := strings.ToLower(word)

		sortedWord := []rune(lowerCaseWord)
		slices.Sort(sortedWord)

		structureSlice[key] = struct {
			word       string
			sortedWord string
		}{
			word:       lowerCaseWord,
			sortedWord: string(sortedWord),
		}
	}

	tempAnagramMap := make(map[string][]string)

	for _, structure := range structureSlice {
		tempAnagramMap[structure.sortedWord] = append(tempAnagramMap[structure.sortedWord], structure.word)
	}

	return transformAnagramMap(tempAnagramMap)
}

func convertArrayPointerToSlice(arrayPointer unsafe.Pointer) []string {
	const StringDescriptorSize = unsafe.Sizeof("")

	var (
		stringSlice []string
		index       uintptr
	)

	for {
		offset := index * StringDescriptorSize
		value := *(*string)(unsafe.Pointer(uintptr(arrayPointer) + offset))

		if value == "" {
			break
		}

		firstSymbol := []rune(value[:2])[0]

		if !unicode.Is(unicode.Cyrillic, firstSymbol) {
			break
		}

		stringSlice = append(stringSlice, value)

		index++
	}

	return stringSlice
}

func transformAnagramMap(tempAnagramMap map[string][]string) *map[string]*[]string {
	outputAnagramMap := make(map[string]*[]string)

	for _, anagramSlice := range tempAnagramMap {
		if len(anagramSlice) == 1 {
			continue
		}

		tempAnagramSlice := anagramSlice[1:]

		slices.SortFunc(tempAnagramSlice, func(a, b string) int {
			return strings.Compare(a, b)
		})

		outputAnagramMap[anagramSlice[0]] = &tempAnagramSlice
	}

	return &outputAnagramMap
}
