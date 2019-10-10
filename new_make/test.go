package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// ret := f2()
	// fmt.Println(ret)
	// ret1 := funcName("321")
	// fmt.Println(ret1)
	// var iBuffer [10]int
	// slice := []byte{1, 2, 3, 4, 5, 5, 6}
	// for i := 0; i < len(slice); i++ {
	// 	slice[i] = byte(i)
	// }
	// slice[4] = slice[4] + 100
	// fmt.Println("before", slice)
	// AddOneToEachElement(slice)
	// fmt.Println("after", slice)

	// var iBuffer [19]int
	// slice := iBuffer[0:0]
	// for i := 0; i < 20; i++ {
	// 	slice = Extend(slice, i)
	// 	fmt.Println(slice)
	// }

	// slice1 := []int{0, 1, 2, 3, 4}
	// slice2 := []int{55, 66, 77}
	// fmt.Println(slice1)
	// slice1 = append(slice1, slice2...) // The '...' is essential!
	// fmt.Println(slice1)

	// b := [30]string{"Penn", "Teller"}
	// usr := "/usr/ken"
	// slice := []byte(usr)
	// slice[2] = 'p'
	// fmt.Printf("%+v\n", len(b))

	i := int(1)
	j := int32(1)
	f := float32(3.14)
	fmt.Println(unsafe.Sizeof(i), unsafe.Sizeof(j), unsafe.Sizeof(f))

}
func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}
func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i] = slice[i] + 100
	}
}

func f2() int {
	t := 5
	tp := &t
	defer func() {
		*tp = *tp + 5
	}()
	return t
}

func funcName(a interface{}) string {
	v, ok := a.(string)
	if !ok {
		return ""
	}
	return v
}
