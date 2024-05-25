package stackqueue

import "strconv"

type rpnQueue struct {
	data []string
}

func evalRPN(tokens []string) int {
	var res int
	if len(tokens) == 1 {
		res, _ = strconv.Atoi(tokens[0])
	}
	var queue rpnQueue
	for _, v := range tokens {
		if !isMathToken(v) {
			queue.push(v)
		} else {
			num1, _ := strconv.Atoi(queue.pop())
			num2, _ := strconv.Atoi(queue.pop())
			switch v {
			case "+":
				res = num1 + num2
			case "-":
				res = num2 - num1
			case "*":
				res = num1 * num2
			case "/":
				res = num2 / num1
			}
			queue.push(strconv.Itoa(res))
		}
	}
	return res
}

func (r *rpnQueue) push(n string) {
	r.data = append(r.data, n)
}

func (r *rpnQueue) empty() bool {
	return len(r.data) == 0
}

func (r *rpnQueue) pop() string {
	res := r.data[len(r.data)-1]
	r.data = r.data[0 : len(r.data)-1]
	return res
}

func isMathToken(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}
