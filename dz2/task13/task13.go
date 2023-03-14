package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//13. Постфиксная запись
//Ограничение времени	1 секунда
//Ограничение памяти	256Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//В постфиксной записи (или обратной польской записи) операция записывается после двух операндов. Например, сумма двух чисел A и B записывается как A B +. Запись B C + D * обозначает привычное нам (B + C) * D, а запись A B C + D * + означает A + (B + C) * D. Достоинство постфиксной записи в том, что она не требует скобок и дополнительных соглашений о приоритете операторов для своего чтения.
//Формат ввода
//В единственной строке записано выражение в постфиксной записи, содержащее цифры и операции +, -, *. Цифры и операции разделяются пробелами. В конце строки может быть произвольное количество пробелов.
//Формат вывода
//Необходимо вывести значение записанного выражения.

type StackElement struct {
	val         int
	prevElement *StackElement
}

func main() {
	dat, _ := os.ReadFile("input.txt")

	line := strings.Trim(string(dat), "\n")

	slc := strings.Split(line, " ")

	fmt.Println(calculate(slc))
}

func calculate(slc []string) (res int) {
	stack := &StackElement{}
	stack = nil
	for _, elem := range slc {
		switch elem {
		case "+":
			first := stack.val
			stack = stack.pop()
			second := stack.val
			stack.val = first + second
		case "-":
			first := stack.val
			stack = stack.pop()
			second := stack.val
			stack.val = second - first
		case "*":
			first := stack.val
			stack = stack.pop()
			second := stack.val
			stack.val = first * second
		default:
			num, _ := strconv.Atoi(elem)
			stack = stack.push(num)
		}
	}
	res = stack.val
	return
}

func (element *StackElement) push(n int) (newElement *StackElement) {
	newElement = &StackElement{
		val:         n,
		prevElement: element,
	}
	return
}

func (element *StackElement) pop() (newElement *StackElement) {
	if element != nil {
		if element.prevElement != nil {
			newElement = element.prevElement
			return
		}
		return
	}
	return nil
}
