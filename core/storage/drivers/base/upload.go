package base

// StreamFile
type StreamFile struct {
	Name    string
	Size    int64
	Content []byte
}

// GetName
//
//	@receiver s
//	@return string
func (s StreamFile) GetName() string {
	return s.Name
}

// GetSize
//
//	@receiver s
//	@return int64
func (s StreamFile) GetSize() int64 {
	return s.Size
}

// GetContent
//
//	@receiver s
//	@return []byte
func (s StreamFile) GetContent() []byte {
	return s.Content
}

// PreUploadRes
type PreUploadRes struct {
	Code        int
	BlockStatus []bool
}
