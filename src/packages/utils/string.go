package utils

import (
	"fmt"
	"strings"
)

func Stringify(data []interface{}) string {
	var strs []string
	for _, item := range data {
		switch v := item.(type) {
		case int:
			strs = append(strs, fmt.Sprintf("%d", v))
		case bool:
			strs = append(strs, fmt.Sprintf("%t", v))
		case string:
			strs = append(strs, v)
		default:
			strs = append(strs, fmt.Sprintf("%v", v))
		}
	}
	return strings.Join(strs, "\n")
}
