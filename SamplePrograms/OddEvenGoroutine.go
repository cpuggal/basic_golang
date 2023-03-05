// Write a program to print ‘n’ even and odd numbers using goroutines.

package main

import (
	"fmt"
	"time"
)

var number int = 0

func printNumber(n int) {

	for i := 1; i <= n; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Number:", i)
	}
}

func printOddNumber(n int) {

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			time.Sleep(1 * time.Second)
			fmt.Println("Odd:", i)
		}
	}
}

func printEvenNumber(n int) {

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			time.Sleep(1 * time.Second)
			fmt.Println("Even:", i)
		}
	}
}

func main() {
	go printOddNumber(8)
	go printEvenNumber(8)
	printNumber(10)

}
