package main

import "fmt"

// Request limit preview.
func main() {
	limit := 60
	snapshot := limit
	limit = 45
	fmt.Println("current limit:", limit)
	fmt.Println("saved limit:", snapshot)
	reserve(limit, 12)
	fmt.Println("after ignored call — limit:", limit)
	fmt.Println("after ignored call — snapshot:", snapshot)
	proposed := reserve(limit, 12)

	fmt.Println("after preview — limit:", limit)
	fmt.Println("after preview — snapshot:", snapshot)
	fmt.Println("after preview — proposed:", proposed)

	limit = proposed
	proposed++
	fmt.Println("after increment — proposed:", proposed)
	fmt.Println("after assignment — limit:", limit)

}
func reserve(limit int, amount int) int {
	limit -= amount
	fmt.Println("inside:", limit)
	return limit
}
