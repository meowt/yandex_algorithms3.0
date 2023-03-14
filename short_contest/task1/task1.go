package main

import (
	"fmt"
	"strconv"
	"strings"
)

//A. Подземная доставка
//Ограничение времени	2 секунды
//Ограничение памяти	256Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//Для ускорения работы служб доставки под городом Длинноградом был прорыт тоннель, по которому ходит товарный поезд, останавливающийся на промежуточных станциях возле логистических центров. На станциях к концу поезда могут быть присоединены вагоны с определенными товарами, а также от его конца может быть отцеплено некоторое количество вагонов или может быть проведена ревизия, во время которой подсчитывается количество вагонов с определенным товаром.
//Обработайте операции в том порядке, в котором они производились, и ответьте на запросы ревизии.
//
//Формат ввода
//В первой строке вводится число N (1 ≤ N ≤ 100000) — количество операций, произведенных над поездом.
//В каждой из следующих N строк содержится описание операций. Каждая операция может иметь один из трех типов:
//add <количество вагонов> <название товара> —
//добавить в конец поезда <количество вагонов> с грузом <название товара>. Количество вагонов не может превышать 10^9, название товара — одна строка из строчных латинских символов длиной до 20.
//delete <количество вагонов> —
//отцепить от конца поезда <количество вагонов>. Количество отцепляемых вагонов не превосходит длины поезда.
//get <название товара> —
//определить количество вагонов с товаром <название товара> в поезде. Название товара — одна строка из строчных латинских символов длиной до 20.
//
//Формат вывода
//На каждый запрос о количестве вагонов с определенным товаром выведите одно число — количество вагонов с таким товаром. Запросы надо обрабатывать в том порядке, как они поступали.

type car struct {
	value   int
	product string
	prev    *car
}

func main() {
	var n int
	var res []string
	fmt.Scan(&n)
	tail := &car{}
	tail = nil
	for i := 0; i < n; i++ {
		var operation string
		fmt.Scan(&operation)
		switch operation {
		case "add":
			var (
				amount int
				prod   string
			)
			fmt.Scan(&amount, &prod)
			tail = tail.add(amount, prod)

		case "delete":
			var (
				amount int
			)
			fmt.Scan(&amount)
			tail = tail.delete(amount)

		case "get":
			var (
				prod string
			)
			fmt.Scan(&prod)
			res = append(res, tail.get(prod))
		}
	}
	printResult(res)
}

func (element *car) add(val int, prod string) (newElement *car) {
	newElement = &car{
		value:   val,
		product: prod,
		prev:    element,
	}
	return
}

func (element *car) delete(amount int) (newElement *car) {
	newElement = element
	for newElement != nil {
		amount, newElement.value = amount-newElement.value, newElement.value-amount
		if amount > 0 {
			newElement = newElement.prev
		} else {

			break
		}
	}
	return
}

func (element *car) get(prod string) string {
	var res int
	tempElement := element
	for tempElement != nil {
		if tempElement.product == prod {
			res += tempElement.value
		}
		tempElement = tempElement.prev
	}
	return strconv.Itoa(res)
}

func printResult(res []string) {
	fmt.Println(strings.Join(res, "\n"))
}
