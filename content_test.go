package main

import (
	"go-ofd/util"
	"os"
	"path/filepath"
)

var (
	contentXml string
)

func init() {
	fileBytes, _ := os.ReadFile(filepath.Join(util.Pwd, "doc", "OFD.xml"))
	contentXml = string(fileBytes)
}
