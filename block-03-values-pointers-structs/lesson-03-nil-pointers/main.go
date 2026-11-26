package main

import "fmt"

func main() {
	stock := 0
	var counter *int
	res1 := recordDelivery(counter, 5)
	fmt.Printf("Без счётчика: %t; На складе: %d\n", res1, stock)
	counter = &stock
	observer := counter
	res2 := recordDelivery(counter, 5)
	fmt.Printf("С подключённым счётчиком: %t; На складе: %d\n", res2, stock)
	counter = nil
	res3 := recordDelivery(counter, 3)
	fmt.Printf("После отключения: %t; На складе: %d; Через наблюдателя: %d\n", res3, stock, *observer)
}

func recordDelivery(target *int, amount int) bool {
	if target == nil {
		return false
	} else {
		*target += amount
		return true
	}
}
