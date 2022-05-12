package str

import (
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"strings"
	"unsafe"
)

// SliceSplit  把切片如[]string []int  分割成用指定符号分割的字符串
func SliceSplit(slice []string, separator string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", separator, -1)
}

// ByteToString byte 转字符串
func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Plural 转为复数 user -> users
func Plural(word string) string {
	return pluralize.NewClient().Singular(word)
}

// Singular 转为单数 users -> user
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

// Snake 转为 snake_case，如 TopicComment -> topic_comment
func Snake(s string) string {
	return strcase.ToSnake(s)
}

// Camel 转为 CamelCase，如 topic_comment -> TopicComment
func Camel(s string) string {
	return strcase.ToCamel(s)
}

// LowerCamel 转为 lowerCamelCase，如 TopicComment -> topicComment
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
