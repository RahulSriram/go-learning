package main

import "fmt"

func temp(i ...int) (int, int) {
	return i[1], len(i)
}

func temp2(i ...int) (int, int) {
        return i[2], len(i)
}

func main() {
	m := map[int]int{0 : 0, 1 : 10, 2 : 20, 3 : 30, 4 : 40}

	for i, v := range "hey" {
		fmt.Println(i, "->", v);
	}

	for i, v := range m {
		fmt.Println(i, "->", v);
	}

	a, b := temp(111, 222, 444, 333);
	fmt.Println(a, b);
}
