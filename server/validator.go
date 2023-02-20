package server

// Chceking messages and names function
func isValidStr(str string) bool {
	for _, w := range str {
		if w != ' ' && w != '\n' {
			return true
		}
	}
	return false
}
