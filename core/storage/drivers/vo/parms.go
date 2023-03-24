package vo

import (
	"time"
)

type IDir interface {
	GetID() string
	GetSize() int64
	GetName() string
	ModTime() time.Time
	IsDir() bool
	GetPath() string
}

type IStreamFile interface {
	GetSize() int64
	GetName() string
	GetContent() []byte
}
