package goseg

import (
	"goseg/dict"
	"strings"
)

type Segmenter struct {
	dictionary *dict.Dictionary
}

func NewSegmenter(dictionary *dict.Dictionary) *Segmenter {
	return &Segmenter{dictionary: dictionary}
}

func (seg *Segmenter) Segment(text string) TermList {
	var prefixTokens []*dict.Token
	chars := []rune(strings.ToLower(text))
	terms := &QuickSortTermLink{}
	for cursor := range chars {
		prefixTokens = seg.processCurrentChar(chars, cursor, prefixTokens, terms)
	}

	pathMap := Arbitrate(terms, true)

	result := make(TermList, 0)
	for index := 0; index < len(chars); {
		path, ok := pathMap[index]
		if ok && path != nil {
			term := path.PollFirst()
			for term != nil {
				term.text = string(chars[term.begin:term.GetEndPosition()])
				result = append(result, term)
				index = term.GetEndPosition()
				term = path.PollFirst()
			}
		} else {
			index++
		}
	}

	return result
}

func (seg *Segmenter) processCurrentChar(chars []rune, currentIndex int, prefixTokens []*dict.Token, terms *QuickSortTermLink) []*dict.Token {
	if len(prefixTokens) > 0 {
		// 优先处理缓存的 token
		tmpTokens := make([]*dict.Token, len(prefixTokens))
		copy(tmpTokens, prefixTokens)
		for _, token := range tmpTokens {
			token = seg.dictionary.SearchWithToken(chars, currentIndex, token)
			if token.IsMatch() {
				// 输出当前匹配到的词条
				term := &Term{begin: token.Begin(), length: currentIndex - token.Begin() + 1}
				terms.Add(term)

				if !token.IsPrefix() {
					// 不是前缀，移除 prefix token
					prefixTokens = removeToken(prefixTokens, token)
				}
			} else if token.IsMismatch() {
				prefixTokens = removeToken(prefixTokens, token)
			}
		}
	}

	// 对当前指针位置进行单字匹配
	token := seg.dictionary.SearchInMain(chars, currentIndex, 1)
	if token.IsMatch() {
		term := &Term{begin: currentIndex, length: 1}
		terms.Add(term)

		if token.IsPrefix() {
			prefixTokens = append(prefixTokens, token)
		}
	} else if token.IsPrefix() {
		prefixTokens = append(prefixTokens, token)
	}

	return prefixTokens
}

func removeToken(tokens []*dict.Token, token *dict.Token) []*dict.Token {
	if len(tokens) == 0 {
		return tokens
	}

	j := 0
	for i, t := range tokens {
		if t == token {
			j = i
			break
		}
	}
	if j == 0 {
		return tokens[1:]
	}
	if j == len(tokens)-1 {
		return tokens[:j]
	}

	return append(tokens[:j], tokens[j+1:]...)
}
