package statement

import "strings"

func BatchStmt(str ...string) string {
	var sb strings.Builder
	sb.WriteString("begin batch ")
	for _, s := range str {
		sb.WriteString(s + "; ")
	}
	sb.WriteString(" apply batch;")
	return sb.String()
}
