package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
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

Требования:
1. Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
2. Выходные данные: ссылка на мапу множеств анаграмм
3. Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
4. Массив должен быть отсортирован по возрастанию.
5. Множества из одного элемента не должны попасть в результат.
6. Все слова должны быть приведены к нижнему регистру.
7. В результате каждое слово должно встречаться только один раз.



*/

func main() {
	words := []string{"пятак", "Пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	lower_words := sliceToLower(words)
	fmt.Println(*sliceUnique(lower_words))
	// for i := range *sliceUnique(lower_words) {
	// 	fmt.Println(i)
	// }
}

func sliceToLower(words []string) []string {
	lower_words := make([]string, 0, len(words))
	mapWords := make(map[string]bool, len(words))
	for _, v := range words {
		lower_v := strings.ToLower(v)
		if _, ok := mapWords[lower_v]; !ok {
			mapWords[lower_v] = true
			lower_words = append(lower_words, lower_v)
		}
	}
	return lower_words
}

func wordToMap(word string) map[string]int {
	tempMap := make(map[string]int)
	for _, v := range strings.Split(word, "") {
		tempMap[v]++
	}
	return tempMap
}

func sliceUnique(words []string) *map[string][]string {
	mapa := make(map[string][]string)

	words = sliceToLower(words)

	for _, word := range words {
		mapWord := wordToMap(word)
		f := true
		for k := range mapa {
			if reflect.DeepEqual(mapWord, wordToMap(k)) {
				mapa[k] = append(mapa[k], word)
				f = false
				break
			}
		}
		if f {
			mapa[word] = append(mapa[word], word)
		}
	}
	for k := range mapa {
		sort.Strings(mapa[k])
	}

	return &mapa
}
