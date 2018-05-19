// Code generated by go-bindata.
// sources:
// _templates/.dirlist.html.swp
// _templates/dirlist.html
// _templates/notebook.css
// _templates/notebook.html
// _templates/page.html
// _templates/style.css
// _templates/table.html
// DO NOT EDIT!

package main

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

var _DirlistHtmlSwp = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4b\x8c\x1c\x47\x19\xfe\xc2\x3e\xe6\xb5\x5e\xa3\x70\x20\x17\xc8\xef\xce\x2a\x99\x61\x77\xbb\xc7\x5e\x3f\x36\xf6\xcc\xa0\xc4\x71\x1e\x38\x7e\xc4\x71\x62\xc7\xab\x90\xd4\x76\xd7\x4c\xd7\xba\xa6\xbb\xb7\xab\x7a\x76\x87\xcd\x90\x48\x91\x90\xb8\x20\x1e\x52\x90\x72\xe0\x04\x07\x0e\x1c\xe0\x00\x07\x2c\x94\x2b\xc9\x11\xc4\xe3\x10\x24\x10\x04\x01\x12\x02\x24\xc4\x11\x54\xdd\x3d\xe3\xf1\xc6\xc1\xbe\x40\x14\xa9\x3f\x69\xb6\xba\xaa\xfe\xfa\xff\xef\x7f\x54\x75\x8d\x76\x36\x9b\xcf\x3f\x75\x8e\xd6\xed\x26\x00\x7c\x1c\x78\xf7\x1c\x36\xde\xb9\xef\x21\xbc\xfd\x12\x90\x6c\x26\x81\x4e\x70\x47\x6c\x26\xae\xcf\x02\x16\x34\x0f\xff\x77\xb9\x2f\x66\x0a\x1d\x15\xbb\x4e\x4f\x68\x3f\xd9\xb4\xdd\xb0\xef\xe4\xcb\xb9\xd3\x67\x89\xc7\x9d\x97\x34\xef\x47\x92\x69\xae\x1c\x4f\xc4\x52\x28\x6d\xfb\xba\x2f\xef\x4c\xa3\x40\x81\x02\x1f\x84\x44\x77\x57\xd7\x0f\x60\xed\xc8\xe1\x74\xab\x3f\x60\x1d\xa2\x4f\xdc\xfb\xdc\x87\xcd\xaa\x40\x81\x02\x05\x0a\x14\x28\x50\xa0\x40\x81\x02\xff\x47\xe8\xe8\x1e\xbc\x0a\xe0\x63\x79\xff\xcb\x79\x7b\xcf\xbe\xb6\x40\x81\x02\x05\x0a\x14\x28\x50\xa0\x40\x81\x02\x05\x0a\x14\x28\xf0\xd1\x05\xf3\x80\x37\xee\x01\x7e\x31\x93\xfd\xff\x7f\xfc\xfd\xff\x6f\x07\x81\x3f\x1d\x04\x7e\x77\x10\xb8\x71\x10\x78\xe1\x20\xf0\x8d\x45\xe0\xeb\x8b\x80\xbb\x08\x6c\x2c\x02\x97\x17\x81\x87\x17\x81\xc5\x45\xe0\x85\x03\xc0\xd9\x03\xc0\xe9\x03\x40\xf5\x00\x50\x39\x00\xfc\x65\x01\xf8\xf9\x02\xf0\xd6\x02\x20\x17\x80\x97\x17\x80\xab\x0b\xc0\x33\x0b\xc0\xc5\x05\xa0\xbe\x00\x3c\xb4\x00\xdc\xbf\x00\xbc\x57\x03\xbe\x5d\x03\x5e\xab\x01\xaf\xd6\x80\x97\x6a\xc0\xe7\x6b\xc0\x85\x1a\xf0\x68\x0d\x78\xa4\x06\xdc\x5f\x03\x0e\xd6\x80\xc5\x1a\xf0\xfb\x2a\x70\xa3\x0a\x7c\xb7\x0a\xbc\x59\x05\xbe\x59\x05\xde\xa8\x02\xba\x0a\xa8\x2a\xb0\x51\x05\xae\x55\x81\xc7\xab\xc0\xa9\x2a\x70\xa2\x0a\x1c\xaf\x02\x8b\x55\xe0\x9f\x15\xe0\xdd\x0a\xf0\xeb\x0a\xf0\xab\x0a\xf0\x76\x05\xf8\x51\x05\xf8\x61\x05\xf8\x56\x05\x78\xbd\x02\xec\x56\x00\x5d\x01\x54\x05\xd8\xa8\x00\xa7\x2b\xc0\x72\x05\xb8\xb7\x02\x2c\x56\x80\x5a\x05\xa8\x56\x80\x4a\x05\xf8\x6b\x19\x78\xb7\x0c\xfc\xa0\x0c\xbc\x59\x06\xbe\x54\x06\xfc\x32\x70\xb5\x0c\x3c\x5a\x06\xd6\xca\xc0\x72\x19\xf8\x64\x19\xa8\x96\x81\xd9\x32\x80\x32\xf0\xef\x12\xf0\xaf\x12\xf0\xcb\x12\xf0\x4e\x09\x78\xab\x04\xdc\x28\x01\x3f\x2e\x01\x5f\x2d\x01\xbb\x25\xc0\x2b\x01\x67\x4b\xc0\x89\x12\xf0\x99\x12\xb0\x54\x02\xa8\x04\x7c\xba\x04\x7c\xaa\x04\xfc\x63\x1e\xf8\xed\x3c\xf0\xb3\x79\xe0\xed\x79\xe0\xa7\xf3\xc0\x2b\xf3\x80\x33\x0f\xbc\x37\x07\xdc\x98\x03\xbe\x36\x07\xbc\x3e\x07\x04\x73\xc0\xe3\x73\xc0\xf2\x1c\x70\xdf\x1c\xf0\xe7\x59\xe0\x0f\xb3\xc0\x4f\x66\x81\xef\xcc\x02\xaf\xcd\x02\xcf\xcf\x02\x97\x66\x81\x0b\xb3\xc0\x99\x59\xe0\xb1\x59\xe0\xd4\x2c\xf0\xf7\x19\xe0\x8f\x33\xc0\x7b\x33\xc0\x6f\x66\xb2\x7a\xf9\xfe\x0c\xf0\xbd\x19\xe0\x2b\x33\x1f\x72\x11\x7f\x14\xd1\x72\x7c\xdd\x97\x1d\xb4\x9c\xcd\xc9\x63\x6b\x6a\x30\xf4\x86\x1d\x23\xd5\x0d\x43\xcd\xe3\x0e\x88\x5a\x51\xcb\xb9\xa3\x48\xe7\x62\xb8\xc3\xe3\xbb\x12\xe2\x1e\x6d\x0e\xa9\xc5\xc8\x8f\x79\xb7\x6d\xf9\x5a\x47\xea\xa4\xf3\xc1\xbf\xcc\xb1\x3a\x69\xd3\x72\x58\xa7\xe5\x44\x1d\xb4\xc6\x2a\xd1\x72\x94\x1b\x8b\x48\x77\x30\x3a\x85\x8a\xf9\x54\x44\x97\xea\x2e\x93\xf2\x7c\xb8\xd3\xa0\x6e\x12\xb8\x36\x8b\x22\x39\xac\xbb\x61\xa0\xf9\xae\x5e\x21\x16\xf7\x54\xc3\x48\x6a\xd1\xe7\x61\xa2\xa9\x4d\x8a\xeb\xcb\x59\xa7\x2e\x99\xe6\xf1\x0a\xed\x30\xa1\x53\x21\x57\x72\x16\x8f\x27\xf3\x15\xe9\xc4\x80\xc5\x94\xdb\xa1\x36\x89\x7e\x9f\x7b\x82\x69\x4e\x0f\x3e\x48\x87\x72\x39\x23\x96\x72\x4a\x49\x1d\x9a\xc8\xdc\x81\xd7\x14\xb1\x20\x91\x72\x6c\x2c\x65\x46\xed\x74\xad\x16\x61\x50\x6f\xd0\xde\x98\x47\xa6\x83\xda\xa4\x7d\xa1\x32\x55\xd4\x36\x4d\xd2\xe7\x81\x56\xa7\x50\x89\xb9\x4e\xe2\x60\xdf\x62\xb3\x76\xc2\x75\x3c\x45\x1e\xdf\x0c\x93\xc0\xe5\x75\x33\x92\x85\x62\x85\xa6\xc8\xef\xc1\x71\x48\x72\xe6\x89\xa0\x47\xdc\xeb\xf1\x15\x12\x81\xd2\x9c\x79\x14\x76\x49\xfb\x9c\x74\xcc\x84\x14\x41\xcf\x36\x92\xe7\xa9\x2f\xa4\x14\x8a\xbb\x61\xe0\x29\x9b\x9e\xea\xd2\xcb\x13\x6d\x2f\x93\x50\x14\x31\xa5\xb8\xb7\x42\x3a\x16\xbd\x1e\x8f\x53\x15\x13\x36\x61\x60\xfa\x46\xd1\x26\x1f\x4b\x70\xcf\xa6\xcb\xd3\x42\x3b\x42\x4a\x33\x6f\x12\xc2\x3d\x62\x5d\x13\x2a\xa1\x49\xe9\x30\x52\xb4\xc9\x0d\xd3\x7c\xae\x1b\xc6\x46\xd9\xa5\x34\x1e\x8a\xd8\x44\xc9\x0a\x69\x9f\x99\x3c\x28\x92\x61\xd0\x33\xad\xd0\x69\x68\x45\x90\x70\x45\x3a\x34\x16\x44\x30\x08\xaf\x1b\xb2\xa9\xc9\x20\xd4\xc0\x68\x85\x8e\x37\x9b\x8d\x06\x88\x14\x67\xb1\xeb\x53\x9b\x2c\x0b\x44\x26\xbc\xdb\x69\x79\x99\x51\x2c\xd5\x2d\xb3\x29\xac\x86\x7d\x9d\x0f\xa3\x98\x2b\x55\xbf\x25\xd4\x69\x5e\xf8\xc0\x04\x18\x23\x10\x99\xcf\xf8\x2f\x51\x96\xc0\xbc\xb3\x54\xe7\xb2\x61\x77\x43\x37\x51\xf5\x46\x3a\x66\x2a\x2c\xde\xb5\x35\x57\xba\x1e\xb0\x3e\x6f\x18\x35\x66\xc2\x90\xe0\x92\xda\xc4\xa5\xda\x10\x2f\x4e\xc6\x8c\x90\xa9\x30\xd6\xe7\xf9\x78\x37\x8c\xa9\x6e\xa6\x04\xb5\xa9\x79\x8a\x04\xb5\xb2\x69\x5b\xf2\xa0\xa7\xfd\x53\x24\x96\x97\x53\x76\x63\x56\x13\x4e\x23\x72\x99\x76\xfd\x29\x93\xf1\xae\x51\xce\x77\xe8\x12\xef\x9d\xd9\x8d\xea\x59\x0c\x0c\x57\x1d\x0f\xa7\x75\xdc\xa5\x7f\x37\x05\xf7\x89\x8e\x85\x5d\x29\xdc\xeb\x13\x61\x13\x8e\x5e\x38\x8e\x81\xe9\x19\x4f\x6c\xa5\x59\xac\xd5\x15\xa1\xfd\xfa\xf6\xff\x32\x42\x64\xea\x46\x85\x92\xdb\x32\xec\xd5\xb7\x1b\x53\x9b\x2b\xcc\x22\x51\xdf\x5e\xa1\x8c\x20\x30\x6a\x4c\xa2\xb1\x54\xb7\x1e\xc8\x04\xac\x5b\xfd\xe7\x03\x3b\x8a\xf9\x80\x07\xfa\x31\xde\x65\x89\xd4\xe9\xb8\x71\x8c\x0f\xec\x1d\x5f\x98\xaa\x6b\xd3\xd1\x13\x99\x75\xc7\x21\xcb\xb1\xe8\x3a\x1f\xa6\x9c\xf3\x02\x7c\x7f\xde\xe8\x26\x9f\xac\x31\x9b\x30\xe1\xb7\x53\x7d\x78\x2d\x4f\xfd\xbe\x15\x53\x65\xbf\xdc\x26\xf7\xa6\x11\x97\xda\xe4\xda\x3a\x7c\xda\x1c\xf9\xa7\x99\xe2\x13\xc6\x4c\x46\x3e\x1b\x17\xab\x9b\xe7\x21\x3d\xc7\xa8\x4d\xcf\xea\xd8\x9c\x1c\xdd\x38\xec\x9f\xf6\x59\x7c\x3a\xf4\xf8\x84\x47\xe3\x36\x2e\xe4\x3c\xfb\x5c\xb3\xb3\x7c\x98\x73\x9c\x8e\xfe\xd4\xe2\xdb\xec\xc0\xf7\x6d\x3c\x43\x92\x4b\x65\x47\x89\xf2\x4d\x61\x81\xf2\x14\xa7\x03\xe6\x98\x6d\xe4\x6c\xf3\x23\x37\x2b\x3f\xd3\xa9\x37\xf6\xb9\xbb\x54\xb7\xec\xfc\xa7\xa4\x24\x05\x31\xab\x61\x73\xe6\xfa\x37\x8d\x8a\x15\xe2\x32\xb5\x6b\x14\x4e\x9f\x1e\xa6\x9f\xc7\xe9\xd6\x8d\xe4\x6c\x3c\xb2\x7a\xed\x45\xa7\x91\x2d\xe1\xd2\x1c\xf4\x1b\x2f\x62\x5c\xb0\x79\x17\xad\xec\xc5\x48\x7a\x18\xf1\xb6\x65\xd8\x39\x5b\x6c\xc0\xb2\x51\xab\x93\x06\x29\x0e\x95\x0a\x63\xd1\x13\x41\xdb\x62\x41\x18\x0c\xfb\x61\xa2\xac\xce\xcd\x97\x2a\x91\x08\x34\xef\xc5\x42\x0f\xdb\x96\xf2\xd9\x91\x63\xc7\x57\xd7\xb8\x17\xf7\x87\xc9\x33\xcd\x9d\xe3\xc7\xba\xeb\xbd\xee\xa3\x6a\x3b\xdc\xf9\xc2\xd6\xe7\xf8\x11\x71\xee\x78\xd0\x0c\xce\xba\xe2\xe2\x73\xd1\xfa\x70\xf9\xc4\x99\xb6\x39\x04\x55\xec\xde\x7c\xc5\xbb\xa1\xc7\xed\xad\xed\x84\xc7\xc3\xf4\x3d\x9f\x3d\xae\xae\xd9\x6b\xf6\x61\x5b\x49\xd1\xb7\xfb\x22\xb0\xb7\x94\x35\xe6\x0f\xb4\x3c\x31\x20\x57\x32\xa5\xda\x56\xcc\x99\xd7\xe7\x56\x67\x6f\x8f\xec\x4b\xe9\x33\x8d\x46\x2d\xc7\x13\x83\xf4\x2e\x90\xb6\x44\x2d\x27\x91\xa6\xdd\xdb\x23\x1e\x78\x34\xca\x0a\xa6\x25\x45\x67\x72\xe5\x30\x0a\x9e\xbc\x74\xe6\x71\x1a\x8d\xac\x4e\x4b\x8c\xf5\x77\x19\x75\xd9\xaa\x99\xbc\x3c\x8c\x78\x36\xe9\x88\xd4\xdc\x79\x96\x1b\x33\x77\x10\x29\x72\xfd\x31\x0b\x7a\x9c\xec\xa7\x4d\x82\x53\x3b\xad\x44\x8e\x95\xe5\x89\xb7\x3a\xb7\xb8\x90\x8f\xae\xee\xc4\x2c\x8a\x78\x9c\x65\xa2\x25\x82\x28\xd1\x24\xbc\xb6\x95\x6f\xfe\xa9\xbc\x59\xa4\xf4\x50\x72\xb3\x54\x45\x92\x0d\x4f\x06\x61\xc0\x2d\x72\xa6\x5d\xdd\xef\x6c\xe6\xee\xde\x9e\xe8\xd2\x92\x08\x3c\xbe\x3b\x1a\xb5\x54\xc4\x82\x31\x0b\x37\x4e\xfa\x9b\xab\x8a\x47\x56\xc7\x69\x39\x66\xa6\xb3\xb7\xc7\x03\x6f\x34\xda\xdb\x33\x3b\x6a\x89\x4b\x6e\x6e\x0e\xf6\x45\xa6\x7d\xe3\xf6\x54\xe0\xf6\xcf\xa5\xe9\x98\x0c\x4e\xc5\xc9\x50\x92\xca\xf4\x6e\x23\x30\xe1\x3b\x8e\xe6\x54\x3c\x33\xca\x2b\x93\x35\x74\xb2\x4d\xf6\x45\x73\x72\xbf\x2f\xc8\x9b\xa6\x24\x52\x6f\x94\x09\x74\x76\xe9\x6c\x39\x3e\x67\x5e\x1a\x1f\x29\x82\xeb\x14\x73\xd9\xb6\xd2\x20\x2a\x9f\x73\x6d\x4d\x15\xc1\xb3\xe9\x68\xea\x05\xee\x24\x3f\x2e\xe2\x44\x71\xbb\x1b\x06\x9a\xed\x70\x15\xf6\x79\x5a\xc8\x31\x97\x9c\x29\xae\x9c\xc1\x31\xbb\x69\x1f\x5e\x73\x5c\xa5\x1c\x26\xa5\xed\x2a\x65\xed\xdb\x45\x6b\xeb\x47\x57\x1f\x3b\x7f\xe1\xc9\x6b\xc7\xd7\x9f\x5b\xf7\xaf\x75\xcf\x5e\xbd\x10\xeb\xad\x2b\x83\xad\xdd\x44\x3d\x11\x3e\x7c\xe5\x99\x20\x3e\xbf\x7b\x44\x6d\x3f\xd1\xd4\x5d\xd5\xf3\x1f\x19\xe8\xe7\xe5\xa5\x2b\x6b\x7a\x70\xfd\xea\x95\x6b\xfe\xb1\xf5\xf3\x0f\x6f\x45\xd6\x07\xee\xde\xbb\xf5\xc2\x78\xa0\xec\x5e\x18\xf6\x24\x67\x91\x50\xa9\x1b\xae\x52\x9f\xed\xb2\xbe\x90\xc3\xf6\xc5\xcb\xcb\xcf\xb2\x40\x9d\x3c\xda\x6c\xae\x1c\x6d\x36\x85\x66\x52\xb8\x2b\x27\x9a\xcd\xcc\x84\x16\x5a\xf2\x74\x67\x5c\x36\x4f\x34\x1a\xd1\x2b\x74\x2e\xbb\xa6\x67\x73\x68\x65\x59\x68\xe5\xdf\x08\x0e\x79\xa1\x6b\xaa\x9a\xb2\xfe\x7f\x02\x00\x00\xff\xff\xe3\x10\x8f\x04\x00\x30\x00\x00")

