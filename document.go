package main

import (
	"encoding/xml"
	"time"
)

type Document struct {
	XMLName      xml.Name     `xml:"Document"`
	Text         string       `xml:",chardata"`
	Ofd          string       `xml:"ofd,attr"`
	CommonData   CommonData   `xml:"CommonData"`
	Pages        Pages        `xml:"Pages"`
	Outlines     string       `xml:"Outlines"`
	Permissions  Permissions  `xml:"Permissions"`
	Actions      Actions      `xml:"Actions"`
	VPreferences VPreferences `xml:"VPreferences"`
	Bookmarks    Bookmarks    `xml:"Bookmarks"`
	Attachments  string       `xml:"Attachments"`
	Annotations  string       `xml:"Annotations"`
	CustomTags   string       `xml:"CustomTags"`
	Extensions   string       `xml:"Extensions"`
}

type CommonData struct {
	Text         string   `xml:",chardata"`
	MaxUnitID    string   `xml:"MaxUnitID"`
	PageArea     PageArea `xml:"PageArea"`
	PublicRes    string   `xml:"PublicRes"`
	DocumentRes  string   `xml:"DocumentRes"`
	TemplatePage string   `xml:"TemplatePage"`
	DefaultCS    string   `xml:"DefaultCS"`
}

type PageArea struct {
	Text           string `xml:",chardata"`
	PhysicalBox    string `xml:"PhysicalBox"`
	ApplicationBox string `xml:"ApplicationBox"`
	ContentBox     string `xml:"ContentBox"`
	BleedBox       string `xml:"BleedBox"`
}

type Pages struct {
	Text string `xml:",chardata"`
	Page []Page `xml:"Page"`
}

type Page struct {
	Text     string   `xml:",chardata"`
	ID       string   `xml:"ID,attr"`
	BaseLoc  string   `xml:"BaseLoc,attr"`
	Area     PageArea `xml:"Area"`
	Template Template `xml:"Template"`
	ZOrder   string   `xml:"ZOrder"`
	PageRes  string   `xml:"PageRes"`
	Content  string   `xml:"Content"`
	Layer    string   `xml:"Layer"`
	Actions  Actions
}

type Template struct {
	Text       string `xml:",chardata"`
	TemplateID string `xml:"TemplateID,attr"`
	ZOrder     string `xml:"ZOrder,attr"`
}

type Permissions struct {
	Text       string       `xml:",chardata"`
	Permission []Permission `xml:"Permission"`
}

type Permission struct {
	Text        string      `xml:",chardata"`
	Edit        bool        `xml:"Edit"`
	Annot       bool        `xml:"Annot"`
	Export      bool        `xml:"Export"`
	Signature   bool        `xml:"Signature"`
	Watermark   bool        `xml:"Watermark"`
	PrintScreen bool        `xml:"PrintScreen"`
	Print       Print       `xml:"Print"`
	ValidPeriod ValidPeriod `xml:"ValidPeriod"`
}

type Print struct {
	Text      string `xml:",chardata"`
	Printable bool   `xml:"Printable,attr"`
	Copies    int    `xml:"Copies,attr"`
}

type ValidPeriod struct {
	Text      string    `xml:",chardata"`
	StartDate time.Time `xml:"StartDate,attr"`
	EndDate   time.Time `xml:"EndDate,attr"`
}

type Actions struct {
	Text   string   `xml:",chardata"`
	Action []Action `xml:"Action"`
}
type Action struct {
	Text string `xml:",chardata"`
}

type VPreferences struct {
	Text          string  `xml:",chardata"`
	PageMode      string  `xml:"PageMode"`
	PageLayout    string  `xml:"PageLayout"`
	TabDisplay    string  `xml:"TabDisplay"`
	HideToolbar   bool    `xml:"HideToolbar"`
	HideMenubar   bool    `xml:"HideMenubar"`
	HideWindowsUI bool    `xml:"HideWindowsUI"`
	ZoomMode      string  `xml:"ZoomMode"`
	Zoom          float64 `xml:"Zoom"`
}

type Bookmarks struct {
	Text     string     `xml:",chardata"`
	Bookmark []Bookmark `xml:"bookmark"`
}

type Bookmark struct {
	Text string `xml:",chardata"`
	Name string `xml:"Name,attr"`
	Dest string `xml:"Dest"`
}
