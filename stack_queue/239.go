package stackqueue

type slidStack struct {
	data []int
}

func maxSlidingWindow(nums []int, k int) []int {
	var myslidQueue slidStack
	var res []int
	for _, v := range nums {
		myslidQueue.push(v)
		if myslidQueue.getQueueLen() == k {
			res = append(res, myslidQueue.getMaxInQueue())
			myslidQueue.pop()

		}
	}

	return res
}

func (m *slidStack) push(s int) {
	m.data = append(m.data, s)
}

func (m *slidStack) getMaxInQueue() int {
	// 注释掉的这个性能不好,会超时
	// tempData := make([]int, len(m.data))
	// copy(tempData, m.data)
	// sort.Slice(tempData, func(i, j int) bool { return tempData[i] > tempData[j] })
	// return tempData[0]
	var max int
	for index := range m.data {
		if index == 0 {
			max = m.data[index]
		}
		if m.data[index] > max {
			max = m.data[index]
		}
	}
	return max
}

func (m *slidStack) pop() {
	m.data = m.data[1:]
}

func (m *slidStack) getQueueLen() int {
	return len(m.data)
}
