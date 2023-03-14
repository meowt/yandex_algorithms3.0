package main

import (
	"fmt"
	"sort"
)

//31. Поиск в глубину
//Все языки	Python 3.6
//Ограничение времени	2 секунды	5 секунд
//Ограничение памяти	256Mb	256Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//Дан неориентированный граф, возможно, с петлями и кратными ребрами. Необходимо построить компоненту связности, содержащую первую вершину.
//
//Формат ввода
//В первой строке записаны два целых числа N (1 ≤ N ≤ 103) и M (0 ≤ M ≤ 5 * 105) — количество вершин и ребер в графе. В последующих M строках перечислены ребра — пары чисел, определяющие номера вершин, которые соединяют ребра.
//
//Формат вывода
//В первую строку выходного файла выведите число K — количество вершин в компоненте связности. Во вторую строку выведите K целых чисел — вершины компоненты связности, перечисленные в порядке возрастания номеров.

var res []int

var visited = map[int]bool{}

func main() {
	var v, e int
	fmt.Scan(&v, &e)

	graph := make([][]int, v+1)
	for i := 0; i < e; i++ {
		var first, second int
		fmt.Scan(&first, &second)
		graph[first] = append(graph[first], second)
		graph[second] = append(graph[second], first)
	}

	//DFS for each Vertice and finding max adjacency list by length
	dfs(graph, 1)
	sort.Ints(res)

	printResult(res)
}

func dfs(graph [][]int, now int) {
	if visited[now] {
		return
	}
	visited[now] = true
	res = append(res, now)
	for _, neig := range graph[now] {
		if !visited[neig] {
			dfs(graph, neig)
		}
	}
	return
}

func printResult(res []int) {
	//	Длина компонента
	fmt.Println(len(res))

	// Итоговая строка с вершинами
	s := fmt.Sprint(res)
	fmt.Println(s[1 : len(s)-1])
}
