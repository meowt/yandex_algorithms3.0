package main

import (
	"fmt"
	"sync"
)

//9. Сумма в прямоугольнике
//Ограничение времени	3 секунды
//Ограничение памяти	256Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//Вам необходимо ответить на запросы узнать сумму всех элементов числовой матрицы N×M в прямоугольнике с левым верхним углом (x1, y1) и правым нижним (x2, y2)
//
//Формат ввода
//В первой строке находится числа N, M размеры матрицы (1 ≤ N, M ≤ 1000) и K — количество запросов (1 ≤ K ≤ 100000). Каждая из следующих N строк содержит по M чисел`— элементы соответствующей строки матрицы (по модулю не превосходят 1000). Последующие K строк содержат по 4 целых числа, разделенных пробелом x1 y1 x2 y2 — запрос на сумму элементов матрице в прямоугольнике (1 ≤ x1 ≤ x2 ≤ N, 1 ≤ y1 ≤ y2 ≤ M)
//
//Формат вывода
//Для каждого запроса на отдельной строке выведите его результат — сумму всех чисел в элементов матрице в прямоугольнике (x1, y1), (x2, y2)

func main() {
	var height, width, k int
	fmt.Scan(&height, &width, &k)

	matrix := scanMatrix(height, width)

	//var coordSlc [][]int
	coordSlc := make([][4]int, k)
	for i := 0; i < k; i++ {
		var x1, x2, y1, y2 int
		fmt.Scan(&x1, &y1, &x2, &y2)
		tempSlc := [4]int{x1, y1, x2, y2}
		coordSlc[i] = tempSlc
	}

	calcSum(matrix, coordSlc)
}

func scanMatrix(height, width int) [][]int {
	matrix := make([][]int, height)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}
	for i := 0; i < height; i++ {
		var tempSlc []int
		for j := 0; j < width; j++ {
			var tempInt int
			fmt.Scan(&tempInt)
			tempSlc = append(tempSlc, tempInt)
		}
		matrix[i] = tempSlc
	}
	return matrix
}

func calcSum(matrix [][]int, coordSlc [][4]int) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(len(coordSlc))
	var resSlc []int
	//resSlc := make([]int, len(coordSlc))
	for range coordSlc {
		resSlc = append(resSlc, 0)
	}
	for i := range coordSlc {
		go func(i int) {
			defer wg.Done()
			var res int
			for y := coordSlc[i][1] - 1; y <= coordSlc[i][3]-1; y++ {
				for x := coordSlc[i][0] - 1; x <= coordSlc[i][2]-1; x++ {
					res += matrix[x][y]
				}
			}
			mu.Lock()
			resSlc[i] = res
			//fmt.Println(res)
			mu.Unlock()

		}(i)
	}
	wg.Wait()

	for _, v := range resSlc {
		fmt.Println(v)
	}
}
