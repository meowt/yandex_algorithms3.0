package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//3. Коллекционер Диего
//Ограничение времени	2 секунды
//Ограничение памяти	256Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//
//Диего увлекается коллекционированием наклеек. На каждой из них написано число,
//и каждый коллекционер мечтает собрать наклейки со всеми встречающимися числами.
//Диего собрал N наклеек, некоторые из которых, возможно, совпадают. Как-то раз к нему пришли K коллекционеров.
//i-й из них собрал все наклейки с номерами не меньшими, чем pi. Напишите программу,
//которая поможет каждому из коллекционеров определить, сколько недостающих ему наклеек есть у Диего.
//Разумеется, гостей Диего не интересуют повторные экземпляры наклеек.
//
//Формат ввода
//В первой строке содержится единственное число N (0 ≤ N ≤ 100 000) — количество наклеек у Диего.
//В следующей строке содержатся N целых неотрицательных чисел (не обязательно различных) — номера наклеек Диего.
//Все номера наклеек не превосходят 109.
//В следующей строке содержится число K (0 ≤ K ≤ 100 000) — количество коллекционеров, пришедших к Диего.
//В следующей строке содержатся K целых чисел pi (0 ≤ pi ≤ 109),
//где pi — наименьший номер наклейки, не интересующий i-го коллекционера.
//
//Формат вывода
//Для каждого коллекционера в отдельной строке выведите количество различных чисел на наклейках,
//которые есть у Диего, но нет у этого коллекционера.

func main() {
	var stickersAmount, collectorsAmount int
	var collectorsMinSticker, sortedDiegoStickers []int
	diegoStickersMap := map[int]struct{}{}
	fmt.Scan(&stickersAmount)
	for i := 0; i < stickersAmount; i++ {
		var temp int
		fmt.Scan(&temp)
		diegoStickersMap[temp] = struct{}{}
	}
	for i, _ := range diegoStickersMap {
		sortedDiegoStickers = append(sortedDiegoStickers, i)
	}
	sort.Ints(sortedDiegoStickers)

	fmt.Scan(&collectorsAmount)
	for i := 0; i < collectorsAmount; i++ {
		var temp int
		fmt.Scan(&temp)
		collectorsMinSticker = append(collectorsMinSticker, temp)
	}

	var resSlc []string
	for _, minSticker := range collectorsMinSticker {
		resSlc = append(resSlc, calculate(sortedDiegoStickers, minSticker))
	}
	printResult(resSlc)
}

func calculate(diegoStickers []int, minSticker int) string {
	for i, diegoSticker := range diegoStickers {
		if minSticker <= diegoSticker {
			return strconv.Itoa(i)
		}
	}
	return strconv.Itoa(len(diegoStickers))
}

func printResult(res []string) {
	fmt.Println(strings.Join(res, "\n"))
}
