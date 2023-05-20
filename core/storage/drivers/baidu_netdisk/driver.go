// package
package baidu_netdisk

// import (
// 	"bytes"
// 	"context"
// 	"crypto/md5"
// 	"encoding/hex"
// 	"fmt"
// 	"path"
// 	"strconv"
// 	"strings"

// 	"github.com/alist-org/alist/v3/pkg/utils"
// 	"github.com/pipikai/yun/common/logger"
// 	"github.com/pipikai/yun/common/util"
// 	"github.com/pipikai/yun/core/storage/drivers/vo"
// 	"github.com/pipikai/yun/core/storage/models"
// )

// // BaiduNetdisk
type BaiduNetdisk struct {
	Addition
	AccessToken string
}

// // Config
// //
// //	@receiver d
// //	@return vo.Config
// func (d *BaiduNetdisk) Config() vo.Config {
// 	return config
// }

// // GetAddition
// //
// //	@receiver d
// //	@return vo.Addition
// func (d *BaiduNetdisk) GetAddition() vo.Addition {
// 	return &d.Addition
// }

// // Init
// //
// //	@receiver d
// //	@param ctx
// //	@return error
// func (d *BaiduNetdisk) Init(ctx context.Context) error {
// 	return d.refreshToken()
// }

// // func (d *BaiduNetdisk) Drop(ctx context.Context) error {
// // 	return nil
// // }

