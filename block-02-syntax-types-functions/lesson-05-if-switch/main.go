package main

func main() {
	var plan string = "free"
	var current, incoming int = 2, 1
	var limit int8

	switch plan {
	case "free":
		limit = 3
	case "pro":
		limit = 10
	case "team":
		limit = 25
	default:
		limit = 0
	}
	var next uint8 = 0

	// TODO: раздели неизвестный тариф и неверные числа, чтобы вывести точную причину отказа.
	if limit != 0 && incoming > 0 && current >= 0 {
		if res := uint8(current) + uint8(incoming); res > 0 {
			next = res
		}
	} else {
		println("Your plan is defined or number not valid")
		println("check complete")
		return
	}

	if next < uint8(limit) {
		println("Можно добавить, место ещё останется ", next)
	} else if next == uint8(limit) {
		println("Можно добавить, лимит будет достигнут ", next)
	} else {
		println("Невозможно добавить, лимит превышен")
	}

	println("check complete")
}
