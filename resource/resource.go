// Code generated by go-bindata.
// sources:
// views/index.html
// views/static/favicon.ico
// views/static/xas.png
// config/seelog.xml
// config/app.ini
// DO NOT EDIT!

package resource

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _viewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8e\xbf\xce\xc3\x20\x0c\xc4\xf7\x48\x79\x07\x7f\xec\x5f\x58\x3b\x38\x2c\xfd\xb3\xb6\x03\x1d\x3a\xba\x80\x08\x12\xd0\x2a\x78\x68\x15\xe5\xdd\xab\x50\xba\xd8\xb2\xef\x7e\xba\xc3\xbf\xc3\x79\xaf\x6f\x97\x23\x4c\x9c\xa2\xea\x3b\xdc\x36\x44\xca\x7e\x14\x2e\x8b\xfa\x71\x64\x55\xdf\x01\x00\x60\x72\x4c\x60\x26\x9a\x8b\xe3\x51\x5c\xf5\xe9\x7f\x27\x7e\x1a\x07\x8e\x4e\xe9\x6d\xa2\xfc\x1e\x7d\x87\xb2\xe1\x78\x7f\xd8\x77\xb3\x2e\xcb\x60\x89\x69\x5d\x01\x43\xf2\x50\x66\x33\x8a\xc2\xc4\xc1\xc8\x17\x95\xe1\x99\x7d\x0d\x96\x0d\x41\x59\xcb\x7d\x02\x00\x00\xff\xff\x71\x1f\x93\x15\xac\x00\x00\x00")

func viewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsIndexHtml,
		"views/index.html",
	)
}