// // func (d *BaiduNetdisk) List(ctx context.Context, dir model.Obj, args model.ListArgs) ([]model.Obj, error) {
// // 	files, err := d.getFiles(dir.GetPath())
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return utils.SliceConvert(files, func(src File) (model.Obj, error) {
// // 		return fileToObj(src), nil
// // 	})
// // }

// // Link
// //
// //	@receiver d
// //	@param ctx
// //	@param file
// //	@return *models.Link
// //	@return error
// func (d *BaiduNetdisk) Link(ctx context.Context, file vo.IDir) (*models.Link, error) {
// 	return d.linkOfficial(file)
// }

// // Remove
// //
// //	@receiver d
// //	@param ctx
// //	@param file
// //	@return error
// func (d *BaiduNetdisk) Remove(ctx context.Context, file vo.IDir) error {
// 	// data := []string{file.GetPath()}
// 	// _, err := d.manage("delete", data)
// 	return nil
// }

// // Upload
// //
// //	@receiver d
// //	@param ctx
// //	@param stream
// //	@return *models.Link
// //	@return error
// func (d *BaiduNetdisk) Upload(ctx context.Context, stream vo.IStreamFile) (*models.Link, error) {
// 	logger.Logger.Info(stream.GetName(), stream.GetSize(), len(stream.GetContent()))

// 	var Default int64 = 4 * 1024 * 1024
// 	var SliceSize int64 = 256 * 1024
// 	// cal md5
// 	h1 := md5.New()
// 	h2 := md5.New()
// 	block_list := make([]string, 0)
// 	content_md5 := ""
// 	slice_md5 := ""
// 	content := stream.GetContent()
// 	for i := int64(0); i < stream.GetSize(); i += Default {
// 		var byteData []byte

// 		if i+Default >= stream.GetSize() {
// 			byteData = content[i:]
// 		} else {
// 			byteData = content[i : i+Default]
// 		}
// 		h1.Write(byteData)
// 		h2.Write(byteData)
// 		block_list = append(block_list, fmt.Sprintf("\"%s\"", hex.EncodeToString(h2.Sum(nil))))
// 		h2.Reset()
// 	}

// 	content_md5 = hex.EncodeToString(h1.Sum(nil))

// 	if stream.GetSize() <= SliceSize {
// 		slice_md5 = content_md5
// 	} else {
// 		sliceData := content[:SliceSize]
// 		h2.Write(sliceData)
// 		slice_md5 = hex.EncodeToString(h2.Sum(nil))
// 	}
// 	rawPath := path.Join(DefaultDir, stream.GetName())
// 	path := encodeURIComponent(rawPath)
// 	block_list_str := fmt.Sprintf("[%s]", strings.Join(block_list, ","))
// 	data := fmt.Sprintf("path=%s&size=%d&isdir=0&autoinit=1&block_list=%s&content-md5=%s&slice-md5=%s",
// 		path, stream.GetSize(),
// 		block_list_str,
// 		content_md5, slice_md5)
// 	params := map[string]string{
// 		"method": "precreate",
// 	}
// 	var precreateResp PrecreateResp
// 	res, err := d.post("/xpan/file", params, data, &precreateResp)
// 	logger.Logger.Debug("res", string(res), "  ", precreateResp)

// 	if err != nil {
// 		logger.Logger.Error(err)

// 		return nil, err
// 	}
// 	if precreateResp.ReturnType == 2 {
// 		link, err := d.Link(ctx, models.File{
// 			FileMeta: models.FileMeta{},
// 			ID:       util.Json.Get(res, "info", "fs_id").ToString(),
// 			Name:     stream.GetName(),
// 		})
// 		return link, err
// 	}
// 	params = map[string]string{
// 		"method":       "upload",
// 		"access_token": d.AccessToken,
// 		"type":         "tmpfile",
// 		"path":         path,
// 		"uploadid":     precreateResp.Uploadid,
// 	}

// 	// left = stream.GetSize()

// 	upload_size := 0
// 	for _, partseq := range precreateResp.BlockList {
// 		if utils.IsCanceled(ctx) {
// 			logger.Logger.Error(err)

// 			return nil, ctx.Err()
// 		}
// 		var byteData []byte

// 		block := int64(partseq)
// 		if Default >= int64(len(content[block*Default:])) {
// 			byteData = content[block*Default:]
// 		} else {
// 			byteData = content[block*Default : (block+1)*Default]
// 		}

// 		upload_size += len(byteData)
// 		u := "https://d.pcs.baidu.com/rest/2.0/pcs/superfile2"
// 		params["partseq"] = strconv.Itoa(partseq)
// 		res, err := util.RestyClient.R().
// 			SetContext(ctx).
// 			SetQueryParams(params).
// 			SetFileReader("file", stream.GetName(), bytes.NewReader(byteData)).
// 			Post(u)
// 		if err != nil {
// 			logger.Logger.Error(err)

// 			return nil, err
// 		}
// 		logger.Logger.Debugln(res.String())

// 	}
// 	logger.Logger.Info(upload_size)
// 	logger.Logger.Info(block_list_str)
// 	create_res, err := d.create(rawPath, stream.GetSize(), 0, precreateResp.Uploadid, block_list_str)
// 	logger.Logger.Info(string(create_res))

// 	if err != nil {
// 		logger.Logger.Error(err)
// 		return nil, err
// 	}

// 	link, err := d.Link(ctx, models.File{
// 		ID:       util.Json.Get(create_res, "fs_id").ToString(),
// 		FileMeta: models.FileMeta{},
// 		Name:     stream.GetName(),
// 	})

// 	return link, err
// }

// // MakeDir
// //
// //	@receiver d
// //	@param ctx
// //	@param parentDir
// //	@param dirName
// //	@return error
// func (d *BaiduNetdisk) MakeDir(ctx context.Context, parentDir vo.IDir, dirName string) error {
// 	// _, err := d.create(path.Join(parentDir.GetPath(), dirName), 0, 1, "", "")
// 	return nil
// }

// // func (d *BaiduNetdisk) Move(ctx context.Context, srcObj, dstDir model.Obj) error {
// // 	data := []base.Json{
// // 		{
// // 			"path":    srcObj.GetPath(),
// // 			"dest":    dstDir.GetPath(),
// // 			"newname": srcObj.GetName(),
// // 		},
// // 	}
// // 	_, err := d.manage("move", data)
// // 	return err
// // }

// // func (d *BaiduNetdisk) Rename(ctx context.Context, srcObj model.Obj, newName string) error {
// // 	data := []base.Json{
// // 		{
// // 			"path":    srcObj.GetPath(),
// // 			"newname": newName,
// // 		},
// // 	}
// // 	_, err := d.manage("rename", data)
// // 	return err
// // }

// // func (d *BaiduNetdisk) Copy(ctx context.Context, srcObj, dstDir model.Obj) error {
// // 	data := []base.Json{
// // 		{
// // 			"path":    srcObj.GetPath(),
// // 			"dest":    dstDir.GetPath(),
// // 			"newname": srcObj.GetName(),
// // 		},
// // 	}
// // 	_, err := d.manage("copy", data)
// // 	return err
// // }
