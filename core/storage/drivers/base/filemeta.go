package base

type BlockInfo struct {
	Md5 string
}

func (b BlockInfo) GetMd5() string {
	return b.Md5
}
