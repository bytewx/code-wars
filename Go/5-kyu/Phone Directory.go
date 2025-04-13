package kata

import (
	"regexp"
	"strings"
)

func Phone(dir, num string) string {
	lines := strings.Split(dir, "\n")
	matchingLines := []string{}

	for _, line := range lines {
		if strings.Contains(line, "+"+num) {
			matchingLines = append(matchingLines, line)
		}
	}

	if len(matchingLines) > 1 {
		return "Error => Too many people: " + num
	}
  
	if len(matchingLines) == 0 {
		return "Error => Not found: " + num
	}

	line := matchingLines[0]

	nameRegex := regexp.MustCompile(`<([^>]+)>`)
	nameMatch := nameRegex.FindStringSubmatch(line)
	name := ""
	if len(nameMatch) > 1 {
		name = nameMatch[1]
	}

	line = strings.ReplaceAll(line, "<"+name+">", "")
	line = strings.ReplaceAll(line, "+"+num, "")

	addr := line
	addr = regexp.MustCompile(`[^a-zA-Z0-9\.-]`).ReplaceAllString(addr, " ")
	addr = regexp.MustCompile(`\s+`).ReplaceAllString(addr, " ")
	addr = strings.TrimSpace(addr)

	return "Phone => " + num + ", Name => " + name + ", Address => " + addr
}
