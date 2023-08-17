package fileutil

import (
	"fmt"
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
func GetFileExt(fstr string) (fext string) {
	return path.Ext(fstr)
}

// Check file or path exist
func CheckExist(fstr string) (fexist bool) {
	_, err := os.Stat(fstr)
	return !os.IsNotExist(err)
	// if os.IsNotExist(err) {
	// 	return false
	// }
	// return true
}

// Check file or path Permission
func CheckPermission(fsrc string) (fpermission bool) {
	_, err := os.Stat(fsrc)
	return !os.IsPermission(err)
}

// mkdir folder if this folder not exist
func MakeDirNotExist(pstr string) error {

	if !CheckExist(pstr) {
		if CheckPermission(pstr) {
			err := os.MkdirAll(pstr, 0766)
			return err
		} else {
			return fmt.Errorf("permissions error")
		}

	}
	return nil

}
