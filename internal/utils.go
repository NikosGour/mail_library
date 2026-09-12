package internal

import "strings"

func Unescape(v string) string {
	escaped_v := strings.ReplaceAll(v, "\r", "\\r")
	escaped_v = strings.ReplaceAll(escaped_v, "\n", "\\n")
	return escaped_v
}
