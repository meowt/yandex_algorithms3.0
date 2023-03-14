package main

import "fmt"

//22. Кузнечик
//Ограничение времени	1 секунда
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//У одного из студентов в комнате живёт кузнечик, который очень любит прыгать по клетчатой одномерной доске. Длина доски — N клеток. К его сожалению, он умеет прыгать только на 1, 2, …, k клеток вперёд.
//
//Однажды студентам стало интересно, сколькими способами кузнечик может допрыгать из первой клетки до последней. Помогите им ответить на этот вопрос.
//
//Формат ввода
//В первой и единственной строке входного файла записано два целых числа — N и k (1 ≤ N ≤ 30, 1 ≤ N ≤ 10).
//
//Формат вывода
//Выведите одно число — количество способов, которыми кузнечик может допрыгать из первой клетки до последней.

var cache = [100]int{0, 1, 1}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(f(n, k))
}

func f(n, k int) (res int) {
	if cache[n] != 0 {
		return cache[n]
	}

	for i := 1; i <= k; i++ {
		if n-i >= 0 {
			cache[n] += f(n-i, k)
		}
	}
	return cache[n]
}