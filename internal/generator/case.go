package generator

import "strings"

// CamelCase Todo 如果开头是数字会有问题
func CamelCase(name string, upper bool) string {
	var b strings.Builder
	//b.Grow(len(s))
	for i := 0; i < len(name); i++ {
		c := name[i]
		if !isValid(c) {
			upper = true
			continue
		}
		if upper && 'a' <= c && c <= 'z' {
			c -= 'a' - 'A'
			upper = false
		}
		b.WriteByte(c)
	}
	return b.String()
}

func isValid(r uint8) bool {
	if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9') {
		return true
	}
	return false
}
