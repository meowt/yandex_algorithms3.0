package main

import "fmt"

//20. Пирамидальная сортировка
//Ограничение времени	2 секунды
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//Отсортируйте данный массив. Используйте пирамидальную сортировку.
//
//Формат ввода
//Первая строка входных данных содержит количество элементов в массиве N, N ≤ 105. Далее задаются N целых чисел, не превосходящих по абсолютной величине 109.
//
//Формат вывода
//Выведите эти числа в порядке неубывания.

func main() {
	var n int
	var s1 []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		s1 = append(s1, temp)
	}

	i := 0
	tmp := 0

	for i = len(s1)/2 - 1; i >= 0; i-- {
		s1 = heapSort(s1, i, len(s1))
	}

	for i = len(s1) - 1; i >= 1; i-- {
		tmp = s1[0]
		s1[0] = s1[i]
		s1[i] = tmp
		s1 = heapSort(s1, 0, i)
	}
	for _, v := range s1 {
		fmt.Print(v, " ")
	}
}

func heapSort(s1 []int, i int, s1Len int) []int {
	done := false

	tmp := 0
	maxChild := 0

	for (i*2+1 < s1Len) && (!done) {
		if i*2+1 == s1Len-1 || s1[i*2+1] > s1[i*2+2] {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}

		if s1[i] < s1[maxChild] {
			tmp = s1[i]
			s1[i] = s1[maxChild]
			s1[maxChild] = tmp
			i = maxChild
		} else {
			done = true
		}
	}

	return s1
}

/*	heapify(slc, 0)
	//heapSort(slc)
	for _, v := range slc {
		fmt.Print(v, " ")
	}
}

func heapSort(slc []int) {
	delta := 1
	maxI := (len(slc) - 2) / 2
	heapify(slc, maxI)
	fmt.Println(slc)
	for i := 0; i < len(slc)-1; i++ {
		slc[0], slc[len(slc)-delta] = slc[len(slc)-delta], slc[0]
		maxI = (maxI - 2) / 2
		heapify(slc[:len(slc)-delta], maxI)
		fmt.Println(slc)
		delta++
	}
}

func heapify(slc []int, i int) {
	if len(slc) == 1 {
		return
	}
	if slc[i] < slc[i*2+1] {
		slc[i], slc[i*2+1] = slc[i*2+1], slc[i]
		heapify(slc, i*2+1)
		return
	}
	if i*2+2 >= len(slc) {
		return
	}
	if slc[i] < slc[i*2+2] {
		slc[i], slc[i*2+2] = slc[i*2+2], slc[i]
		heapify(slc, i*2+2)
		return
	}
}
*/
