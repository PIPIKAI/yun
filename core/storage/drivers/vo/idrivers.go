package vo

import (
	"context"
	"net/http"
)

type ILink interface {
	GetPath() string
	GetHeader() http.Header
}

type Driver interface {
	GetAddition() Addition
	Init(context.Context) error
	Readder
	Writter
}

type Readder interface {
	Link(context.Context, IDir) (ILink, error)
}

type Writter interface {
	Upload(context.Context, IStreamFile) (ILink, error)
	Remove(context.Context, IDir) error
}

type Addition interface{}
