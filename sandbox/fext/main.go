package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	extType := flag.String("type", "go", "extension type to locate")
	flag.Parse()
	println("Locating files with extension", *extType)

	// walk all files in directory
	searchDir, err := os.Getwd()
	if err != nil {
		println(err)
		return
	}

	fileList := []string{}
	err = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+*extType {
			fileList = append(fileList, path)
		}

		return nil
	})
	if err != nil {
		println(err)
		return
	}

	for _, file := range fileList {
		fmt.Println(file)
	}
	println("Done!")
}
