package main

/*
type Deq struct {
	val         int
	prevElement *Deq
	nextElement *Deq
}

func main() {
	var command string
	front := &Deq{}
	back := &Deq{}
	front, back = nil, nil
	for {
		fmt.Scan(&command)

		switch command {
		case "push_front":
			var n int
			fmt.Scan(&n)
			if front == nil {
				front, back = pushStart(n)
			} else {
				front = front.pushFront(n)
			}

		case "push_back":
			var n int
			fmt.Scan(&n)
			if front == nil {
				front, back = pushStart(n)
			} else {
				back = back.pushBack(n)
			}

		case "pop_front":
			front, back = popFront(front, back)

		case "pop_back":
			front, back = popBack(front, back)

		case "front":
			front.value()

		case "back":
			back.value()

		case "size":
			fmt.Println(front.size())

		case "clear":
			front, back = clear()

		case "exit":
			fmt.Println("bye")
			os.Exit(0)
		}
	}
}

func pushStart(n int) (newFront, newBack *Deq) {
	newFront = &Deq{
		val: n,
	}
	newBack = newFront
	fmt.Println("ok")
	return
}

func (element *Deq) pushFront(n int) (newElement *Deq) {
	newElement = &Deq{
		val:         n,
		prevElement: element,
	}
	element.nextElement = newElement
	fmt.Println("ok")
	return
}

func (element *Deq) pushBack(n int) (newElement *Deq) {
	newElement = &Deq{
		val:         n,
		nextElement: element,
	}
	element.prevElement = newElement
	fmt.Println("ok")
	return
}

func popFront(front, back *Deq) (newFront, newBack *Deq) {
	newFront, newBack = front, back
	if front != nil {
		fmt.Println(front.val)
		if front == back {
			return nil, nil
		}
		if front.prevElement != nil {
			if front.prevElement == back {
				newBack.nextElement = nil
			}
			newFront = front.prevElement
		}
		return
	}
	fmt.Println("error")
	return nil, nil
}

func popBack(front, back *Deq) (newFront, newBack *Deq) {
	newFront, newBack = front, back
	if back != nil {
		fmt.Println(back.val)
		if back == front {
			return nil, nil
		}
		if back.nextElement != nil {
			if back.nextElement == front {
				newFront.prevElement = nil
			}
			newBack = back.nextElement
		}
		return
	}
	fmt.Println("error")
	return nil, nil
}

func (element *Deq) value() {
	if element != nil {
		fmt.Println(element.val)
		return
	}
	fmt.Println("error")
	return
}

func (element *Deq) size() (resSize int) {
	for element != nil {
		resSize++
		if element.prevElement == nil {
			break
		}
		element = element.prevElement
	}
	return
}

func clear() (front, back *Deq) {
	fmt.Println("ok")
	return nil, nil
}
*/
