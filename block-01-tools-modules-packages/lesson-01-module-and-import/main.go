package main

import (
	"fmt"
	. "lesson00/greeting"
)

func main() {
	fmt.Println("Hello, World!")
	h := 2
	h += 1
	fmt.Println(h)
	say(SayHoola())
}
