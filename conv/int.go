package conv

// IntToBool :
func IntToBool(v int) bool {
	if v == 0 {
		return false
	}
	return true
}

// BoolToInt :
func BoolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}
