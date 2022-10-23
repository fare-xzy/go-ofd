package main

import (
	"encoding/xml"
	"go-ofd/util"
	"os"
	"path/filepath"
	"testing"
)

var (
	ofdXmlContent string
)

func init() {
	fileBytes, _ := os.ReadFile(filepath.Join(util.Pwd, "doc", "OFD.xml"))
	ofdXmlContent = string(fileBytes)
}
func TestOFD(t *testing.T) {
	var ofd OFD
	if err := xml.Unmarshal([]byte(ofdXmlContent), &ofd); err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", ofd)
	}
}
