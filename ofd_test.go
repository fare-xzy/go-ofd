package main

import (
	"encoding/xml"
	"testing"
)

const ofdXmlContent = `<?xml version="1.0" encoding="UTF-8" ?>
<ofd:OFD xmlns:ofd="http://www.ofdspec.org/2016" DocType="OFD" Version="1.1">
    <ofd:DocBody>
        <ofd:DocInfo>
            <ofd:DocID>966aa7300b1a11eca4a663f264a663f2</ofd:DocID>
            <ofd:CreationDate>2021-09-01</ofd:CreationDate>
            <ofd:ModDate>2021-09-01T07:48:56+08:00</ofd:ModDate>
            <ofd:Creator>suwell-pdf2ofd</ofd:Creator>
            <ofd:CreatorVersion>1.0.21.0203</ofd:CreatorVersion>
            <ofd:CustomDatas>
                <ofd:CustomData Name="VersionDetail">13915-Feb 3 2021, 16:57:48</ofd:CustomData>
            </ofd:CustomDatas>
        </ofd:DocInfo>
        <ofd:DocRoot>Doc_0/Document.xml</ofd:DocRoot>
        <ofd:Signatures>Doc_0/Signs/Signatures.xml</ofd:Signatures>
    </ofd:DocBody>
</ofd:OFD>
`

func TestOFD(t *testing.T) {
	var ofd OFD
	if err := xml.Unmarshal([]byte(ofdXmlContent), &ofd); err != nil {
		t.Logf("%v", err)
	} else {
		t.Logf("%v", ofd)
	}
}
