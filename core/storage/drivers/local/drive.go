package local

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/storage/drivers/vo"
	"github.com/pipikai/yun/core/storage/models"
	"github.com/pipikai/yun/pb"
	"github.com/shirou/gopsutil/disk"
)

var temp_path = "temp/_tmp/"

type Addition struct {
	Cap      int64
	RootPath string
}

var mkdirPerm = 0777

type Local struct {
	Addition
}

func (d *Local) GetCap(ctx context.Context) (int64, error) {
	return d.Cap, nil
}
func (d *Local) Init(ctx context.Context) error {
	if !util.Exists(d.RootPath) {
		if err := os.MkdirAll(d.RootPath, fs.FileMode(mkdirPerm)); err != nil {
			return err
		}
	}

	path := d.RootPath
	if runtime.GOOS == "windows" {
		path = "C:"
	}
	v, err := disk.Usage(path)
	if err != nil {
		return err
	}
	if v.Path != path {
		return err
	}
	d.Cap = int64(v.Free)
	fmt.Println(v)
	if !filepath.IsAbs(d.RootPath) {
		abs, err := filepath.Abs(d.RootPath)
		if err != nil {
			return err
		}
		d.RootPath = abs
	}

	return nil
}

// func (d *Local) Upload(ctx context.Context, file vo.IStreamFile) (vo.FileMete, error) {

// 	path := "/" + path.Ext(file.GetName()) + "/" + util.GetYMDStylePath() + "/"

// 	if err := os.MkdirAll(d.RootPath+path, fs.FileMode(mkdirPerm)); err != nil {
// 		return nil, err
// 	}
// 	f, err := os.OpenFile(d.RootPath+path+util.GetUnixString()+"-"+file.GetName(), os.O_CREATE|os.O_RDWR, fs.FileMode(mkdirPerm))
// 	if err != nil {
// 		return nil, err
// 	}
// 	_, err = f.Write(file.GetContent())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return base.BlockInfo{
// 		Scheme: "http",
// 		Path:   path + util.GetUnixString() + "-" + file.GetName(),
// 	}, nil
// }

// PreUpload(context.Context, *pb.PreUploadRequest) (*pb.PreUploadReply, error)
// Upload(context.Context, *pb.UploadRequest) (*pb.UploadReply, error)
// CreateFile(context.Context, *pb.MergeRequest) (*pb.MergeReply, error)
func (d *Local) PreUpload(ctx context.Context, req *pb.PreUploadRequest) (*pb.PreUploadReply, error) {
	block_status := make([]bool, 0)
	for _, v := range req.BlockMd5 {
		_, err := os.Stat(temp_path + v)
		block_status = append(block_status, !os.IsNotExist(err))
	}
	return &pb.PreUploadReply{
		Code:        1,
		Blockstatus: block_status,
	}, nil
}

func (d *Local) Upload(ctx context.Context, file *pb.UploadRequest) (*pb.UploadReply, error) {

	hash := md5.New()
	_, err := hash.Write(file.RawData)
	if err != nil {
		return nil, err
	}
	md5_string := hex.EncodeToString(hash.Sum(nil))

	if err := os.MkdirAll(temp_path, fs.FileMode(mkdirPerm)); err != nil {
		return nil, err
	}
	// _, err = os.Stat(temp_path + md5_string)
	// if !os.IsNotExist(err) {
	// 	// 若文件已经存在
	// 	os.
	// 	return &pb.UploadReply{
	// 		Md5: md5_string,
	// 	}, nil
	// }

	f, err := os.OpenFile(temp_path+md5_string, os.O_CREATE|os.O_RDWR, fs.FileMode(mkdirPerm))
	if err != nil {
		return nil, err
	}
	_, err = f.Write(file.RawData)
	f.Close()
	if err != nil {
		return nil, err
	}
	return &pb.UploadReply{
		Md5: md5_string,
	}, nil
}

func (d *Local) CreateFile(ctx context.Context, file *models.File) (*pb.MergeReply, error) {
	f, err := os.OpenFile(d.RootPath+"/"+file.GetPath(), os.O_CREATE|os.O_RDWR, fs.FileMode(mkdirPerm))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// toDo 这里可以考虑用协程
	for _, filename := range file.BlockMd5 {
		content, err := os.ReadFile(temp_path + filename)
		if err != nil {
			return nil, err
		}
		f.Write(content)
	}
	for _, filename := range file.BlockMd5 {
		err = os.Remove(temp_path + filename)
		if err != nil {
			return nil, err
		}
	}
	return &pb.MergeReply{
		Path: file.GetPath(),
	}, nil
}

func (d *Local) Link(ctx context.Context, dir vo.IDir) (vo.ILink, error) {
	return nil, nil
}

func (d *Local) Remove(ctx context.Context, dir vo.IDir) error {

	return nil
}
