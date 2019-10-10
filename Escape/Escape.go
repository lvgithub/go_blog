package main

type person struct {
	name string
	age  int
}

func main() {
	makePerson(32, "艾玛·斯通")
	showPerson(33, "杨幂")

	fun := closure()
	fun()
	fun()
	fun()
}

func makePerson(age int, name string) *person {
	maliya := person{name, age}
	return &maliya
}

func showPerson(age int, name string) person {
	yangmi := person{name, age}
	return yangmi
}

func closure() func() int {
	i := 100
	return func() int {
		i++
		return i
	}
}
