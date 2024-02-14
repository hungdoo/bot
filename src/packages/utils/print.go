package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(data interface{}) string {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("PrettyPrint Error: %v\n", err)
		return ""
	}
	fmt.Printf("%s", string(b))
	return string(b)
}
