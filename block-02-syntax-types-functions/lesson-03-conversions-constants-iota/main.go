package main

import "fmt"

func main() {

	var number1 int = 24
	var number2 int32 = int32(number1)
	fmt.Printf("type numb1:%T value:%d\n", number1, number1)
	fmt.Printf("type numb2:%T value:%d\n", number2, number2)

	number3 := int32(number1)
	number4 := int(number3)
	number5 := int32(40)
	number6 := number4 + int(number5)
	fmt.Printf("type: %T value: %d\n", number6, number6)

	const step = 5
	var number7 int = step
	var number8 int32 = step
	
	fmt.Printf("type numb1:%T value:%d\n", number7, number7)
	fmt.Printf("type numb2:%T value:%d\n", number8, number8)

	const (
		First int = iota
		Second
		Third
	)

	fmt.Println(First, " ", Second, " ", Third)

	const (
		One   int = iota //0
		Two              // 1
		Three            // 2
	)

	fmt.Println(One, " ", Two, " ", Three)

	var n int = 12
	var converted int32 = int32(n)
	fmt.Println(converted)

	/*
		Объяви int со значением 24, явно преобразуй его в int32 и выведи тип и значение обеих переменных.
		Используй знакомые форматы %T и %d.

		Возьми исходное значение int из предыдущего пункта и преобразуй его в int32.
		Преобразуй получившееся значение обратно в int, сохранив результат в отдельную переменную.
		Подготовь второе значение — 40 — другого числового типа.
		Сложи первое и второе значения. Перед сложением приведи одно из них к типу другого, чтобы оба операнда имели одинаковый тип.
		Выведи тип и значение результата. До запуска запиши, каким ожидаешь результат.

		Объяви нетипизированную константу const step = 5.
		Используй её для переменных типов int и int32. Выведи их значения.

		Создай именованную группу из трёх типизированных
		констант int с iota, например уровни приоритета.
		Выведи каждое значение.
		Добавь вторую группу с iota. Запиши прогноз её первого
		значения до запуска и проверь, что счётчик начинается заново в новой группе.
	*/
}
