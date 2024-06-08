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
	array := [...]string{"Пятак", "Пятка", "Тяпка", "Листок", "Слиток", "Палка", "Столик"}

	response := *searchAnagramFromArray(unsafe.Pointer(&array))

	for key, pointer := range response {
		fmt.Println(key, convertArrayPointerToSlice(pointer))
	}
}

func searchAnagramFromArray(arrayPointer unsafe.Pointer) *map[string]unsafe.Pointer {
	const ElementArrayMaxCount = 50

	type tempSlice struct {
		value       string
		sortedValue string
	}

	stringSlice := convertArrayPointerToSlice(arrayPointer)

	tempSlices := make([]tempSlice, len(stringSlice))

	for key, value := range stringSlice {
		lowerCaseString := strings.ToLower(value)

		sortedString := []rune(lowerCaseString)
		slices.Sort(sortedString)

		tempSlices[key] = tempSlice{
			value:       lowerCaseString,
			sortedValue: string(sortedString),
		}
	}

	slices.SortFunc(tempSlices, func(a, b tempSlice) int {
		return strings.Compare(a.value, b.value)
	})

	tempMap := make(map[string][]string)

	for _, tempSlice := range tempSlices {
		tempMap[tempSlice.sortedValue] = append(tempMap[tempSlice.sortedValue], tempSlice.value)
	}

	outputMap := make(map[string]unsafe.Pointer)

	for _, slice := range tempMap {
		if len(slice) == 1 {
			continue
		}

		tempArray := [ElementArrayMaxCount]string{}

		for key, value := range slice[1:] {
			tempArray[key] = value
		}

		outputMap[slice[0]] = unsafe.Pointer(&tempArray)
	}

	return &outputMap
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
