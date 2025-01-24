package kata

import (
	"strings"
)

func RakeGarden(garden string) string {
	items := strings.Split(garden, " ")

	for i, item := range items {
		if item != "rock" && item != "gravel" {
			items[i] = "gravel"
		}
	}

	return strings.Join(items, " ")
}
