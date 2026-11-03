package main

import "fmt"

func main() {
	operation := increment
	fmt.Println(apply(3, operation)) //4
	operation = func(s int) int { return s * 2 }
	fmt.Println(apply(3, operation), "\n") //6
	first := makeCounter(0)
	fmt.Println(first()) //1
	fmt.Println(first()) //2
	fmt.Println(first()) //3
	fmt.Println(first()) //4
	fmt.Println(first()) //5
	fmt.Println(first()) //6
	fmt.Println(first()) //7
	fmt.Println(first()) //8
	second := makeCounter(10)
	fmt.Println(second()) //11
	copyFirst := first
	fmt.Println(copyFirst()) //9
	fmt.Println(first)
	fmt.Println(copyFirst)
}

func increment(n int) int {
	return n + 1
}

func apply(s int, f func(int) int) int {
	return f(s)
}

func makeCounter(start int) func() int {
	count := start

	return func() int {
		count += 1
		return count
	}
}
