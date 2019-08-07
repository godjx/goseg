package goseg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadLines(filepath string, lineHandler func(string)) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("fail to load: %s\n", filepath)
		return
	}
	defer func() { _ = file.Close() }()

	buffer := bufio.NewReader(file)
	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\n")
		line = strings.TrimSpace(line)
		lineHandler(line)
	}
	fmt.Printf("load file [%s] complete\n", filepath)
}
