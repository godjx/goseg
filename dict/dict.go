package dict

import "strings"

type Dictionary struct {
	main *DictNode
	stop *DictNode
}

func New() *Dictionary {
	return &Dictionary{
		main: &DictNode{},
		stop: &DictNode{},
	}
}

func (dict *Dictionary) AddWord(word, pos string) {
	w := strings.ToLower(strings.TrimSpace(word))
	if len(w) > 0 {
		dict.main.AddWord(w, pos)
		dict.stop.StopWord(w, pos)
	}
}

func (dict *Dictionary) AddWords(wordPairs [][]string) {
	if len(wordPairs) > 0 {
		for _, wordPair := range wordPairs {
			w := strings.ToLower(strings.TrimSpace(wordPair[0]))
			pos := strings.TrimSpace(wordPair[1])
			if len(w) > 0 {
				dict.main.AddWord(w, pos)
				dict.stop.StopWord(w, pos)
			}
		}
	}
}

func (dict *Dictionary) StopWord(word, pos string) {
	w := strings.ToLower(strings.TrimSpace(word))
	if len(w) > 0 {
		dict.stop.AddWord(w, pos)
		dict.main.StopWord(w, pos)
	}
}

func (dict *Dictionary) StopWords(wordPairs [][]string) {
	if len(wordPairs) > 0 {
		for _, wordPair := range wordPairs {
			word := strings.ToLower(strings.TrimSpace(wordPair[0]))
			pos := strings.TrimSpace(wordPair[1])
			if len(word) > 0 {
				dict.stop.AddWord(word, pos)
				dict.main.StopWord(word, pos)
			}
		}
	}
}

func (dict *Dictionary) SearchInMain(chars []rune, begin, length int) *Token {
	return dict.main.Search(chars, begin, length, nil)
}

func (dict *Dictionary) SearchWithToken(chars []rune, currentIndex int, token *Token) *Token {
	dn := token.GetPrefixNode()
	if dn != nil {
		return dn.Search(chars, currentIndex, 1, token)
	}
	token.SetMismatch()
	return token
}

func (dict *Dictionary) IsStopped(chars []rune, begin, length int) bool {
	return dict.stop.Search(chars, begin, length, nil).IsMatch()
}
