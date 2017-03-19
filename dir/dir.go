package dir

import (
	"os"
	"path/filepath"
)

type FileInfo struct {
	FileName string
	FileSize int64
}

func WalkDir(dirPath, filter string) (files []FileInfo, err error) {
	files = make([]FileInfo, 0, 30)
	//filter = strings.ToUpper(filter)
	walkFunc := func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		/* 根据filter过滤文件和目录 */
		fileInfo := FileInfo{filename, fi.Size()}
		files = append(files, fileInfo)
		return nil
	}
	err = filepath.Walk(dirPath, walkFunc)
	return files, err
}
