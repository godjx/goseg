package dict

import "strings"

const (
	mismatch = 0x00000000
	match    = 0x00000001
	prefix   = 0x00000010
)

type DictNode struct {
	char     rune
	isWord   bool
	pos      string
	children map[rune]*DictNode
}

type Token struct {
	state      int
	begin      int
	end        int
	pos        string
	prefixNode *DictNode
}

func (token *Token) Begin() int                   { return token.begin }
func (token *Token) End() int                     { return token.end }
func (token *Token) Pos() string                  { return token.pos }
func (token *Token) SetBegin(begin int)           { token.begin = begin }
func (token *Token) SetEnd(end int)               { token.end = end }
func (token *Token) SetPos(pos string)            { token.pos = pos }
func (token *Token) IsMatch() bool                { return (token.state & match) > 0 }
func (token *Token) SetMatch()                    { token.state = token.state | match }
func (token *Token) IsPrefix() bool               { return (token.state & prefix) > 0 }
func (token *Token) SetPrefix()                   { token.state = token.state | prefix }
func (token *Token) IsMismatch() bool             { return token.state == mismatch }
func (token *Token) SetMismatch()                 { token.state = mismatch }
func (token *Token) GetPrefixNode() *DictNode     { return token.prefixNode }
func (token *Token) SetPrefixNode(node *DictNode) { token.prefixNode = node }
func (token *Token) Equals(other *Token) bool {
	if other == nil {
		return false
	}
	return token.begin == other.begin && token.end == other.end
}

func (node *DictNode) IsParent() bool {
	return len(node.children) > 0
}

func (node *DictNode) Search(chars []rune, begin, length int, token *Token) *Token {
	if token == nil {
		token = &Token{begin: begin}
	} else {
		token.SetMismatch()
	}
	token.end = begin
	keyChar := chars[begin]
	if node.children != nil {
		dn, ok := node.children[keyChar]
		if ok && dn != nil {
			if length > 1 {
				// 词未匹配完，继续向下搜索
				return dn.Search(chars, begin+1, length-1, token)
			} else if length == 1 {
				if dn.isWord {
					token.SetMatch()
					token.pos = dn.pos
				}
				if dn.IsParent() {
					token.SetPrefix()
					token.SetPrefixNode(dn)
				}
				return token
			}
		}
	}
	return token
}

func (node *DictNode) AddWord(word, pos string) {
	word = strings.TrimSpace(word)
	if len(word) == 0 {
		return
	}
	chars := []rune(word)
	node.insert(chars, 0, len(chars), true, pos)
}

func (node *DictNode) StopWord(word, pos string) {
	word = strings.TrimSpace(word)
	if len(word) == 0 {
		return
	}
	chars := []rune(word)
	node.insert(chars, 0, len(chars), false, pos)
}

func (node *DictNode) insert(chars []rune, begin, length int, isWord bool, pos string) {
	keyChar := chars[begin]
	dn := node.getChildOrCreate(keyChar, isWord)
	if dn != nil {
		if length > 1 {
			dn.insert(chars, begin+1, length-1, isWord, pos)
		} else if length == 1 {
			dn.isWord = isWord
			dn.pos = pos
		}
	}
}

func (node *DictNode) getChildOrCreate(keyChar rune, isWord bool) *DictNode {
	if node.children == nil {
		node.children = make(map[rune]*DictNode)
	}
	dn, ok := node.children[keyChar]
	if ok && dn != nil {
		return dn
	} else if isWord {
		dn = &DictNode{char: keyChar}
		node.children[keyChar] = dn
		return dn
	}
	return nil
}
