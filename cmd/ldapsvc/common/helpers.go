package common

import "os"

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}
