package main

import (
	"fmt"
	"os"
)

//17. Игра в пьяницу
//Ограничение времени	1 секунда
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//В игре в пьяницу карточная колода раздается поровну двум игрокам. Далее они вскрывают по одной верхней карте, и тот, чья карта старше, забирает себе обе вскрытые карты, которые кладутся под низ его колоды. Тот, кто остается без карт – проигрывает. Для простоты будем считать, что все карты различны по номиналу, а также, что самая младшая карта побеждает самую старшую карту ("шестерка берет туза"). Игрок, который забирает себе карты, сначала кладет под низ своей колоды карту первого игрока, затем карту второго игрока (то есть карта второго игрока оказывается внизу колоды). Напишите программу, которая моделирует игру в пьяницу и определяет, кто выигрывает. В игре участвует 10 карт, имеющих значения от 0 до 9, большая карта побеждает меньшую, карта со значением 0 побеждает карту 9.
//
//Формат ввода
//Программа получает на вход две строки: первая строка содержит 5 чисел, разделенных пробелами — номера карт первого игрока, вторая – аналогично 5 карт второго игрока. Карты перечислены сверху вниз, то есть каждая строка начинается с той карты, которая будет открыта первой.
//
//Формат вывода
//Программа должна определить, кто выигрывает при данной раздаче, и вывести слово first или second, после чего вывести количество ходов, сделанных до выигрыша. Если на протяжении 106 ходов игра не заканчивается, программа должна вывести слово botva.

type Queue struct {
	val  int
	next *Queue
}

func main() {
	var ok bool
	first := &Queue{}
	second := &Queue{}
	first, second = nil, nil

	for i := 0; i < 5; i++ {
		var temp int
		fmt.Scan(&temp)
		first = first.push(temp)
	}
	for i := 0; i < 5; i++ {
		var temp int
		fmt.Scan(&temp)
		second = second.push(temp)
	}

	for i := 0; i < 1000000; i++ {
		first, second, ok = compare(first, second)
		if ok {
			fmt.Print(i + 1)
			os.Exit(0)
		}
	}
	fmt.Println("botva")
}

func compare(first, second *Queue) (newFirst, newSecond *Queue, res bool) {
	newFirst, newSecond = first, second
	f, s := newFirst.front(), newSecond.front()
	newFirst = newFirst.pop()
	newSecond = newSecond.pop()

	if f == 9 && s == 0 {
		newSecond = newSecond.push(f)
		newSecond = newSecond.push(s)
		if newFirst.size() == 0 {
			fmt.Print("second ")
			res = true
			return
		}
		return
	}
	if s == 9 && f == 0 {
		newFirst = newFirst.push(f)
		newFirst = newFirst.push(s)
		if newSecond.size() == 0 {
			fmt.Print("first ")
			res = true
			return
		}
		return
	}
	if f > s {
		newFirst = newFirst.push(f)
		newFirst = newFirst.push(s)
		if newSecond.size() == 0 {
			fmt.Print("first ")
			res = true
			return
		}
		return
	}
	if s > f {
		newSecond = newSecond.push(f)
		newSecond = newSecond.push(s)
		if newFirst.size() == 0 {
			fmt.Print("second ")
			res = true
			return
		}
		return
	}
	return
}

func (elem *Queue) push(n int) (newElement *Queue) {
	newElement = &Queue{
		val:  n,
		next: elem,
	}
	return
}

func (elem *Queue) pop() (newElem *Queue) {
	if elem != nil {
		newElement := elem
		head := elem.getHead()
		if newElement == head {
			return nil
		}
		for newElement.next != head {
			newElement = newElement.next
		}
		newElement.next = nil
		return elem
	}
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
		topVal = head.val
		return
	}
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
