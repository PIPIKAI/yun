// driver local , store file in local disk
package local

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/storage/drivers/vo"
	"github.com/pipikai/yun/pb"
	"github.com/shirou/gopsutil/disk"
)

var tempPath = "temp/_tmp/"

// Addition
type Addition struct {
	Cap      int64
	RootPath string
}

var mkdirPerm = 0777

// Local
type Local struct {
	Addition
}

// GetCap 获取容量
//
//	@receiver d
//	@param ctx
//	@return int64
//	@return error
func (d *Local) GetCap(ctx context.Context) (int64, error) {
	return d.Cap, nil
}

// Init local dirver configs disk infos
//
//	@receiver d
//	@param ctx
//	@return error
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

func (d *Local) Upload(ctx context.Context, file *pb.UploadRequest) (*pb.UploadReply, error) {

	fileName := "/" + file.Md5

	// if err := os.MkdirAll(d.RootPath+fileName, fs.FileMode(mkdirPerm)); err != nil {
	// 	return nil, err
	// }
	f, err := os.OpenFile(d.RootPath+fileName, os.O_CREATE|os.O_RDWR, fs.FileMode(mkdirPerm))
	if err != nil {
		return nil, err
	}
	_, err = f.Write(file.RawData)
	f.Close()
	if err != nil {
		return nil, err
	}
	return &pb.UploadReply{Md5: file.Md5, Code: 1}, nil
}

// PreUpload(context.Context, *pb.PreUploadRequest) (*pb.PreUploadReply, error)
// Upload(context.Context, *pb.UploadRequest) (*pb.UploadReply, error)
// CreateFile(context.Context, *pb.MergeRequest) (*pb.MergeReply, error)

// Upload 本地上传分块文件，就是将文件保存在本地目录
//
//	@receiver d
//	@param ctx
//	@param file
//	@return *pb.UploadReply
//	@return error
// func (d *Local) Upload(ctx context.Context, file *pb.UploadRequest) (*pb.UploadReply, error) {
// 	res := &pb.UploadReply{
// 		Code: 1,
// 		Md5:  file.Md5,
// 	}
// 	_, err := leveldb.GetOne[models.Block](file.Md5)
// 	if err == nil {
// 		res.Code = 2
// 		return res, nil
// 	}
// 	err = leveldb.UpdataOne(models.Block{
// 		ID:          file.Md5,
// 		Md5:         file.Md5,
// 		CreatedTime: time.Now(),
// 		Size:        int64(len(file.RawData)),
// 		Content:     file.RawData,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func (d *Local) Download(ctx context.Context, req *pb.DownloadRequest) (*pb.DownloadReply, error) {
	// 获取文件锁
	res, err := os.ReadFile(d.RootPath + "/" + req.Md5)
	if err != nil {
		return nil, err
	}
	return &pb.DownloadReply{
		Content: res,
	}, nil
}

// Link 获取文件真实链接
//
//	@receiver d
//	@param ctx
//	@param dir
//	@return vo.ILink
//	@return error
func (d *Local) Link(ctx context.Context, dir vo.IDir) (vo.ILink, error) {
	return nil, nil
}

// Remove 删除文件
//
//	@receiver d
//	@param ctx
//	@param dir
//	@return error
func (d *Local) Remove(ctx context.Context, dir vo.IDir) error {

	return nil
}
