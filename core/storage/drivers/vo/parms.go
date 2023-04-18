package vo

import (
	"time"
)

type IDir interface {
	GetID() string
	GetSize() int64
	GetName() string
	ModTime() time.Time
}

type IStreamFile interface {
	GetSize() int64
	GetName() string
	GetContent() []byte
}
