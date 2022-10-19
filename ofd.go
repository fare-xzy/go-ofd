package main

//OFD.xml 文档结构

import "encoding/xml"

type OFD struct {
	XMLName xml.Name  `xml:"OFD"`
	Text    string    `xml:",chardata"`
	Ofd     string    `xml:"ofd,attr"`
	DocType string    `xml:"DocType,attr"`
	Version string    `xml:"Version,attr"`
	DocBody []DocBody `xml:"DocBody"`
}

type DocBody struct {
	Text       string     `xml:",chardata"`
	DocInfo    DocInfo    `xml:"DocInfo"`
	DocRoot    DocRoot    `xml:"DocRoot"`
	Version    string     `xml:"Version"`
	Signatures Signatures `xml:"Signatures"`
}

type DocInfo struct {
	Text           string      `xml:",chardata"`
	DocID          string      `xml:"DocID"`
	Title          string      `xml:"Title"`
	Author         string      `xml:"Author"`
	Subject        string      `xml:"Subject"`
	Abstract       string      `xml:"Abstract"`
	CreationDate   string      `xml:"CreationDate"`
	ModDate        string      `xml:"ModDate"`
	DocUsage       string      `xml:"DocUsage"`
	Cover          string      `xml:"Cover"`
	Keywords       Keywords    `xml:"Keywords"`
	Creator        string      `xml:"Creator"`
	CreatorVersion string      `xml:"CreatorVersion"`
	CustomDatas    CustomDatas `xml:"CustomDatas"`
}

type DocRoot struct {
	Text string `xml:",chardata"`
}

type Signatures struct {
	Text string `xml:",chardata"`
}

type Keywords struct {
	Text    string    `xml:",chardata"`
	Keyword []Keyword `xml:"Keyword"`
}

type Keyword struct {
	Text string `xml:",chardata"`
}

type CustomDatas struct {
	Text       string       `xml:",chardata"`
	CustomData []CustomData `xml:"CustomData"`
}

type CustomData struct {
	Text string `xml:",chardata"`
	Name string `xml:"Name,attr"`
}
