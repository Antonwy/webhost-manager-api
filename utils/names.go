package util

import "strings"

func WordPressContainerName(name string) string {
	lowercased := strings.ToLower(name)
	underscored := strings.ReplaceAll(lowercased, " ", "_")

	return "wp_" + underscored
}
