package util

func IF(cond bool, a, b interface{}) interface{} {
	if cond {
		return a
	}
	return b
}
