package storage

import (
	"runtime"
	"testing"

	"github.com/shirou/gopsutil/disk"
)

func TestDiskUsage(t *testing.T) {
	path := "./"
	if runtime.GOOS == "windows" {
		path = "C:"
	}
	v, err := disk.Usage(path)
	if err != nil {
		t.Errorf("error %v", err)
	}
	if v.Path != path {
		t.Errorf("error %v", err)
	}

}
