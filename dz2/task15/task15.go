package main

import (
	"fmt"
	"strings"
)

//15. Великое Лайнландское переселение
//Ограничение времени	1 секунда
//Ограничение памяти	256Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//Лайнландия представляет из себя одномерный мир, являющийся прямой, на котором располагаются N городов, последовательно пронумерованных от 0 до N - 1 . Направление в сторону от первого города к нулевому названо западным, а в обратную — восточным.
//Когда в Лайнландии неожиданно начался кризис, все были жители мира стали испытывать глубокое смятение. По всей Лайнландии стали ходить слухи, что на востоке живётся лучше, чем на западе.
//Так и началось Великое Лайнландское переселение. Обитатели мира целыми городами отправились на восток, покинув родные улицы, и двигались до тех пор, пока не приходили в город, в котором средняя цена проживания была меньше, чем в родном.
//
//Формат ввода
//В первой строке дано одно число N ( 2 ≤ N ≤ 10^5 ) — количество городов в Лайнландии. Во второй строке дано N чисел ai
//(0 ≤ ai ≤ 10^9) — средняя цена проживания в городах с нулевого по (N - 1)-ый соответственно.
//
//Формат вывода
//Для каждого города в порядке с нулевого по (N - 1)-ый выведите номер города, в который переселятся его изначальные жители. Если жители города не остановятся в каком-либо другом городе, отправившись в Восточное Бесконечное Ничто, выведите -1 .

type StackElement struct {
	index       int
	val         int
	prevElement *StackElement
}

func main() {
	var n int
	fmt.Scan(&n)

	slc := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slc[i])
	}

	fmt.Println(calculate(slc))
}

func calculate(slc []int) (res string) {
	strSlc := make([]string, len(slc))
	stack := &StackElement{-1, -1, nil}

	for i, v := range slc {
		for stack.val > v {
			strSlc[stack.index] = fmt.Sprint(i)
			stack = stack.pop()
		}
		stack = stack.push(i, v)
	}

	for stack.index != -1 {
		strSlc[stack.index] = "-1"
		stack = stack.pop()
	}
	res = strings.Join(strSlc, " ")
	return
}

func (element *StackElement) push(i, v int) (newElement *StackElement) {
	newElement = &StackElement{
		index:       i,
		val:         v,
		prevElement: element,
	}
	return
}

func (element *StackElement) pop() (newElement *StackElement) {
	newElement = element.prevElement
	return
}
