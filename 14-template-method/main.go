package main

import "fmt"

type Downloader interface {
	Download(uri string)
}

type Template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func NewTemplate(impl implement) *Template {
	return &Template{
		implement: impl,
	}
}

func (t *Template) Download(uri string) {
	t.uri = uri
	fmt.Println("start downloading")
	t.implement.download() // 父类持有子类的引用，并将实现延迟到子类
	t.implement.save()
	fmt.Println("downloading end")
}

func (t *Template) save() {
	fmt.Println("default save") // 提供一个默认的save方法
}

type FtpDownloader struct {
	*Template
}

func (f *FtpDownloader) download() {
	fmt.Println("ftp downloading")
}

func (f *FtpDownloader) save() {
	fmt.Println("ftp saving")
}

func NewFtpDownloader() Downloader {
	downloader := &FtpDownloader{}
	template := NewTemplate(downloader)
	downloader.Template = template
	return downloader
}

type HttpDownloader struct {
	*Template
}

func (h *HttpDownloader) download() {
	fmt.Println("http downloading")
}

func (h *HttpDownloader) save() {
	fmt.Println("http save")
}

func NewHttpDownloader() Downloader {
	downloader := &HttpDownloader{}
	template := NewTemplate(downloader)
	downloader.Template = template
	return downloader
}

func main() {
	downloader := NewFtpDownloader()
	downloader.Download("ftp://gxy.com")

	downloader = NewHttpDownloader()
	downloader.Download("http://gxy.com")
}