func DirlistHtmlSwpBytes() ([]byte, error) {
	return bindataRead(
		_DirlistHtmlSwp,
		".dirlist.html.swp",
	)
}

func DirlistHtmlSwp() (*asset, error) {
	bytes, err := DirlistHtmlSwpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".dirlist.html.swp", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1526746657, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dirlistHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x51\x73\xdc\xb6\x11\x7e\x16\x7f\xc5\x9a\xd5\x64\x78\xa3\x13\x78\xb6\x6c\x59\xb1\xc8\xeb\xa4\x8e\x9b\xb4\x4e\x14\x45\x76\xaa\x8c\x3c\x9e\x09\x8e\x5c\x92\x90\x40\x80\x02\xc0\x3b\x5d\xaf\xfc\xef\x9d\x05\x49\xdd\xf5\xa2\xb6\x7e\xc9\xbd\x90\x58\x7c\xf8\x76\xf7\xc3\x2e\x88\x4b\x9e\xe5\x3a\x73\xeb\x06\xa1\x72\xb5\x9c\x07\xc9\xf8\x40\x9e\xcf\x03\x80\xc4\x09\x27\x71\xbe\xd9\x00\xfb\x48\x6f\xd0\x75\xf0\x2f\xf8\x91\xb7\x39\x26\x71\x3f\x47\x28\x29\xd4\x1d\x18\x94\x69\x68\xdd\x5a\xa2\xad\x10\x5d\x08\x95\xc1\x22\x0d\x2b\xe7\x1a\xfb\x26\x8e\x0b\xad\x9c\x65\xa5\xd6\xa5\x44\xde\x08\xcb\x32\x5d\xc7\x99\xb5\x7f\x2e\x78\x2d\xe4\x3a\xbd\xfc\x78\xf4\x81\x2b\xfb\xe6\xe5\x6c\x36\x7d\x39\x9b\x09\xc7\xa5\xc8\xa6\xaf\x67\xb3\xf0\x8b\x5d\xb4\x16\x19\xb9\xe1\x2b\xb4\xba\x46\xef\xc1\xa0\x44\x6e\xd1\xc6\xcb\x57\x6c\xc6\x9e\x9f\x90\xcb\x98\x4b\xc9\x32\x6b\x43\x10\xca\x61\x69\x84\x5b\xa7\xa1\xad\xf8\xc9\xd9\xcb\xe3\x6f\x2f\x7e\xfa\xfe\xe6\xf4\xec\x97\xb3\xea\xa6\x78\xff\xeb\x4f\xc6\xdd\x5e\x2f\x6f\x1f\x5a\xfb\x9d\xfe\xfa\xfa\x67\x65\x2e\x1e\x5e\xd8\xfb\xef\x66\xae\xb0\x65\xf5\xcd\xd2\xfd\x43\x5e\x5d\x9f\xb8\xe5\xdd\xaf\xd7\x37\xd5\xab\xb3\x8b\xaf\x6f\x9b\x10\x32\xa3\xad\xd5\x46\x94\x42\xa5\x21\x57\x5a\xad\x6b\xdd\xda\x70\x1e\xfc\xbf\x34\x48\xe5\x0f\xde\x0a\x5d\x17\xce\x83\x24\xee\xb7\x21\x59\xe8\x7c\xed\x45\x68\x25\x64\x92\x5b\x9b\x86\x0b\x83\x3c\xcf\x4c\x5b\x2f\xac\xd7\x07\x60\xb3\x01\xc3\x55\x89\x70\x28\x54\x8e\x0f\x53\x38\x44\x89\x35\x2a\x07\x6f\x52\x60\x97\xdc\x38\xa2\xf5\x50\x1f\xc7\x7c\xb3\x11\xc5\x00\xee\xba\xc4\x36\x5c\x8d\xe4\x9e\xf7\xd8\x62\x13\xce\xe3\x24\xa6\x99\xf9\x66\x83\x2a\xef\xba\xcd\x06\x68\xd1\xc0\xcc\x2e\xb9\xab\xa0\xeb\x12\xbe\xcd\x60\x7f\x2e\x9c\xef\x1a\x2f\x78\x4d\x35\x94\xc4\x9c\xcc\x28\x2d\x8d\x9e\x00\xd0\xa4\xca\x3d\x52\x8a\xc7\xfc\x7a\x13\x09\x11\xb7\xd2\x0b\x22\x54\xd3\x3a\x10\x79\x1a\x5a\xe4\x26\xab\x42\xa0\x62\x4e\x43\x87\x0f\x2e\x04\x2f\x71\x1a\xe6\xc2\x36\x92\xaf\xdf\x28\xad\x30\x84\x78\x1e\x04\x49\x2e\x96\x63\xb2\xb9\x30\x52\x58\x77\xbc\x32\xbc\x69\xd0\x84\x7b\x42\x0f\xd3\xde\xfc\x28\x31\xfb\x41\x58\x37\xaa\x49\x5a\xee\x2a\xc0\xbe\xbf\x7a\xf7\x57\x9f\x79\x22\x46\x96\x82\x43\xc1\x8f\x7d\x1b\x51\xb3\xf9\xc9\x58\xf8\xbe\xda\x91\x64\x4c\xf6\xf7\xa9\x26\x71\x2e\x96\x7b\x81\x53\x05\xd4\xe8\xf5\x65\x57\xfe\xdd\xd3\x0c\x40\x9b\x19\xd1\xb8\x00\xc0\x9a\x6c\xdb\x21\x99\xce\x91\xdd\xde\xb7\x68\xd6\xbe\x3b\xfa\xd7\xe3\x13\x76\xc2\x9e\x33\x2b\x45\xcd\x6a\xa1\xd8\xad\x0d\x03\xd8\xeb\x8d\x17\xaf\x4e\x8f\x4f\x30\x37\xf5\xba\xfd\x79\xb6\x3a\x7d\x55\x9c\x95\xc5\x5f\xec\xbd\x5e\xfd\xf3\xf6\xef\xf8\x42\xfc\x78\xaa\x66\xea\x7d\x26\x2e\x7f\x69\xce\xd6\x47\xaf\xdf\xa5\xc4\xf0\xdf\x3a\x21\x89\xfb\xe8\xb6\x71\xee\x6c\x5b\x7c\xcb\x97\xbc\xb7\x52\xcf\x2c\xb9\x01\xc5\x6b\xb4\x90\xc2\xa7\xcf\x7e\x88\x72\x18\xf8\x11\x97\x4d\xc5\xcd\x03\xa4\xa0\x70\x05\x57\x58\xbe\x7b\x68\xa2\xf8\xd3\x37\xc7\x37\x9f\xe3\x89\x47\xf4\xa5\x01\x29\x84\x61\x10\x1c\x46\x21\x1b\xf6\x14\xa4\x00\x1e\x4e\x18\xf2\xac\x8a\x8a\x56\x65\x4e\x68\x15\x89\x29\xa0\x9c\xc0\x26\x00\xa0\xc5\x14\x12\xa4\x70\x18\xa1\x9c\x30\x1a\x44\x13\xe6\xf4\x0f\x7a\x85\xe6\x2d\xb7\x18\x4d\x02\xe8\xe3\x63\x4d\x6b\xab\x88\x10\x64\x42\x39\x18\x50\x4e\x82\x6e\xe2\xfd\x52\x1b\x87\x13\x76\x87\xeb\xc6\xa0\xb5\x5b\x97\xb8\x24\x7f\x24\x98\x56\x56\x4b\x64\x52\x97\x11\x2e\xd9\xaa\x12\x59\x35\xa1\x09\x51\x00\x19\x6a\x74\xfc\x3d\xae\xfb\xe8\x00\x0c\xba\xd6\xa8\x00\xa0\x0b\x86\x68\x33\x48\xe1\x83\x33\x42\x95\xac\x30\xba\x7e\x5b\x71\xf3\x56\xe7\xb8\x43\xd6\x73\x0d\xa2\x31\x87\xd6\x45\xd9\x64\x24\xa4\xe5\xd9\xef\xd2\xf3\xec\x83\x88\x47\x29\x64\x01\x40\xae\xfb\x71\xd4\x3f\x76\x63\xf4\x7e\x20\x4d\xe1\xf9\xc9\x48\xbb\x87\x9e\x82\x33\x2d\x4e\x9e\x48\x21\x8e\x21\x8c\x43\xb8\xc3\x35\x14\x7a\xdc\xb8\x27\xa8\x5f\xbe\x1e\xa9\x71\xc9\x1a\x83\x4b\x54\xee\x5b\x2c\x78\x2b\x5d\xd4\xf3\x1e\x46\xe1\x9f\x86\x23\x61\xc2\x0a\x9d\xb5\x76\x48\x85\xf6\x22\x18\x95\xdf\x46\x76\x3f\x85\x52\xf7\xa4\xbb\x9b\x70\x4f\x8b\x28\x94\x88\xd4\x15\x90\xc2\xec\x1c\x04\x24\xc3\x96\x4b\x54\xa5\xab\xce\x41\x1c\x1d\x8d\x01\x8d\xf5\x4a\xf5\x48\x98\x4f\xe2\xf3\xa3\x1d\x25\xa4\x54\x19\xa3\x8d\xd2\x22\x10\xb3\x8e\x8e\xe7\x6b\xe1\xaa\xe8\xfe\x71\x33\xfa\xf9\x31\xaa\xfe\xd7\x97\x61\x26\x45\x76\x37\x24\xfa\x1f\x22\xd2\x6f\x3c\xe2\x7b\xe8\x36\xf5\x3d\x60\x37\x6a\xee\xcc\x7a\x27\xf4\xfd\x46\x1a\xf7\x17\xa0\x83\x8c\xbb\xac\x7a\xb2\xf4\xfe\x50\x85\xc6\x32\x25\xf0\x8e\x38\x5f\x94\x5f\xf7\x64\xe3\xe5\xb8\xd0\xad\xca\x70\xbf\x03\xfb\x18\xee\x21\xdd\x56\xde\xee\xd9\xd1\x4d\xe1\x74\x36\x9b\x4c\x82\x20\x8e\xe1\xca\xbb\xb2\xc0\x61\x24\x99\x82\xab\xb8\x9b\x02\xb7\x20\xb5\x2a\xe9\x29\x1c\x55\x93\x13\xaa\x45\x0b\x4e\xc3\x02\x41\xa8\xa5\xbe\xc3\x7c\x0a\x2b\x21\x25\x28\xed\x88\x6c\x81\xe0\x8c\x28\x4b\x34\x98\x33\xf8\x58\xe1\x23\x69\x0f\x5b\x20\x64\x5c\x4a\xcc\x81\x17\x0e\x0d\xf1\x5a\xa7\x1b\x0b\x0b\x14\xaa\x1c\xe7\x0a\x6d\x88\xec\x02\x6a\x21\xa5\xb0\x98\x69\x95\x5b\x06\x7f\x2b\xe0\x37\x51\xd7\x98\x0b\xee\xf0\x37\x10\x16\x1a\x6e\x2d\x85\x30\xf8\x04\xb7\xeb\x50\x2b\x1a\x13\x91\x44\x9e\x13\x3d\xe6\x25\x4e\x41\x28\xeb\x90\xe7\xa0\x0b\x8f\x77\x86\x0b\x49\xe7\xcc\x4e\x33\xed\xca\x3a\x85\x15\x17\x6e\x0a\x8f\x9e\x49\xe0\x03\x7f\xa2\x8a\x1a\x75\xeb\xce\x83\x83\x7e\xbb\x1e\x5d\x47\x1e\xe2\x31\xa4\x5a\x7f\xf0\xba\x4a\xd8\x29\x70\x53\xd2\x99\xcf\x4d\xd9\xd2\x35\xc1\x9e\x0f\x38\xc9\x49\x8e\x74\x9f\xe2\x60\xf0\x41\x25\xd6\x4a\x49\xe8\x03\x2a\xa5\x67\x3b\xd1\xd0\x12\xc6\x9b\x46\xae\xa3\xc1\x5b\xef\x66\x42\xe8\x6e\x74\x40\xd2\x5e\xe8\x15\xa4\xdb\x44\xe0\xab\xaf\xe0\xd9\x36\x89\x83\x4c\x22\x37\x1f\xfb\x71\x34\xd8\x3d\xc9\x36\x08\x8b\x6e\x04\xf8\x88\x7b\x71\x3c\x88\xc2\x1a\x9c\xfc\xef\x98\xba\xf3\xa0\x3b\x0f\x76\x3f\x9d\x85\xd6\x0e\x8d\xbf\xaf\x34\xf3\x4b\x3a\xb2\x31\x87\xc5\x1a\x1e\x2f\x23\xe3\x57\xbf\x14\xae\x6a\x17\xfe\x83\xbf\x68\xb3\x8a\x2b\x8e\x71\x4d\x77\xf8\x70\x5e\xf7\x57\x79\xba\x7b\x34\x74\xcd\x18\x39\x83\x24\xee\x6f\x9d\x49\xdc\xff\x25\xf8\x77\x00\x00\x00\xff\xff\xcb\x6b\x11\x38\x2a\x0c\x00\x00")

func dirlistHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dirlistHtml,
		"dirlist.html",
	)
}

