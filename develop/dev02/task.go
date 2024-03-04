package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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

func unpuck(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}
	res := ""
	if unicode.IsLetter(rune(str[0])) {
		str := strings.Split(str, "")
		for i := 0; i < len(str); i++ {
			if str[i] == `\` {
				res += str[i+1]
				i++
			} else {
				if n, err := strconv.Atoi(string(str[i])); err == nil {
					for j := 0; j < n-1; j++ {
						res += string(str[i-1])
					}
				} else {
					res += str[i]
				}
			}

		}
	} else {
		return "", errors.New("некорректная строка")
	}

	return res, nil
}

func main() {
	str := []string{"a4bc2d5e", "abcd", "45", ""}
	// str := []string{`qwe\4\5`, `qwe\45`, `qwe\\5`}
	unpuckStr := make([]string, 0, len(str))
	for _, v := range str {
		unpuckV, err := unpuck(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Errror by unpucking : %s\n", err)
		}
		unpuckStr = append(unpuckStr, unpuckV)
	}
	fmt.Println(unpuckStr)
}
