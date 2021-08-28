package utils

func Plural(n int32) string {
	if n > 1 {
		return "s"
	}
	return ""
}
