// util
package util

// Tnt2Bool
//
//	@param a
//	@return bool
func Tnt2Bool(a int64) bool {
	return a != 0
}

// Bool2Int
//
//	@param a
//	@return int64
func Bool2Int(a bool) int64 {
	if a {
		return 1
	}
	return 0
}
