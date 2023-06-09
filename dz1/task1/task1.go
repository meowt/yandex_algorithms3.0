package main

import (
	"fmt"
	"os"
	"sort"
)

//1. Гистограмма
//Ограничение времени	1 секунда
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//
//Вовочка ломает систему безопасности Пентагона. Для этого ему понадобилось узнать,
//какие символы в секретных зашифрованных посланиях употребляются чаще других.
//Для удобства изучения Вовочка хочет получить графическое представление встречаемости символов.
//Поэтому он хочет построить гистограмму количества символов в сообщении.
//Гистограмма — это график, в котором каждому символу, встречающемуся в сообщении хотя бы один раз,
//соответствует столбик, высота которого пропорциональна количеству этих символов в сообщении.
//
//Формат ввода
//Входной файл содержит зашифрованный текст сообщения. Он содержит строчные и прописные латинские буквы,
//цифры, знаки препинания («.», «!», «?», «:», «-», «,», «;», «(», «)»), пробелы и переводы строк.
//Размер входного файла не превышает 10000 байт. Текст содержит хотя бы один непробельный символ.
//Все строки входного файла не длиннее 200 символов.
//Для каждого символа c кроме пробелов и переводов строк выведите столбик из символов «#»,
//количество которых должно быть равно количеству символов c в данном тексте.
//Под каждым столбиком напишите символ, соответствующий ему. Отформатируйте гистограмму так,
//чтобы нижние концы столбиков были на одной строке, первая строка и первый столбец были непустыми.
//Не отделяйте столбики друг от друга. Отсортируйте столбики в порядке увеличения кодов символов.
//
//Формат вывода
//Для каждого символа c кроме пробелов и переводов строк выведите столбик из символов «#»,
//количество которых должно быть равно количеству символов c в данном тексте.
//Под каждым столбиком напишите символ, соответствующий ему.
//Отформатируйте гистограмму так, чтобы нижние концы столбиков были на одной строке,
//первая строка и первый столбец были непустыми. Не отделяйте столбики друг от друга.
//Отсортируйте столбики в порядке увеличения кодов символов.

const (
	space   = byte(10)
	newLine = byte(32)
)

func main() {
	dat, _ := os.ReadFile("input.txt")
	histogram(string(dat))
}

func histogram(text string) {
	byteSlc := []byte(text)
	byteMap := map[byte]int{}
	var keySlc []int
	var maxCount int
	for _, b := range byteSlc {
		if b == space || b == newLine {
			continue
		}
		byteMap[b]++
		if byteMap[b] > maxCount {
			maxCount = byteMap[b]
		}
	}

	for key := range byteMap {
		keySlc = append(keySlc, int(key))
	}
	sort.Ints(keySlc)

	for i := maxCount; i != 0; i-- {
		for _, key := range keySlc {
			if byteMap[byte(key)] == i {
				fmt.Print("#")
				byteMap[byte(key)]--
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	for _, key := range keySlc {
		fmt.Print(string(rune(key)))
	}
}