func viewsIndexHtml() (*asset, error) {
	bytes, err := viewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/index.html", size: 172, mode: os.FileMode(438), modTime: time.Unix(1534605757, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsStaticFaviconIco = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\xd7\x6f\x6c\x13\xe7\x1d\x07\xf0\x03\x52\x9b\x35\x19\x61\x83\x02\xcb\x5c\xe2\x32\x94\xb2\x68\x2d\x15\x83\x2d\xb4\x69\xe3\xb9\xc1\x1d\xe9\x36\xa6\x32\x8d\x95\xb5\xa5\xa8\x88\x12\x16\x52\xa7\x04\x37\x09\xa6\x49\x20\x09\x84\x24\x36\x4e\xec\x4b\xec\x74\x76\x8c\x88\x14\x2f\xaf\xaa\xda\xd4\xff\xb0\x4d\x11\xe9\xab\xbd\x42\xbc\xc8\x8b\x4a\xae\xaa\xde\xf3\xef\xbc\x05\x59\x51\xa9\xbc\xf0\x9b\xee\x62\x93\x18\x6c\x13\x51\x78\xa4\xaf\x22\x25\xce\xf3\xf9\xde\xdd\xf3\xf8\xee\x38\x6e\x19\xb7\x8c\x53\xab\xa5\x9f\x6a\x6e\x62\x35\xc7\xad\xe3\x38\x6e\x0b\xc7\x71\x6a\x8e\xe3\x34\xdc\xfc\xef\xe5\xb1\x9a\x7b\x64\x03\xf5\x28\xd4\xa8\xf7\x09\x9f\xe0\x1d\xf0\x09\x5e\x93\xfa\xd1\xcd\xfc\x00\xb7\x7d\xd5\x4a\x74\xba\xd8\x88\x7a\x94\xb3\xe8\xdc\x13\x20\x78\x4d\x52\x66\x05\xaf\xc9\x28\x78\x4d\x2b\x1f\xa7\x2d\xb4\xad\xd1\xa1\x8f\x4b\xa7\xd1\xe9\x62\x40\x3d\x4a\x40\xbd\x4a\x10\x7c\x96\x4c\x07\x29\xd3\x82\x77\x40\xf7\xa8\x5d\xd4\xa4\x52\xa1\x13\x1b\x26\x51\xeb\x5a\x40\x1f\x97\xc2\xbc\xff\x23\x40\xfd\xc5\x80\x2e\x0f\x82\xe0\xbb\xb0\xb8\x83\x94\x49\xc1\x6b\x52\xfd\x50\x57\xa8\xaf\x50\xa0\x86\x4d\x06\x41\xff\x74\x12\x9d\xf8\x19\xdc\xf5\xcf\xfc\x18\x70\x6f\x29\xe0\xb4\x9f\xab\xc3\xb7\x5e\x53\x52\xf0\x9a\x0c\x82\xdf\xaa\x78\x28\xfb\x50\xa5\x16\xbd\xff\xec\x4d\xd4\xb0\x09\x90\xfe\x69\xb8\xeb\xb7\xff\x04\x50\xcf\x1a\xc0\xe7\xb3\xfd\x9c\xe7\xe1\xf3\x21\x40\x7e\xeb\x4d\xc1\x6f\xd5\x2e\xd5\xc5\x47\x77\x95\xe1\x06\xcd\x38\x3a\xf4\x1c\xa0\xf7\x9f\x85\x2c\xbf\x6d\x1d\xe0\xce\xf5\x80\xf3\xf8\x59\x1d\xe6\xed\xc5\x19\x17\xfc\xd6\xb2\x42\x36\xa9\xaf\xd3\xe3\xfa\x5d\xb7\xc8\xb1\xdf\x01\xfe\xa0\x1a\xd0\x91\x2d\x0b\xbe\xa1\x0c\xd0\xa9\x9f\x03\x3e\x9d\xe5\xbb\xd1\xe5\x21\xf7\xbd\x1d\xee\x71\x17\xe7\x16\xf2\x5b\x1b\xf3\x1e\xfb\x65\xeb\x1c\xee\x3a\x08\xe4\xe8\x2b\x40\x9a\x5e\x06\x7c\xbc\x0a\xd0\xb1\xcd\x80\x9a\xca\x01\xb7\xaa\x01\x67\xfc\xee\x35\x37\x50\xef\x6a\x0d\x32\x95\x2c\x97\xaf\x95\xcf\xa2\x41\x3e\xcb\x0d\xd9\x0e\xf0\x80\x83\xc3\x80\x02\xb6\x7c\x1d\xe6\xf2\xfa\xc1\x91\x14\x0e\x3b\x00\xff\xab\x1b\x70\xeb\xef\x81\x1c\xaf\x02\xd2\xba\x03\xf0\x89\x8a\x8c\x9f\xc4\x9d\xeb\x9b\x71\xf7\x53\xf2\x9a\xc2\xa6\x62\x05\x1e\x28\xb9\x8a\xcf\xaf\xfa\x15\xf2\x5d\x50\x20\xbf\xad\x19\x07\xf9\xa4\xe4\x17\xe8\x90\x2a\xe8\x07\x47\x00\x07\x86\x41\xee\xf1\x49\x23\x10\xe3\x8b\x40\xda\x77\x00\x6e\x7d\xc6\x83\x8d\xaa\xac\x3d\x85\x4d\xc5\x45\xb8\xbf\x24\x85\xcf\x97\xa6\x70\xbf\xca\x4c\xfa\x7e\x51\x82\x03\xbc\x0a\x05\x78\x4f\x81\x0e\x0f\xf6\xd3\x21\x61\x07\x10\xbf\x15\xc8\x45\xbd\x25\xe7\xe7\x25\x7f\xa0\x24\x85\x4d\x1b\x80\xf4\x6f\x04\x72\x6e\xd3\xb7\xa4\xbb\xe2\x6f\xd2\xdf\x50\x80\xaf\xc5\x41\x7e\x3a\x47\x87\x25\xf9\xb2\x9d\x9d\x10\x09\x3b\x2a\xb2\x3e\x6f\x2e\x2d\xc2\x96\xa7\x52\xc4\x5c\x96\xf1\x81\x74\x57\x00\xe9\xac\xbc\x42\x4e\x3d\xbf\x05\xf9\x6d\x0a\x1c\xe0\x5b\x70\x70\x78\x76\x51\x87\x07\xfa\x39\xec\x4c\xbe\x23\x61\xc7\x59\x1c\x72\x3c\x29\x7d\x9e\x0c\xad\x2f\x22\x96\x0d\xf7\xfb\x1d\x95\x80\x8d\x5b\xbf\x27\x2d\xdb\x7a\x49\x4b\xd5\x93\x38\xc0\x6f\xc4\x01\xfe\xd3\x74\x87\x42\xfe\x5c\x01\xfb\x6e\x70\xd8\xf1\x35\x0e\xd9\xff\x44\x2e\x94\x14\x11\xcb\xfa\x9c\x3e\x31\x6e\x05\xd2\xf2\x6b\x20\xcd\xbf\xfd\x86\xe8\xab\xdf\x90\xe7\x0f\xf0\x75\x28\x60\xfb\x2a\xaf\x1f\xb2\xbf\x4b\x42\x76\xf2\xc0\x0e\x11\xa7\x1c\x1c\x72\x7c\x46\x72\x9d\xff\x6c\x1f\x88\xbe\x1a\x48\x43\xcd\xe7\xa4\x5e\xbb\x19\x07\xf8\xbc\xf7\x48\xfa\x8f\x3d\x2b\x49\xc8\xbe\x9a\x84\xec\xc3\x24\xec\xc8\x7d\x2e\x22\x4e\xa0\x51\x97\x1c\x12\x75\x7f\x4f\x2e\xac\x1d\xc2\xe6\xb2\xd4\x12\x7c\x20\xf5\xaf\x7e\x47\x0e\xeb\x3a\xf3\xfa\x47\xff\xf8\x15\x3d\x56\x57\x97\x3e\x17\xdb\x49\xd8\xf1\x65\x3e\x7b\x3e\xce\x14\xbd\x36\x59\x84\xcd\x65\xdb\x70\xff\xc6\xeb\x4b\xf0\x01\x1f\xd6\xe5\xfd\xfe\xa1\xc7\xf6\xa4\x68\xe3\x6e\xa0\xfa\xd7\x3e\xa5\xc7\x6b\x37\x92\xe0\xc8\x72\x12\xb2\x1f\xc6\x61\x47\xe2\x7e\x5b\x4e\x8a\x46\x5d\x45\xf2\x5a\xec\x53\x2f\x27\x67\x37\x1d\x22\x5d\x15\xac\x90\x4f\x0e\xeb\xf2\xae\x3f\xd6\x76\x20\x45\xf5\x7f\x00\xda\xa4\x03\xda\xac\x9d\xa5\x1f\xd5\xb4\xd1\xb6\x97\x14\x24\xe2\x5c\x4b\x23\xce\x51\x1a\x75\xce\xe5\xf3\x33\x83\x74\x55\xfc\x94\x74\x54\xda\x89\x71\xeb\xdc\x7d\xfe\x91\x57\x81\x1c\xd9\x9d\xd7\x17\x4f\x1e\x48\x31\xe3\x5b\x40\x3f\xdc\x2d\xf9\x40\x5b\x6a\x80\x76\x68\xa7\x69\x8f\xa6\x56\x3e\x3f\x11\x67\x15\x8d\x3a\xff\x5d\xc8\xcf\x0c\x3c\xd6\x60\xc0\x67\x6a\x81\x7c\xb8\x63\xc1\x6f\xd0\x01\x35\xbc\x59\xd0\x17\x4f\xbd\x0d\xac\xfd\x4d\xa0\x86\x5d\x40\x4f\xe9\x80\x9d\xae\x05\xd6\xa3\x01\x76\xbe\xda\xc3\x06\x76\xaa\x68\xe4\x9f\x45\x34\xe2\x3c\x4a\xa3\xce\xff\xe6\xf2\x69\xd4\xa5\xc2\xc1\x61\x0f\x96\xee\x43\xd2\x5e\x75\x36\x01\xd1\xbf\x0c\xa4\x51\x0b\xb4\x75\x1f\xb0\x25\xf8\x62\xfb\x7e\x10\xcf\xbe\x05\xac\xa3\x6e\xb1\x0f\x6c\xa0\x2a\x49\x2d\xdb\x9b\x99\x6d\xbb\x82\x46\x9c\xeb\x68\xd4\xe9\xa6\x11\x97\x22\xed\x2a\x68\xd4\xd5\x4c\xa3\xae\xa4\xbc\x37\xd3\xdf\xff\x72\x02\x3c\x50\xeb\x07\xc0\x0c\x7f\x29\xe8\xb3\x93\x07\xe6\x44\xe3\xdb\x20\x76\xbd\x0b\x89\xde\x77\x40\xec\xdb\x0f\xac\x4b\xb7\xe0\x9b\xaa\x80\x0d\x6e\x07\x66\x7d\xe1\x06\x1b\x79\x5e\xb3\x70\xcc\x4e\x0d\x8d\xba\x6e\x2c\x5e\x1b\x77\x3b\x84\x46\x80\xc4\xdc\x40\xbf\xb8\x04\xf4\xb3\x21\xa0\x1d\xef\xe5\x5d\xff\xe2\xc9\x03\x8d\x62\xe7\x7b\xb7\x12\xdd\x07\x65\x3f\xd1\xf7\x77\x10\x4d\x7f\x05\x76\x4e\x9b\xed\xdb\x5e\x00\x36\xf2\x1c\x50\xc7\x2f\xdd\x34\x3a\xe6\xce\xb1\x2f\xa4\xbd\x09\x34\x3a\x06\xf4\xea\xc5\x85\xc4\xdc\xb7\xe8\xb5\x4b\x79\x9f\x3f\xa4\x91\x38\x73\xa8\x2c\xd1\x7d\x70\x3c\xe3\x27\xcc\xfb\x40\x1c\xdc\x0b\xac\xef\x95\xfb\x7c\xe6\xa8\xcc\xb5\x27\xe7\x13\x1b\x03\x76\x7d\x42\x4e\xda\x1f\x27\x57\x2f\x16\x7c\xfe\xca\xea\xd1\xfb\x8e\x36\xd1\xb7\xff\xa6\xe4\x27\x06\xf7\x82\x68\xdb\x03\xcc\xfc\x22\x30\xcb\x12\x7c\xd9\xf6\x80\x38\x25\xe7\x26\xbb\x3e\xb1\xe4\xe7\xcf\xac\x0e\x03\xfb\x15\x09\xf3\x3e\x43\x62\x70\x6f\x32\xc1\xff\x19\x44\xfb\xeb\xc0\x86\x7e\x23\x5d\xff\xfc\x7e\xfa\xb8\xc5\xa9\x89\xa4\x38\xe5\x31\xb0\x29\xcf\x43\x3d\x7f\x67\xf5\x18\x7a\x43\x95\xe0\xf7\x4c\x26\xec\xaf\x83\xf8\xc9\x6b\xc0\xf8\x6d\x40\xe7\xaf\x7f\x6e\xfb\xfa\xc4\xa4\x38\xe5\xf9\xc1\xef\x1f\xf7\x0e\x66\xd5\xea\xc4\x11\xcd\x34\x1b\x7d\x09\xa8\x75\x33\x50\x5b\x79\xfa\xde\x60\x07\x72\x65\x54\x5a\xe3\xd3\xec\xda\xa5\x47\xfe\xfe\xb5\x78\x88\xf6\x9a\x95\xe2\x68\xb5\x91\x39\x77\xce\x52\xfe\x99\xb4\x3f\x3a\x4b\xbf\xb8\x64\x64\xd7\xc6\x1f\xeb\xfb\x67\x56\x0f\xd7\x4e\x35\x1b\xdb\xe1\x23\x57\x46\x7d\x24\x36\xf6\xd0\xef\xdf\x70\xa7\x1d\xe0\x7f\x35\x00\x33\xa5\x00\xb1\x15\x00\x1d\xcb\xe0\x0e\xc7\xb5\xdf\xe6\xb8\xf2\x19\x8e\x2b\x8d\xaf\x88\x29\xe3\xa5\x7a\x65\xbc\x7c\x83\x32\x5e\xb3\x4a\x19\x2f\x5f\x71\x36\xae\xe4\xbe\x8c\xaf\xe0\x3e\x9a\xe1\xb8\xf2\xdb\x1c\xa7\xfc\x0f\x17\x53\x7e\xd3\x31\xa3\xfc\x3a\x76\x5b\x19\xd7\xdf\x51\xc6\x55\xb7\x95\xf1\x55\x33\x4a\xe9\x7f\xa5\x39\xa4\xb9\xa4\x39\xa5\xb9\x65\x43\xb2\x24\xf3\x4e\x3b\xfc\x3f\x00\x00\xff\xff\x4e\xe9\x99\xf4\xbe\x10\x00\x00")

