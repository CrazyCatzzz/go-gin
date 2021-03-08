package util

import (
	"os"
)

func FileExist(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.Mode().IsRegular()
}

func DirExist(dir string) bool {
	fi, err := os.Stat(dir)
	return err == nil && fi.Mode().IsDir()
}
