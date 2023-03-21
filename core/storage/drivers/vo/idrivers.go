package vo

import "context"

type ILink interface {
	GetPath() string
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
