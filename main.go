package main

import (
	"fmt"
	inLog "go-ofd/log"
	"go-ofd/util"
	"os"
	"path"
	"path/filepath"
	"time"
)

var (
	localTempFolder = filepath.Join(os.TempDir(), "ofd")
)

func main() {
	args := os.Args
	ofdFilePath := args[1]
	fmt.Println(len(args))
	fmt.Println(args[1])

	_, f := filepath.Split(ofdFilePath)
	// 获取文件名
	fileprefix := f[0 : len(f)-len(path.Ext(f))]

	unPackZipFolder := filepath.Join(localTempFolder, fileprefix)
	fmt.Println(unPackZipFolder)
	err := util.UnzipFile(ofdFilePath, unPackZipFolder)
	if err != nil {
		inLog.Errorf("Error , %+v", err)
	}
	time.Sleep(time.Duration(2) * time.Minute)
}
