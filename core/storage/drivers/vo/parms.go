package vo

// IDir
type IDir interface {
	GetID() string
	GetSize() int64
	GetName() string
}

// IStreamFile
type IStreamFile interface {
	GetSize() int64
	GetName() string
	GetContent() []byte
}

// ICreateReq
type ICreateReq interface {
	GetContentMd5() string
	GetSlickMd5() string
	GetBlockMd5() []string
	GetSize() string
	GetName() string
	GetDir() string
}
