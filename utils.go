package sbapi

import (
	"strings"
	"time"
)

// Helper function to convert timestamp to a human-readable string
func formatTimestamp(ms int64) string {
	return time.UnixMilli(ms).Format("2006-01-02 15:04:05")
}

func containsPrefix(str string, skippable []string) bool {
	for _, skip := range skippable {
		if strings.HasPrefix(str, skip) {
			return true
		}
	}
	return false
}
