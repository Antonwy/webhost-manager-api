package util

import "strings"

func LoweredAndUnderscored(text string) string {
	lowercased := strings.ToLower(text)
	underscored := strings.ReplaceAll(lowercased, " ", "_")

	return underscored
}
