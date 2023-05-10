package models

var FileInfoDB = "file_info_db"

type FileMeta struct {
	Size    int64  `json:"size"`
	ModTime int64  `json:"modtime"`
	Md5     string `json:"md5"`
}

type FileInfo struct {
	FileMeta
	ID          string `json:"id"`
	Name        string `json:"name"`
	CreatedTime int64  `json:"create_time"`
	Storage     string `json:"storage"`
}

func (d FileInfo) GetDB() string {
	return FileInfoDB
}
func (d FileInfo) GetID() string {
	return d.ID
}
func (d FileInfo) GetSize() int64 {
	return d.Size
}
func (d FileInfo) GetName() string {
	return d.Name
}
func (d FileInfo) GetModTime() int64 {
	return d.ModTime
}
