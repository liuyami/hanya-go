package str

import (
	"fmt"
	"strings"
	"unsafe"
)

// SliceSplit  把切片如[]string []int  分割成用指定符号分割的字符串
func SliceSplit(slice []string, separator string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", separator, -1)
}

func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
