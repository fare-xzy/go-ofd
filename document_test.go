package main

import (
	"go-ofd/util"
	"os"
	"path/filepath"
)

var (
	documentXml string
)

func init() {
	fileBytes, _ := os.ReadFile(filepath.Join(util.Pwd, "doc", "Document.xml"))
	documentXml = string(fileBytes)
}
