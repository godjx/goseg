package goseg

import "github.com/emirpasic/gods/sets/linkedhashset"

type TermList []*Term

type Term struct {
	begin  int
	length int
	text   string
}

func (term *Term) GetLength() int {
	return term.length
}

func (term *Term) GetBeginPosition() int {
	return term.begin
}

func (term *Term) GetEndPosition() int {
	return term.begin + term.length
}

func (term *Term) GetText() string {
	return term.text
}

func (term *Term) SetText(text string) {
	term.text = text
}

func (term *Term) Equals(other *Term) bool {
	if other == nil {
		return false
	}
	if term == other {
		return true
	}
	if term.begin == other.begin && term.length == other.length {
		return true
	}
	return false
}

func (term *Term) Append(other *Term) bool {
	if other != nil && term.GetEndPosition() == other.GetBeginPosition() {
		term.length += other.length
		return true
	}
	return false
}

func (term *Term) CompareTo(o interface{}) int {
	other := o.(*Term)
	if term.begin < other.begin {
		return -1
	} else if term.begin == other.begin {
		if term.length > other.length {
			return -1
		} else if term.length == other.length {
			return 0
		} else {
			return 1
		}
	} else {
		return 1
	}
}

func (termList TermList) Words() []string {
	words := make([]string, 0)
	for _, term := range termList {
		words = append(words, term.text)
	}
	return words
}

func (termList TermList) WordSet() *linkedhashset.Set {
	wordSet := linkedhashset.New()
	for _, term := range termList {
		wordSet.Add(term.text)
	}
	return wordSet
}
