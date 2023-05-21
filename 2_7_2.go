package main

import (
	"fmt"
	"sort"
)

func main() {
	var input_string int
	fmt.Println("Set value for slice")
	fmt.Scan(&input_string)
	ww := make_slice(input_string)
	isSorted(ww)
}

func make_slice(n int) []string {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	return a
}

func isSorted(ww []string) bool {
	sorted_slice := sort.Strings(ww)
	if sorted_slice == ww {
		return true
	}

}

/*
task: Реализуйте функцию func isSorted(ww []string) bool которой на вход подается непустой слайс состоящий из строк,
необходимо вернуть true тогда и только тогда, когда все слова в слайсе отсортированы лексикографически по возрастанию относительно друг друга.
 В ином случае вывести false. Продемонстрировать использование функции
*/
