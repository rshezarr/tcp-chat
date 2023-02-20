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

// Chceking port function
func isValidPort(port string) bool {
	for _, v := range port {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}
