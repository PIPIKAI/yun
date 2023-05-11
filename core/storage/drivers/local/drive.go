// driver local , store file in local disk
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

// PreUpload 上传文件之前校验文件
//
//	@receiver d
//	@param ctx
//	@param req
//	@return *pb.PreUploadReply
//	@return error
func (d *Local) PreUpload(ctx context.Context, req *pb.PreUploadRequest) (*pb.PreUploadReply, error) {
	block_status := make([]bool, 0)
	for _, v := range req.BlockMd5 {
		_, err := os.Stat(tempPath + v)
		block_status = append(block_status, !os.IsNotExist(err))
	}
	return &pb.PreUploadReply{
		Code:        1,
		Blockstatus: block_status,
	}, nil
}

// Upload 本地上传分块文件，就是将文件保存在本地目录
//
//	@receiver d
//	@param ctx
//	@param file
//	@return *pb.UploadReply
//	@return error
func (d *Local) Upload(ctx context.Context, file *pb.UploadRequest) (*pb.UploadReply, error) {

	hash := md5.New()
	_, err := hash.Write(file.RawData)
	if err != nil {
		return nil, err
	}
	md5String := hex.EncodeToString(hash.Sum(nil))

	if err := os.MkdirAll(tempPath, fs.FileMode(mkdirPerm)); err != nil {
		return nil, err
	}
	// _, err = os.Stat(tempPath + md5String)
	// if !os.IsNotExist(err) {
	// 	// 若文件已经存在
	// 	os.
	// 	return &pb.UploadReply{
	// 		Md5: md5String,
	// 	}, nil
	// }

	f, err := os.OpenFile(tempPath+md5String, os.O_CREATE|os.O_RDWR, fs.FileMode(mkdirPerm))
	if err != nil {
		return nil, err
	}
	_, err = f.Write(file.RawData)
	f.Close()
	if err != nil {
		return nil, err
	}
	return &pb.UploadReply{
		Md5: md5String,
	}, nil
}

// CreateFile 合并分块文件
//
//	@receiver d
//	@param ctx
//	@param file
//	@return *pb.MergeReply
//	@return error
func (d *Local) CreateFile(ctx context.Context, file *models.File) (*pb.MergeReply, error) {
	f, err := os.OpenFile(d.RootPath+"/"+file.GetPath(), os.O_CREATE|os.O_RDWR, fs.FileMode(mkdirPerm))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// toDo 这里可以考虑用协程
	for _, filename := range file.BlockMd5 {
		content, err := os.ReadFile(tempPath + filename)
		if err != nil {
			return nil, err
		}
		f.Write(content)
	}
	for _, filename := range file.BlockMd5 {
		err = os.Remove(tempPath + filename)
		if err != nil {
			return nil, err
		}
	}
	return &pb.MergeReply{
		Path: file.GetPath(),
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
