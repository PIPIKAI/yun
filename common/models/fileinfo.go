// package db models
package models

// FileInfoDB
var FileInfoDB = "file_info_db"

// FileMeta
type FileMeta struct {
	Size    int64  `json:"size"`
	ModTime int64  `json:"modtime"`
	Md5     string `json:"md5"`
}

// FileInfo
type FileInfo struct {
	FileMeta
	ID          string `json:"id"`
	Name        string `json:"name"`
	CreatedTime int64  `json:"create_time"`
	Storage     string `json:"storage"`
}

// GetDB
//
//	@receiver d
//	@return string
func (d FileInfo) GetDB() string {
	return FileInfoDB
}

// GetID
//
//	@receiver d
//	@return string
func (d FileInfo) GetID() string {
	return d.ID
}

// GetSize
//
//	@receiver d
//	@return int64
func (d FileInfo) GetSize() int64 {
	return d.Size
}

// GetName
//
//	@receiver d
//	@return string
func (d FileInfo) GetName() string {
	return d.Name
}

// GetModTime
//
//	@receiver d
//	@return int64
func (d FileInfo) GetModTime() int64 {
	return d.ModTime
}