func viewsStaticFaviconIcoBytes() ([]byte, error) {
	return bindataRead(
		_viewsStaticFaviconIco,
		"views/static/favicon.ico",
	)
}

func viewsStaticFaviconIco() (*asset, error) {
	bytes, err := viewsStaticFaviconIcoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/static/favicon.ico", size: 4286, mode: os.FileMode(438), modTime: time.Unix(1534605757, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsStaticXasPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\xc5\x10\x3a\xef\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x28\x00\x00\x00\x28\x08\x06\x00\x00\x00\x8c\xfe\xb8\x6d\x00\x00\x0a\xb9\x69\x43\x43\x50\x50\x68\x6f\x74\x6f\x73\x68\x6f\x70\x20\x49\x43\x43\x20\x70\x72\x6f\x66\x69\x6c\x65\x00\x00\x48\x89\x95\x96\x07\x50\x54\x59\x16\x86\xcf\x7b\x9d\x13\xa9\x1b\x04\x24\x34\xb1\x09\x82\x04\x81\x06\x24\xc7\x26\x4b\x06\x13\x4d\x77\x03\x4d\x68\xdb\x0e\x28\x98\x10\x11\x47\x70\x0c\xa8\x88\x60\x18\xc1\x51\x10\x05\xc7\x00\xc8\x18\x10\x03\xa6\x41\x30\x81\x71\x40\x06\x05\x75\x1d\x0c\x98\x50\xd9\x02\x96\x66\x67\xb7\x76\xb7\xf6\xaf\x7a\xef\x7e\x75\xde\xb9\xff\x3d\xf7\xbd\x77\xab\x0e\x00\xb5\x9b\x2b\x16\x67\xa2\x2a\x00\x59\x22\x99\x24\x32\xc0\x9b\x19\x9f\x90\xc8\x24\xf4\x02\x1e\x68\xa0\x0a\x16\xe0\xca\xe5\x49\xc5\x5e\x11\x11\x21\x00\x00\x93\xe3\x5f\xf5\xe1\x1e\x20\x63\xe3\x6d\xab\x31\xaf\x7f\x7f\xfe\x5f\xa5\xca\x17\x48\x79\x00\x48\x04\x00\x24\xf3\xa5\xbc\x2c\x00\xe4\x04\x00\xd2\xc2\x13\x4b\x64\x00\x18\x19\x00\x18\x2d\x91\x89\xc7\xb8\x0c\x00\x18\x92\xf8\x84\x44\x00\xcc\xe1\x31\x4e\x9d\xe0\x96\x31\x4e\x9e\xe0\x3b\xe3\x39\xd1\x91\x3e\x00\x98\x01\x00\x22\x95\xcb\x95\xa4\x02\x50\xde\x03\x00\x33\x9b\x97\x2a\x03\xa0\x32\x00\xc0\x46\xc4\x17\x8a\x00\xa8\xbe\x00\xe0\xce\x4b\xe3\xf2\x01\xa8\x05\x00\x30\x23\x2b\x6b\xd1\x18\x1f\x05\x00\xb3\xe4\x7f\xf2\x49\xfd\x8b\x67\xb2\xc2\x93\xcb\x4d\x55\xf0\xc4\x5e\xc6\x45\xf4\x15\x4a\xc5\x99\xdc\x9c\xff\xf3\x75\xfc\x6f\x65\x65\xca\x27\xd7\x30\x04\x00\x6a\x9a\x24\x30\x12\x00\x34\x01\x90\xea\x8c\x45\xc1\x0a\x16\x25\x87\x85\x4f\xb2\x90\x3f\x9e\x3f\xce\x69\xf2\xc0\x98\x49\xe6\x49\x7d\x12\x27\x99\xcf\xf5\x0d\x9e\x64\x79\x46\x8c\xd7\x24\x73\x25\x53\x73\x85\x32\x4e\xf4\x24\x4b\x16\x45\x2a\xfc\x45\x99\x61\x21\x0a\x7f\x01\x47\xc1\x02\xa9\x5f\xd4\x24\xa7\x08\xfd\x39\x93\x9c\x9b\x16\x1d\x37\xc9\xd9\xc2\xd8\xb0\x49\x96\x66\x44\x05\x4f\xe5\xf8\x28\xe2\x12\x79\xa4\xa2\xe6\x14\x89\xbf\x62\x8f\x59\xd2\xa9\xda\x78\xdc\xa9\xb5\x64\x69\xd1\x81\x53\x35\xc4\x2b\xea\xe1\x0b\x7c\xfd\x14\x71\x51\x8c\x22\x5f\x2c\xf3\x56\x78\x8a\x33\x23\xa6\xea\xcf\x0c\x50\xc4\xa5\xd9\x51\x8a\xb9\x32\x49\xb4\x22\x9e\xce\x0d\x8a\x98\xf2\x89\x50\xbc\x1f\x08\x83\x08\xb0\x05\x1b\xb0\x01\x26\x44\x81\x00\x24\x20\x04\x01\x48\x65\x82\xa5\x63\xff\x34\xf8\x2c\x12\xe7\x48\x84\xa9\x69\x32\xa6\x97\x58\x9c\x29\x60\x72\x44\x3c\xeb\x19\x4c\x3b\x1b\x5b\x27\x80\xb1\x33\x38\xf1\x89\x1f\xac\x9d\x38\x5b\x9f\xab\xa6\x62\x61\x5d\x00\xa6\xdd\x00\xd0\x37\x15\x5b\xfb\x14\xe0\x7a\x25\x00\xbd\x79\x2a\x66\xa8\x0b\x40\x2b\x04\x38\xbb\x99\x27\x97\x64\x4f\xc4\xb0\x63\x37\x1c\x90\x41\x19\x18\xa0\x05\x7a\x60\x04\x66\x60\x05\x76\xe0\x08\xae\xe0\x09\x7e\x10\x04\xe1\x10\x0d\x09\xb0\x00\x78\x90\x06\x59\x20\x81\x25\xb0\x1c\x56\x43\x21\x14\xc3\x66\xd8\x0e\xe5\xb0\x17\xaa\xa0\x1a\x8e\xc0\x31\x68\x84\xd3\x70\x1e\x2e\xc3\x75\xe8\x80\xbb\xf0\x10\x7a\xa0\x1f\x5e\xc2\x10\x7c\x80\x11\x04\x41\x08\x08\x0d\xa1\x23\x5a\x88\x3e\x62\x82\x58\x22\x76\x08\x1b\x71\x47\xfc\x90\x10\x24\x12\x49\x40\x92\x90\x54\x44\x84\xc8\x91\xe5\xc8\x1a\xa4\x18\x29\x41\xca\x91\x7d\x48\x0d\xf2\x0b\x72\x0a\x39\x8f\x5c\x45\x3a\x91\xfb\x48\x2f\x32\x88\xbc\x45\xbe\xa0\x18\x94\x8a\x32\x50\x5d\xd4\x14\x9d\x89\xb2\x51\x2f\x34\x18\x8d\x46\xe7\xa3\xa9\xe8\x62\x34\x17\x2d\x40\x37\xa2\x65\x68\x25\x7a\x18\x6d\x40\xcf\xa3\xd7\xd1\xbb\x68\x0f\xfa\x12\x1d\xc6\x00\x86\x82\xd1\xc0\x18\x60\xac\x30\x6c\x8c\x0f\x26\x1c\x93\x88\x49\xc1\x48\x30\x2b\x31\x45\x98\x52\x4c\x25\xa6\x0e\xd3\x8c\x69\xc3\xdc\xc6\xf4\x60\x5e\x61\x3e\x63\xf1\x58\x3a\x96\x89\xb5\xc2\xba\x62\x03\xb1\x31\x58\x1e\x76\x31\x76\x25\x76\x03\xb6\x1c\x5b\x8d\x6d\xc0\x5e\xc4\xde\xc6\xf6\x62\x87\xb0\xdf\x71\x34\x9c\x0e\xce\x12\xe7\x82\xe3\xe0\xe2\x71\xa9\xb8\x25\xb8\x42\x5c\x29\xee\x00\xee\x24\xee\x12\xee\x2e\xae\x1f\xf7\x01\x8f\xc7\x6b\xe0\x59\x78\x27\x7c\x20\x3e\x01\x9f\x8e\x5f\x86\xdf\x80\xdf\x8d\xaf\xc7\xb7\xe0\x3b\xf1\x7d\xf8\x61\x02\x81\xa0\x45\xb0\x24\xb8\x11\xc2\x09\x5c\x82\x8c\x50\x48\xd8\x49\x38\x4c\x38\x47\xb8\x45\xe8\x27\x7c\x22\x52\x88\xfa\x44\x3b\xa2\x3f\x31\x91\x28\x22\xe6\x13\x4b\x89\x87\x88\x67\x89\xb7\x88\xcf\x89\x23\x24\x15\x92\x09\xc9\x85\x14\x4e\xe2\x93\x72\x48\x9b\x48\xfb\x49\xcd\xa4\x9b\xa4\x7e\xd2\x08\x59\x95\xcc\x22\xbb\x91\xa3\xc9\xe9\xe4\xd5\xe4\x32\x72\x1d\xf9\x12\xf9\x11\xf9\x1d\x85\x42\x31\xa4\x38\x53\xe6\x50\x84\x94\x3c\x4a\x19\xe5\x28\xe5\x0a\xa5\x97\xf2\x99\xaa\x46\xb5\xa0\xfa\x50\xe7\x51\xe5\xd4\x8d\xd4\x83\xd4\x16\xea\x7d\xea\x3b\x1a\x8d\x66\x4a\xf3\xa4\x25\xd2\x64\xb4\x8d\xb4\x1a\xda\x05\xda\x13\xda\x27\x25\xba\x92\xb5\x12\x47\x89\xaf\xb4\x4a\xa9\x42\xa9\x41\xe9\x96\xd2\x6b\x65\x92\xb2\x89\xb2\x97\xf2\x02\xe5\x5c\xe5\x52\xe5\xe3\xca\x37\x95\x5f\xa9\x90\x54\x4c\x55\x7c\x54\xb8\x2a\x2b\x55\x2a\x54\x4e\xa9\x74\xa9\x0c\xab\xd2\x55\x6d\x55\xc3\x55\xb3\x54\x37\xa8\x1e\x52\xbd\xaa\x3a\xa0\x46\x50\x33\x55\xf3\x53\xe3\xab\x15\xa8\x55\xa9\x5d\x50\xeb\xa3\x63\xe8\x46\x74\x1f\x3a\x8f\xbe\x86\xbe\x9f\x7e\x89\xde\xcf\xc0\x33\x58\x0c\x0e\x23\x9d\x51\xcc\x38\xc2\x68\x67\x0c\xa9\xab\xa9\xcf\x52\x8f\x55\x5f\xaa\x5e\xa1\x7e\x46\xbd\x47\x03\xa3\x61\xaa\xc1\xd1\xc8\xd4\xd8\xa4\x71\x4c\xe3\x9e\xc6\x97\x69\xba\xd3\xbc\xa6\x09\xa6\xad\x9f\x56\x37\xed\xd6\xb4\x8f\x9a\xd3\x35\x3d\x35\x05\x9a\x45\x9a\xf5\x9a\x77\x35\xbf\x68\x31\xb5\xfc\xb4\x32\xb4\xb6\x68\x35\x6a\x3d\xd6\xc6\x6a\x5b\x68\xcf\xd1\x5e\xa2\xbd\x47\xfb\x92\xf6\xab\xe9\x8c\xe9\xae\xd3\x79\xd3\x8b\xa6\x1f\x9b\xfe\x40\x07\xd5\xb1\xd0\x89\xd4\x59\xa6\x53\xa5\x73\x43\x67\x58\x57\x4f\x37\x40\x57\xac\xbb\x53\xf7\x82\xee\x2b\x3d\x0d\x3d\x4f\xbd\x74\xbd\x6d\x7a\x67\xf5\x06\xf5\xe9\xfa\xee\xfa\x42\xfd\x6d\xfa\xe7\xf4\x5f\x30\xd5\x99\x5e\xcc\x4c\x66\x19\xf3\x22\x73\xc8\x40\xc7\x20\xd0\x40\x6e\xb0\xcf\xa0\xdd\x60\xc4\x90\x65\x18\x63\x98\x6f\x58\x6f\xf8\xd8\x88\x6c\xc4\x36\x4a\x31\xda\x66\xd4\x6a\x34\x64\xac\x6f\x1c\x6a\xbc\xdc\xb8\xd6\xf8\x81\x09\xc9\x84\x6d\x92\x66\xb2\xc3\xa4\xcd\xe4\xa3\x29\xcb\x34\xce\x74\x9d\x69\xa3\xe9\x00\x4b\x93\xc5\x61\xe5\xb2\x6a\x59\x8f\xcc\x68\x66\x1e\x66\x8b\xcd\x2a\xcd\xee\x98\xe3\xcd\xd9\xe6\x19\xe6\xbb\xcd\x3b\x2c\x50\x0b\x07\x8b\x34\x8b\x0a\x8b\x9b\x96\xa8\xa5\xa3\xa5\xd0\x72\xb7\x65\xe7\x0c\xdc\x0c\xe7\x19\xa2\x19\x95\x33\xba\xac\xa8\x56\x5e\x56\xd9\x56\xb5\x56\xbd\xd6\x1a\xd6\x21\xd6\xf9\xd6\x8d\xd6\xaf\x67\x1a\xcf\x4c\x9c\xb9\x65\x66\xdb\xcc\xef\x36\x0e\x36\x99\x36\xfb\x6d\x1e\xda\xaa\xd9\x06\xd9\xe6\xdb\x36\xdb\xbe\xb5\xb3\xb0\xe3\xd9\x55\xd8\xdd\xb1\xa7\xd9\xfb\xdb\xaf\xb2\x6f\xb2\x7f\x33\xcb\x72\x96\x60\xd6\x9e\x59\xdd\x0e\x74\x87\x50\x87\x75\x0e\xad\x0e\xdf\x1c\x9d\x1c\x25\x8e\x75\x8e\x83\x4e\xc6\x4e\x49\x4e\xbb\x9c\xba\xd8\x0c\x76\x04\x7b\x03\xfb\x8a\x33\xce\xd9\xdb\x79\x95\xf3\x69\xe7\xcf\x2e\x8e\x2e\x32\x97\x63\x2e\x7f\xba\x5a\xb9\x66\xb8\x1e\x72\x1d\x98\xcd\x9a\x2d\x98\xbd\x7f\x76\x9f\x9b\xa1\x1b\xd7\x6d\x9f\x5b\x8f\x3b\xd3\x3d\xc9\xfd\x27\xf7\x1e\x0f\x03\x0f\xae\x47\xa5\xc7\x53\x4f\x23\x4f\xbe\xe7\x01\xcf\xe7\x5e\xe6\x5e\xe9\x5e\x87\xbd\x5e\x7b\xdb\x78\x4b\xbc\x4f\x7a\x7f\xf4\x71\xf1\x59\xe1\xd3\xe2\x8b\xf1\x0d\xf0\x2d\xf2\x6d\xf7\x53\xf3\x8b\xf1\x2b\xf7\x7b\xe2\x6f\xe8\x9f\xea\x5f\xeb\x3f\x14\xe0\x10\xb0\x2c\xa0\x25\x10\x17\x18\x1c\xb8\x25\xb0\x8b\xa3\xcb\xe1\x71\x6a\x38\x43\x41\x4e\x41\x2b\x82\x2e\x06\x53\x83\xa3\x82\xcb\x83\x9f\x86\x58\x84\x48\x42\x9a\x43\xd1\xd0\xa0\xd0\xad\xa1\x8f\xc2\x4c\xc2\x44\x61\x8d\xe1\x10\xce\x09\xdf\x1a\xfe\x38\x82\x15\xb1\x38\xe2\xd7\x39\xf8\x39\x11\x73\x2a\xe6\x3c\x8b\xb4\x8d\x5c\x1e\xd9\x16\x45\x8f\x5a\x18\x75\x28\xea\x43\xb4\x77\xf4\xa6\xe8\x87\x31\x66\x31\xf2\x98\xd6\x58\xe5\xd8\x79\xb1\x35\xb1\x1f\xe3\x7c\xe3\x4a\xe2\x7a\xe2\x67\xc6\xaf\x88\xbf\x9e\xa0\x9d\x20\x4c\x68\x4a\x24\x24\xc6\x26\x1e\x48\x1c\x9e\xeb\x37\x77\xfb\xdc\xfe\x79\x0e\xf3\x0a\xe7\xdd\x9b\xcf\x9a\xbf\x74\xfe\xd5\x05\xda\x0b\x32\x17\x9c\x59\xa8\xbc\x90\xbb\xf0\x78\x12\x2e\x29\x2e\xe9\x50\xd2\x57\x6e\x38\xb7\x92\x3b\x9c\xcc\x49\xde\x95\x3c\xc4\xf3\xe1\xed\xe0\xbd\xe4\x7b\xf2\xb7\xf1\x07\x05\x6e\x82\x12\xc1\xf3\x14\xb7\x94\x92\x94\x81\x54\xb7\xd4\xad\xa9\x83\x69\x1e\x69\xa5\x69\xaf\x84\x3e\xc2\x72\xe1\x9b\xf4\xc0\xf4\xbd\xe9\x1f\x33\xc2\x33\x0e\x66\x8c\x66\xc6\x65\xd6\x67\x11\xb3\x92\xb2\x4e\x89\xd4\x44\x19\xa2\x8b\x8b\xf4\x16\x2d\x5d\xd4\x29\xb6\x14\x17\x8a\x7b\x16\xbb\x2c\xde\xbe\x78\x48\x12\x2c\x39\x20\x45\xa4\xf3\xa5\x4d\x32\x86\x4c\x2c\xbb\x21\x37\x93\xaf\x95\xf7\x66\xbb\x67\x57\x64\x7f\x5a\x12\xbb\xe4\xf8\x52\xd5\xa5\xa2\xa5\x37\x72\x2c\x72\xd6\xe7\x3c\xcf\xf5\xcf\xfd\x79\x19\x76\x19\x6f\x59\xeb\x72\x83\xe5\xab\x97\xf7\xae\xf0\x5a\xb1\x6f\x25\xb2\x32\x79\x65\xeb\x2a\xa3\x55\x05\xab\xfa\xf3\x02\xf2\xaa\x57\x93\x57\x67\xac\xfe\x2d\xdf\x26\xbf\x24\xff\xfd\x9a\xb8\x35\xcd\x05\xba\x05\x79\x05\x7d\x6b\x03\xd6\xd6\x16\x2a\x15\x4a\x0a\xbb\xd6\xb9\xae\xdb\xfb\x03\xf6\x07\xe1\x0f\xed\xeb\xed\xd7\xef\x5c\xff\xbd\x88\x5f\x74\xad\xd8\xa6\xb8\xb4\xf8\xeb\x06\xde\x86\x6b\x3f\xda\xfe\x58\xf6\xe3\xe8\xc6\x94\x8d\xed\x9b\x1c\x37\xed\xd9\x8c\xdf\x2c\xda\x7c\x6f\x8b\xc7\x96\xea\x12\xd5\x92\xdc\x92\xbe\xad\xa1\x5b\x1b\xb6\x31\xb7\x15\x6d\x7b\xbf\x7d\xe1\xf6\xab\xa5\xb3\x4a\xf7\xee\x20\xef\x90\xef\xe8\x29\x0b\x29\x6b\xda\x69\xbc\x73\xf3\xce\xaf\xe5\x69\xe5\x77\x2b\xbc\x2b\xea\x77\xe9\xec\x5a\xbf\xeb\xe3\x6e\xfe\xee\x5b\x7b\x3c\xf7\xd4\xed\xd5\xdd\x5b\xbc\xf7\xcb\x4f\xc2\x9f\xba\xf7\x05\xec\x6b\xa8\x34\xad\x2c\xad\xc2\x57\x65\x57\x3d\xdb\x1f\xbb\xbf\xed\x67\xf6\xcf\x35\x07\xb4\x0f\x14\x1f\xf8\x76\x50\x74\xb0\xa7\x3a\xb2\xfa\x62\x8d\x53\x4d\xcd\x21\x9d\x43\x9b\x6a\xd1\x5a\x79\xed\xe0\xe1\x79\x87\x3b\x8e\xf8\x1e\x69\xaa\xb3\xaa\xdb\x57\xaf\x51\x5f\x7c\x14\x8e\xca\x8f\xbe\xf8\x25\xe9\x97\x7b\xc7\x82\x8f\xb5\x1e\x67\x1f\xaf\x3b\x61\x72\x62\xd7\x49\xfa\xc9\xa2\x06\xa4\x21\xa7\x61\xa8\x31\xad\xb1\xa7\x29\xa1\xa9\xf3\x54\xd0\xa9\xd6\x66\xd7\xe6\x93\xbf\x5a\xff\x7a\xf0\xb4\xc1\xe9\x8a\x33\xea\x67\x36\x9d\x25\x9f\x2d\x38\x3b\x7a\x2e\xf7\xdc\x70\x8b\xb8\xe5\xd5\xf9\xd4\xf3\x7d\xad\x0b\x5b\x1f\x5e\x88\xbf\x70\xe7\xe2\x9c\x8b\xed\x97\x82\x2f\x5d\xb9\xec\x7f\xf9\x42\x9b\x57\xdb\xb9\x2b\x6e\x57\x4e\x5f\x75\xb9\x7a\xea\x1a\xfb\x5a\xe3\x75\xc7\xeb\x0d\x37\x1c\x6e\x9c\xfc\xcd\xe1\xb7\x93\xed\x8e\xed\x0d\x37\x9d\x6e\x36\x75\x38\x77\x34\x77\xce\xee\x3c\x7b\xcb\xe3\xd6\xf9\xdb\xbe\xb7\x2f\xdf\xe1\xdc\xb9\x7e\x37\xec\x6e\xe7\xbd\x98\x7b\xdd\x5d\xf3\xba\x7a\xba\xf9\xdd\x03\xf7\x33\xef\xbf\x79\x90\xfd\x60\xe4\x61\xde\x23\xdc\xa3\xa2\xc7\x2a\x8f\x4b\x9f\xe8\x3c\xa9\xfc\xdd\xfc\xf7\xfa\x1e\xc7\x9e\x33\xbd\xbe\xbd\x37\x9e\x46\x3d\x7d\xd8\xc7\xeb\x7b\xf9\x87\xf4\x8f\xaf\xfd\x05\xcf\x68\xcf\x4a\x9f\xeb\x3f\xaf\x19\xb0\x1b\x38\x3d\xe8\x3f\xd8\xf1\x62\xee\x8b\xfe\x97\xe2\x97\x23\xaf\x0a\xff\xa6\xfa\xb7\x5d\xaf\xcd\x5e\x9f\xf8\xd3\xf3\xcf\x1b\x43\xf1\x43\xfd\x6f\x24\x6f\x46\xdf\x6e\x78\xa7\xf5\xee\xe0\xfb\x59\xef\x5b\x87\x23\x86\x9f\x7c\xc8\xfa\x30\xf2\xb1\xe8\x93\xd6\xa7\xea\xcf\xec\xcf\x6d\x5f\xe2\xbe\x3c\x1f\x59\xf2\x95\xf0\xb5\xec\x9b\xf9\xb7\xe6\xef\xc1\xdf\x1f\x8d\x66\x8d\x8e\x8a\xb9\x12\xee\x78\x2b\x80\x01\x00\x34\x25\x05\xe0\xed\x41\x00\x5a\x02\x00\xbd\x03\x80\xac\x34\xd1\x23\x8f\x0b\x99\xe8\x3d\xc6\x09\xfe\x13\x4f\xf4\xd1\xe3\x72\x04\xa8\x06\x80\xb1\x56\x2d\xca\x13\xa0\x22\x0f\xc0\x18\x00\x94\x00\x20\xd4\x13\x20\x31\x0f\x50\x7b\x7b\xc5\xf5\x0f\x49\x53\xec\xed\x26\xbc\x74\x45\x00\xca\x1a\xa3\xa3\xef\x3a\x00\xd4\x1f\x00\xbc\x65\x8d\x8e\x8e\x30\x46\x47\xbf\x77\x03\x20\x2c\x80\x6e\xfe\x44\x6f\x3e\x26\x76\x15\x40\xce\xf6\x31\xba\xd0\xfc\x34\xef\x5f\x7b\xe4\xbf\x03\xbc\x9e\x07\x28\xc3\xeb\x55\x8f\x00\x00\x00\x20\x63\x48\x52\x4d\x00\x00\x65\xa6\x00\x00\x87\x1c\x00\x00\xf2\xd8\x00\x00\x8e\x48\x00\x00\x72\x06\x00\x00\xf4\xb3\x00\x00\x2d\xd4\x00\x00\x13\xe3\x49\x96\x8b\xb2\x00\x00\x05\x9b\x49\x44\x41\x54\x58\xc3\xcd\x98\x7f\x68\xd4\x65\x1c\xc7\x6f\xd3\x4d\xf2\xfb\x3c\xcf\xf7\xfb\xcd\xca\x1f\x11\x92\x8a\xc9\x08\x21\x85\x20\x28\x12\x29\xff\x28\xbc\xd6\x1a\xa3\x31\x1a\x36\x1c\xad\x5a\xad\x98\x31\x3c\xea\xac\x19\x97\x5d\x78\xc2\xe0\x5c\x47\x1d\x5c\x79\x76\x71\xc8\xb1\xe3\xea\xb4\x0b\x92\xac\xfe\x68\x91\x61\x39\x5d\x19\x0a\x25\xa6\x04\x8b\xa9\x7b\x3e\x8f\x36\x7d\xe2\x71\x9f\xd3\x67\xe7\xf7\x7b\xdf\x9b\x9a\xec\x8f\x17\xdf\x3d\x9f\xcf\xeb\xf3\x79\xde\x7f\xec\xb6\xbb\xf3\x49\x29\x7d\xd3\x19\xdf\xb4\x0f\x28\x0a\x46\x39\xee\x10\x05\xe3\x90\x28\x18\x23\xa2\x60\xfc\xfd\x3f\xa0\xf6\x0e\x8a\x82\x41\xdc\x03\xee\x31\xca\xb1\x4d\xec\x31\xe4\x4d\x20\xe8\x1e\x30\x6f\xb8\xb1\x42\xe4\x8d\x11\x91\x37\xa4\xc8\x1b\x07\x44\xde\x78\x43\xe4\x8d\x90\x03\x9b\x45\xde\xd8\x22\xf2\xc6\x30\xba\x45\x86\xb4\xbe\xd3\x9c\xda\x77\x10\xdd\x13\x22\x6f\x2c\x73\x0e\x98\x23\x6e\xec\x14\x39\x22\x45\x8e\x9c\x15\x39\xb2\xa6\x8c\xa7\x58\x2e\x72\x64\x04\xfd\x22\x27\x45\x8e\x2c\xf3\x98\xf3\x8b\x1c\xe1\xe8\xc7\x9c\x03\x66\x89\x13\x8f\x8a\x2c\x11\x22\x4b\xa4\xc8\x92\x4f\x5c\x9c\x22\x55\x22\x4b\x76\xa2\x7b\x5e\x64\xc9\x9b\x22\x4b\x2e\xe0\xf9\x43\x8f\x59\xc5\x00\xba\xa3\x22\x4b\x1e\xbc\x3a\x60\x86\x94\x52\x2b\x32\xa4\x20\x32\x44\x8a\x0c\x19\x15\x19\x72\x8f\x83\xa3\xf3\x08\xba\x8a\xf7\x45\x86\xdc\x22\x32\x64\x07\x9e\xc7\x45\x86\x3c\xe4\x31\xbf\x5c\x64\x08\x47\x3f\x23\x32\x64\xc6\xe4\x80\x69\x52\xca\xd3\x22\x4d\x24\x12\x74\xe8\xeb\xd4\x8a\x34\x19\x44\xf7\xb8\x48\x93\x3a\xac\xdf\x27\xd2\xe4\x14\xd6\xf7\x89\x34\x99\xe1\xb6\x43\x85\x10\x69\x12\xd6\xee\x7c\x62\x72\xc0\x14\xd1\x21\x22\x45\x7e\x13\x29\x22\x45\x8a\x0c\x89\x14\x59\x50\xd2\x2f\xe5\x05\x74\x15\x9b\xb4\x7a\x95\x48\x91\x2d\x5a\x6f\x9d\xdb\x8e\x4b\x01\x53\x64\xa1\x48\x91\xdf\xd1\xfd\x45\xa4\xc8\xec\x2b\x01\x93\x44\xe7\x75\x91\x24\x52\x24\xc9\x45\x91\x24\xcf\x95\xf4\x4a\x99\x27\x92\x64\x18\xfd\x21\x91\x24\x77\x96\xf4\x17\x89\x24\x39\x82\xfd\x83\x22\x49\x6c\xa7\x3d\xc5\x20\x22\x49\x5e\xc1\x7b\x95\xbf\xe1\x72\x40\x48\xd0\x22\x77\x43\x82\x9e\x80\x04\x95\x90\xa0\xdf\x40\x82\xce\xd1\x7a\x4e\x6c\x46\x57\xf1\xbc\x8b\xb3\x41\x73\x36\x3a\x39\xc5\x20\x90\xa0\x73\x21\x41\x07\xd1\xfd\x13\x12\xf4\xae\x89\x80\x71\x5a\x64\x3b\xc4\xa9\x84\x38\x05\x88\xd3\x06\xad\xee\x44\x1d\xc4\xe9\x5f\xe8\x7f\x07\x71\x3a\xcb\xc5\x33\x20\x4e\x7f\x44\xef\x38\xc4\xe9\x92\x52\x47\xff\x7d\x83\x38\x7d\x06\xe2\x54\xa0\xdf\x37\x11\x30\x46\x15\x0f\x40\x8c\x9e\x81\x18\x95\x10\xa3\x59\xac\xb9\x31\x13\x62\xb4\x1f\xdd\x71\x88\xd1\x55\x1e\xfe\x63\xe8\x2a\xb6\x42\x8c\x56\xeb\xfd\x49\x01\x27\x7a\x05\x74\x4f\x43\x8c\xde\xef\x83\x28\x35\x20\x4a\x77\x41\x94\x4a\x64\x08\xa2\x34\x0f\x51\xfa\x85\x03\x7b\x20\x4a\xf7\x42\x94\x9e\x43\x97\x43\x94\x16\xb0\xee\xe6\x7f\x09\x51\x2a\x34\x7f\x6f\x19\xff\x73\x88\xd2\xc3\x5a\x96\x5d\x3e\xe8\xa3\x4b\xa0\x8f\x8e\x41\x1f\x95\xd3\x90\x31\x1f\x44\x58\x0d\x44\xd8\x6b\x10\x61\x17\x20\xc2\x24\x3e\xf7\x41\x84\xed\x80\x08\x4b\x6b\x7c\x04\x11\x76\x14\x1d\xc5\xf7\x0e\x8e\x1b\xca\xdb\xaf\xcd\xfe\x0a\x11\xf6\xb1\x83\xb3\x4f\xcb\x31\x0e\x11\xb6\xc1\x07\x61\xa6\x98\x09\x61\xd6\x0f\x61\x26\x21\xcc\x2e\x42\x98\x0d\x40\x98\x51\x08\xb3\x1a\x08\xb3\x5a\x64\x01\x84\xd9\x01\x74\x86\x21\xcc\x16\x43\x98\x55\x6b\xfd\x72\x28\x6f\x05\x84\xd9\x1f\x38\x3f\x08\x61\x36\x47\xeb\xd7\xe0\x7d\x03\x78\xbf\x72\xb6\xab\x5c\x3e\x08\xb1\x22\xb3\x20\xc4\x3e\x83\x10\x93\x10\x62\x17\x20\xc4\x42\x5a\x4f\x11\xc0\xba\xea\xbf\x5a\xd2\xab\x94\xb7\xb5\xfd\x2f\xab\xda\xe5\x17\x48\x88\xbd\xa3\xed\xcf\xa9\x3c\x13\xaf\xe2\x5e\xa6\x33\x1f\x7a\xd9\x4f\xd0\xcb\x24\xd2\x89\xf5\x65\xd0\xcb\x8e\x61\x4d\xf5\x67\x97\xcc\x55\xca\xed\xd0\xcb\x8e\xe0\x9e\x61\xe8\x65\x0b\x2f\x85\x9b\xb8\xa7\x78\xe7\x7e\xe8\x65\xf3\xae\xfc\xa1\x0e\xb2\x52\x96\x43\x90\x1d\x85\x20\x93\x10\x64\x63\x10\x64\x4f\x42\x90\x6d\xc5\xb3\xc2\xef\x30\x33\x15\xd6\x69\xbb\x82\x10\x64\x6b\x21\xc8\x38\x9e\xd5\xbd\xf7\x4e\xfa\x5f\x0c\x01\xd3\x89\xc7\x21\x60\xfe\x03\x01\x53\x42\xc0\x3c\x0b\x01\xf3\x0c\xfe\x3c\x00\x01\xb3\xca\x65\xa6\x52\x6a\x21\x60\x7e\x8d\xfb\xd4\x1d\xa7\xf1\xe7\x11\x75\xef\x55\x6f\xb7\xa0\xc7\x74\xe3\x25\xe8\x31\xa5\xc6\x69\xe8\x31\x57\x96\xf1\xa7\xc2\x2a\xe8\x31\x79\xc9\xfe\x4e\xc7\x37\xac\xd0\x6d\x96\xe3\x3d\xe8\x36\x25\xf2\x03\x74\x9b\xb7\x79\xf8\x95\x32\x17\xba\xcd\x9f\xb5\xdd\x61\xd7\xcf\x24\xd0\x65\x96\xa3\x06\xba\xcc\x14\x74\x99\x12\xf9\x00\xba\xcc\x6a\x8f\x19\x2f\xd4\x7c\x5c\xdb\xb9\x53\xdd\xe3\x1e\xb0\xd3\xf4\xc2\x86\x4e\x73\x2f\x74\x9a\x12\xd9\x54\xc1\x4c\x39\xde\xd2\x76\x7d\xa5\xf6\x97\xfd\x5c\xcc\x3b\xac\x4a\xa8\xe3\x1d\xd6\x61\xde\x61\x49\xde\x61\xfd\xcb\x3b\xac\x67\x2b\x9c\x2b\xa5\x0d\xe7\x25\xee\xab\xf3\xfc\xe0\xce\xdb\xad\x4a\x59\xc3\xdb\xad\x51\xde\x6e\x49\xde\x6e\x9d\xe4\xed\xd6\xc3\x53\x98\x55\xac\xe2\xed\xd6\x29\x9c\x1f\xc5\x7d\xde\xdf\x2c\xf0\x36\x6b\x2a\xb4\xf2\x36\x6b\x9c\xb7\x59\x92\xb7\x59\x43\xbc\xcd\x5a\x54\xe1\xdc\x62\xf4\x25\xce\xb7\x16\x7b\xde\x01\x5b\xad\xa9\x12\xe0\xad\x96\x44\xbe\xe5\xad\x96\xed\xe1\xdb\xe8\x15\x67\x36\xea\x7d\xef\x80\x2d\xd6\x54\xa9\xe2\x2d\x56\x3f\x6f\xb1\x24\x92\xe2\x2d\x56\xb5\x8b\xab\xea\x9f\x6a\x6e\x3f\xce\x5f\x76\xbc\x03\x36\x5b\xd7\xc2\xad\xbc\xd9\xda\xcd\x9b\x2d\x89\x6c\xe1\xcd\x56\x55\x89\xa3\xce\xef\x6a\xce\x6e\x9c\x9b\xb4\xcb\x3b\x60\x93\x7d\xad\x2c\xe5\x4d\xf6\x21\xde\x64\x4b\xde\x64\x03\x6f\xb2\xd7\x97\xf4\xd7\x63\x5d\xa2\xb7\xd4\x69\x8f\x77\xc0\x46\xfb\x7a\x58\xc9\x1b\xed\x53\xbc\xd1\x96\xbc\xd1\x1e\xe1\x8d\xf6\x6a\xac\xaf\xc6\xb3\xc4\xfe\x4a\xb7\x1d\xde\x01\x1b\xec\xeb\xc5\xcf\x1b\xec\xf3\xbc\xc1\x96\xbc\xc1\x3e\xc6\x1b\xec\xa7\xf0\xa9\xce\xe7\x78\x83\xbd\xb6\xdc\xbc\x77\xc0\x7a\xfb\x46\xf0\x22\xaf\xb7\xc7\x79\xbd\x2d\x79\xbd\x7d\x11\x9f\xe3\x58\x2f\x3b\xeb\x1d\xd0\x6f\xdf\x08\x66\x72\xbf\xbd\x8d\xfb\x6d\xa9\xb1\x0d\xeb\x65\x67\x6f\x56\x40\x05\xe3\x7e\x3b\x87\xe1\x72\x78\xf6\x4d\xa7\x80\x8a\xf9\xdc\x6f\xc7\xf1\xe9\xbb\x21\x01\xa7\xfb\xb7\xfc\xff\x01\x50\xa5\x03\xb8\x44\xf7\x8d\x2d\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\xbb\xb2\x0e\x4d\xc5\x10\x00\x00")

