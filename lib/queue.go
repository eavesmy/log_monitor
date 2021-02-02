package lib

type Queue []string

func (q *Queue) Push(v string) {
	*q = append(*q, v)
}

func (q *Queue) Drop() { // 放弃掉前100个
	*q = (*q)[1000:]
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) Search(str string) {
	// 后向前查找对应行
}
