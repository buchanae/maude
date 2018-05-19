// Code generated by go-bindata.
// sources:
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

var _dirlistHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x51\x73\xdb\xc6\x11\x7e\x16\x7e\xc5\x1a\xd5\x64\xc0\x11\x75\xa0\x2d\x5b\x56\x2c\x80\x9d\xd4\x71\x93\xd6\x89\xa2\xc8\x4a\x95\xb1\xc7\x33\x39\x02\x0b\xe0\xa4\xc3\x1d\x74\x77\x20\xc5\xb2\xf8\xef\x9d\x3d\x00\x22\x4b\xab\x6d\x5e\xc2\x17\xe2\xf6\x16\xdf\x7e\xfb\xed\xee\xf1\x98\x3c\xcb\x75\xe6\xd6\x0d\x42\xe5\x6a\x39\x0f\x92\xf1\x0b\x79\x3e\x0f\x00\x12\x27\x9c\xc4\xf9\x66\x03\xec\x9a\x9e\xa0\xeb\xe0\x5f\xf0\x23\x6f\x73\x4c\xe2\x7e\x8f\xbc\xa4\x50\x77\x60\x50\xa6\xa1\x75\x6b\x89\xb6\x42\x74\x21\x54\x06\x8b\x34\xac\x9c\x6b\xec\x9b\x38\x2e\xb4\x72\x96\x95\x5a\x97\x12\x79\x23\x2c\xcb\x74\x1d\x67\xd6\xfe\xb9\xe0\xb5\x90\xeb\xf4\xf2\xfa\xe8\x03\x57\xf6\xcd\xcb\xd9\x6c\xfa\x72\x36\x13\x8e\x4b\x91\x4d\x5f\xcf\x66\xe1\xef\x0e\xd1\x5a\x64\x14\x86\xaf\xd0\xea\x1a\x7d\x04\x83\x12\xb9\x45\x1b\x2f\x5f\xb1\x19\x7b\x7e\x42\x21\x63\x2e\x25\xcb\xac\x0d\x41\x28\x87\xa5\x11\x6e\x9d\x86\xb6\xe2\x27\x67\x2f\x8f\xbf\xbd\xf8\xe9\xfb\x8f\xa7\x67\xbf\x9c\x55\x1f\x8b\xf7\xbf\xfe\x64\xdc\xed\xcd\xf2\xf6\xa1\xb5\xdf\xe9\xaf\x6f\x7e\x56\xe6\xe2\xe1\x85\xbd\xff\x6e\xe6\x0a\x5b\x56\xdf\x2c\xdd\x3f\xe4\xd5\xcd\x89\x5b\xde\xfd\x7a\xf3\xb1\x7a\x75\x76\xf1\xf5\x6d\x13\x42\x66\xb4\xb5\xda\x88\x52\xa8\x34\xe4\x4a\xab\x75\xad\x5b\x1b\xce\x83\xff\x97\x06\xa9\xfc\xc1\x5b\xa1\xeb\x7a\xff\xcd\x06\x0c\x57\x25\x02\xfb\x41\x58\x07\x5d\x17\x00\x78\xab\x28\x00\xef\x81\x5d\x53\xe9\xc2\x42\xcb\x1c\x4d\x38\x6e\xef\x04\x69\x0c\x16\xe8\xb2\x6a\x37\xc4\xf7\x57\xef\xfe\x0a\x5d\x17\x7b\x5d\x3d\x18\xaa\xbc\x7f\x77\xfb\x9c\xc4\x7d\x0b\x24\x0b\x9d\xaf\x7d\x01\x5a\x09\x99\xe4\xd6\xa6\xe1\xc2\x20\xcf\x33\xd3\xd6\x0b\xbb\xc5\xe8\x69\x1e\x0a\x95\xe3\xc3\x14\x0e\x51\x62\x8d\xca\xc1\x9b\x14\xd8\x25\x37\xce\x8e\xe4\x3c\xbd\xf9\x66\x23\x8a\xc1\xb9\xeb\x12\xdb\x70\x35\x82\x7b\xdc\x63\x8b\x4d\x38\x8f\x93\x98\x76\xe6\x9b\x0d\xaa\xbc\xeb\xfa\xac\x47\x64\x76\xc9\x5d\x05\x5d\x97\xf0\x6d\x6a\xfb\x7b\xe1\x7c\xd7\x78\xc1\x6b\xea\xdf\x24\xe6\x64\x46\x69\x69\xf5\x84\xc3\xa3\x08\x49\x2c\xc5\x97\x1a\x25\x71\x2b\xbd\x20\x42\x35\xad\x03\x91\xa7\xa1\x45\x6e\x48\x63\x1a\xa4\x34\x74\xf8\xe0\x42\xf0\xe5\x4d\xc3\x5c\xd8\x46\xf2\xf5\x1b\xa5\x15\x86\x10\xcf\x83\x20\xc9\xc5\x72\x4c\x36\x17\x46\x0a\xeb\x8e\x57\x86\x37\x0d\x9a\x70\x4f\xe8\x61\xdb\x9b\x9f\xee\x04\xd2\x72\x57\x81\xb1\xb8\xe1\x3c\x11\x23\x4a\xc1\xa1\xe0\xc7\x7e\x84\xa9\x5b\xfc\x66\x2c\xfc\x4c\xef\x48\x32\x26\xfb\x65\xaa\x49\x9c\x8b\xe5\x1e\x71\xea\x80\x1a\xbd\xbe\xec\xca\x3f\x7b\x98\xc1\xd1\x66\x46\x34\x2e\x00\xb0\x26\xdb\x4e\x67\xa6\x73\x64\xb7\xf7\x2d\x9a\xb5\x9f\xcc\xfe\xf1\xf8\x84\x9d\xb0\xe7\xcc\x4a\x51\xb3\x5a\x28\x76\x6b\xc3\x00\xf6\xe6\xf2\xc5\xab\xd3\xe3\x13\xcc\x4d\xbd\x6e\x7f\x9e\xad\x4e\x5f\x15\x67\x65\xf1\x17\x7b\xaf\x57\xff\xbc\xfd\x3b\xbe\x10\x3f\x9e\xaa\x99\x7a\x9f\x89\xcb\x5f\x9a\xb3\xf5\xd1\xeb\x77\x29\x21\xfc\xb7\x29\x4c\xe2\x9e\xdd\x96\xe7\x4e\xd9\xe2\x5b\xbe\xe4\xbd\x95\xe6\x6f\xc9\x0d\x28\x5e\xa3\x85\x14\x3e\x7d\xf6\x4b\x94\xc3\xc2\xaf\xb8\x6c\x2a\x6e\x1e\x20\x05\x85\x2b\xb8\xc2\xf2\xdd\x43\x13\xc5\x9f\xbe\x39\xfe\xf8\x39\x9e\x78\x8f\xbe\x35\x20\x85\x30\x0c\x82\xc3\x28\x64\x43\x4d\x41\x0a\xe0\xe1\x84\x21\xcf\xaa\xa8\x68\x55\xe6\x84\x56\x91\x98\x02\xca\x09\x6c\x02\x00\x7a\x99\x28\x41\x0a\x87\x11\xca\x09\xa3\x45\x34\x61\x4e\xff\xa0\x57\x68\xde\x72\x8b\xd1\x24\x80\x9e\x1f\x6b\x5a\x5b\x45\xe4\x41\x26\x94\x83\x01\xe5\x24\xe8\x26\x3e\x2e\x8d\x71\x38\x61\x77\xb8\x6e\x0c\x5a\xbb\x0d\x89\x4b\x8a\x47\x82\x69\x65\xb5\x44\x26\x75\x19\xe1\x92\xad\x2a\x91\x55\x13\xda\x10\x05\x90\xa1\x46\xc7\xdf\xe3\xba\x67\x07\x60\xd0\xb5\x46\x05\x00\x5d\x30\xb0\xcd\x20\x85\x0f\xce\x08\x55\xb2\xc2\xe8\xfa\x6d\xc5\xcd\x5b\x9d\xe3\x0e\x58\x8f\x35\x88\xc6\x1c\x5a\x17\x65\x93\x11\x90\x5e\xcf\xbe\x48\xcf\xa3\x0f\x22\x1e\xa5\x90\x05\x00\xb9\xee\xd7\x51\xff\xb5\xcb\xd1\xc7\x81\x34\x85\xe7\x27\x23\xec\x9e\xf7\x14\x9c\x69\x71\xf2\x44\x0a\x71\x0c\x61\x1c\xc2\x1d\xae\xa1\xd0\x63\xe1\x9e\x80\x7e\xf9\x7a\x84\xc6\x25\x6b\x0c\x2e\x51\xb9\x6f\xb1\xe0\xad\x74\x51\x8f\x7b\x18\x85\x7f\x1a\x8e\x84\x09\x2b\x74\xd6\xda\x21\x15\xaa\x45\x30\x2a\xbf\x65\x76\x3f\x85\x52\xf7\xa0\xbb\x45\xb8\xa7\x97\x88\x4a\x44\xea\x0a\x48\x61\x76\x0e\x02\x92\xa1\xe4\x12\x55\xe9\xaa\x73\x10\x47\x47\x23\xa1\xb1\x5f\xa9\x1f\xc9\xe7\x93\xf8\xfc\x68\x47\x09\x29\x75\xc6\x68\xa3\xb4\xc8\x89\x59\x47\xc7\xf3\x8d\x70\x55\x74\xff\x58\x8c\x7e\x7f\x64\xd5\x7f\xfa\x36\xcc\xa4\xc8\xee\x86\x44\xff\x43\x44\xfa\x8c\x47\x7c\xef\xba\x4d\x7d\xcf\xb1\x1b\x35\x77\x66\xbd\x43\x7d\x7f\x90\xc6\xfa\x02\x74\x90\x71\x97\x55\x4f\xb6\xde\x1f\xaa\xd0\xd8\xa6\xe4\xbc\x23\xce\xef\xca\xaf\x7b\x72\xf0\x72\x5c\xe8\x56\x65\xb8\x3f\x81\x3d\x87\x7b\x48\xb7\x9d\xb7\x7b\x76\x74\x53\x38\x9d\xcd\x26\x93\x20\x88\x63\xb8\xf2\xa1\x2c\x70\x18\x41\xa6\xe0\x2a\xee\xa6\xc0\x2d\x48\xad\x4a\xfa\x16\x8e\xba\xc9\x09\xd5\xa2\x05\xa7\x61\x81\x20\xd4\x52\xdf\x61\x3e\x85\x95\x90\x12\x94\x76\x04\xb6\x40\x70\x46\x94\x25\x1a\xcc\x19\x5c\x57\xf8\x08\xda\xbb\x2d\x10\x32\x2e\x25\xe6\xc0\x0b\x87\x86\x70\xad\xd3\x8d\x85\x05\x0a\x55\x8e\x7b\x85\x36\x04\x76\x01\xb5\x90\x52\x58\xcc\xb4\xca\x2d\x83\xbf\x15\xf0\x9b\xa8\x6b\xcc\x05\x77\xf8\x1b\x08\x0b\x0d\xb7\x96\x28\x0c\x31\xc1\xed\x06\xd4\x8a\xd6\x04\x24\x91\xe7\x04\x8f\x79\x89\x53\x10\xca\x3a\xe4\x39\xe8\xc2\xfb\x3b\xc3\x85\xa4\x73\x66\x67\x98\x76\x65\x9d\xc2\x8a\x0b\x37\x85\xc7\xc8\x24\xf0\x81\x3f\x51\x45\x8d\xba\x75\xe7\xc1\x41\x5f\xae\xc7\xd0\x91\x77\xf1\x3e\xa4\x5a\x7f\xf0\xba\x4a\xd8\x29\x70\x53\xd2\x99\xcf\x4d\xd9\xd2\x35\xc1\x9e\x0f\x7e\x92\x93\x1c\xe9\x3e\xc4\xc1\x10\x83\x5a\xac\x95\x92\xbc\x0f\xa8\x95\x9e\xed\xb0\xa1\x57\x18\x6f\x1a\xb9\x8e\x86\x68\x7d\x98\x09\x79\x77\x63\x00\x92\xf6\x42\xaf\x20\xdd\x26\x02\x5f\x7d\x05\xcf\xb6\x49\x1c\x64\x12\xb9\xb9\xee\xd7\xd1\x60\xf7\x20\x5b\x12\x16\xdd\xe8\xe0\x19\xf7\xe2\x78\x27\xa2\x35\x04\xf9\xdf\x9c\xba\xf3\xa0\x3b\x0f\x76\x7f\x3a\x0b\xad\x1d\x1a\x7f\x5f\x69\xe6\x97\x74\x64\x63\x0e\x8b\x35\x3c\x5e\x46\xc6\x5f\xfd\x52\xb8\xaa\x5d\xf8\x1f\xfc\x45\x9b\x55\x5c\x71\x8c\x6b\xfa\xff\x10\xce\xeb\xfe\x6f\x04\xdd\x3d\x1a\xba\x66\x8c\x98\x41\x12\xf7\xb7\xce\x24\xee\xff\x8e\xfc\x3b\x00\x00\xff\xff\x1f\x24\xf3\xe1\xa6\x0c\x00\x00")

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

	info := bindataFileInfo{name: "dirlist.html", size: 3238, mode: os.FileMode(436), modTime: time.Unix(1526747625, 0)}
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

