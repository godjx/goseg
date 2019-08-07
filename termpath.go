package goseg

type TermPath struct {
	QuickSortTermLink
	begin         int
	end           int
	payloadLength int
}

func NewTermPath() *TermPath {
	return &TermPath{begin: -1, end: -1, payloadLength: 0}
}

func (path *TermPath) Begin() int {
	return path.begin
}

func (path *TermPath) End() int {
	return path.end
}

func (path *TermPath) Length() int {
	return path.end - path.begin
}

func (path *TermPath) AddCrossTerm(term *Term) bool {
	if path.IsEmpty() {
		path.Add(term)
		path.begin = term.begin
		path.end = term.begin + term.length
		path.payloadLength += term.length
		return true
	} else if path.HasCrossWith(term) {
		path.Add(term)
		if term.begin+term.length > path.end {
			path.end = term.begin + term.length
		}
		path.payloadLength = path.end - path.begin
		return true
	} else {
		return false
	}
}

func (path *TermPath) AddNotCrossTerm(term *Term) bool {
	if path.IsEmpty() {
		path.Add(term)
		path.begin = term.begin
		path.end = term.begin + term.length
		path.payloadLength += term.length
		return true
	} else if path.HasCrossWith(term) {
		return false
	} else {
		path.Add(term)
		path.payloadLength += term.length
		head := path.PeekFirst()
		path.begin = head.begin
		tail := path.PeekLast()
		path.end = tail.begin + tail.length
		return true
	}
}

func (path *TermPath) RemoveTail() *Term {
	tail := path.PollLast()
	if path.IsEmpty() {
		path.begin = -1
		path.end = -1
		path.payloadLength = 0
	} else {
		path.payloadLength -= tail.length
		newTail := path.PeekLast()
		path.end = newTail.begin + newTail.length
	}
	return tail
}

func (path *TermPath) HasCrossWith(term *Term) bool {
	return (term.begin >= path.begin && term.begin < path.end) || (path.begin >= term.begin && path.begin < term.begin+term.length)
}

// 核心的权重函数
// term 长度积
func (path *TermPath) TermLengthWeight() int {
	product := 1
	cursor := path.head
	for cursor != nil && cursor.content != nil {
		product *= cursor.content.length
		cursor = cursor.next
	}
	return product
}

// term 位置权重
func (path *TermPath) TermPositionWeight() int {
	weight := 0
	position := 0
	cursor := path.head
	for cursor != nil && cursor.content != nil {
		position++
		weight += position * cursor.content.length
		cursor = cursor.next
	}
	return weight
}

func (path *TermPath) CompareTo(o interface{}) int {
	other := o.(*TermPath)
	// 比较有效文本长度
	if path.payloadLength > other.payloadLength {
		return -1
	} else if path.payloadLength < other.payloadLength {
		return 1
	} else {
		// term 个数越少越好
		if path.size < other.size {
			return -1
		} else if path.size > other.size {
			return 1
		} else {
			// 路径跨度越大越好
			if path.Length() > other.Length() {
				return -1
			} else if path.Length() < other.Length() {
				return 1
			} else {
				// term 位置越靠后越好
				if path.end > other.end {
					return -1
				} else if path.end > other.end {
					return 1
				} else {
					// term 长度越平均越好
					if path.TermLengthWeight() > other.TermLengthWeight() {
						return -1
					} else if path.TermLengthWeight() < other.TermLengthWeight() {
						return 1
					} else {
						// term 位置权重越大越好
						if path.TermPositionWeight() > other.TermPositionWeight() {
							return -1
						} else if path.TermPositionWeight() < other.TermPositionWeight() {
							return 1
						}
					}
				}
			}
		}
	}
	return 0
}

func (path *TermPath) Copy() *TermPath {
	cloned := NewTermPath()
	cloned.begin = path.begin
	cloned.end = path.end
	cloned.payloadLength = path.payloadLength
	cursor := path.head
	for cursor != nil && cursor.content != nil {
		cloned.Add(cursor.content)
		cursor = cursor.next
	}
	return cloned
}
