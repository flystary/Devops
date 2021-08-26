package main

import (
	"fmt"
	"time"
)

func sequence() {
	const n = 40
	starttime := time.Now()
	fibN := fib(n)
	endtime := time.Now()
	costTime := endtime.Sub(starttime)
	fmt.Println(costTime)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func multiplication() {
	start := time.Now()
	var i, j int
	for i = 1; i <= 9; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d \t", i, j, i*j)
		}
		fmt.Println()
	}
	//end := time.Now()
	tc := time.Since(start)
	//tc := end.Sub(start)
	fmt.Printf("耗时是%v", tc)
}

func getMaximumCommonDivisor(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else if a < b {
			b = b - a
		}
	}
	return a
}

func isPalindrome(a string) bool {
	var i, j int
	var b bool
	for i = 0; i < len(a)/2-1; i++ {
		j = len(a) - 1
		if a[i] != a[j] {
			b = false
		}
		b = true
		j--
	}
	return b
}
func isDaffodilNum(num int) bool {
	a := num / 100
	b := (num / 10) % 10
	c := num % 10
	result := a*a*a + b*b*b + c*c*c
	if num == result {
		return true
	}
	return false
}