var _styleCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x51\xae\xdb\x2a\x10\xfd\xf7\x2a\x46\xf7\xea\xfd\x3d\x22\xe7\xa6\x69\x12\xb2\x89\x4a\xed\x06\x30\x8c\x1d\x14\x0c\x08\x70\xe3\xb4\xca\xde\x2b\x30\x76\x6c\xdf\x54\x55\x3e\x22\xc6\x33\x73\xce\xcc\x39\x50\x19\x71\x87\xdf\x05\x40\x6d\x74\x20\x35\x6b\xa5\xba\x53\x78\xfb\xf6\x03\xbe\x33\xed\xdf\xfe\x07\xcf\xb4\x27\x1e\x9d\xac\xcf\xc5\xa3\xd8\x08\xe9\x94\xf4\x81\xdc\x1c\xb3\x16\x5d\x2a\x6d\x99\x6b\xa4\xa6\xb0\x2f\x6d\x0f\xac\x0b\xe6\x5c\x00\xdc\xa4\x08\x17\x0a\x5f\xcb\xff\xe6\x75\x29\x3f\x35\xf0\xe1\xae\x90\x84\xbb\x45\x0a\xda\x68\x5c\x64\x29\x99\x12\x2b\xe3\x04\x3a\xe2\x98\x90\x9d\xa7\xb0\xb3\xfd\x2a\x8b\x5e\xcc\xcf\x4c\xa2\x62\xfc\xda\x38\xd3\x69\x41\xb8\x51\xc6\x51\x78\xc7\x63\xfc\xad\x1b\xb3\x94\x1e\xb0\x0f\x44\x20\x37\x8e\x05\x69\xf4\xc8\x01\x60\x2c\x2e\xcb\xd3\x49\xf0\x18\x11\xd2\x5b\xc5\xee\x14\x2a\x65\xf8\x35\x46\x2c\x13\x42\xea\x86\xc2\xc1\xf6\xf0\x51\x7e\xa2\x05\x72\x53\x0f\x30\x53\x66\x09\xfb\x98\x06\xd0\x4a\x4d\xf2\x6e\x86\xca\x27\x64\x75\x88\xbf\xf3\xc8\x8e\x29\xd9\x68\x0a\x1c\x75\x40\x47\x8b\x47\x11\xbb\x12\xec\x03\x3a\xcd\x14\x51\x52\x5f\x09\x53\xe1\xa9\x9e\x97\xbf\x90\xc2\x66\x8f\x6d\xe2\x53\x39\x64\x82\xbb\xae\xad\xfc\xb8\xd0\xbf\x6c\x7e\x59\x7f\x8a\xf5\xb3\xb1\xa5\x56\x32\xeb\x33\x6f\xb9\x49\xff\xc4\xa3\x7d\x39\xe9\xa3\x28\x36\x31\xbb\xc5\xf4\x79\x69\x88\x78\x1e\x5d\x73\x98\xbb\x66\xd2\x3c\x18\x4b\x61\x6b\x7b\xf0\x46\x49\x01\xef\x88\x38\x7c\x9f\x70\x76\xe3\xde\x2d\x6b\x90\x70\xa3\x03\xea\xb0\xc4\xda\x7f\x59\x61\x95\x33\xa0\x4f\x8d\xac\xc3\xe5\x24\xdb\xac\xcf\x2b\x1b\x02\x44\xe7\xd5\xca\xdc\x48\x4f\x73\xd7\x35\x17\xd9\x36\xa9\x61\x66\x73\xda\xa7\xab\x50\xd4\xc6\x84\x6c\xda\xa7\xf6\x42\x88\xd7\xc2\x67\xb2\xc6\xcb\xc1\xa7\xb5\xec\x51\x8c\xab\x0a\xc1\xb4\x14\xca\xe1\xe8\x64\x73\x09\xd3\x29\x4f\x41\x72\x74\x97\x67\x59\x4a\x7d\x1c\xac\x32\x52\x62\x2b\x52\xc7\x03\xab\x38\x4f\x19\x41\xac\x55\x3e\x64\x95\xdf\x3d\x32\xc7\x2f\xb9\xf4\x25\xd1\x25\xb3\x24\x6d\xb9\x14\xe6\x63\xa2\x37\xdd\x8d\x29\xf2\x7a\xfd\x63\x7c\xe1\x92\xbc\xc5\x68\x74\x8d\xe4\x82\x03\xee\x76\x98\x73\x21\xfa\xcc\xa5\xf3\xb7\x81\xd6\x86\x77\xfe\x5f\x0f\xca\x13\x7b\xbc\x3f\xa6\x0b\x11\xf2\xf9\x94\xfd\x09\x00\x00\xff\xff\x78\xdf\xba\xb1\x58\x05\x00\x00")

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

	info := bindataFileInfo{name: "style.css", size: 1368, mode: os.FileMode(436), modTime: time.Unix(1526767673, 0)}
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

