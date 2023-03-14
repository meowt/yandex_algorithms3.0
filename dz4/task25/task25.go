package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

//25. Гвоздики
//Ограничение времени	1 секунда
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//В дощечке в один ряд вбиты гвоздики. Любые два гвоздика можно соединить ниточкой. Требуется соединить некоторые пары гвоздиков ниточками так, чтобы к каждому гвоздику была привязана хотя бы одна ниточка, а суммарная длина всех ниточек была минимальна.
//
//Формат ввода
//В первой строке входных данных записано число N — количество гвоздиков (2 ≤ N ≤ 100). В следующей строке заданы N чисел — координаты всех гвоздиков (неотрицательные целые числа, не превосходящие 10000).
//
//Формат вывода
//Выведите единственное число — минимальную суммарную длину всех ниточек.

var dp = [100]int{}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)

	if len(arr) == 2 {
		fmt.Println(arr[1] - arr[0])
		os.Exit(0)
	}
	fmt.Println(f(arr))
}

func f(arr []int) int {
	dp[1] = arr[1] - arr[0]
	dp[2] = arr[2] - arr[0]
	for i := 3; i < len(arr); i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + arr[i] - arr[i-1]
	}
	return dp[len(arr)-1]
}

func min(a, b int) (res int) {
	return int(math.Min(float64(a), float64(b)))
}