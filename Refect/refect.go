// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type myType struct {
// 	text string
// }

// func main() {

// 	text := myType{"我爱golang"}
// 	// sl := []int{1, 2, 3}
// 	slType := reflect.TypeOf(text)
// 	copy(slType, slType)
// }

// func copy(dst, src reflect.Type) {
// 	dn := dst.Name()
// 	fmt.Println(dn)
// 	fmt.Println("dk")
// 	dk := dst.Kind()
// 	fmt.Println(dk)
// }
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Email  string `mcl:"email"`
	Name   string `mcl:"name"`
	Age    int    `mcl:"age"`
	Github string `mcl:"github" default:"a8m"`
}

func main() {
	u := &User{Name: "Ariel Mashraki"}
	// Elem returns the value that the pointer u points to.
	v := reflect.ValueOf(u).Elem()
	f := v.FieldByName("Github")
	// make sure that this field is defined, and can be changed.
	if !f.IsValid() || !f.CanSet() {
		return
	}
	if f.Kind() != reflect.String || f.String() != "" {
		return
	}
	f.SetString("a8m1")
	fmt.Printf("Github username was changed to: %#v\n", u)
}
