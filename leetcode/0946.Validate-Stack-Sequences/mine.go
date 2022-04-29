package leetcode

func validate(pushed []int, popped []int) bool {
	var s stack
	var i, j = 0, 0

	for {

		if j == len(popped) {
			return true
		}

		if len(s) > 0 && s.peek() == popped[j] {
			s.pop()
			j += 1
			continue
		}

		if i < len(pushed) {
			s.push(pushed[i])
			i++
			continue
		}

		return false
	}

}

type stack []int

func (s stack) push(i int) {
	s = append(s, i)
}

func (s stack) pop() {
	s = s[0 : len(s)-1]

}

func (s stack) peek() int {
	return s[len(s)-1]
}
