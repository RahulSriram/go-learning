package main

import "fmt"

func main() {
	var a string = "there"
	fmt.Println("hey " + a)
	var b, c int = 25/5, 20/5
	fmt.Println("25/5 : 20/5->", b, c)
	var d = true
	fmt.Println("Yep, this is", d)

	x := "here-> "
	xy := "\n256/8 ="
	y := 256/8
	xz := "\n3.14/2 ="
	z := 3.14/2
	fmt.Println(x, xy, y, xz, z)
}
