package base

type StreamFile struct {
	Name    string
	Size    int64
	Content []byte
}

func (s StreamFile) GetName() string {
	return s.Name
}
func (s StreamFile) GetSize() int64 {
	return s.Size
}
func (s StreamFile) GetContent() []byte {
	return s.Content
}

type PreUploadRes struct {
	Code        int
	BlockStatus []bool
}
