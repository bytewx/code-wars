package kata

import "unicode"

type MyString string

func (s MyString) IsUpperCase() bool {
    for _, r := range s {
        if unicode.IsLower(r) {
            return false
        }
    }
    return true
}
