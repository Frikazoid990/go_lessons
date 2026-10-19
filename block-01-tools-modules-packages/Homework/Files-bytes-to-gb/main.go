package main

import "fmt"

func main() {
	sizeFilesInBytes := 4_831_838_208

	var sizeFilesInGigabytes float64 = float64(sizeFilesInBytes) / 1024.0 / 1024.0 / 1024.0

	fmt.Printf("%.2f", sizeFilesInGigabytes)
}
