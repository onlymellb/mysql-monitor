package base

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	nsmap      = make(map[string]string)
	sfilepaths []string
)

func ListFunc(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}
	ok := strings.HasSuffix(path, ".sock")
	if ok {
		sfilepaths = append(sfilepaths, path)
		file := filepath.Base(path)
		if strings.HasPrefix(file, "mysql") {
			//fmt.Println("the socket file is", file)
			if len(file)-11 > 0 {
				appname := file[6 : len(file)-5]
				//fmt.Println("the appname is", appname)
				nsmap[appname] = path
				//fmt.Println("the app name is", appname)
			}
		}
	}
	return nil
}

func GetSockFileList(path string) (err error) {
	err = filepath.Walk(path, ListFunc)
	if err != nil {
		return
	}
	return
}

func GetAppSockMap(path string) (map[string]string, error) {
	err := GetSockFileList(path)
	if err != nil {
		return nil, err
	}
	return nsmap, err
}
