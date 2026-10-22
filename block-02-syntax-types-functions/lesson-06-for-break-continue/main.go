package main

func main() {
	total := 8
	skipEvery := 3
	maxProcessed := 4
	processed := 0
	idSum := 0

	for i := 1; i <= total; i++ {
		if i%skipEvery == 0 {
			println("Multiple of three: ", i)
			continue
		}

		processed++
		idSum += i
		println("process: ", i)

		if processed >= maxProcessed {
			println("limit reached")
			break
		}
	}

	println("Processed: ", processed)
	println("ID Sum: ", idSum)
	println("\n=======================\n")

	processed = 0
	idSum = 0
	var i int = 1

	for i <= total {
		if i%skipEvery == 0 {
			println("Multiple of three: ", i)
			i++
			continue
		}

		processed++
		idSum += i
		println("process: ", i)
		if processed >= maxProcessed {
			println("limit reached")
			break
		}
		i++
	}

	println("Processed: ", processed)
	println("ID Sum: ", idSum)
	println("batch complete")
}
