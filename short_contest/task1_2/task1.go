package main

import (
	"fmt"
	"strconv"
	"strings"
)

//B. Эффективный менеджмент
//Ограничение времени	2 секунды
//Ограничение памяти	256Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//В компании принята следующая практика: новая задача поручается последнему нанятому сотрудинку. Задача номер i характеризуется временем начала ai и временем Ti, которое требуется для ее выполнения. Сотрудник может работать одновременно только над одной задачей. Закончив выполнять задачу, сотрудник сразу же может приступить к выполнению другой задачи. Если последний нанятый сотрудник не может выполнить задачу из-за пересечения с уже порученными, то компания нанимает нового сотрудника и поручает задачу ему.
//
//Эффективный Менеджер выделил N задач на следующий период продолжительностью W. Задачи могут поручаться в любом порядке. Помогите Эффективному Менеджеру переупорядочить поручение задач таким образом, чтобы количество нанятых сотрудников было минимально.
//
//На рисунке ниже приведен пример выполнения трех задач двумя сотрудниками (рисунок соответствует примеру).
//
//
//Формат ввода
//В первой строке входного файла записаны числа N (1 ≤ N ≤ 100000) и W (1 ≤ W ≤ 109).
//
//В следующих N строках записаны пары чисел ai и Ti (1 ≤ ai ≤ W – Ti + 1).
//
//Формат вывода
//Выведите минимальное количество сотрудников, которые будут наняты компанией.
//
//Затем выведите переупорядоченную последовательность поручения задач, приводящую к найму наименьшего количества сотрудников. Задачи нумеруются натуральными числами от 1 до N в том порядке, в котором они указаны во входных данных. Если возможных вариантов несколько, выведите любой из них.

func main() {
	var n int
	var amountStack []int
	var productStack, res []string
	fmt.Scan(&n)
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
			amountStack = append(amountStack, amount)
			productStack = append(productStack, prod)

		case "delete":
			var (
				amount int
			)
			fmt.Scan(&amount)
			for j := len(amountStack) - 1; amount > 0 || j > 0; j-- {
				if amountStack[j]-amount < 0 {
					amount -= amountStack[j]
					amountStack = amountStack[:len(amountStack)-1]
					productStack = productStack[:len(productStack)-1]
				} else {
					amountStack[j], amount = amountStack[j]-amount, amount-amountStack[j]
				}
			}

		case "get":
			var (
				prod    string
				tempRes int
			)
			fmt.Scan(&prod)
			for j, p := range productStack {
				if p == prod {
					tempRes += amountStack[j]
				}
			}
			res = append(res, strconv.Itoa(tempRes))
		}
	}
	printResult(res)
}

func printResult(res []string) {
	fmt.Println(strings.Join(res, "\n"))
}
