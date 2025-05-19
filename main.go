package main

import "fmt"

/**
- (if/else) реализовать функцию, которая принимает целое число и возвращает сообщение "чётное" или "не чётное"
func printFullNumber(a int) string {
	if a%2 == 0 {
		return "Чётное число"
	}
	return "Не чётное число"
}

func main() {
	result := printFullNumber(91)
	fmt.Println(result)
}**/

/**func printFullNumber2(a int) string {
	if a%2 == 0 {
		return "Чётное число"
	} else {
		return "Нечётное число"
	}
}

func main() {
	fmt.Println(printFullNumber2(10))
}**/

//- (switch) написать функцию, которая принимает целое число от 1 до 5 и возвращает текстовое сообщение (one, two, three, four, five) соответственно

/**func fiveDays(day int) {
	switch day {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	}
}

func main() {
	fiveDays(3)
}**/

//- (for) написать функцию, которая принимает целое число 10, в цикле от 1 до 10 печает сообщение "current number is: <number>", где number -- текущее число в цикле. должно быть 10 сообщений

/**func cycle10(a int) {
	var b int
	for b = 1; b <= a; b++ { // b = 1 - начинаем переменную с 1, b <= a - b меньше или равно а, b++ -увеличиваем i на 1 после каждой итерации
		fmt.Println("current number is:", b)
	}
}

func main() {
	cycle10(10) // передаём 10 в функцию
}**/

//- (while) сделать то же самое, используя конструкцию while - невозможно, while отсутствует в Go

/**- (for + array) написать функцию, которая принимает входной массив
myArray := []int{1, 2, 3, 4, 5}
и с помощью цикла for выводить каждое значение в консоль**/

func printArrayElements(arr []int) { //arr []int — функция принимает срез целых чисел.
	for i := 0; i < len(arr); i++ {  //Цикл for i := 0 (инициализация); i < len(arr) (условие продолжения. Цикл продолжается до тех пор, пока это условие истинно (true) len(arr) — возвращает длину массива (или среза), т.е..); i++ (обновление переменной. После каждой итерации (прохода цикла) i увеличивается на 1 (i = i + 1).)проходит по всем элементам.
		fmt.Println(arr[i])
	}
}

func main() {
	myArray := []int{1, 2, 3, 4, 5}
	printArrayElements(myArray) //fmt.Println(arr[i]) — печатает текущее значение.
}

/**- (map) написать функцию, которая принимает map
myMap := map[string]int{
"apple":  5,
"banana": 3,
"cherry": 7,
}
и выводит в консоль все связки ключ + значение**/
