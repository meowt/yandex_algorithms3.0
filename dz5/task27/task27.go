package main

import (
	"fmt"
	"math"
)

//27. Вывести маршрут максимальной стоимости
//Ограничение времени	1 секунда
//Ограничение памяти	256Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//В левом верхнем углу прямоугольной таблицы размером N × M
// находится черепашка. В каждой клетке таблицы записано некоторое число. Черепашка может перемещаться вправо или вниз, при этом маршрут черепашки заканчивается в правом нижнем углу таблицы.
//Подсчитаем сумму чисел, записанных в клетках, через которую проползла черепашка (включая начальную и конечную клетку). Найдите наибольшее возможное значение этой суммы и маршрут, на котором достигается эта сумма.
//
//Формат ввода
//В первой строке входных данных записаны два натуральных числа N и M, не превосходящих 100 — размеры таблицы. Далее идет N строк, каждая из которых содержит M чисел, разделенных пробелами — описание таблицы. Все числа в клетках таблицы целые и могут принимать значения от 0 до 100.
//Формат вывода
//Первая строка выходных данных содержит максимальную возможную сумму, вторая — маршрут, на котором достигается эта сумма. Маршрут выводится в виде последовательности, которая должна содержать N-1 букву D, означающую передвижение вниз и M-1 букву R, означающую передвижение направо. Если таких последовательностей несколько, необходимо вывести ровно одну (любую) из них.

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	matrix := makeMatrix(n, m)
	prev := makePrev(n, m)
	fmt.Println(f(matrix, prev))
	printPrev(prev)
}

func f(matrix [][]int, prev [][]string) int {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if i == 1 && j == 1 {
				continue
			}
			maximum := max(matrix[i-1][j], matrix[i][j-1])
			matrix[i][j] += maximum
			if maximum == matrix[i-1][j] {
				prev[i][j] = "D"
			} else {
				prev[i][j] = "R"
			}
		}
	}
	return matrix[len(matrix)-1][len(matrix[0])-1]
}

func max(i, j int) int {
	return int(math.Max(float64(i), float64(j)))
}

func makeMatrix(n, m int) [][]int {
	n++
	m++
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		matrix[0][i] = -1
	}
	for i := 1; i < n; i++ {
		matrix[i][0] = -1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	return matrix
}

func printPrev(prev [][]string) {
	i, j := len(prev)-1, len(prev[0])-1
	for prev[i][j] != "" {
		switch prev[i][j] {
		case "D":
			defer fmt.Print("D ")
			i--
		case "R":
			defer fmt.Print("R ")
			j--
		}
	}
}

func makePrev(n, m int) [][]string {
	n++
	m++
	matrix := make([][]string, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]string, m)
	}
	return matrix
}
