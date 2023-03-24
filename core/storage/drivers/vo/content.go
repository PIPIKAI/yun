package vo

import "net/http"

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
	Scheme string
	Path   string
	Header http.Header `json:"header"` // needed header
}

func (l Link) GetScheme() string {
	return l.Scheme
}
func (l Link) GetPath() string {
	return l.Path
}
func (l Link) GetHeader() http.Header {
	return l.Header
}