func dirlistHtml() (*asset, error) {
	bytes, err := dirlistHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dirlist.html", size: 3114, mode: os.FileMode(436), modTime: time.Unix(1526746657, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _notebookCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x96\xcf\x6f\xa3\x38\x14\xc7\xef\xf9\x2b\x2c\x8d\xe6\x32\x12\x19\x60\x27\x49\x87\x9c\x66\x9a\x76\x3b\xda\xee\x74\x34\xad\xb4\x67\x63\x3f\xc0\x8a\xb1\x2d\xdb\x34\x89\xaa\xfe\xef\x2b\x42\x68\xb0\x03\x94\xdd\x72\x68\x42\xde\xe7\xfb\x7e\xf2\x4c\x2a\xe9\x01\xbd\xcc\x10\x4a\x31\xd9\xe6\x5a\x56\x82\x06\x44\x72\xa9\x13\xf4\x21\x4a\xeb\x6b\x3d\x7b\x9d\xcd\xb5\x94\x36\x20\x52\x58\x10\xf6\x68\xbe\x63\xd4\x16\x09\x5a\x2d\x3e\xae\x67\x08\x95\x58\xe7\x4c\x24\x28\x44\xb8\xb2\xf2\x48\x10\xe0\x3c\x30\xb2\xd2\x04\xd0\x9c\x14\x5a\x96\xf8\x08\x2a\x4c\x29\x13\x79\x82\xe2\x50\xed\x51\x14\xaa\x7d\x2d\x90\x4a\x4d\x41\x07\x1a\x53\x56\x99\x04\xfd\x51\xdf\x6d\x45\x64\x65\x55\xd5\x78\x6d\xfd\x2c\xd4\x1e\x85\x35\xb7\x2b\x98\x85\xc0\x28\x4c\x20\x41\x4a\x43\xb0\xd3\x58\xd5\x3f\xc8\x67\xd0\x19\x97\xbb\x60\x9f\x20\x43\xb4\xe4\x7c\xdd\x9f\x24\x00\x0c\x05\xd0\x09\xb6\x89\xf3\x75\x66\x71\xca\xa1\xa9\x57\x03\x10\xc9\x39\x56\x06\x12\xd4\x7e\x3a\xab\x25\x28\x52\x7b\x64\x24\x67\x14\x7d\xc0\x71\x7d\xf5\x07\x71\xcc\x62\x3d\xab\xf5\xa9\x5b\xa4\xc5\x5b\x8d\x5e\x67\xb6\x18\x6a\xd4\x29\x87\x5e\xea\xf8\xf7\xf9\x13\xfa\xfe\x86\xa1\x4f\x9f\xcf\x1d\x41\xad\x46\x76\x95\x5d\x65\xf1\xba\x4f\x3e\x5e\xc5\x57\x71\x8c\x5e\x6b\x99\x1b\xad\xa5\xee\x2a\xcc\x41\xeb\x8e\xcc\xd7\x65\x18\x2e\xc2\x5e\x99\x08\xc2\x30\x0a\x1b\x99\x7b\x26\xe0\xa9\xae\xe5\xd3\xc6\x11\xe3\xa2\xae\x00\x7a\x06\x6d\x19\xc1\x3c\xc0\x9c\xe5\x22\x41\x56\xaa\xf5\x39\xbd\x70\x7d\x1e\xb8\xf5\x5b\xad\xc3\xb5\x27\xed\x0b\x37\xad\x6b\x1b\x57\xcf\x4c\xab\xf6\xbe\xf2\x69\xdc\x8f\xd3\xfd\x36\x5b\xed\x77\xca\x8c\xe2\xf8\x90\xa0\x94\x4b\xb2\xed\x84\x71\xc7\xf2\x82\xb3\xbc\xb0\x4e\x28\x05\x47\x2f\x17\xcc\xc9\x41\x14\x86\x1f\x3b\x02\x3f\xab\x32\x05\x6d\x7a\xd3\x41\x2f\xa7\x58\x03\x5d\xbb\x48\x50\x38\xff\x02\x65\x37\x99\xe6\x4e\xfb\xff\x52\xd6\x53\xfc\x9f\x82\x7f\xc1\x61\x27\xb5\x33\x55\xf3\x6d\x67\x22\x96\x4b\xfa\x15\x32\xc7\xf6\x5a\x0a\x63\xb1\x70\xcb\xb2\x25\xe3\xd0\x06\x08\xc7\x1a\x5b\x26\x85\xcb\xd1\x71\xee\x27\x2e\xe1\xb8\x20\x5c\x4a\x74\x67\xff\x6b\xbc\x5c\xc5\x0e\xf5\xcb\x40\x45\xa5\x8b\xa8\x71\x47\xbf\xc1\x80\x7e\x06\xaf\x12\x7a\x1c\x7a\x3a\x28\x2f\x30\x3b\x00\xd4\x79\x7c\xb3\x56\xb3\xb4\xb2\x2e\x22\xba\xcf\x31\x5e\x42\x1c\xc3\x19\xb9\xe6\xd8\xb8\x9d\x16\x64\xcc\xbc\xaf\x35\x42\x8e\xc4\xb4\x01\x22\x35\xb6\xde\x56\x10\x74\xc4\xc9\xcd\x9e\x80\xba\x68\xa4\x80\x11\xe4\xb6\x12\xe4\x92\xc8\x46\x88\x07\x5b\x80\x17\xd3\x7e\xc4\xfc\x09\xe7\xae\xb1\x1d\x18\x90\x7b\x66\x41\x63\xee\x3e\x3c\x5d\x61\xb8\x8a\xb2\xcc\xb1\xdd\x60\xaf\x63\xbc\x5b\x1d\x58\xd2\x74\xf5\xc5\x01\x1e\xad\x66\xc2\x8d\xc7\x4c\x21\xbe\x65\x19\xdb\xbb\x18\x9e\xc2\xd5\x87\x83\x65\x64\xeb\xa2\xe9\x14\xf4\xba\xc0\x6e\x95\x0d\x99\x82\x6d\x80\xb3\xb2\xbe\xe1\xb0\x94\x4f\x62\x25\x71\x3d\x4e\xaa\xe6\x46\x56\xfe\x16\x35\xf1\x14\xf0\xc6\x10\xec\x3d\xa4\x06\xc6\x1b\xde\x80\x77\xa0\x81\xfa\xb1\x16\x53\x5c\xfe\x10\x16\xb4\x92\xee\x90\x19\x36\x05\xbd\x1c\x7b\xb3\x9f\xc2\xfd\x86\x1c\xbc\xd1\xd1\x53\xb8\x47\x26\x72\xbf\xac\xd1\x24\xf0\x50\xa6\x7e\x86\xef\x0c\x79\x73\x6e\x39\x48\x39\xde\x88\x86\xf8\xce\xdc\xbd\x51\xa6\x53\xa8\x5b\x2e\xb1\xbb\x07\xcb\x6c\x0a\x77\xe7\xd5\xb1\x2c\xa6\x50\x75\xcb\x73\x3f\x39\xf6\x1f\xc8\x7b\xe9\x6d\x0c\xf6\xce\x56\x6a\xe8\x07\xe2\xe5\x28\x07\xa8\x07\x05\x97\x6b\x5e\x0e\x2c\xc9\xd6\xf8\x1f\xff\xcd\x40\xee\x06\x88\x6b\x59\x96\xe0\x9d\x3b\xdd\x4d\xb2\x5a\xac\xa2\x05\x38\xb6\x77\xd8\x14\x29\xf6\xb2\x26\xc5\x38\xf4\x77\xc5\x2d\xe3\x4c\xb8\x23\x4b\xca\x71\xaa\x67\xca\x49\xf4\x0e\xa2\x80\x30\xef\x9c\x20\x66\x9c\xf9\xa5\x41\x69\x6f\x63\x10\x35\x89\xb9\x65\x7e\x7c\x2a\x1b\x00\xff\x04\x01\x9a\x91\x0d\x70\xb0\xde\xeb\x4a\x4e\x07\xda\x73\x62\x6e\x4a\x55\xb8\x40\xbd\x09\x33\x29\x6c\x60\xec\x81\x43\x82\x98\xc5\x9c\x11\x07\xfa\x21\x0c\xe8\x0b\x4f\x6c\xe0\x30\x3e\x41\x8f\x56\xfb\x03\x9d\x9b\xd6\xd7\x0e\x9a\x37\xd4\x54\x72\xea\x52\x55\x5a\x00\xa6\xfe\xe1\x99\x57\x7d\xa5\xf8\x37\x00\x00\xff\xff\xb5\xae\x1e\x33\xfc\x0e\x00\x00")

