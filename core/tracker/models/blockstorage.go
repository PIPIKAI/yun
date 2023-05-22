package models

var block_storage_name = "block_storage"

type BlockStorage struct {
	ID   string
	Mark []Storage
}

func (BlockStorage) GetDB() string {
	return block_storage_name
}
func (b BlockStorage) GetID() string {
	return b.ID
}
