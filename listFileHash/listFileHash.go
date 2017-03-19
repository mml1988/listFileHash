package main

import (
	"dir"
	"flag"
	"fmt"
	"sha1"
)

//http://www.yiibai.com/go/golang-reading-files.html

var root *string = flag.String("r", "/root", "Root directory for search")
var filter *string = flag.String("f", "", "Specify which files to skip ")
var workers *int = flag.Int("w", 8, "Specify how many workers to work")

func main() {
	flag.Parse()

	if root != nil {
		fmt.Println("root =", *root, "filter =", *filter, "workers =", *workers)
		files, _ := dir.WalkDir(*root, *filter)
		for _, file := range files {
			var sha1Str, _ = sha1.SHA1File(file.FileName)
			fmt.Printf("%s,%s,%d\r\n", file.FileName, sha1Str, file.FileSize)
		}
	}
}
