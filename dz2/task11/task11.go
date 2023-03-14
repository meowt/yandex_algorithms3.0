package main

import (
	"fmt"
	"os"
)

//11. Стек с защитой от ошибок
//Ограничение времени	1 секунда
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//Научитесь пользоваться стандартной структурой данных stack для целых чисел. Напишите программу, содержащую описание стека и моделирующую работу стека, реализовав все указанные здесь методы. Программа считывает последовательность команд и в зависимости от команды выполняет ту или иную операцию. После выполнения каждой команды программа должна вывести одну строчку. Возможные команды для программы:
//
//push n
//Добавить в стек число n (значение n задается после команды). Программа должна вывести ok.
//
//pop
//Удалить из стека последний элемент. Программа должна вывести его значение.
//
//back
//Программа должна вывести значение последнего элемента, не удаляя его из стека.
//
//size
//Программа должна вывести количество элементов в стеке.
//
//clear
//Программа должна очистить стек и вывести ok.
//
//exit
//Программа должна вывести bye и завершить работу.
//
//Перед исполнением операций back и pop программа должна проверять, содержится ли в стеке хотя бы один элемент. Если во входных данных встречается операция back или pop, и при этом стек пуст, то программа должна вместо числового значения вывести строку error.
//
//Формат ввода
//Вводятся команды управления стеком, по одной на строке
//
//Формат вывода
//Программа должна вывести протокол работы стека, по одному сообщению на строке

type StackElement struct {
	val         int
	prevElement *StackElement
}

func main() {
	var command string
	stack := &StackElement{}
	stack = nil
	for {
		fmt.Scan(&command)

		switch command {
		case "push":
			var n int
			fmt.Scan(&n)
			stack = stack.push(n)

		case "pop":
			stack = stack.pop()

		case "back":
			stack.back()

		case "size":
			fmt.Println(stack.size())

		case "clear":
			stack.clear()

		case "exit":
			fmt.Println("bye")
			os.Exit(0)
		}
	}
}

func (element *StackElement) push(n int) (newElement *StackElement) {
	newElement = &StackElement{
		val:         n,
		prevElement: element,
	}
	fmt.Println("ok")
	return
}

func (element *StackElement) pop() (newElement *StackElement) {
	if element != nil {
		fmt.Println(element.val)
		if element.prevElement != nil {
			newElement = element.prevElement
			return
		}
		return
	}
	fmt.Println("error")
	return nil
}

func (element *StackElement) back() (topVal int) {
	if element != nil {
		fmt.Println(element.val)
		return
	}
	fmt.Println("error")
	return
}

func (element *StackElement) size() (resSize int) {
	for element != nil {
		resSize++
		if element.prevElement == nil {
			break
		}
		element = element.prevElement
	}
	return
}

func (element *StackElement) clear() {
	element = &StackElement{}
	element = nil
}
