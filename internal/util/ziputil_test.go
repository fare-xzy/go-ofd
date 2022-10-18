package util

import "testing"

func TestGzipFile(t *testing.T) {
	GzipFile("C:\\Users\\12691\\Desktop\\test", "C:\\Users\\12691\\Desktop\\1.tar.gz")
}

func TestUnGzipFile(t *testing.T) {
	UnGzipFile("C:\\Users\\12691\\Desktop\\1.tar.gz", "C:\\Users\\12691\\Desktop\\1tar")
}

func TestZipFile(t *testing.T) {
	ZipFile("C:\\Users\\12691\\Desktop\\test", "C:\\Users\\12691\\Desktop\\1.zip")
}

func TestUnzipFile(t *testing.T) {
	UnzipFile("C:\\Users\\12691\\Desktop\\1.zip", "C:\\Users\\12691\\Desktop\\1zip")
}
