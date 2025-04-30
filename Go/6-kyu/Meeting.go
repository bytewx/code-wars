package kata

import (
    "sort"
    "strings"
)

func Meeting(s string) string {
    s = strings.ToUpper(s)

    people := strings.Split(s, ";")

    var entries []string
    for _, person := range people {
        names := strings.Split(person, ":")
        entries = append(entries, "(" + names[1] + ", " + names[0] + ")")
    }

    sort.Strings(entries)

    return strings.Join(entries, "")
}
