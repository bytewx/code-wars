package kata

import (
    "regexp"
    "strings"
)

func Catalog(s string, article string) string {
    re := regexp.MustCompile(`<prod><name>(.*?)</name><prx>(.*?)</prx><qty>(.*?)</qty></prod>`)
    matches := re.FindAllStringSubmatch(s, -1)

    var results []string

    for _, m := range matches {
        name := m[1]
        price := m[2]
        qty := m[3]

        if strings.Contains(name, article) {
            line := name + " > prx: $" + price + " qty: " + qty
            results = append(results, line)
        }
    }

    if len(results) == 0 {
        return "Nothing"
    }

    return strings.Join(results, "\n")
}
