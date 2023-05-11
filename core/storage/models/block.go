// package
package models

import "time"

// BlockDB
var BlockDB = "block_db"

// k: md5
type Block struct {
	Md5         string
	CreatedTime time.Time
	Size        int64
	Status      string
	Path        string
}

func (d Block) GetDB() string {
	return BlockDB
}
func (d Block) GetID() string {
	return d.Md5
}
