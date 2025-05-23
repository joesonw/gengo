package gengo

import (
	"fmt"
	"go/format"
	"strings"
)

func FormatSource(src string) (string, error) {
	output, err := format.Source([]byte(src))
	if err != nil {
		lines := strings.Split(src, "\n")
		length := len(lines)
		lineFormat := fmt.Sprintf("%%%dd: %%s\n", len(fmt.Sprint(length)))
		for i, line := range lines {
			fmt.Printf(lineFormat, i+1, line)
		}
		fmt.Println("")
		return "", err
	}
	return string(output), nil
}
