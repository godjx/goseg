package goseg

import (
	"github.com/emirpasic/gods/sets/hashset"
)

type TermList []*Term

type Term struct {
	begin  int
	length int
	text   string
	pos    string
}

func (term *Term) Length() int {
	return term.length
}

func (term *Term) BeginPosition() int {
	return term.begin
}

func (term *Term) EndPosition() int {
	return term.begin + term.length
}

func (term *Term) Text() string {
	return term.text
}

func (term *Term) SetText(text string) {
	term.text = text
}

func (term *Term) Pos() string {
	return term.pos
}

func (term *Term) SetPos(pos string) {
	term.pos = pos
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
	if other != nil && term.EndPosition() == other.BeginPosition() {
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

type Word struct {
	text string
	pos  string
}

func (word *Word) Text() string {
	return word.text
}

func (word *Word) Pos() string {
	return word.pos
}

func (termList TermList) Words() []*Word {
	words := make([]*Word, 0)
	for _, term := range termList {
		words = append(words, &Word{term.text, term.pos})
	}
	return words
}

func (termList TermList) WordSet() *hashset.Set {
	wordSet := hashset.New()
	for _, term := range termList {
		wordSet.Add(Word{term.text, term.pos})
	}
	return wordSet
}
