package models

// StorageFileDb
const StorageFileDb = "tracker_file_db"

type FileMeta struct {
	Size    int64  `json:"size"`
	ModTime int64  `json:"modtime"`
	Md5     string `json:"md5"`
	Type    string `json:"type"`
}

// File
type File struct {
	ID       string `json:"id"`
	PreID    string `json:"pre_id"`
	FileMeta `json:"file_meta"`
	Group    string `json:"group"`
	Name     string `json:"name"`
	// 0: 正在上传 1:上传完成 -1:已经删除
	Status     int      `json:"status"`
	Path       string   `json:"path"`
	BlockSize  int64    `json:"block_size"`
	BlockMd5   []string `json:"block_md5"`
	Dir        bool     `json:"dir"`
	Link       *Link    `json:"link"`
	UpdataTime int64    `json:"updata_time"`
}

func (d File) GetPath() string {
	return d.Path
}

func (d File) GetDB() string {
	return StorageFileDb
}

func (d File) GetID() string {
	return d.ID
}

func (d File) GetMd5() string {
	return d.Md5
}

func (d File) GetName() string {
	return d.Name
}
