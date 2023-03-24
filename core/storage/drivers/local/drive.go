package local

import (
	"context"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/storage/drivers/vo"
)

type Addition struct {
	RootPath string
}

var mkdirPerm = 0777

type Local struct {
	Addition
}

func (d *Local) Init(ctx context.Context) error {
	if !util.Exists(d.RootPath) {
		if err := os.MkdirAll(d.RootPath, fs.FileMode(mkdirPerm)); err != nil {
			return err
		}
	}

	if !filepath.IsAbs(d.RootPath) {
		abs, err := filepath.Abs(d.RootPath)
		if err != nil {
			return err
		}
		d.RootPath = abs
	}

	return nil
}

func (d *Local) Upload(ctx context.Context, file vo.IStreamFile) (vo.ILink, error) {

	path := "/" + path.Ext(file.GetName()) + "/" + util.GetYMDStylePath() + "/"

	if err := os.MkdirAll(d.RootPath+path, fs.FileMode(mkdirPerm)); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(d.RootPath+path+util.GetUnixString()+"-"+file.GetName(), os.O_CREATE|os.O_RDWR, fs.FileMode(mkdirPerm))
	if err != nil {
		return nil, err
	}
	_, err = f.Write(file.GetContent())
	if err != nil {
		return nil, err
	}
	return vo.Link{
		Scheme: "http",
		Path:   path + util.GetUnixString() + "-" + file.GetName(),
	}, nil
}
func (d *Local) Link(ctx context.Context, dir vo.IDir) (vo.ILink, error) {
	return nil, nil
}

func (d *Local) Remove(ctx context.Context, dir vo.IDir) error {

	return nil
}
