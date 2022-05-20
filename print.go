package main

import "fmt"

func main() {
	num := 0
	for i := 1; i <= 100; i++ {
		num = i
		if (num%3) == 0 && (num%15) != 0 {
			fmt.Println("Fizz")
		} else if (num%5) == 0 && (num%15) != 0 {
			fmt.Println("Buzz")
		} else if (num % 15) == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(num)
		}
	}
}
