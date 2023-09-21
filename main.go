package main

import "fmt"

//if28//func main() {
//	var year uint
//	fmt.Println("Введите год: ")
//	fmt.Scan(&year)
//	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
//		fmt.Println("Этот год является  Високосным!")
//	} else {
//		fmt.Println("Этот год является обычним!")
//	}
//
//}

// if29
//func main() {
//	var number int
//	fmt.Println("Введите число: ")
//	fmt.Scan(&number)
//	if number > 0 && number%2 == 0 {
//		fmt.Println("Положительное четное-число!")
//	} else if number < 0 && number%2 == 0 {
//		fmt.Println("Отрицательное четное-число!")
//	} else if number < 0 && number%2 != 0 {
//		fmt.Println("Отрицательное нечетное-число!")
//	} else if number > 0 && number%2 != 0 {
//		fmt.Println("Положитеоное  нечетное-число!")
//	} else {
//		fmt.Println("Нулевое чило!")
//	}
//}

// if30
//func main() {
//	var number int
//	fmt.Println("Введите число от 1 до 999: ")
//	fmt.Scan(&number)
//	if number > 0 && number < 10 && number%2 == 0 {
//		fmt.Println("четное однозначное число!")
//	} else if number > 0 && number < 10 && number%2 != 0 {
//		fmt.Println("нечетное однозначное число!")
//	} else if number > 9 && number < 100 && number%2 != 0 {
//		fmt.Println("нечетное двухзначное число!")
//	} else if number > 9 && number < 100 && number%2 == 0 {
//		fmt.Println("четное двухзначное число!")
//	} else if number > 99 && number < 1000 && number%2 != 0 {
//		fmt.Println("нечетное трехзначное число!")
//	} else if number > 99 && number < 1000 && number%2 == 0 {
//		fmt.Println("четное трехзначное число!")
//	} else {
//		fmt.Println("Неверный ввод")
//	}
//}

//for39

//func main() {
//	var first_num int
//	var second_num int
//	fmt.Println("Введите 1-гое число: ")
//	fmt.Scan(&first_num)
//	fmt.Println("Введите 2-гое число: ")
//	fmt.Scan(&second_num)
//	for i := first_num; i <= second_num; i++ {
//		for j := 1; j <= i; j++ {
//			fmt.Println(i)
//		}
//	}
//}

// 40
func main() {
	var first_num int
	var second_num int
	fmt.Println("Введите 1-гое число: ")
	fmt.Scan(&first_num)
	fmt.Println("Введите 2-гое число: ")
	fmt.Scan(&second_num)
	for i := first_num; i <= second_num; i++ {
		for j := first_num; j <= i; j++ {
			fmt.Println(i)
		}
	}
}
