package string

import (
	"strconv"
	"strings"
)

func StartsWith(a, b string) bool {
	return strings.HasPrefix(a, b)
}

func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func StringToPointer(v string) *string {
	return &v
}
