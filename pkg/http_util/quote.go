package http_util

import "strings"

// CleanQuotedString returns s quoted per quoted-string in RFC 7230 with invalid
// bytes and invalid UTF8 replaced with _.
func CleanQuotedString(s string) string {
	var result strings.Builder
	result.Grow(len(s) + 2) // optimize for case where no \ are added.

	result.WriteByte('"')
	for _, r := range s {
		if (r < ' ' && r != '\t') || r == 0x7f || r == 0xfffd {
			r = '_'
		}
		if r == '\\' || r == '"' {
			result.WriteByte('\\')
		}
		result.WriteRune(r)
	}
	result.WriteByte('"')
	return result.String()
}
