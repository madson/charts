package string_helper

func MaxString(s string, max int) string {
	if len(s) > max {
		r := 0
		for i := range s {
			r++
			if r > max {
				return s[:i-3] + "..."
			}
		}
	}
	return s
}
