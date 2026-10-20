package main

import (
	"flag"
	"fmt"
)

func main() {

	//Division Lab

	posOrNegNumb := flag.Bool("morp", false, "Sign 1 Task")
	a := 19
	flag.Parse()
	if *posOrNegNumb {
		a = -19
	}

	var b int = 4

	fmt.Println("Sum: ", a+b, "Difference: ", a-b, " Product: ", a*b)

	c := float64(a) / float64(b)
	d := float64(a / b)
	fmt.Printf("type: %T value: %.2f\n", c, c)
	fmt.Printf("type: %T value: %.1f\n", d, d)

	//Capacity Check

	current := 9
	capacity := 12

	for i := 2; i <= 4; i++ {
		fmt.Println("Значение I = ", i)
		res := current + i
		if res < capacity {
			fmt.Println("Значение поместилось в размерность")
		} else if res == capacity {
			fmt.Println("будущий размер равен пределу")
		} else {
			fmt.Println("Значение превысело предел")
		}
	}

	// Access Rule

	active := false
	paid := false
	trial := true
	blocked := false

	allowed := active && (paid || trial) && !blocked

	allowed1 := active && paid || trial // в случае если есть скобки, то тогда выполняется || первым, потом уже &&
	// в ином случае у нас выполняется первым && потом || тем самым получаются разные результаты

	fmt.Println(allowed1)

	fmt.Println(allowed)

	//safe divisibility
	total := 24
	divisor := 0
	divisible := divisor != 0 && total%divisor == 0
	fmt.Println(divisible)
}
