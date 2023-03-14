package main

import (
	"fmt"
	"sort"
)

//8. Минимальный прямоугольник
//Ограничение времени	1 секунда
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//На клетчатой плоскости закрашено K клеток. Требуется найти минимальный по площади прямоугольник, со сторонами, параллельными линиям сетки, покрывающий все закрашенные клетки.
//
//Формат ввода
//Во входном файле, на первой строке, находится число K (1 ≤ K ≤ 100). На следующих K строках находятся пары чисел Xi и Yi – координаты закрашенных клеток (|Xi|, |Yi| ≤ 109).
//
//Формат вывода
//Выведите в выходной файл координаты левого нижнего и правого верхнего углов прямоугольника.

func main() {
	var n int
	var xSlc, ySlc []int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var tempX, tempY int
		fmt.Scan(&tempX, &tempY)
		xSlc = append(xSlc, tempX)
		ySlc = append(ySlc, tempY)
	}

	fmt.Println(calc(xSlc, ySlc))
}

func calc(xSlc, ySlc []int) (res string) {
	xMin, xMax := MinMax(xSlc)
	yMin, yMax := MinMax(ySlc)
	res = fmt.Sprintf("%v %v %v %v", xMin, yMin, xMax, yMax)
	return
}

func MinMax(v []int) (int, int) {
	sort.Ints(v)
	return v[0], v[len(v)-1]
}
