package fs

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type FileInfo struct {
	Name string    `json:"name"`
	Abs  string    `json:"abs"`
	Size int64     `json:"size"`
	Date time.Time `json:"date"`
	Dir  bool      `json:"dir"`
}

func NewFileInfo(info os.FileInfo, dir string) FileInfo {
	if dir == "/" {
		dir = ""
	} else if len(dir) > 0 {
		if dir[len(dir)-1:] == "/" {
			dir = dir[:len(dir)-1]
		}

		if dir[0:1] != "/" {
			dir = "/" + dir
		}
	}

	result := FileInfo{
		Name: info.Name(),
		Abs: fmt.Sprintf("%s/%s", dir, info.Name()),
		Size: info.Size(),
		Date: info.ModTime(),
		Dir:  info.IsDir(),
	}

	return result
}

type FileInfos []FileInfo

func (infos FileInfos) Len() int      { return len(infos) }
func (infos FileInfos) Swap(i, j int) { infos[i], infos[j] = infos[j], infos[i] }
func (infos FileInfos) Less(i, j int) bool {
	return strings.Compare(infos[i].Abs, infos[j].Abs) == -1
}

func NewFileInfos(infos []os.FileInfo, dir string) FileInfos {
	var result FileInfos

	for _, info := range infos {
		result = append(result, NewFileInfo(info, dir))
	}

	return result
}