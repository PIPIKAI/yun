package vo

import (
	"context"
	"net/http"

	"github.com/pipikai/yun/pb"
)

// 不处理数据库，只负责文件的处理
type FileMete interface {
	GetMd5() string
}

// ILink
type ILink interface {
	GetPath() string
	GetHeader() http.Header
}

// Driver
type Driver interface {
	GetAddition() Addition
	Init(context.Context) error
	Readder
	Writter
}

// Readder
type Readder interface {
	Link(context.Context, IDir) (ILink, error)
	GetCap(context.Context) (int64, error)
}

// Writter
type Writter interface {
	Upload(context.Context, *pb.UploadRequest) (*pb.UploadReply, error)
	Download(context.Context, *pb.DownloadRequest) (*pb.DownloadReply, error)
	Remove(context.Context, IDir) error
}

// Addition
type Addition interface{}
