package baidu_netdisk

type TokenErrResp struct {
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}

type File struct {
	//TkbindId     int    `json:"tkbind_id"`
	//OwnerType    int    `json:"owner_type"`
	//Category     int    `json:"category"`
	//RealCategory string `json:"real_category"`
	FsId        int64 `json:"fs_id"`
	ServerMtime int64 `json:"server_mtime"`
	//OperId      int   `json:"oper_id"`
	//ServerCtime int   `json:"server_ctime"`
	Thumbs struct {
		//Icon string `json:"icon"`
		Url3 string `json:"url3"`
		//Url2 string `json:"url2"`
		//Url1 string `json:"url1"`
	} `json:"thumbs"`
	//Wpfile         int    `json:"wpfile"`
	//LocalMtime     int    `json:"local_mtime"`
	Size int64 `json:"size"`
	//ExtentTinyint7 int    `json:"extent_tinyint7"`
	Path string `json:"path"`
	//Share          int    `json:"share"`
	//ServerAtime    int    `json:"server_atime"`
	//Pl             int    `json:"pl"`
	//LocalCtime     int    `json:"local_ctime"`
	ServerFilename string `json:"server_filename"`
	//Md5            string `json:"md5"`
	//OwnerId        int    `json:"owner_id"`
	//Unlist int `json:"unlist"`
	Isdir int `json:"isdir"`
}

// func fileToObj(f File) *model.ObjThumb {
// 	return &model.ObjThumb{
// 		Object: model.Object{
// 			ID:       strconv.FormatInt(f.FsId, 10),
// 			Name:     f.ServerFilename,
// 			Size:     f.Size,
// 			Modified: time.Unix(f.ServerMtime, 0),
// 			IsFolder: f.Isdir == 1,
// 		},
// 		Thumbnail: model.Thumbnail{Thumbnail: f.Thumbs.Url3},
// 	}
// }

type ListResp struct {
	Errno     int    `json:"errno"`
	GuidInfo  string `json:"guid_info"`
	List      []File `json:"list"`
	RequestId int64  `json:"request_id"`
	Guid      int    `json:"guid"`
}

type DownloadResp struct {
	Errmsg string `json:"errmsg"`
	Errno  int    `json:"errno"`
	List   []struct {
		Dlink string `json:"dlink"`
	} `json:"list"`
	//Names struct {
	//} `json:"names"`
	RequestId string `json:"request_id"`
}

type DownloadResp2 struct {
	Errno int `json:"errno"`
	Info  []struct {
		Dlink string `json:"dlink"`
	} `json:"info"`
	RequestID int64 `json:"request_id"`
}

type PrecreateResp struct {
	Path       string `json:"path"`
	Uploadid   string `json:"uploadid"`
	ReturnType int    `json:"return_type"`
	BlockList  []int  `json:"block_list"`
	Errno      int    `json:"errno"`
	RequestId  int64  `json:"request_id"`
}
