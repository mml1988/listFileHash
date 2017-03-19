package main

import (
	"dir"
	"flag"
	"fmt"
	"sha1"
)

var root *string = flag.String("r", "G:\\MML", "Root directory for search")
var filter *string = flag.String("f", "*", "Specify which files to skip ")
var workers *int = flag.Int("w", 16, "Specify how many workers to work")

func worker(FirstIndex int, LastIndex int, files []dir.FileInfo, result chan bool) {
	for i := FirstIndex; i < LastIndex; i++ {
		files[i].FileSha1, _ = sha1.SHA1File(files[i].FileName)
	}
	result <- true
	fmt.Printf("files from %d to %d have down.\r\n", FirstIndex, LastIndex)
}
func main() {
	flag.Parse()

	if root != nil {
		//fmt.Println("root =", *root, "filter =", *filter)
		files, _ := dir.WalkDir(*root, *filter)

		results := make([]chan bool, *workers)
		for i := 0; i < *workers; i++ {
			results[i] = make(chan bool)
			go worker(i*(len(files)/(*workers)), (i+1)*(len(files)/(*workers)), files, results[i])
		}
		for _, result := range results {
			<-result
		}
		for _, fileInfo := range files {
			fmt.Printf("%s,%s,%d\r\n", fileInfo.FileName, fileInfo.FileSha1, fileInfo.FileSize)
		}
	}
}
