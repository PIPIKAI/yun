package vo

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

type Link struct {
	Path string
}

func (l Link) GetPath() string {
	return l.Path
}
