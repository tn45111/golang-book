package main

import "fmt"

func main() {

	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	fmt.Println("8")
	fmt.Println("9")
	fmt.Println("10")

	for count := 0; count < 10; count++ {
		fmt.Println(count)
	}

	i := 90
	for i <= 100 {
		fmt.Println(i)
		i = i + 1
	}
	i=200
	for {
		fmt.Println(i)
		i = i + 1
		if i>210 { 
			break}
	}

}
