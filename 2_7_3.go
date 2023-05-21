package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	word := "Hello Go"
	stringStat(word)
}

func stringStat(word string) {
	tt := make(map[string]int)
	for _, c := range strings.ToLower(word) {
		if unicode.IsLetter(c) {
			char := strconv.QuoteRune(c)
			if val, ok := tt[char]; ok {
				val++
				tt[char] = val
			} else {
				tt[char] = 1
			}
		}
	}
	fmt.Println(fmt.Sprintf("Map: %v", tt))
}

/*
task:
Реализуйте функцию func stringStat(word string) string которая используя
тип данных map посчитает сколько раз встречается каждый символ во входной
строке и вернет строку в формате символ - количество для каждого символа новая строка.
Вывод инфорации о символе должен осуществляться в том же порядке,
в каком он встречается во входной фразе в первый раз.
Все символы переводить в нижний регистр, пробелы не считать.

Пример: Пример: Для строки “Старт и финиш” вывод будет следующий
с - 1
т - 2
а - 1
р - 1
и - 3
ф - 1
н - 1
ш - 1

roadmap:
	split string to structure
	del same value from struct
	count values
*/
