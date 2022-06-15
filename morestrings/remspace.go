package test

import (
	"bytes"
)

func Remspace(s string) string {
	var buffer bytes.Buffer

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			buffer.WriteByte(s[i])
		}
	}

	return buffer.String()
}
