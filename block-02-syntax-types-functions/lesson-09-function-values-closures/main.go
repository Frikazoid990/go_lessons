package main

func main() {
	first := makecounter(1)
	println(first(5))
	println(first(5))
}

func makecounter(s int) func(int) int {
	t := s
	return func(i int) int {
		t += i
		return t
	}
}
