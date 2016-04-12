package main

import "fmt"

func count() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	a := count()

	fmt.Println(a())
	fmt.Println(a())

	b := count()

	fmt.Println(b())
        fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
