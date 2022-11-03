package util

import (
	"fmt"
	"github.com/beevik/etree"
)

func ParseXml(path string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		panic(err)
	}
	root := doc.SelectElement("ofd:OFD")
	root1 := doc.SelectElement("OFD")
	fmt.Println("ROOT element:", root.Tag)
	fmt.Println("ROOT element:", root1.Tag)

	for _, book := range root.SelectElements("book") {
		fmt.Println("CHILD element:", book.Tag)
		if title := book.SelectElement("title"); title != nil {
			lang := title.SelectAttrValue("lang", "unknown")
			fmt.Printf("  TITLE: %s (%s)\n", title.Text(), lang)
		}
		for _, attr := range book.Attr {
			fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
		}
	}
}
