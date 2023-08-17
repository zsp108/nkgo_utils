package fileutil

import (
	"fmt"
	"os"
	"testing"
)

var (
	fpath    string = "/tmp/file_test"
	fpathstr string = "./getfilesize_test.md"
)

func TestMakeDirNotExist(t *testing.T) {
	err := MakeDirNotExist(fpath)
	if err != nil {
		t.Errorf("t:%v", err)
	}
}

func TestGetFileSize(t *testing.T) {

	ftsize := 137

	// file, err := os.OpenFile(fpathstr, os.O_CREATE|os.O_APPEND, 0666)
	file, err := os.Open(fpathstr)
	if err != nil {
		t.Errorf("Open File Error,Msg:%e\n", err)
	}
	defer file.Close()

	fsize, err := GetFileSize(file)
	if err != nil {
		t.Errorf("Get file size err:MSG:%v\n", err)
	}

	if ftsize != fsize {
		t.Errorf("File size is %d, not %d", fsize, ftsize)
	}

	fmt.Printf("fsize: %v\n", fsize)
	// os.Remove(fpathstr)

}

func TestGetFileExt(t *testing.T) {

	fext := ".md"
	fext2 := GetFileExt(fpathstr)
	if fext != fext2 {
		t.Errorf(" GetFileExt(%v) is not %v", fpathstr, fext)
	}

}

func TestCheckExist(t *testing.T) {
	fpath1 := "./getfilesize_test"
	fpath2 := "./getfilesize_test.md"
	fexist1 := CheckExist(fpath1)
	if fexist1 {
		t.Error(fexist1)
	}

	fexist2 := CheckExist(fpath2)
	if !fexist2 {
		t.Error(fexist2)
	}

}

func TestCheckPermission(t *testing.T) {
	if !CheckPermission(fpathstr) {
		t.Errorf("is not permission %v", CheckPermission(fpathstr))
	}
}
