package main

import "fmt"

func main() {
	plan := "free"
	current := 2
	incoming := 1

	var limit int8
	switch plan {
	case "free":
		limit = 3
	case "pro":
		limit = 10
	case "team":
		limit = 25
	default:
		limit = 0
	}

	if limit == 0 {
		fmt.Println("Неизвестный тариф")
		fmt.Println("check complete")
		return
	}
	if current < 0 || incoming <= 0 {
		fmt.Println("Неверные числа")
		fmt.Println("check complete")
		return
	}

	next := uint8(current) + uint8(incoming)
	fmt.Println("Plan:", plan, "Limit:", limit, "Next:", next)

	switch {
	case next < uint8(limit):
		fmt.Println("Можно добавить, место ещё останется")
	case next == uint8(limit):
		fmt.Println("Можно добавить, лимит будет достигнут")
	default:
		fmt.Println("Невозможно добавить, лимит превышен")
	}

	fmt.Println("check complete")
}
