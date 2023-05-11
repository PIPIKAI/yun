// base parms
package base

// BlockInfo
type BlockInfo struct {
	Md5 string
}

// GetMd5
//
//	@receiver b
//	@return string
func (b BlockInfo) GetMd5() string {
	return b.Md5
}
