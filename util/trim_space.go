package util

import "strings"

func TrimSpace(str string) string {
	//str = strings.ReplaceAll(str, " ", "")

	return strings.ReplaceAll(str, "\n", "")
	//return str
}
