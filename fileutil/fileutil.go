package fileutil

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// Get file size
func GetFileSize(f multipart.File) (filesize int, err error) {
	fs, err := ioutil.ReadAll(f)
	filesize = len(fs)
	// fmt.Println(filesize)
	return
}

// Get file name extension
func GetFileExt(f string) (fext string) {
	fext = path.Ext(f)
	return
}

// Check file or path exist
func CheckNotExist(f string) (fexist bool) {
	_, err := os.Stat(f)
	fexist = os.IsNotExist(err)
	return
}