func viewsStaticXasPngBytes() ([]byte, error) {
	return bindataRead(
		_viewsStaticXasPng,
		"views/static/xas.png",
	)
}

func viewsStaticXasPng() (*asset, error) {
	bytes, err := viewsStaticXasPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/static/xas.png", size: 4293, mode: os.FileMode(438), modTime: time.Unix(1534605757, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configSeelogXml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\x4d\x4b\x1b\x41\x18\xc7\xef\x82\xdf\xe1\xe9\xc0\x42\x52\x1a\xd7\x42\x5b\xa8\xcd\xe6\x54\x0b\x05\xa5\x97\xd2\x8b\x78\x58\x37\xcf\xae\x03\xbb\x33\x61\x5f\x44\x7b\x2a\x35\x9a\xb5\x15\x62\xab\xf5\x25\x08\xb5\xa0\x44\x90\xbc\x1c\x44\x34\xa6\xe9\x97\x71\x76\x93\x93\x5f\xa1\x98\xdd\xa4\xdb\x78\x71\x4e\x33\xff\xc9\x33\xff\xff\xef\xc9\xc3\x66\x1d\x44\x93\x1b\xe0\xae\x14\x50\x21\xaa\xb3\xc2\x34\x97\x5a\x68\x13\xe8\xef\x29\x73\xd1\x5e\x52\x4d\x85\x3c\x9f\xec\x2f\x02\x16\x65\x26\x2e\xa1\xa9\x90\x3c\x2e\x78\x06\x01\x4b\x5d\x8e\x05\xb4\x6d\x6e\x93\xdc\xf8\xd8\xf8\x18\x00\x40\x16\x97\x35\x2c\xb8\x94\x33\x27\x17\x29\x7d\xf5\x51\x26\x93\xb8\xca\x41\xb0\xdf\xb8\xb9\x6c\x85\x1b\x57\xa2\x5e\x09\x2b\xc5\xf0\xf4\xab\x68\x95\x83\xdd\xd2\xcd\xf5\x45\x4a\xa7\x26\x16\x54\xd7\x45\x9b\xa5\x03\x7f\x57\x94\x7e\x07\x3f\x9a\x29\xdd\x63\xda\x40\xed\xd6\x3b\x77\x75\x1b\x57\x41\xfd\x4b\x58\x29\x06\x7b\x27\xe2\xcf\x5e\xb7\x5a\x14\xfe\x01\x64\x32\x49\xdf\xa1\x25\x24\x5e\x55\xc8\x63\x4b\xa5\x6c\xc2\xe0\x49\xb2\x08\x44\x8e\xab\xb3\xf2\x7f\x20\xb1\xc8\x3d\xb7\xe0\xb9\x0e\xe8\xdc\xb6\x54\x97\xe6\x15\xa2\x71\xcb\xe2\x8c\xdc\x63\xd5\x38\x73\xb8\x89\x39\x08\x8e\x4a\x61\xad\xd3\xfd\x75\x1a\x1e\xb7\xba\x9d\x6d\x51\x6a\x09\xbf\x19\x5e\xfb\xe1\x59\x63\x24\x6b\x5c\x22\xdf\x7b\x6b\xc1\xd3\x75\xb4\x31\x9f\x03\xd1\x5c\x8f\x60\xc5\x9a\x2f\x6a\xfb\xe2\xf0\x54\xac\xaf\x89\xda\xfe\xcd\x65\xed\xb6\xbd\x29\xea\x95\xe0\xf0\xa7\x58\x3f\x10\x6b\x27\x51\x2f\x6f\xdb\x9b\xbd\x4f\x9f\xc5\x96\x1f\x57\x5d\x5d\x88\xf2\xb7\x5e\xa9\xdc\xed\xac\x8a\xe3\x6a\xe0\xef\xc2\xdb\x77\x10\x34\x76\xba\x9d\xd5\xf0\xbc\x2a\xda\x47\x61\xa5\x28\x0e\x5b\x62\xcb\x07\x87\x7e\xc4\x29\x08\xdb\xdb\x77\x3e\xc7\x55\xd1\x2c\xbf\x02\xdd\xf4\x9c\xc5\x02\xda\x94\xe7\x07\x57\xbd\xbd\xf3\x5e\x65\x27\x15\x34\xce\xc2\xea\xf7\xf4\x08\x91\x4e\x4d\x17\x6d\xe8\x37\xd8\x51\x08\x65\x3a\x7f\x32\x9c\x17\x48\xac\x21\xe2\xfd\xce\xf6\x83\x28\xe4\x69\x34\x88\x89\x04\x91\x36\xfa\xd2\xc0\x17\xa1\xa0\xba\x8b\x0a\x31\xb9\xe1\xc8\x26\x37\x0c\xb4\x27\x4c\x6e\x10\x79\xd4\x59\x1e\x76\x37\x11\x5c\x8e\x92\x0f\x67\x21\xfe\xdf\x07\xe7\x28\xe4\xe8\x80\x27\xc1\xfb\x3f\x80\x24\x46\x24\x29\x44\x7a\xad\xba\x08\xd2\x7b\x6a\x21\x48\xd3\x8e\x36\x9b\x7a\xf6\x22\x3d\x27\xcd\x4c\x7f\x98\x8f\x8f\x2f\xd3\xd1\x66\x32\x0d\x73\xd2\x1b\x6a\xe2\x94\x34\x43\x19\xce\xdf\x9d\x3c\xa6\xcd\x83\x34\xeb\x18\x12\x23\x90\x64\x19\x69\xfc\xc3\xfc\x23\xdb\x07\xba\x64\xe5\x7f\xd8\x59\x39\xfa\x84\xe4\xc6\xc7\xfe\x06\x00\x00\xff\xff\xe3\xf7\x64\x25\x4d\x04\x00\x00")

