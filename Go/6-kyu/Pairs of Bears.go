package kata

func CheckPairs(nbBear int, bearList string) (bearCouple string, hasEnoughCouple bool) {
    pairs := ""
    runes := []rune(bearList)

    for i := 0; i < len(runes)-1; i++ {
        // check for 'B8' or '8B'
        if (runes[i] == 'B' && runes[i+1] == '8') || (runes[i] == '8' && runes[i+1] == 'B') {
            pairs += string(runes[i]) + string(runes[i+1])
            i++ // skip next bear since it's used
        }
    }

    hasEnoughCouple = len(pairs) >= nbBear
    return pairs, hasEnoughCouple
}
