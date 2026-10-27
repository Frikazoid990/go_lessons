package main

import "fmt"

func main() {
	var p, s int
	var ok bool

	p, s, ok = runBatch(8, 3, 4)
	printReport("First test", p, s, ok)
	p, s, ok = runBatch(0, 3, 4)
	printReport("Second test", p, s, ok)
	p, s, ok = runBatch(8, 0, 4)
	printReport("Third test", p, s, ok)
	fmt.Println("batch complete")
}

func shouldSkip(id int, skipEvery int) bool {
	if id%skipEvery == 0 {
		return true
	}
	return false
}

func runBatch(total, skipEvery, maxProccessed int) (int, int, bool) {
	processed := 0
	idSum := 0
	if skipEvery <= 0 || total < 0 || maxProccessed <= 0 {
		return 0, 0, false
	} else {
		for i := 1; i <= total; i++ {
			if shouldSkip(i, skipEvery) {
				continue
			}
			processed++
			idSum += i
			if processed >= maxProccessed {
				break
			}
		}
	}
	return processed, idSum, true
}

func printReport(label string, processed, idSum int, ok bool) {
	if !ok {
		fmt.Println(label, "invalid input")
		return
	}
	fmt.Println("Label: ", label)
	fmt.Println("Processed: ", processed)
	fmt.Println("idSum: ", idSum)
}
