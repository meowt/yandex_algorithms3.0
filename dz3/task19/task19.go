package main

import "fmt"

//19. Хипуй
//Ограничение времени	2 секунды
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//В этой задаче вам необходимо самостоятельно (не используя соответствующие классы и функции стандартной библиотеки) организовать структуру данных Heap для хранения целых чисел, над которой определены следующие операции: a) Insert(k) – добавить в Heap число k ; b) Extract достать из Heap наибольшее число (удалив его при этом).
//
//Формат ввода
//В первой строке содержится количество команд N (1 ≤ N ≤ 100000), далее следуют N команд, каждая в своей строке. Команда может иметь формат: “0 <число>” или “1”, обозначающий, соответственно, операции Insert(<число>) и Extract. Гарантируется, что при выполнении команды Extract в структуре находится по крайней мере один элемент.
//
//Формат вывода
//Для каждой команды извлечения необходимо отдельной строкой вывести число, полученное при выполнении команды Extract.

const (
	insert = iota
	extract
)

func main() {
	var heap []int
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var action int
		fmt.Scan(&action)
		switch action {
		case insert:
			var num int
			fmt.Scan(&num)
			heap = append(heap, num)
			currentPos := len(heap) - 1

			for currentPos != 0 {
				parentPos := (currentPos - 1) / 2
				if heap[parentPos] > heap[currentPos] {
					break
				}

				heap[parentPos], heap[currentPos], currentPos = heap[currentPos], heap[parentPos], parentPos
			}

		case extract:
			res := heap[0]
			heap[0] = heap[len(heap)-1]
			pos := 0
			for pos*2+2 < len(heap) {
				max_son_indx := pos*2 + 1
				if heap[pos*2+2] > heap[max_son_indx] {
					max_son_indx = pos*2 + 2
				}
				if heap[pos] < heap[max_son_indx] {
					heap[pos], heap[max_son_indx] = heap[max_son_indx], heap[pos]
					pos = max_son_indx
				} else {
					break
				}
			}
			heap = heap[:len(heap)-1]
			fmt.Println(res)
		}
	}
}
