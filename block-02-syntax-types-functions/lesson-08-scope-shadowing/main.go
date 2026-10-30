package main

func main() {

	mode := "second"
	limit := 10
	println("Limit until if: ", limit) // 10
	if mode == "first" {
		limit, extra := 4, 2
		println("Local limit in if:", limit, " and extra in if: ", extra) // 4
	} else {
		limit = 9
	}
	println("Limit after if: ", limit) // если mode == first 10, если что-то другое 9

	var incValue int

	for i := 1; i <= limit; i++ {
		var incValueInsideFor int
		incValueInsideFor += i
		incValue += 1
	}

	// incValueInsideFor доступна только внутри тела цикла и заново объявляется
	// на каждой итерации; в последней итерации она получает значение limit.
	println("Count after for: ", incValue)
	limit, extra := 15, 22
	println("Limit: ", limit, " and extra: ", extra, "after if")
}
