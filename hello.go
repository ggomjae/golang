package main

import "fmt"

func main() {
	fmt.Println("hello")
	var a int = 1
	b := "oh"
	fmt.Println(a, b)

	sum := 0
	for index := 0; index <= 10; index++ {
		sum += index
	}

	fmt.Println(sum)

	i := 0
	var secondSum int = 0
	for i <= 10 {
		secondSum += i
		i++
	}

	fmt.Println(secondSum)
}
