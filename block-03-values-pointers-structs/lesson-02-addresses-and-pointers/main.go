package main

import "fmt"

func main() {
	freeSeats := 24
	before := freeSeats
	counter := &freeSeats
	display := counter
	book(counter, 5)
	fmt.Printf("До бронирования: %d\nСвободно сейчас: %d\nНа табло: %d\n", before, freeSeats, *display) // display хранит копию указателя counter, поэтому через него видно обновлённое число свободных мест.
}

func book(target *int, seats int) {
	*target -= seats
}
