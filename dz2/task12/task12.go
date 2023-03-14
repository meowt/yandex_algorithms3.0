package main

import (
	"fmt"
	"strings"
)

//12. Правильная скобочная последовательность
//Ограничение времени	1 секунда
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//Рассмотрим последовательность, состоящую из круглых, квадратных и фигурных скобок. Программа дожна определить, является ли данная скобочная последовательность правильной. Пустая последовательность явлется правильной. Если A – правильная, то последовательности (A), [A], {A} – правильные. Если A и B – правильные последовательности, то последовательность AB – правильная.
//
//Формат ввода
//В единственной строке записана скобочная последовательность, содержащая не более 100000 скобок.
//
//Формат вывода
//Если данная последовательность правильная, то программа должна вывести строку yes, иначе строку no.

type StackElement struct {
	val         string
	prevElement *StackElement
}

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(calculate(text))
}

func calculate(text string) (res string) {
	textSlc := strings.Split(text, "")

	stack := &StackElement{}
	balance := 0

	for _, bracket := range textSlc {
		if bracket == "(" || bracket == "[" || bracket == "{" {
			stack = stack.push(bracket)
			balance++
		}
		if bracket == ")" || bracket == "]" || bracket == "}" {
			var ok bool
			stack, ok = pop(stack, bracket)
			balance--
			if balance < 0 || !ok {
				res = "no"
				return
			}
		}
	}
	if balance != 0 {
		res = "no"
		return
	}
	res = "yes"
	return
}

func (element *StackElement) push(bracket string) (newElement *StackElement) {
	newElement = &StackElement{
		val:         bracket,
		prevElement: element,
	}
	return
}

func pop(element *StackElement, bracket string) (newElement *StackElement, ok bool) {
	ok = true
	if element == nil {
		ok = false
		return
	}
	switch bracket {
	case ")":
		if element.val != "(" {
			ok = false
			return
		}
	case "]":
		if element.val != "[" {
			ok = false
			return
		}
	case "}":
		if element.val != "{" {
			ok = false
			return
		}
	default:
	}
	if element != nil && element.prevElement != nil {
		newElement = element.prevElement
		return
	}
	return
}