func configSeelogXmlBytes() ([]byte, error) {
	return bindataRead(
		_configSeelogXml,
		"config/seelog.xml",
	)
}

func configSeelogXml() (*asset, error) {
	bytes, err := configSeelogXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/seelog.xml", size: 1101, mode: os.FileMode(438), modTime: time.Unix(1535292788, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configAppIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\xc1\x6a\xeb\x30\x10\x45\xf7\x03\xf3\x0f\x82\xf7\xd6\x46\xb1\x94\x58\x14\xbc\x30\xb5\x68\x5c\x70\x64\x24\x19\xd3\x06\x23\x14\x22\x0a\x25\xc5\x8e\x94\xb6\xf4\xef\x8b\xb2\xec\xf6\xde\x99\x73\xcf\xbf\x73\x38\x7d\xbe\x91\x25\x92\x18\x2e\xc1\xa7\x80\xa0\xc7\x83\xeb\x55\x2b\x49\x4d\xee\x25\x02\xc2\xd1\xaf\xeb\x8c\x30\x34\x4f\xd2\x99\xee\x35\x77\x1b\x8a\xf0\x3c\x59\x67\xe4\xa3\x96\x96\xd4\xa4\x64\x8c\x57\xff\x29\xa7\x7c\x53\xde\x7f\x52\x88\x5f\x21\xce\x08\x7b\x6b\x07\x37\x28\x9d\xaf\x04\x15\x02\x41\xcb\xa6\x75\xb6\xeb\xa5\x1a\x73\xb8\xa3\x08\x93\xee\xac\xfc\x93\x21\x1c\xcf\xfe\xe6\x4f\x3e\x85\x19\xc1\xbe\x0c\x79\xf8\xe3\x27\x5d\x2f\x08\xa3\x91\x9a\xd4\x24\x2e\xcb\x2d\x8b\x19\x33\x29\xdd\x66\xaf\x6b\xf9\xcd\x02\x8f\x08\x7b\x65\x32\x88\x6f\x8b\xaa\x2a\x44\x59\x88\xed\x03\x63\x74\x87\x70\x68\xfa\x0c\x7a\x5f\x63\x48\xe9\x37\x00\x00\xff\xff\xcf\xb1\x30\x2d\x02\x01\x00\x00")

func configAppIniBytes() ([]byte, error) {
	return bindataRead(
		_configAppIni,
		"config/app.ini",
	)
}

func configAppIni() (*asset, error) {
	bytes, err := configAppIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/app.ini", size: 258, mode: os.FileMode(438), modTime: time.Unix(1534605755, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"views/index.html": viewsIndexHtml,
	"views/static/favicon.ico": viewsStaticFaviconIco,
	"views/static/xas.png": viewsStaticXasPng,
	"config/seelog.xml": configSeelogXml,
	"config/app.ini": configAppIni,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"config": &bintree{nil, map[string]*bintree{
		"app.ini": &bintree{configAppIni, map[string]*bintree{}},
		"seelog.xml": &bintree{configSeelogXml, map[string]*bintree{}},
	}},
	"views": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{viewsIndexHtml, map[string]*bintree{}},
		"static": &bintree{nil, map[string]*bintree{
			"favicon.ico": &bintree{viewsStaticFaviconIco, map[string]*bintree{}},
			"xas.png": &bintree{viewsStaticXasPng, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

