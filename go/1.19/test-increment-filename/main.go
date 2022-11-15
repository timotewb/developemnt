package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func duplicateFile(f string, i int) (string, error) {
	/*
		i = iteration, number of duplicates, sued for suffix
		f = file name, path and extension - cannot have suffix
		x = extension including leading period
		n = file name excluding path and extension
		nx = fiel name including extension, excluding path
	*/
	x := filepath.Ext(f)
	nx := filepath.Base(f)
	n := strings.ReplaceAll(filepath.Base(f), x, "")
	p := strings.ReplaceAll(f, nx, "")
	i_str := fmt.Sprintf("%04d", i)

	newF := filepath.Join(p, n+"_"+i_str+x)

	if _, err := os.Stat(newF); err == nil {
		newF, err := duplicateFile(f, i+1)
		if err != nil {
			return "", fmt.Errorf("%s", err)
		} else {
			return newF, nil
		}
	} else if errors.Is(err, os.ErrNotExist) {
		return newF, nil
	} else {
		return "", fmt.Errorf("file does or doesnot exist ")
	}
}

func statTimes(name string) (atime, mtime, ctime time.Time, err error) {
	fi, err := os.Stat(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	mtime = fi.ModTime()
	stat := fi.Sys().(*syscall.Stat_t)
	atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	ctime = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	return
}

func main() {

	df, err := duplicateFile("/home/timotewb/development/go/1.19/test-increment-filename/file_name.txt", 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(df)

	atime, mtime, ctime, err := statTimes("/home/timotewb/development/go/1.19/test-increment-filename/file_name_0001.txt")
	fmt.Println(atime, mtime, ctime, err)

}
