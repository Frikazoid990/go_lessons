package main

func main() {
	res1 := traceInt("base", 20) + traceInt("count", 3)*traceInt("unit", 4)     // base count uint
	println("result: ", res1)                                                   // 32
	res2 := (traceInt("base", 20) + traceInt("count", 3)) * traceInt("unit", 4) // base count uint
	println("result: ", res2)                                                   //23 * 4 = 92

	var count int

	for i := 1; i <= 3; i++ {
		ready := traceBool("primary", i == 3) || // ready будет 2 раза true при 2 и 3
			traceBool("backup", i == 2)
		allowed := traceBool("ready", ready) && // соотвественно allowed 1 раз true при i == 3
			traceBool("quota", i == 3)

		if allowed { //скипнется при i == 1 когда все будут false
			println("processed")
			count++
		} else {
			println("skipped")
		}

	}
	println("final count: ", count) // 1 processed attempt

	revision := 10
	revision1 := revision
	advance := func() int { // т.к. тут происходит замыкание
		revision++
		return revision
	}
	total := revision1 + advance() // revision1 remains 10; advance changes revision to 11 and returns 11

	println("total: ", total)
	println("revision: ", revision)
}

func traceInt(label string, value int) int {
	println("label: ", label)
	return value
}

func traceBool(label string, value bool) bool {
	println("label: ", label)
	return value
}
