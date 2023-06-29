package fileutil

import (
	"bufio"
	"os"
	"testing"
)

func TestGetFileSize(t *testing.T) {
	// TODO
	fpathstr := "/tmp/getsizetest.log"
	// ftsize := 20
	file, err := os.Create(fpathstr)
	if err != nil {
		t.Errorf("Open file error,Msg:%v", err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	for i := 0; i < 100000; i++ {
		write.WriteString(fpathstr)
	}
	write.Flush()

	fsize, err := GetFileSize(file)
	if err != nil {
		t.Errorf("Get file size err:MSG:%v", err)
	}
	t.Errorf("%d", fsize)
	os.Remove(fpathstr)

}

func TestGetFileExt(t *testing.T) {
	// TODO
	fstr := "abc.txt"
	fext := ".txt"
	fext2 := GetFileExt(fstr)
	if fext != fext2 {
		t.Errorf(" GetFileExt(%v) is not %v", fstr, fext)
	}

}
