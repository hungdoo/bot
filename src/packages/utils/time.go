package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseCustomDuration(durationString string) (time.Duration, error) {
	// Split the string into number and unit
	re := regexp.MustCompile(`(\d+)(\w)`)
	matches := re.FindStringSubmatch(durationString)

	if len(matches) != 3 { // matches[0] is the durationString
		return 0, fmt.Errorf("invalid duration format: %s -> %v", durationString, matches)
	}

	// Parse the number part
	num, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("failed to parse duration number: %v", err)
	}

	// Map the unit to a time.Duration
	unit := strings.ToLower(matches[2])
	switch unit {
	case "s":
		return time.Duration(num) * time.Second, nil
	case "m":
		return time.Duration(num) * time.Minute, nil
	case "h":
		return time.Duration(num) * time.Hour, nil
	// Add more units as needed

	default:
		return 0, fmt.Errorf("unsupported duration unit: %s", unit)
	}
}
