package main

import "fmt"

func main() {
	fmt.Println("Enter desired Fibonacci sequence length")
	var n int
	fmt.Scanf("%d", &n)
	if n < 1 || n > 92 {
		fmt.Println("Length should be between 1 and 92")
		return
	}
	var a, b int64 = 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%v ", b)
		a, b = b, a
		b += a
	}
}
