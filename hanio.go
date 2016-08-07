package main

import "fmt"

func hanio(a, b, c string, n int) {
	if n == 1 {
		fmt.Println(a, "=>", c)
		return
	}
	hanio(a, c, b, n-1)
	fmt.Println(a, "=>", c)
	hanio(b, a, c, n-1)
}
func main() {
	hanio("a", "b", "c", 3)
}