func notebookCssBytes() ([]byte, error) {
	return bindataRead(
		_notebookCss,
		"notebook.css",
	)
}

func notebookCss() (*asset, error) {
	bytes, err := notebookCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "notebook.css", size: 3836, mode: os.FileMode(436), modTime: time.Unix(1526236113, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _notebookHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xcf\x6a\xf5\x20\x10\xc5\xf7\x79\x8a\xc1\x7d\x14\xbe\x8f\x6e\x8a\xc9\xa2\x85\xd2\x55\xbb\x68\x5f\xc0\xab\x73\x6f\x4c\xbd\x9a\x3a\x93\x4b\x43\xc8\xbb\x17\x0d\xf4\xff\x4a\x38\x33\xbf\x73\xce\xa8\x07\x34\xae\x6f\x00\x74\xf0\xf1\x05\x32\x86\x4e\x10\x2f\x01\x69\x40\x64\x01\x43\xc6\x63\x27\xd6\x15\xe4\x53\x55\x61\xdb\x44\xdf\x94\x7d\xb2\xd9\x4f\x0c\x94\x6d\x27\x06\xe6\x89\xae\x95\xb2\x2e\x8e\x24\x6d\x48\xb3\x3b\x06\x93\x51\xda\x74\x56\x66\x34\x6f\x2a\xf8\x03\xa9\x8c\xaf\xb3\xcf\x28\x47\x52\xff\xe4\x7f\x79\xf5\x21\x9c\x7d\x94\x23\x89\x5e\xab\xdd\xb5\x6f\xb4\xda\x8b\x35\xfa\x90\xdc\x52\x5e\xe7\x2f\x60\x83\x21\xea\x44\x4e\x89\x5b\x9b\x22\x63\x64\xd1\x37\xeb\x0a\xd9\xc4\x13\x82\x7c\xb8\x91\xb7\x18\x42\x69\xf9\x0d\xb0\x18\x02\x94\x23\xca\xf4\x79\x99\x10\xb6\xad\x2d\xa2\xa8\xa7\xff\xd8\x6c\x29\xcd\xd9\xa2\xe8\x0b\x71\xef\x4f\x43\xf0\xa7\x81\xd1\xc1\xb6\x69\xe5\xfc\xa5\x30\x9f\x99\x8f\x33\x4f\x33\xd7\xc8\x3f\xac\x52\x9d\xd6\x98\x02\xb5\x20\xef\x52\x3e\x1b\x2e\x76\xed\x8e\x7c\xb1\xc4\xe8\x6a\xf5\x5d\xfa\x25\x34\x5a\xed\xbf\xf1\x1e\x00\x00\xff\xff\x94\xc7\xc8\xcd\xb6\x01\x00\x00")

func notebookHtmlBytes() ([]byte, error) {
	return bindataRead(
		_notebookHtml,
		"notebook.html",
	)
}

func notebookHtml() (*asset, error) {
	bytes, err := notebookHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "notebook.html", size: 438, mode: os.FileMode(436), modTime: time.Unix(1526236113, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pageHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x41\x6f\xd4\x3c\x14\xbc\xe7\x57\xf8\xb3\x7a\xfb\xda\x38\xd0\x16\xb6\x95\x63\x84\x40\x2a\x12\x62\xbb\xd0\xc2\x56\xbd\x79\x93\x97\xc4\xad\x63\x07\x3f\x67\xbb\x51\xc8\x7f\x47\x76\x48\xe1\x06\xa7\x24\xef\x8d\x66\x26\xe3\x31\xff\xaf\xb4\x85\x1f\x3a\x20\x8d\x6f\xb5\x48\xf8\xf2\x00\x59\x8a\x84\x10\xee\x95\xd7\x20\xc6\x91\xa4\xb7\xe1\x8d\x4c\x13\xf9\x41\x3e\xc9\xbe\x04\xce\xe6\x5d\x40\x69\x65\x1e\x89\x03\x9d\x53\xf4\x83\x06\x6c\x00\x3c\x25\x8d\x83\x2a\xa7\x8d\xf7\x1d\x5e\x32\x56\x59\xe3\x31\xad\xad\xad\x35\xc8\x4e\x61\x5a\xd8\x96\x15\x88\x6f\x2a\xd9\x2a\x3d\xe4\x9b\xdb\xff\x6f\xa4\xc1\xcb\xb3\x2c\x3b\x3e\xcb\x32\xe5\xa5\x56\xc5\xf1\xeb\x2c\xa3\xff\x2c\xd1\x23\xa4\x41\x46\x3e\x01\xda\x16\xa2\x82\x03\x0d\x12\x01\xd9\xfe\x3c\xcd\xd2\x17\xa7\x41\x92\x49\xad\xd3\x02\x91\x12\x65\x3c\xd4\x4e\xf9\x21\xa7\xd8\xc8\xd3\xd5\xd9\xc9\xfb\xf5\xf5\x87\xfb\x57\xab\xaf\xab\xe6\xbe\xfa\x78\x77\xed\xfc\xc3\x76\xff\x70\xe8\xf1\xca\x5e\x6c\x3f\x1b\xb7\x3e\xbc\xc4\xef\x57\x99\xaf\xb0\x6e\xde\xee\xfd\x37\xfd\x65\x7b\xea\xf7\x8f\x77\xdb\xfb\xe6\x7c\xb5\xbe\x78\xe8\x28\x29\x9c\x45\xb4\x4e\xd5\xca\xe4\x54\x1a\x6b\x86\xd6\xf6\x48\x45\xf2\xb7\xdf\x08\x29\xdf\xc4\x29\x99\x26\x2a\x12\xce\xe6\x63\xe0\x3b\x5b\x0e\x31\x84\x5e\x93\x42\x4b\xc4\x9c\xee\x1c\xc8\xb2\x70\x7d\xbb\xc3\x98\x0f\x21\xe3\x48\x9c\x34\x35\x90\x23\x65\x4a\x38\x1c\x93\x23\xd0\xd0\x82\xf1\xe4\x32\x27\xe9\x46\x3a\x1f\x68\x23\x34\xfa\x10\xe3\xa8\xaa\x5f\xe0\x69\xe2\xd8\x49\xb3\x90\x47\xde\x13\x84\x8e\x0a\xc6\x59\xd8\x88\x71\x04\x53\x4e\x13\x97\xbf\xbd\x2e\xfc\xe9\x46\xfa\x26\x3a\xfe\x73\xb8\x96\x6d\x68\x0b\x67\x52\x70\xa6\xd5\xb3\x47\x30\xe5\x6c\x83\xb3\x5e\x8b\x24\xe1\xa5\xda\x2f\xba\x9d\xac\xe1\xa4\xb0\xc6\x83\xf1\x91\x2d\x7d\x37\x7f\x44\xa2\x52\xed\x03\xbe\xb2\xd6\x83\x8b\x71\x74\x62\x63\x9f\xc0\x41\x49\x76\x03\x79\xb6\xb6\xb4\xa1\x56\xbe\xe9\x77\xb1\x04\xbb\xbe\x68\xa4\x91\xc0\xda\xd0\x5c\x2a\xda\xb9\xc0\xc1\x5a\x17\x72\x5e\x38\x13\xce\xe6\xac\x39\x9b\x2f\xc2\xcf\x00\x00\x00\xff\xff\x92\x39\x04\x42\x20\x03\x00\x00")

func pageHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pageHtml,
		"page.html",
	)
}

func pageHtml() (*asset, error) {
	bytes, err := pageHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "page.html", size: 800, mode: os.FileMode(436), modTime: time.Unix(1526056397, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _styleCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x5b\x8e\xdb\x30\x0c\xfc\xf7\x29\x88\x0d\xfa\x57\x05\xde\x0d\xd2\x24\xca\x25\x0a\xb4\x17\x90\x25\xda\x21\x22\x4b\x82\x24\x77\x9d\x16\xb9\x7b\x61\x5b\x7e\x65\xb3\x58\xf8\x23\x10\x45\x72\x86\x9c\x51\x0a\xab\x6e\xf0\x2f\x03\x28\xad\x89\xac\x14\x35\xe9\x1b\x87\x97\x9f\xbf\xe1\x97\x30\xe1\xe5\x3b\x04\x61\x02\x0b\xe8\xa9\x3c\x67\xf7\x6c\xab\xc8\x6b\x0a\x91\xbd\x7b\xe1\x1c\xfa\xbe\xb4\x16\xbe\x22\xc3\x61\x9f\xbb\x16\x44\x13\xed\x39\x03\x78\x27\x15\x2f\x1c\x7e\xe4\xdf\x96\x75\x7d\x7e\xdf\x20\xc4\x9b\x46\x16\x6f\x0e\x39\x18\x6b\x70\x95\xa5\xa9\x4f\x2c\xac\x57\xe8\x99\x17\x8a\x9a\xc0\x61\xe7\xda\x87\x2c\x7e\xb1\x7f\x12\x89\x42\xc8\x6b\xe5\x6d\x63\x14\x93\x56\x5b\xcf\x61\x83\xc7\xee\x7b\x6c\x2c\xfa\xf4\x88\x6d\x64\x0a\xa5\xf5\x22\x92\x35\x23\x07\x80\xb1\x38\xcf\x4f\x27\x25\xbb\x88\xa2\xe0\xb4\xb8\x71\x28\xb4\x95\xd7\x2e\xe2\x84\x52\x64\x2a\x0e\x07\xd7\xc2\x5b\xfe\x81\x16\xd0\xb6\x1c\x60\xa6\xcc\x1c\xf6\x5d\x1a\x40\x4d\x86\xa5\xdd\x0c\x95\x33\x64\x71\xe8\xbe\xf3\xc8\x4e\x68\xaa\x0c\x07\x89\x26\xa2\xe7\xd9\x3d\xeb\xba\x32\x6c\x23\x7a\x23\x34\xd3\x64\xae\x4c\xe8\x38\xab\x17\xe8\x2f\x72\xd8\xee\xb1\xee\xf9\x14\x1e\x85\x92\xbe\xa9\x8b\x30\x2e\xf4\x93\xcd\xaf\xeb\x4f\x5d\xfd\x62\x6c\x32\x9a\x92\x3e\xcb\x96\xdb\xfe\x97\x05\x74\x4f\x27\xbd\x67\xd9\xb6\xcb\xae\xb1\xbf\x5e\x1b\xa2\x3b\x8f\xae\x39\x2c\x5d\x33\x69\x1e\xad\xe3\xf0\xea\x5a\x08\x56\x93\x82\x0d\x22\x0e\xf7\x13\xce\x6e\xdc\xbb\x13\x15\x32\x69\x4d\x44\x13\xbf\xc0\xca\x17\x40\x1f\x1a\x65\xa5\xb5\x31\xb9\x69\x16\x45\x29\xf5\x5c\x91\xd4\xc5\x06\x1a\x0c\x54\x52\x8b\x6a\x9c\x21\x46\x5b\x73\xc8\x87\xa3\xa7\xea\x12\xa7\x53\xc2\x65\x29\xba\x4b\x26\x58\x6b\x70\x1c\x34\x1c\x29\x89\x07\x52\xc7\x83\x28\xa4\xec\x33\xa2\x7a\x5c\xff\x21\x8d\xb3\x09\x28\xbc\xbc\xa4\xd2\xa7\x44\xd7\xcc\xfa\x9d\xe7\xeb\x8d\xbd\x4d\xf4\x26\xd3\x4e\x91\x67\xcf\x73\x8e\xaf\xe4\x4b\x5b\xec\x1c\x68\x90\x5d\x70\xc0\x7d\x1d\xe6\x5c\xa9\xb1\xb0\xcf\xf2\xd1\xf2\xd2\xca\x26\x7c\xf6\xd2\x65\x59\xee\xca\xf2\x9c\xcd\xd8\xa3\xb1\x6d\x13\x3b\xc8\xf9\x3f\xe6\x7f\x00\x00\x00\xff\xff\xc3\x26\x13\x83\xf1\x04\x00\x00")

func styleCssBytes() ([]byte, error) {
	return bindataRead(
		_styleCss,
		"style.css",
	)
}

func styleCss() (*asset, error) {
	bytes, err := styleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "style.css", size: 1265, mode: os.FileMode(436), modTime: time.Unix(1526701016, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tableHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x41\x6f\xd4\x3c\x14\xbc\xe7\x57\xf8\xb3\x7a\xfb\xda\x38\xd0\x16\xb6\x95\x63\x84\x40\x2a\x12\x62\xbb\xd0\xc2\x56\xbd\x79\x93\x97\xc4\xad\x63\x07\x3f\x67\xbb\x51\xc8\x7f\x47\x76\x48\xe1\x06\xa7\x24\xef\x8d\x66\x26\xe3\x31\xff\xaf\xb4\x85\x1f\x3a\x20\x8d\x6f\xb5\x48\xf8\xf2\x00\x59\x8a\x84\x10\xee\x95\xd7\x20\xc6\x91\xa4\xb7\xe1\x8d\x4c\x13\xf9\x41\x3e\xc9\xbe\x04\xce\xe6\x5d\x40\x69\x65\x1e\x89\x03\x9d\x53\xf4\x83\x06\x6c\x00\x3c\x25\x8d\x83\x2a\xa7\x8d\xf7\x1d\x5e\x32\x56\x59\xe3\x31\xad\xad\xad\x35\xc8\x4e\x61\x5a\xd8\x96\x15\x88\x6f\x2a\xd9\x2a\x3d\xe4\x9b\xdb\xff\x6f\xa4\xc1\xcb\xb3\x2c\x3b\x3e\xcb\x32\xe5\xa5\x56\xc5\xf1\xeb\x2c\xa3\xff\x2c\xd1\x23\xa4\x41\x46\x3e\x01\xda\x16\xa2\x82\x03\x0d\x12\x01\xd9\xfe\x3c\xcd\xd2\x17\xa7\x41\x92\x49\xad\xd3\x02\x91\x12\x65\x3c\xd4\x4e\xf9\x21\xa7\xd8\xc8\xd3\xd5\xd9\xc9\xfb\xf5\xf5\x87\xfb\x57\xab\xaf\xab\xe6\xbe\xfa\x78\x77\xed\xfc\xc3\x76\xff\x70\xe8\xf1\xca\x5e\x6c\x3f\x1b\xb7\x3e\xbc\xc4\xef\x57\x99\xaf\xb0\x6e\xde\xee\xfd\x37\xfd\x65\x7b\xea\xf7\x8f\x77\xdb\xfb\xe6\x7c\xb5\xbe\x78\xe8\x28\x29\x9c\x45\xb4\x4e\xd5\xca\xe4\x54\x1a\x6b\x86\xd6\xf6\x48\x45\xf2\xb7\xdf\x08\x29\xdf\xc4\x29\x99\x26\x2a\x12\xce\xe6\x63\xe0\x3b\x5b\x0e\x31\x84\x5e\x93\x42\x4b\xc4\x9c\xee\x1c\xc8\xb2\x70\x7d\xbb\xc3\x98\x0f\x21\xe3\x48\x9c\x34\x35\x90\x23\x65\x4a\x38\x1c\x93\x23\xd0\xd0\x82\xf1\xe4\x32\x27\xe9\x46\x3a\x1f\x68\x23\x34\xfa\x10\xe3\xa8\xaa\x5f\xe0\x69\xe2\xd8\x49\xb3\x90\x47\xde\x13\x84\x8e\x0a\xc6\x59\xd8\x88\x71\x04\x53\x4e\x13\x97\xbf\xbd\x2e\xfc\xe9\x46\xfa\x26\x3a\xfe\x73\xb8\x96\x6d\x68\x0b\x67\x52\x70\xa6\xd5\xb3\x47\x30\xe5\x6c\x83\xb3\x5e\x8b\x24\xe1\xa5\xda\x2f\xba\x9d\xac\xe1\xa4\xb0\xc6\x83\xf1\x91\x2d\x7d\x37\x7f\x44\xa2\x52\xed\x03\xbe\xb2\xd6\x83\x8b\x71\x74\x62\x63\x9f\xc0\x41\x49\x76\x03\x79\xb6\xb6\xb4\xa1\x56\xbe\xe9\x77\xb1\x04\xbb\xbe\x68\xa4\x91\xc0\xda\xd0\x5c\x2a\xda\xb9\xc0\xc1\x5a\x17\x72\x5e\x38\x13\xce\xe6\xac\x39\x9b\x2f\xc2\xcf\x00\x00\x00\xff\xff\x92\x39\x04\x42\x20\x03\x00\x00")

func tableHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tableHtml,
		"table.html",
	)
}

func tableHtml() (*asset, error) {
	bytes, err := tableHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "table.html", size: 800, mode: os.FileMode(436), modTime: time.Unix(1526695644, 0)}
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
	".dirlist.html.swp": DirlistHtmlSwp,
	"dirlist.html": dirlistHtml,
	"notebook.css": notebookCss,
	"notebook.html": notebookHtml,
	"page.html": pageHtml,
	"style.css": styleCss,
	"table.html": tableHtml,
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
	".dirlist.html.swp": &bintree{DirlistHtmlSwp, map[string]*bintree{}},
	"dirlist.html": &bintree{dirlistHtml, map[string]*bintree{}},
	"notebook.css": &bintree{notebookCss, map[string]*bintree{}},
	"notebook.html": &bintree{notebookHtml, map[string]*bintree{}},
	"page.html": &bintree{pageHtml, map[string]*bintree{}},
	"style.css": &bintree{styleCss, map[string]*bintree{}},
	"table.html": &bintree{tableHtml, map[string]*bintree{}},
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

