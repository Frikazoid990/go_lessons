package main

import "fmt"

func main() {
	var original int
	fmt.Printf("type: %T, value: %d\n", original, original)

	original = 7
	snapshot := original

	original, fresh := 15, 20

	fmt.Printf("original: %d snapshot: %d fresh: %d", original, snapshot, fresh)

}
