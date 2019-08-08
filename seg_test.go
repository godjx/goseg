package goseg

import (
	"fmt"
	"github.com/godjx/goseg/dict"
	"strings"
	"testing"
)

func TestSegmenter_Segment(t *testing.T) {
	filepath := "/Users/jack/Develop/data/words/words.txt"
	dictionary := dict.New()
	LoadLines(filepath, func(line string) {
		wordPair := strings.SplitN(line, ",", 2)
		if len(wordPair[0]) > 0 {
			dictionary.AddWord(wordPair[0], wordPair[1])
		}
	})
	segmenter := NewSegmenter(dictionary)
	text := "衬衣女2019新款长袖秋季气质早秋白色雪纺女式女士ol职业白衬衫女POLO领"

	result := segmenter.Segment(text)
	for _, term := range result {
		fmt.Printf("begin: %d, length: %d, text: %s, pos: %s\n", term.begin, term.length, term.text, term.pos)
	}

	wordSet := result.WordSet()
	for _, e := range wordSet.Values() {
		word := e.(Word)
		fmt.Printf("text: %s      pos: %s\n", word.text, word.pos)
	}
}
