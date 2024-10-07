package util

import (
	"fmt"
	"time"
)

func ParsedDate(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}
	parsedDate, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %s", err)
	}
	return &parsedDate, nil
}

func ParsedTime(timeStr string) (*time.Time, error) {
	if timeStr == "" {
		return nil, nil
	}
	parsedTime, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		return nil, fmt.Errorf("invalid time format: %s", err)
	}
	return &parsedTime, nil
}