package main

import (
	"fmt"
	"math/big"
	"os"
)

func sum(x, y int) int {
	result := x + y
	return result
}

func multiply(x, y int) int {
	result := x * y
	return result
}

func prime(input int) []int {
	var result []int
	count := 1
	i := 1

	for {
		if count <= input {
			if big.NewInt(int64(i)).ProbablyPrime(0) {
				result = append(result, i)
				count++
			}
		} else {
			break
		}
		i++
	}

	return result
}

func fibonacci(input int) []int {
	var result []int
	for i := 0; i < input; i++ {
		result = append(result, fib(i))
	}

	return result
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var input int
	var x, y int
	for {
		fmt.Print("\n===== Kitabisa Technical Test =====\n")
		fmt.Println("Menus :")
		menus := []string{"1.Sum", "2.Multiply", "3.Prime Number", "4.Fibonacci", "5.Exit"}
		for _, menu := range menus {
			fmt.Println(" ", menu)
		}
		fmt.Print("Please enter a menu option : ")
		fmt.Scanf("%d", &input)
		switch input {
		case 1:
			fmt.Print("Input x, y : ")
			fmt.Scanf("%d", &x)
			fmt.Scanf("%d", &y)
			result := sum(x, y)
			fmt.Printf("Output %d + %d : %d\n", x, y, result)

		case 2:
			fmt.Print("Input x, y : ")
			fmt.Scanf("%d", &x)
			fmt.Scanf("%d", &y)
			result := multiply(x, y)
			fmt.Printf("Output %d x %d : %d\n", x, y, result)

		case 3:
			var input int
			fmt.Print("Input : ")
			fmt.Scanf("%d", &input)
			result := prime(input)
			fmt.Printf("Output : %v\n", result)

		case 4:
			fmt.Print("Input : ")
			fmt.Scanf("%d", &input)
			result := fibonacci(input)
			fmt.Printf("Output : %v\n", result)

		case 5:
			os.Exit(0)
		default:
			fmt.Print("Oops!, Menu not found!\n")
		}
	}

}
