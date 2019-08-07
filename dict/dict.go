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

func (dict *Dictionary) AddWord(word string) {
	w := strings.ToLower(strings.TrimSpace(word))
	if len(w) > 0 {
		dict.main.AddWord(w)
		dict.stop.StopWord(w)
	}
}

func (dict *Dictionary) AddWords(words []string) {
	if len(words) > 0 {
		for _, word := range words {
			w := strings.ToLower(strings.TrimSpace(word))
			if len(w) > 0 {
				dict.main.AddWord(w)
				dict.stop.StopWord(w)
			}
		}
	}
}

func (dict *Dictionary) StopWord(word string) {
	w := strings.ToLower(strings.TrimSpace(word))
	if len(w) > 0 {
		dict.stop.AddWord(w)
		dict.main.StopWord(w)
	}
}

func (dict *Dictionary) StopWords(words []string) {
	if len(words) > 0 {
		for _, word := range words {
			w := strings.ToLower(strings.TrimSpace(word))
			if len(w) > 0 {
				dict.stop.AddWord(w)
				dict.main.StopWord(w)
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
