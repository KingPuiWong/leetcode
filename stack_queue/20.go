package stackqueue

var (
	leftKuohao       = '('
	rightKuohao      = ')'
	leftZhongKuohao  = '{'
	rightZhongKuohao = '}'
	leftBigKuohao    = '['
	rightBigKuohao   = ']'
)

type myQueue struct {
	data []rune
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	myqueue := myQueue{
		data: make([]rune, 0),
	}
	for _, v := range s {
		pushResult := myqueue.push(v)
		if !pushResult {
			return false
		}
	}

	if len(myqueue.data) != 0 {
		return false
	}

	return true
}

func (myqueue *myQueue) push(v rune) bool {
	if isLeft(v) {
		myqueue.data = append(myqueue.data, v)
	} else {
		firstOne, popResult := myqueue.pop()
		if !popResult {
			return false
		}
		switch firstOne {
		case leftKuohao:
			if v != rightKuohao {
				return false
			}
		case leftZhongKuohao:
			if v != rightZhongKuohao {
				return false
			}
		case leftBigKuohao:
			if v != rightBigKuohao {
				return false
			}
		}
	}

	return true
}

func (myQueue *myQueue) pop() (rune, bool) {
	if len(myQueue.data) == 0 {
		return 0, false
	}
	res := myQueue.data[len(myQueue.data)-1]
	myQueue.data = myQueue.data[0 : len(myQueue.data)-1]
	return res, true
}

func isLeft(v rune) bool {
	if v == leftKuohao || v == leftZhongKuohao || v == leftBigKuohao {
		return true
	}
	return false
}
