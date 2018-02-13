package main

import "fmt"

func main() {

	for count := 1; count <= 100; count++ {
		fmt.Println(count)
		if count%15 == 0 {
			fmt.Println("fizzbuzz\n")
			continue
		}
		if count%3 == 0 {
			fmt.Println("fizz\n")
		}
		if count%5 == 0 {
			fmt.Println("buzz\n")
		}

	}

}
