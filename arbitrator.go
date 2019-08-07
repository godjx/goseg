package goseg

import (
	"github.com/emirpasic/gods/sets/treeset"
	"github.com/emirpasic/gods/stacks/arraystack"
)

func Arbitrate(termLink *QuickSortTermLink, smart bool) map[int]*TermPath {
	pathMap := make(map[int]*TermPath)
	term := termLink.PollFirst()
	crossPath := NewTermPath()
	for term != nil {
		if !crossPath.AddCrossTerm(term) {
			if crossPath.Size() == 1 || !smart {
				if crossPath != nil {
					pathMap[crossPath.Begin()] = crossPath
				}
			} else {
				// 对当前的 crosspath 进行消岐
				head := crossPath.Head()
				judgeResult := judge(head)
				if judgeResult != nil {
					pathMap[judgeResult.Begin()] = judgeResult
				}
			}
			crossPath = NewTermPath()
			crossPath.AddCrossTerm(term)
		}
		term = termLink.PollFirst()
	}

	if crossPath.Size() == 1 || !smart {
		if crossPath != nil {
			pathMap[crossPath.Begin()] = crossPath
		}
	} else {
		head := crossPath.Head()
		judgeResult := judge(head)
		if judgeResult != nil {
			pathMap[judgeResult.Begin()] = judgeResult
		}
	}

	return pathMap
}

func judge(node *TermNode) *TermPath {
	pathOptions := treeset.NewWith(termPathComparator)
	option := NewTermPath()

	termStack := forward(node, option)
	pathOptions.Add(option.Copy())

	var tn *TermNode
	for !termStack.Empty() {
		n, ok := termStack.Pop()
		if ok {
			tn = n.(*TermNode)
			backward(tn.content, option)
			forward(tn, option)
			pathOptions.Add(option.Copy())
		}
	}

	if pathOptions.Empty() {
		return nil
	} else {
		return pathOptions.Values()[0].(*TermPath)
	}
}

func forward(node *TermNode, option *TermPath) *arraystack.Stack {
	conflictStack := arraystack.New()
	cursor := node
	for cursor != nil && cursor.content != nil {
		if !option.AddNotCrossTerm(cursor.content) {
			conflictStack.Push(cursor)
		}
		cursor = cursor.Next()
	}
	return conflictStack
}

func backward(term *Term, option *TermPath) {
	for option.HasCrossWith(term) {
		option.RemoveTail()
	}
}

func termPathComparator(a, b interface{}) int {
	pathA := a.(*TermPath)
	pathB := b.(*TermPath)
	return pathA.CompareTo(pathB)
}
