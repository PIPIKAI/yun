package models

var FileHashDB = "file_hash_db"

type FileHash struct {
	ID  string `json:"id"`
	FID string `json:"fid"`
	Ok  bool   `json:"ok"`
}

func (f FileHash) GetDB() string {
	return FileHashDB
}

func (f FileHash) GetID() string {
	return f.ID
}
