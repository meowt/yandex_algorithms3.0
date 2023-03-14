package main

import "fmt"

//21. Три единицы подряд
//Ограничение времени	1 секунда
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//По данному числу N определите количество последовательностей из нулей и единиц длины N, в которых никакие три единицы не стоят рядом.
//
//Формат ввода
//Во входном файле написано натуральное число N, не превосходящее 35.
//
//Формат вывода
//Выведите количество искомых последовательностей. Гарантируется, что ответ не превосходит 2^31-1.

var cache = [100]int{0, 2, 4, 7}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(f(n))
}

func f(n int) (res int) {
	if cache[n] != 0 {
		return cache[n]
	}
	cache[n] = f(n-1) + f(n-2) + f(n-3)
	return cache[n]
}
