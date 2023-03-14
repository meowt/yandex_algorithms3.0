package main

import (
	"fmt"
	"os"
)

//16. Очередь с защитой от ошибок
//Ограничение времени	1 секунда
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//Научитесь пользоваться стандартной структурой данных queue для целых чисел. Напишите программу, содержащую описание очереди и моделирующую работу очереди, реализовав все указанные здесь методы.
//
//Программа считывает последовательность команд и в зависимости от команды выполняет ту или иную операцию. После выполнения каждой команды программа должна вывести одну строчку.
//
//Возможные команды для программы:
//
//push n
//Добавить в очередь число n (значение n задается после команды). Программа должна вывести ok.
//
//pop
//Удалить из очереди первый элемент. Программа должна вывести его значение.
//
//front
//Программа должна вывести значение первого элемента, не удаляя его из очереди.
//
//size
//Программа должна вывести количество элементов в очереди.
//
//clear
//Программа должна очистить очередь и вывести ok.
//
//exit
//Программа должна вывести bye и завершить работу.
//
//Перед исполнением операций front и pop программа должна проверять, содержится ли в очереди хотя бы один элемент. Если во входных данных встречается операция front или pop, и при этом очередь пуста, то программа должна вместо числового значения вывести строку error.
//
//Формат ввода
//Вводятся команды управления очередью, по одной на строке
//
//Формат вывода
//Требуется вывести протокол работы очереди, по одному сообщению на строке

type Queue struct {
	val  int
	next *Queue
}

func main() {
	var command string
	queue := &Queue{}
	queue = nil

	for {
		fmt.Scan(&command)

		switch command {
		case "push":
			var n int
			fmt.Scan(&n)
			queue = queue.push(n)

		case "pop":
			queue = queue.pop()

		case "front":
			queue.front()

		case "size":
			fmt.Println(queue.size())

		case "clear":
			queue = queue.clear()

		case "exit":
			fmt.Println("bye")
			os.Exit(0)
		}
	}
}

func (elem *Queue) push(n int) (newElement *Queue) {
	newElement = &Queue{
		val:  n,
		next: elem,
	}
	fmt.Println("ok")
	return
}

func (elem *Queue) pop() (newElem *Queue) {
	if elem != nil {
		newElement := elem
		head := elem.getHead()
		if newElement == head {
			fmt.Println(newElement.val)
			return nil
		}
		for newElement.next != head {
			newElement = newElement.next
		}
		newElement.next = nil
		fmt.Println(head.val)
		return elem
	}
	fmt.Println("error")
	return nil
}

func (elem *Queue) getHead() (head *Queue) {
	head = elem
	for head.next != nil {
		head = head.next
	}
	return head
}

func (elem *Queue) front() (topVal int) {
	if elem != nil {
		head := elem.getHead()
		fmt.Println(head.val)
		return
	}
	fmt.Println("error")
	return
}

func (elem *Queue) size() (resSize int) {
	for elem != nil {
		resSize++
		if elem.next == nil {
			break
		}
		elem = elem.next
	}
	return
}

func (elem *Queue) clear() *Queue {
	fmt.Println("ok")
	return nil
}
