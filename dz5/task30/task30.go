package main

import "fmt"

//30. НОП с восстановлением ответа
//Ограничение времени	1 секунда
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//Даны две последовательности, требуется найти и вывести их наибольшую общую подпоследовательность.
//
//Формат ввода
//В первой строке входных данных содержится число N – длина первой последовательности (1 ≤ N ≤ 1000). Во второй строке заданы члены первой последовательности (через пробел) – целые числа, не превосходящие 10000 по модулю.
//
//В третьей строке записано число M – длина второй последовательности (1 ≤ M ≤ 1000). В четвертой строке задаются члены второй последовательности (через пробел) – целые числа, не превосходящие 10000 по модулю.
//
//Формат вывода
//Требуется вывести наибольшую общую подпоследовательность данных последовательностей, через пробел.

func main() {
	var n, m int
	fmt.Scan(&n)

	firstWord := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&firstWord[i])
	}

	fmt.Scan(&m)
	secondWord := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Scan(&secondWord[i])
	}

	dp := makeMatrix(n, m)

	for i := 1; i < len(firstWord); i++ {
		for j := 1; j < len(secondWord); j++ {
			if firstWord[i] == secondWord[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				continue
			}
			if dp[i][j-1] > dp[i-1][j] {
				dp[i][j] = dp[i][j-1]
				continue
			}
			dp[i][j] = dp[i-1][j]
		}
	}
	findWayBack(dp, firstWord)
}

func findWayBack(dp [][]int, firstWord []int) {
	i, j := len(dp)-1, len(dp[0])-1
	for dp[i][j] != 0 {
		if dp[i-1][j] == dp[i][j]-1 && dp[i][j-1] == dp[i][j]-1 {
			defer fmt.Print(firstWord[i], " ")
			i--
			j--
			continue
		}
		for dp[i-1][j] != dp[i][j]-1 {
			i--
		}
		for dp[i][j-1] != dp[i][j]-1 {
			j--
		}
	}
}

func makeMatrix(n, m int) [][]int {
	n++
	m++
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	return matrix
}
