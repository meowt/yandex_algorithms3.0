package main

import (
	"fmt"
	"sync"
)

//2. Красивая строка
//Ограничение времени	1 секунда
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//
//Красотой строки назовем максимальное число идущих подряд одинаковых букв. (красота строки abcaabdddettq равна 3)
//Сделайте данную вам строку как можно более красивой, если вы можете сделать не более k операций замены символа.
//
//Формат ввода
//В первой строке записано одно целое число k (0 ≤ k ≤ 109)
//Во второй строке дана непустая строчка S (|S| ≤ 2 ⋅ 105). Строчка S состоит только из маленьких латинских букв.
//
//Формат вывода
//Выведите одно число — максимально возможную красоту строчки, которую можно получить.

func main() {
	var k int
	var text string
	fmt.Scan(&k, &text)

	fmt.Println(calculateBeauty(k, []byte(text)))
}

func calculateBeauty(k int, byteSlc []byte) (maxBeauty int) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(len(byteSlc))
	for left := 0; left < len(byteSlc); left++ {
		go func(left int) {
			defer wg.Done()

			var tempMax int
			right, count := left, k
			for count >= 0 && right < len(byteSlc) {
				if byteSlc[left] != byteSlc[right] {
					if count == 0 {
						break
					}
					count--
				}
				tempMax++
				right++
			}
			mu.Lock()
			if tempMax > maxBeauty {
				maxBeauty = tempMax
			}
			mu.Unlock()
		}(left)

	}

	wg.Add(len(byteSlc))
	for left := len(byteSlc) - 1; left >= 0; left-- {
		go func(left int) {
			defer wg.Done()

			var tempMax int
			right, count := left, k
			for count >= 0 && right >= 0 {
				if byteSlc[left] != byteSlc[right] {
					if count == 0 {
						break
					}
					count--
				}
				tempMax++
				right--
			}

			mu.Lock()
			if tempMax > maxBeauty {
				maxBeauty = tempMax
			}
			mu.Unlock()
		}(left)
	}
	wg.Wait()
	return
}
