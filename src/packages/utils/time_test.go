package utils

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	durationStrings := []string{"30s", "15m", "2h"}
	for _, durationString := range durationStrings {
		duration, err := ParseCustomDuration(durationString)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Parsed Duration for %s: %v\n", durationString, duration)
		}
	}
}
