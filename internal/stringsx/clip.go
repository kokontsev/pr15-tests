package stringsx

// Clip counts Unicode code points rather than bytes.
func Clip(s string, max int) string {
	if max < 0 {
		max = 0
	}
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return string(r[:max])
}
