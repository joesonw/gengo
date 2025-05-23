package gengo

import (
	"fmt"
	"strings"

	"golang.org/x/tools/imports"
)

func FormatSource(src string) (string, error) {
	output, err := imports.Process("", []byte(src), &imports.Options{
		AllErrors: true,
		Comments:  true,
	})
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
