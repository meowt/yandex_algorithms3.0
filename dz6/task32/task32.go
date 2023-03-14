package main

import (
	"fmt"
)

//32. Компоненты связности
//Все языки	Python 3.6
//Ограничение времени	2 секунды	5 секунд
//Ограничение памяти	256Mb	256Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//Дан неориентированный невзвешенный граф, состоящий из N вершин и M ребер. Необходимо посчитать количество его компонент связности и вывести их.
//
//Формат ввода
//Во входном файле записано два числа N и M (0 < N ≤ 100000, 0 ≤ M ≤ 100000). В следующих M строках записаны по два числа i и j (1 ≤ i, j ≤ N), которые означают, что вершины i и j соединены ребром.
//
//Формат вывода
//В первой строчке выходного файла выведите количество компонент связности. Далее выведите сами компоненты связности в следующем формате: в первой строке количество вершин в компоненте, во второй - сами вершины в произвольном порядке.

var visited = map[int]bool{}

func main() {
	var v, e int
	fmt.Scan(&v, &e)

	graph := make([][]int, v+1)
	for i := 0; i < e; i++ {
		var first, second int
		fmt.Scan(&first, &second)
		if first == second {
			graph[first] = append(graph[first], second)
			continue
		}
		graph[first] = append(graph[first], second)
		graph[second] = append(graph[second], first)
	}

	//visited = make([]bool, v+1)

	//DFS for each Vertice and finding max adjacency list by length
	var res [][]int
	for i := 1; i < v+1; i++ {
		tempRes := dfs(graph, i)
		if tempRes != nil {
			res = append(res, tempRes)
		}
	}

	printResult(res)
}

func dfs(graph [][]int, now int) (res []int) {
	if visited[now] {
		return
	}
	visited[now] = true
	res = append(res, now)
	for _, neig := range graph[now] {
		if !visited[neig] {
			res = append(res, dfs(graph, neig)...)
		}
	}
	return
}

func printResult(res [][]int) {
	//  Количество компонентов связности
	fmt.Println(len(res))
	for _, elem := range res {
		//	Длина компонента
		fmt.Println(len(elem))

		// Итоговая строка с вершинами
		s := fmt.Sprint(elem)
		fmt.Println(s[1 : len(s)-1])
	}
}
