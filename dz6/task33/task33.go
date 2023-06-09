package main

import (
	"errors"
	"fmt"
	"os"
)

//33. Списывание
//Все языки	Python 3.6
//Ограничение времени	2 секунды	5 секунд
//Ограничение памяти	256Mb	256Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//Во время контрольной работы профессор Флойд заметил, что некоторые студенты обмениваются записками. Сначала он хотел поставить им всем двойки, но в тот день профессор был добрым, а потому решил разделить студентов на две группы: списывающих и дающих списывать, и поставить двойки только первым.
//
//У профессора записаны все пары студентов, обменявшихся записками. Требуется определить, сможет ли он разделить студентов на две группы так, чтобы любой обмен записками осуществлялся от студента одной группы студенту другой группы.
//
//Формат ввода
//В первой строке находятся два числа N и M — количество студентов и количество пар студентов, обменивающихся записками (1 ≤ N ≤ 102, 0 ≤ M ≤ N(N−1)/2).
//
//Далее в M строках расположены описания пар студентов: два числа, соответствующие номерам студентов, обменивающихся записками (нумерация студентов идёт с 1). Каждая пара студентов перечислена не более одного раза.
//
//Формат вывода
//Необходимо вывести ответ на задачу профессора Флойда. Если возможно разделить студентов на две группы - выведите YES; иначе выведите NO.

var visited = map[int]int{}

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

	visited[0] = 2
	for i := 1; i < v+1; i++ {
		err := dfs(graph, i, visited[i-1])
		if err != nil {
			fmt.Println("NO")
			os.Exit(0)
		}
	}
	fmt.Println("YES")
}

func dfs(graph [][]int, now, prevColor int) error {
	if visited[now] == prevColor {
		return errors.New("NO")
	}
	visited[now] = 3 - prevColor
	for _, neig := range graph[now] {
		if visited[neig] != 2 {
			err := dfs(graph, neig, 3-prevColor)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
