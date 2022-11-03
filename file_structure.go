package main

type OFDFile struct {
	OwnPath     string
	DocFile     DocFile
	OFDFilePath string
}

type DocFile struct {
	OwnPath           string
	Index             int
	PagesFile         PagesFile
	ResFile           ResFile
	SignsFile         SignsFile
	TagsFile          TagsFile
	DocumentFilePath  string
	PublicResFilePath string
}

type PagesFile struct {
	OwnPath  string
	PageFile []PageFile
}

type PageFile struct {
	OwnPath         string
	Index           int
	ContentFilePath string
}

type ResFile struct {
	OwnPath string
}

type SignsFile struct {
	OwnPath            string
	SignFile           []SignFile
	SignaturesFilePath string
}

type SignFile struct {
	OwnPath             string
	Index               int
	SealFilePath        string
	SignatureFilePath   string
	SignedValueFilePath string
}

type TagsFile struct {
	OwnPath            string
	CustomTagFilePath  string
	CustomTagsFilePath string
}
