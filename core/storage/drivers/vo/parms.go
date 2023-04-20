package vo

type IDir interface {
	GetID() string
	GetSize() int64
	GetName() string
}

type IStreamFile interface {
	GetSize() int64
	GetName() string
	GetContent() []byte
}
