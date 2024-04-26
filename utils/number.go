package utils

func IsContainsNumber(word string) bool {
	for _, char := range word {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}
