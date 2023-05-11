package base

import "net/http"

// Link
type Link struct {
	Scheme string
	Path   string
	Header http.Header `json:"header"` // needed header
}

// GetScheme
//
//	@receiver l
//	@return string
func (l Link) GetScheme() string {
	return l.Scheme
}

// GetPath
//
//	@receiver l
//	@return string
func (l Link) GetPath() string {
	return l.Path
}

// GetHeader
//
//	@receiver l
//	@return http.Header
func (l Link) GetHeader() http.Header {
	return l.Header
}
