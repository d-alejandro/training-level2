package main

import (
	"fmt"
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
	return &map[string]unsafe.Pointer{
		"test": arrayPointer,
	}
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

		stringSlice = append(stringSlice, value)

		index++
	}

	return stringSlice
}
