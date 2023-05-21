package main

type char_count struct {
	value int
	char  string
}

func main() {
	word := "wsdvsdk sd sdf hsdvh sdv sd"
	stringStat(word)
}

func stringStat(word string) string {
	for _, v := range word {
		chars_in_struct := char_count{v}
	}
	return chars_in_struct
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
