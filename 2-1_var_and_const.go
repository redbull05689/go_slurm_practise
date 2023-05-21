package main

import "fmt"

func main() {
	var v1 int
	_ = v1 // Ignore unsed var
	v3 := 5
	fmt.Println(v3)

	const v2 int
	_ = v2 // Ignore unsed const

	_, _ = x, y

}
