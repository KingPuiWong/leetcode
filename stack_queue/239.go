package stackqueue

type slidQueue struct {
	data []int
}

func (s *slidQueue) isEmpty() bool {
	return s.getQueueLen() == 0
}

func (s *slidQueue) back() int {
	return s.data[s.getQueueLen()-1]
}

func (s *slidQueue) getQueueLen() int {
	return len(s.data)
}

func (s *slidQueue) push(val int) {
	for !s.isEmpty() && val > s.back() {
		s.data = s.data[:s.getQueueLen()-1]
	}
	s.data = append(s.data, val)
}

func (s *slidQueue) getFont() int {
	return s.data[0]
}

func (s *slidQueue) pop(val int) {
	if !s.isEmpty() && val == s.getFont() {
		s.data = s.data[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := slidQueue{data: make([]int, 0)}
	for i := 0; i < k; i++ {
		queue.push(nums[i])
	}
	var res []int
	res = append(res, queue.getFont())
	for i := k; i < len(nums); i++ {
		queue.pop(nums[i-k])
		queue.push(nums[i])
		res = append(res, queue.getFont())
	}
	return res
}
