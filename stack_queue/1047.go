package stackqueue

type myQueue struct {
	data []rune
}

func removeDuplicates(s string) string {

	var my myQueue
	my.data = make([]rune, 0)
	for _, v := range s {
		my.push(v)
	}
	return string(my.data)
}

func (m *myQueue) push(s rune) {
	if len(m.data) == 0 {
		m.data = append(m.data, s)
	} else {
		pre := m.getTop()
		if pre == s {
			m.pop()
			return
		} else {
			m.data = append(m.data, s)
		}
	}
}

func (m *myQueue) getTop() rune {
	res := m.data[len(m.data)-1]
	return res
}

func (m *myQueue) pop() {
	m.data = m.data[:len(m.data)-1]
}
