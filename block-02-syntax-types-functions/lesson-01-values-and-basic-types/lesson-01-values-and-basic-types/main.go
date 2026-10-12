package main

import "fmt"

func main() {
	//example
	fmt.Printf("type = %T, value = %v\nNew line\n", 6.25, 6.25)
	//Task 1
	fmt.Println("======================")
	fmt.Println("Task 1")
	fmt.Printf("Int - Type: %T | Value: %d\n", 24, 24)
	fmt.Printf("Float64 - Type: %T | Value: %f\n", 24.0, 24.0)
	fmt.Printf("Float64 - Type: %T | Value: %.2f\n", 3.75, 3.75)
	fmt.Printf("BooleanTrue - Type: %T | Value: %t\n", true, true)
	fmt.Printf("BooleanFalse - Type: %T | Value: %t\n", false, false)
	fmt.Printf("String24 - Type: %T | Value: %v\n", "24", "24")
	fmt.Printf("StringFalse - Type: %T | Value: %q\n", "false", "false")
	fmt.Printf("String - Type: %T | ValueWithSKey: %s ValueWithQKey: %q\n", "Go", "Go", "Go")
	fmt.Printf("StringFalse - Type: %T | ValueWithQKey: %q ValueWiyhoutKey: %s \n", "", "", "")
	fmt.Println("======================")
	//Task 2
	fmt.Println("\n======================")
	fmt.Println("TASK 2")
	fmt.Print("Learning ")
	fmt.Println("Go")

	fmt.Print("Score: ", 15)
	fmt.Println("\nScore:", 15)

	fmt.Printf("\nUser: %s; attempts: %d; active: %t\n", "Nora", 4, true)
	fmt.Printf("2 points after . : %.2f%%",12.375)
	var str string = fmt.Sprintf("\nNora %v", 4)
	fmt.Println(str)

	/*
		24
		24.0
		3.75
		true
		false
		"24"
		"false"
		"Go"
		""
	*/
}
