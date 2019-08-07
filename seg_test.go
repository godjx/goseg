package goseg

import (
	"fmt"
	"github.com/godjx/goseg/dict"
	"testing"
)

func TestSegmenter_Segment(t *testing.T) {
	filepath := "/Users/jack/Develop/data/core-words.txt"
	dictionary := dict.New()
	LoadLines(filepath, func(line string) {
		dictionary.AddWord(line)
	})
	segmenter := NewSegmenter(dictionary)
	text := "衬衣女2019新款长袖秋季气质早秋白色雪纺女式女士ol职业白衬衫女POLO领"

	result := segmenter.Segment(text)
	for _, term := range result {
		fmt.Printf("begin: %d, length: %d, text: %s\n", term.begin, term.length, term.text)
	}

	words := result.WordSet()
	words.Each(func(i int, e interface{}) {
		fmt.Printf("word: %s\n", e)
	})
}
