package baidu_netdisk_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"testing"
)

func TestSlick(t *testing.T) {
	content, err := os.ReadFile("/home/zzk/下载/go1.20.2.linux-amd64.tar.gz")
	steam_size := int64(len(content))
	if err != nil {
		t.Error(err)
	}
	// compute block md5
	var Default int64 = 4 * 1024 * 1024
	var SliceSize int64 = 256 * 1024
	slice_md5 := ""

	count := int64(math.Ceil(float64(steam_size) / float64(Default)))
	content_md5 := md5.New()
	if _, err := content_md5.Write(content); err != nil {
		t.Fatal(err)
	}

	content_md5_str := hex.EncodeToString(content_md5.Sum(nil))

	fmt.Println(content_md5_str)
	block_md5_list := make([]string, 0)

	for i := int64(0); i < int64(steam_size); i += Default {
		var byteData []byte

		if i+Default >= steam_size {
			byteData = content[i:]
		} else {
			byteData = content[i : i+Default]
		}

		h := md5.New()
		h.Write(byteData)
		block_md5_list = append(block_md5_list, hex.EncodeToString(h.Sum(nil)))
		// block_md5_list = append(block_md5_list, fmt.Sprintf("\"%s\"", hex.EncodeToString(h.Sum(nil))))
		t.Log(len(byteData))

	}
	if steam_size <= SliceSize {
		slice_md5 = content_md5_str
	} else {
		sliceData := content[0:SliceSize]
		h := md5.New()
		h.Write(sliceData)
		slice_md5 = hex.EncodeToString(h.Sum(nil))
	}

	t.Log(slice_md5)
	for i := int64(0); i < count; i++ {
		var byteData []byte

		block := int64(i)
		if Default >= int64(len(content[block*Default:])) {
			byteData = content[block*Default:]
		} else {
			byteData = content[block*Default : (block+1)*Default]
		}
		h := md5.New()
		h.Write(byteData)
		md5_str := hex.EncodeToString(h.Sum(nil))
		if md5_str != block_md5_list[i] {
			t.Fatalf("origin %s   new %s \n", md5_str, block_md5_list)
		}
	}
}
