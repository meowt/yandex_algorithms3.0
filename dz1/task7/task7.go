package main

import (
	"fmt"
	"time"
)

//7. SNTP
//Ограничение времени	1 секунда
//Ограничение памяти	64Mb
//Ввод	стандартный ввод или input.txt
//Вывод	стандартный вывод или output.txt
//Для того чтобы компьютеры поддерживали актуальное время, они могут обращаться к серверам точного времени SNTP (Simple Network Time Protocol). К сожалению, компьютер не может просто получить время у сервера, потому что информация по сети передаётся не мгновенно: пока сообщение с текущим временем дойдёт до компьютера, оно потеряет свою актуальность. Протокол взаимодействия клиента (компьютера, запрашивающего точное время) и сервера (компьютера, выдающего точное время) выглядит следующим образом:
//
//1. Клиент отправляет запрос на сервер и запоминает время отправления A (по клиентскому времени).
//
//2. Сервер получает запрос в момент времени B (по точному серверному времени) и отправляет клиенту сообщение, содержащее время B.
//
//3. Клиент получает ответ на свой запрос в момент времени C (по клиентскому времени) и запоминает его. Теперь клиент, из предположения, что сетевые задержки при передаче сообщений от клиента серверу и от сервера клиенту одинаковы, может определить и установить себе точное время, используя известные значения A, B, C.
//
//Вам предстоит реализовать алгоритм, с точностью до секунды определяющий точное время для установки на клиенте по известным A, B и C. При необходимости округлите результат до целого числа секунд по правилам арифметики (в меньшую сторону, если дробная часть числа меньше 1/2, иначе в большую сторону).
//
//Возможно, что, пока клиент ожидал ответа, по клиентскому времени успели наступить новые сутки, однако известно, что между отправкой клиентом запроса и получением ответа от сервера прошло менее 24 часов.
//
//Формат ввода
//Программа получает на вход три временные метки A, B, C, по одной в каждой строке. Все временные метки представлены в формате «hh:mm:ss», где «hh» – это часы, «mm» – минуты, «ss» – секунды. Часы, минуты и секунды записываются ровно двумя цифрами каждое (возможно, с дополнительными нулями в начале числа).
//
//Формат вывода
//Программа должна вывести одну временную метку в формате, описанном во входных данных, – вычисленное точное время для установки на клиенте. В выводе не должно быть пробелов, пустых строк в начале вывода.

func main() {
	var strArr [3]string
	fmt.Scan(&strArr[0])
	fmt.Scan(&strArr[1])
	fmt.Scan(&strArr[2])

	fmt.Println(calc(strArr))
}

func scanTime(strArr [3]string) (arr [3]time.Time) {
	for i, v := range strArr {
		str := fmt.Sprintf("01:01:2000 %v", v)
		arr[i], _ = time.Parse("02:01:2006 15:04:05", str)
	}
	if arr[2].Before(arr[0]) || arr[2].Equal(arr[0]) {
		str := fmt.Sprintf("01:02:2000 %v", strArr[2])
		arr[2], _ = time.Parse("02:01:2006 15:04:05", str)
	}

	return
}

func calc(strArr [3]string) (res string) {
	timeArr := scanTime(strArr)
	sub := timeArr[2].Sub(timeArr[0])

	resTime := timeArr[1].Add(sub / 2)

	res = resTime.Format("15:04:05")
	return
}
