package dir

import (
	"os"
	"path/filepath"
	"regexp"
)

type FileInfo struct {
	FileName string
	FileSize int64
	FileSha1 string
}

func WalkDir(dirPath, filter string) (files []FileInfo, err error) {
	files = make([]FileInfo, 0, 3000)
	//filter = strings.ToUpper(filter)
	walkFunc := func(filename string, fi os.FileInfo, err error) error {
		/* 根据filter过滤文件和目录 */
		match, _ := regexp.MatchString(filter, filename)
		if filter == "" || true != match {
			if fi.IsDir() {
				return nil
			}
			fileInfo := FileInfo{filename, fi.Size(), ""}
			files = append(files, fileInfo)
		}

		return nil
	}
	err = filepath.Walk(dirPath, walkFunc)
	return files, err
}
