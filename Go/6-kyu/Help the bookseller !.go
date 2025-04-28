package kata

import (
	"fmt"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
    if len(listArt) == 0 || len(listCat) == 0 {
        return ""
    }
    
    counts := make(map[string]int)
    
    for _, art := range listArt {
        parts := strings.Split(art, " ")
        if len(parts) != 2 {
            continue
        }
        code := parts[0]
        var qty int
        fmt.Sscanf(parts[1], "%d", &qty)
        
        category := string(code[0])
        counts[category] += qty
    }
    
    var result []string
    for _, cat := range listCat {
        result = append(result, fmt.Sprintf("(%s : %d)", cat, counts[cat]))
    }
    
    return strings.Join(result, " - ")
}
