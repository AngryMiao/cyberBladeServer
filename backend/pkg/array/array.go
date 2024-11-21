package array

import (
	"fmt"
	"strconv"
	"strings"
)

// Todo 后续做个 validate 前置校验

type StringToInts string

func (s StringToInts) String() string {
	return string(s)
}

func (s *StringToInts) ToInts() []int {
	v := strings.Split(s.String(), ",")
	ary := make([]int, len(v))
	for i := range ary {
		ary[i], _ = strconv.Atoi(v[i])
	}
	return ary
}

func IntsToString(v []int, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v)), delim), "[]")

}
