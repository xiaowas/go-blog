package utils

import (
	"strconv"
	"strings"
)

func IndexForOne(i int, p, limit int64) int64 {
	s := strconv.Itoa(i)
	index, _ := strconv.ParseInt(s, 10, 64)
	return (p-1)*limit + index + 1
}

func IndexAddOne(i interface{}) int64 {
	index, _ := ToInt64(i)
	return index + 1
}

func IndexDecrOne(i interface{}) int64 {
	index, _ := ToInt64(i)
	return index - 1
}

func StringReplace(str,old,new string) string {
	return  strings.Replace(str,old,new,-1)
}
