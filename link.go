package goseg

type QuickSortTermLink struct {
	head *TermNode
	tail *TermNode
	size int
}

func (qstl *QuickSortTermLink) Head() *TermNode {
	return qstl.head
}

func (qstl *QuickSortTermLink) Size() int {
	return qstl.size
}

func (qstl *QuickSortTermLink) IsEmpty() bool {
	return qstl.size == 0
}

func (qstl *QuickSortTermLink) Add(term *Term) bool {
	newNode := &TermNode{content: term}
	if qstl.size == 0 {
		qstl.head = newNode
		qstl.tail = newNode
		qstl.size++
		return true
	} else {
		if qstl.tail.CompareTo(newNode) == 0 {
			// 新节点与尾部节点相同，不加入链表
			return false
		} else if qstl.tail.CompareTo(newNode) < 0 {
			// 新节点接入链表尾部
			qstl.tail.next = newNode
			newNode.prev = qstl.tail
			qstl.tail = newNode
			qstl.size++
			return true
		} else if qstl.head.CompareTo(newNode) > 0 {
			// 新节点接入链表头部
			qstl.head.prev = newNode
			newNode.next = qstl.head
			qstl.head = newNode
			qstl.size++
			return true
		} else {
			// 从尾部上逆
			cursor := qstl.tail
			for cursor != nil && cursor.CompareTo(newNode) > 0 {
				cursor = cursor.prev
			}
			if cursor != nil {
				compareResult := cursor.CompareTo(newNode)
				if compareResult == 0 {
					// 新节点与链表中节点重复，不加入链表
					return false
				} else if compareResult < 0 {
					// 新节点插入当前节点后面
					newNode.prev = cursor
					newNode.next = cursor.next
					cursor.next.prev = newNode
					cursor.next = newNode
					qstl.size++
					return true
				}
			}
		}
	}
	return false
}

func (qstl *QuickSortTermLink) PeekFirst() *Term {
	if qstl.head != nil {
		return qstl.head.content
	}
	return nil
}

func (qstl *QuickSortTermLink) PeekLast() *Term {
	if qstl.tail != nil {
		return qstl.tail.content
	}
	return nil
}

func (qstl *QuickSortTermLink) PollFirst() *Term {
	if qstl.size == 1 {
		first := qstl.head.content
		qstl.head = nil
		qstl.tail = nil
		qstl.size--
		return first
	} else if qstl.size > 1 {
		first := qstl.head.content
		qstl.head = qstl.head.next
		qstl.head.prev.next = nil
		qstl.head.prev = nil
		qstl.size--
		return first
	} else {
		return nil
	}
}

func (qstl *QuickSortTermLink) PollLast() *Term {
	if qstl.size == 1 {
		last := qstl.tail.content
		qstl.head = nil
		qstl.tail = nil
		qstl.size--
		return last
	} else if qstl.size > 1 {
		last := qstl.tail.content
		qstl.tail = qstl.tail.prev
		qstl.tail.next.prev = nil
		qstl.tail.next = nil
		qstl.size--
		return last
	} else {
		return nil
	}
}
