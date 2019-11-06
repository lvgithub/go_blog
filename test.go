package main

import (
	"fmt"
)

func main() {
	i := [5]int{1, 3, 4, 56, 3}
	fmt.Println(i)

	j := [5]int{}
	copy(j, i)
	fmt.Println(j)
}
