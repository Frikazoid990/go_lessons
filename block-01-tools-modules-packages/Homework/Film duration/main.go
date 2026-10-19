package main

import "fmt"

func main() {
	durationInSeconds := 8520

	durationHours := durationInSeconds / 3600
	durationMinutes := (durationInSeconds - 3600*durationHours) / 60
	durationSeconds := (durationInSeconds - 3600*durationHours) % 60

	fmt.Printf("%d ч %d мин %d сек", durationHours, durationMinutes, durationSeconds)
}
