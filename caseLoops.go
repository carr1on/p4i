package main

import "fmt"

func main() {
	fmt.Printf("\nВремена года\n\n")
	fmt.Println("Напишите любой месяц года с маленькой буквы, а я подскажу вам, какой сейчас сезон года!")
	fmt.Println("у любой программы ожидающей пользовательский ввод, должно быть \"стоп-слово\".")
	fmt.Println("сегодня это будет \"молоко\"")

firstLoop:
	for {

		var month string
	innerLoop:
		for {
			fmt.Scan(&month)
			switch {
			case month == "декабрь" || month == "январь" || month == "февраль":
				fmt.Println("Зима")

			case month == "март" || month == "апрель" || month == "май":
				fmt.Println("Весна")

			case month == "июнь" || month == "июль" || month == "август":
				fmt.Println("Лето")

			case month == "сентябрь" || month == "октябрь" || month == "ноябрь":
				fmt.Println("Осень")

			case month == "молоко":
				fmt.Println("сворачиваемся! :)")
				break firstLoop
			default:
				fmt.Println("а это точно название месяца?")
				break innerLoop
			}
		}
	}
}
