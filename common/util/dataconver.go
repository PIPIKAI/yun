package util

func Tnt2Bool(a int64) bool {
	return a != 0
}
func Bool2Int(a bool) int64 {
	if a {
		return 1
	}
	return 0
}
