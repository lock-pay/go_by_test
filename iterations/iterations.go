package iterations

import "strings"

func Repeat(value string, count int) string {
	var repeated strings.Builder
	for range count {
		repeated.WriteString(value)
	}
	return repeated.String()
}
