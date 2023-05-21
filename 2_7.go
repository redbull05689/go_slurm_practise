package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Set length of array")
	fmt.Scan(&n)
	doubleDetector(nums []int) 
}

func nums(n int) {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	return
}

doubleDetector(nums []int) bool {
	my_array := nums(n)
	for  k,_ := range my_array {
		 … }

}
/*
Task:
	Реализуйте функцию doubleDetector(nums []int) bool которой на вход подается целочисленный массив nums, 
	на выход должно вывести true, если хотя бы одно значение встречается в массиве как минимум дважды, 
	и false, если каждый элемент различен.
Roadmap:
	1)make array
	2)get array as an input to finction
	3) run doubleDetector
	4) print tesult
*/