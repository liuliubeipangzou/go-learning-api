package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func isAdult(age int) bool {
	return age >= 18
}

func main() {
	name := "tulei"
	age := 25
	fmt.Println(name)
	fmt.Println("age:", age)
	fmt.Println("10 + 20 = ", add(10, 20))

	if isAdult(age) {
		fmt.Println(name, "isadult")
	} else {
		fmt.Println(name, "isminor")
	}

	for i := 0; i <= 5; i++ {
		fmt.Println("count:", i)
	}
}
