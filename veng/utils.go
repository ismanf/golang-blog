package veng

import (
	"strings"
	"os"
	"path/filepath"
	"fmt"
)

func walkDirFiles(dir string) (map[string]string, error) {
	files := make(map[string]string, 0) 
	
	err := filepath.Walk(fmt.Sprintf("./%s", dir), func(path string, fi os.FileInfo, errWF error) error {
		if errWF != nil {
			return errWF
		}
		if !fi.IsDir() {
			files[fi.Name()] = path
		}
		return nil
	})

	return files, err
}

func readDirHtmlFiles(dir string) ([]string, error) {
	return filepath.Glob(dir + "/*.html")
}



func getFileNameFromPath(path string) string {
	return filepath.Base(path)
}

func getFileNameWithoutExt(fileName string) string {
	arr := strings.Split(fileName, ".")
	return arr[0]
}