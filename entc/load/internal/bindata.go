// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package internal

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x51\x4f\xdb\x30\x14\x85\x9f\xe3\x5f\x71\x16\x31\x91\xb0\xe2\x02\x6f\x9b\xd4\x07\x04\x9d\xd4\x69\x83\x49\x45\xda\x03\x43\xc8\x75\x6e\x5a\x8b\xd4\xce\xae\x5d\xb4\xca\xca\x7f\x9f\xec\xb4\x6c\x7b\x4b\x7c\xbe\x7b\xce\xb9\x76\x8c\xd3\x33\x71\xe3\xfa\x3d\x9b\xf5\x26\xe0\xea\xe2\xf2\xe3\x79\xcf\xe4\xc9\x06\x7c\x56\x9a\x56\xce\xbd\x60\x61\xb5\xc4\x75\xd7\x21\x43\x1e\x49\xe7\x57\x6a\xa4\x78\xd8\x18\x0f\xef\x76\xac\x09\xda\x35\x04\xe3\xd1\x19\x4d\xd6\x53\x83\x9d\x6d\x88\x11\x36\x84\xeb\x5e\xe9\x0d\xe1\x4a\x5e\x1c\x55\xb4\x6e\x67\x1b\x61\x6c\xd6\xbf\x2e\x6e\xe6\x77\xcb\x39\x5a\xd3\x11\x0e\x67\xec\x5c\x40\x63\x98\x74\x70\xbc\x87\x6b\x11\xfe\x09\x0b\x4c\x24\xc5\xd9\x74\x18\x84\x88\x11\x0d\xb5\xc6\x12\xca\xad\x32\xb6\xc4\x30\x88\xe9\x14\x37\xa9\xcf\x9a\x2c\xb1\x0a\xd4\x60\xb5\xc7\x29\xd9\xa0\xdf\x8e\x4e\x25\x6e\xef\x71\x77\xff\x80\xf9\xed\xe2\x41\x8a\x5e\xe9\x17\xb5\x26\x24\x0f\x21\xcc\xb6\x77\x1c\x50\x89\xa2\x74\xbe\x14\x45\xb9\xda\x07\x4a\x1f\x31\x22\xd0\xb6\xef\x54\x20\x94\x23\xe5\x73\xa4\x28\xc8\x06\xaf\x37\xb4\x55\x88\x11\x3d\x1b\x1b\x5a\x94\xef\x7f\x95\x90\xdf\x0f\xde\xc3\x20\x6a\x21\x5e\x15\x63\x04\x3d\x66\x78\x7c\x22\x1b\xe4\xc2\x06\xe2\x56\x69\x8a\x29\xe2\x1c\xac\xec\x9a\x70\xf2\x3c\xc1\x89\x55\x5b\xc2\xa7\x19\xe4\x9d\xda\x92\x4f\x1e\xc5\xdf\x28\x99\xe0\xb7\x2c\x1f\x87\xf2\x30\x30\x0c\x93\xd1\x89\x6c\x93\x66\x06\x21\xda\x9d\xd5\x79\xbd\xaa\x46\x14\x45\xaa\xd1\x19\x4b\x1e\x8f\x4f\x8f\x4f\x69\x3f\x51\xb4\x8e\xf1\x3c\x39\xb4\x4b\xa1\x63\x8f\x63\xdb\x28\x8a\x62\x35\x01\x31\x27\xed\x9b\x62\xbf\x51\xdd\x32\x8b\xd5\xc8\xd4\xa2\x28\x4c\x9b\x89\x77\x33\x58\xd3\xe5\x99\xa2\x55\xa6\xab\x88\x39\xc9\xa9\xff\x98\x3b\x83\xea\x7b\xb2\x4d\x95\x7f\x27\x58\xd5\x22\xa9\xce\xcb\x65\x68\xdc\x2e\xc8\x1f\x6c\x02\x55\xf9\xea\xe5\x17\x67\xec\x11\x1c\xeb\x56\xe5\x4f\x5b\xd6\x75\xfd\xb6\xdb\x31\x25\xc5\x3b\xce\x4b\x8e\x5e\xc4\x3c\x7a\x2d\x03\x1b\xbb\x4e\x8c\x9c\x27\xa6\xaa\x3f\x64\x93\x0c\xce\x7f\x9b\x50\x5d\x66\xbb\xff\x5e\x79\xdc\x6c\x7c\xe4\x18\x8f\x17\xfa\x27\x00\x00\xff\xff\x54\xe7\x81\x8f\x3b\x03\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 827, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\x5f\x6f\xe3\x36\x12\x7f\xb6\x3f\xc5\x34\x40\x17\xd6\xc2\x55\x7a\x45\x51\xdc\x79\xcf\x07\x14\xed\x16\xcd\xf5\x36\x5d\x74\x77\xfb\x12\x04\x29\x23\x8d\x12\x6e\x24\xca\x25\xe9\x6c\xd2\x34\xdf\xfd\xc0\x19\x52\xa2\x64\xc9\xf6\xe6\xdf\x4b\xac\x21\x67\x38\xf3\xd3\xfc\xe3\xd8\x87\x87\xf0\x43\xbd\xba\xd5\xf2\xe2\xd2\xc2\x37\x5f\xff\xe3\x5f\x5f\xad\x34\x1a\x54\x16\x7e\x12\x19\x9e\xd7\xf5\x15\x1c\xa9\x2c\x85\xef\xcb\x12\x68\x93\x01\xb7\xae\xaf\x31\x4f\xa7\x87\x87\xf0\xfe\x52\x1a\x30\xf5\x5a\x67\x08\x59\x9d\x23\x48\x03\xa5\xcc\x50\x19\xcc\x61\xad\x72\xd4\x60\x2f\x11\xbe\x5f\x89\xec\x12\xe1\x9b\xf4\xeb\xb0\x0a\x45\xbd\x56\xb9\x13\x21\x15\x6d\xf9\xdf\xd1\x0f\xaf\x8f\xdf\xbd\x86\x42\x96\x18\x68\xba\xae\x2d\xe4\x52\x63\x66\x6b\x7d\x0b\x75\x01\x36\x3a\xcf\x6a\xc4\x74\x3a\x5d\x89\xec\x4a\x5c\x20\x94\xb5\xc8\xa7\x53\x59\xad\x6a\x6d\x61\x36\x9d\x1c\xa0\xca\xea\x5c\xaa\x8b\xc3\x8f\xa6\x56\x07\xd3\xc9\x41\x51\x59\xf7\x4f\x63\x51\x62\x66\x0f\xa6\xd3\xc9\xc1\x85\xb4\x97\xeb\xf3\x34\xab\xab\xc3\xc2\x1b\x7c\x88\x8a\xb6\x8d\x2c\x1d\x9a\xec\x12\x2b\xb1\x7b\xc7\x21\xe6\x17\xb8\xc7\xb6\x42\x62\x99\xef\xb1\x4f\xaa\x1c\x6f\x0e\xa6\xc9\xd4\x81\xf6\x8e\x68\xa0\xd1\xbf\x2e\x03\x42\x01\x2a\x9b\xfa\x05\x7b\x29\x2c\x7c\x12\x86\x50\xc1\x1c\x0a\x5d\x57\x20\x20\xab\xab\x55\x29\xdd\xab\x31\xa8\xc1\x23\x97\x4e\xed\xed\x0a\x83\x48\x63\xf5\x3a\xb3\x70\x37\x9d\x1c\x8b\x0a\xc1\xff\x19\xab\xa5\xba\x80\xfe\xdf\x1f\x0e\xda\xc5\x81\x12\x15\xce\xeb\x4a\x5a\xac\x56\xf6\xf6\xe0\x8f\xe9\xe4\x87\x5a\x15\xd2\xef\x77\x6a\xc5\xcf\x5d\xde\x8c\x56\xba\xdc\xaf\xf3\x0b\x34\x7e\xdb\xc9\xe9\x4b\xf7\x38\x72\xb2\xc3\xd8\x74\x99\x7f\x72\x78\x9a\x86\x99\x1e\x87\x99\x09\xf9\x1e\xf7\x91\x43\xd9\x1f\x7e\x72\xfa\x92\x1e\x87\xb9\x25\xef\xec\xb2\xff\x5c\xd7\x57\x91\xe6\x6f\x6b\x23\xad\xac\xd5\x00\xfb\xa5\xdb\xd9\x65\x7e\x5b\x97\x32\xbb\xdd\x87\x79\x45\x3b\xbb\xdc\xdf\x2b\x55\x5b\xe1\x18\x0c\x54\x62\x75\xc2\xaf\xec\x54\x2a\x8b\xda\xf9\xd3\xdd\x7d\xe0\x16\xed\xce\x8e\x88\x7b\x72\xad\xe6\xd8\x1c\x4d\xa6\xe5\x39\x1a\x10\xb0\x0a\x44\x1f\x99\xec\x93\xde\x73\x1a\x8e\xd6\x77\x22\xdc\xa4\xb2\x00\x87\x87\xc0\x24\xcf\x4f\xd0\x1f\x3a\x0c\xa0\x94\xc6\xa6\xd3\xc9\x1b\x79\x83\xf9\x11\x19\x7b\x5e\xd7\xa5\xe7\x90\x99\xb0\x68\x40\x16\xd1\xa9\x50\x9f\x7f\xc4\x8c\xdd\xbb\x72\x5c\x5f\x49\xc5\x02\xa4\x0a\x87\xf0\x91\x44\x02\x19\x1f\x5c\x11\x89\xcf\x64\x7b\xd9\x41\x36\x23\x89\xe9\x0f\x08\x24\x66\x1c\x8e\xa3\xd1\x48\x1a\x0f\xa5\x23\x55\xd4\xed\xb6\x97\x84\x5c\xfa\xfe\x76\x85\x9d\x05\xcf\xee\x14\xe8\xb2\xbf\x17\xf1\x61\x3b\x4e\xb7\xa2\x17\x89\xef\xe4\x5f\x91\xee\x2f\xa5\xb2\xdf\x7d\x3b\xca\x6d\xe4\x5f\xbd\xc3\x5f\xab\x75\x65\x9a\x6d\x27\xa7\x0c\xca\x1d\x1c\xcf\xe1\xf7\xa0\x4b\xe3\x96\xe8\x36\x77\xf9\x3f\x28\xf9\xe7\xba\x51\x80\xfc\x62\xe0\xcf\xf3\xaf\x69\x73\x57\xc0\xb1\x2c\x4b\x71\x5e\xe2\x5e\x02\x94\xdf\xdc\x15\xf1\xeb\xca\xf9\xb6\x28\xf7\x12\x51\xfb\xcd\x5d\x11\x3f\x62\x21\xd6\xa5\xdd\xcf\x8c\x9c\x37\x0f\x4a\xf8\x5d\x94\x0e\x8e\x38\xa6\xc7\x25\x9c\x5d\xbb\xdd\x3d\x40\x57\xb9\xb0\x18\xf4\xd9\x05\x28\x6d\x3e\x1b\x54\xe8\xa8\xaa\xd6\xb6\x41\x76\x87\x20\x19\x36\x77\x65\xfc\x2e\x4a\x99\x0b\x5b\x6b\x72\x11\x0a\xda\x71\x19\xd7\xcd\xe6\x9e\x87\xda\x5a\x8b\x0b\xfc\x05\x29\x71\xee\xf0\x6f\xc3\x9b\xcf\xae\xf0\xb6\x9f\x7a\xe3\x5c\x3b\x98\x7a\xe3\xec\xcb\xab\x3d\x45\x50\x39\xf2\xf5\x5e\x88\x98\xb0\xb9\x27\x83\x12\x9c\x0b\x6e\xb7\x37\xca\xe2\x1d\xbb\x82\x0c\xda\x7c\xb6\x19\xf2\x71\x25\x80\xb1\x5a\xb0\x5f\x31\xa0\xd2\xbb\x99\x1b\x89\xfc\x80\xd4\x48\x7c\x4f\xd3\x61\x04\x90\x76\xf3\x6e\xcf\x89\x3b\x78\xfb\x09\xf1\x37\x2c\x1a\xad\xb7\xb3\x6a\x2c\xce\x36\xd5\xfe\x0d\x8b\x66\xe3\x60\x63\x13\xf3\x8f\x27\xc3\x11\xf7\xda\x92\x09\x8f\xd4\x35\x6a\xb3\xd5\x39\x9b\xc6\x86\x76\xf6\xf5\xfe\x73\x2d\x35\xe6\xbb\xd9\xb5\xdf\x39\x1e\xa6\x2f\x5d\xd7\x96\x76\x03\x77\x8f\x18\x7d\xaa\x06\x87\x7b\x84\x4d\xa7\x66\xfa\x03\xbc\x9a\x19\x5b\xb7\x8e\x5e\x54\x03\xd5\x96\x37\x13\x75\xbb\x27\x21\xd0\xf7\x6a\x6f\xfb\xbb\x87\xfa\xd9\x08\xe5\xc6\x5d\x77\x00\xcd\x28\x1d\xe3\x27\x72\xcf\x4c\x23\xb5\x60\x42\x05\x44\x9c\x52\x0c\x0b\x7d\xe2\x36\x71\x65\x6b\x9d\x4e\x8b\xb5\xca\x02\xe7\x0c\x73\xff\xa6\x7f\x6c\x76\x24\xde\xe7\xef\xa6\x13\x85\xb0\x58\xc2\x0b\xf7\x78\x37\x9d\xb8\x90\x5c\x34\x9e\x84\x79\xfa\x5e\x5c\xcc\x1d\xf9\x76\x85\x8b\x98\xec\x62\x79\x3a\xa1\xcc\x11\xd3\xdd\xb3\xa3\x33\xf4\x8b\x86\xce\xcf\x6e\xc5\xfb\xff\x22\xac\xf8\x67\xb7\x14\x7c\x7b\xe1\x97\xc2\x33\xaf\x15\xed\x59\xb4\x56\x84\xb3\x5a\x68\x17\xb4\xd4\x3e\xbb\xd5\xc8\x5b\x17\x50\x89\x2b\x9c\x0d\xfb\x6c\x32\x9f\x4e\xee\xa7\x93\xa2\xd6\x70\x36\x07\x61\x1d\x2a\x5a\xa8\x0b\x74\x22\x63\x97\x77\x28\x29\x4c\x45\x9e\xb7\xd4\x99\xb0\x09\xb1\xcb\x02\x34\x16\x8e\x97\x75\x7c\x45\x8f\x5f\x2c\x41\xc9\x32\x70\xba\xd4\xb3\x6c\xde\x8e\xc6\x22\x61\x7a\xe4\x22\x4b\xe0\x7d\x11\x8d\xc4\x6b\xb4\x6b\xad\x40\x61\xeb\x1c\xdc\xe5\x6e\x7a\x07\x39\x21\xbb\x07\x7f\x1c\xf2\x0f\x62\x9e\x15\x79\x68\x67\x63\x0f\x99\xf1\x95\x6d\x0e\xa8\xb5\x7b\xbe\x23\xeb\x50\x6b\x67\x5d\x91\xa7\xaf\xb5\x9e\x25\xaf\x88\x10\xd9\x17\x34\x94\xe5\x1c\x8a\xca\xba\x5d\xb5\x2e\x66\x1c\x13\xf0\xe5\x9f\x0b\xf8\xf2\xfa\x60\xee\xf8\xe9\xf5\x39\x76\x46\xce\x10\x6a\x2f\xe8\xcc\xbb\xbe\x67\x41\xc3\x40\x1e\x54\xd4\xdd\x15\x47\x99\xf7\x9d\x97\x56\xbc\xfb\x52\xff\xbb\x88\x17\x88\xb2\xe1\xa9\xb4\xd4\xfa\x6a\xe8\x5a\x17\xad\x0e\xa1\x35\x9d\x4e\x9a\x86\xb4\x5d\x0d\x14\xb7\xea\x7b\xbb\x45\x2b\x37\x74\x7b\x8c\x16\x9d\x1d\x77\x81\x0b\x3a\xbb\xd3\x17\xb6\x3b\x9b\x36\x6f\xd1\xd8\xdc\xf4\x72\xfd\x10\xa0\xe5\x6e\x10\xb4\x1d\x1e\xad\x97\xa8\x66\x45\x9e\xb6\xd4\x84\x84\x84\x5e\xa8\x39\xa3\xa1\xd0\x72\xd3\x13\x35\x67\x34\x94\x8d\x40\x83\x87\x85\x5a\xb1\x19\x6a\xa6\x18\x0f\x35\x53\xd0\xab\x87\xe5\x6e\xff\xab\xa4\x31\x2e\xeb\x52\xa1\x90\x8e\xc9\x1d\x1f\xbc\xf2\x60\xee\x64\x39\x07\x6b\x65\xbb\x1b\xd7\x62\x09\x74\xd5\x72\x68\xb9\x2b\x58\xf2\x8a\xe9\x5f\x2c\xe1\xeb\xa0\x1d\x5d\xcd\x96\xf0\xc2\x2d\x10\xb3\x2b\x6d\x7c\x4f\xf6\x1d\x3b\xd0\x05\x00\x32\xa1\xe0\x1c\x81\xc6\x61\x98\x83\xad\x69\xcf\x05\x2a\xd4\x82\xa2\xd2\x71\xfe\x54\x6b\xc0\x1b\x51\xad\x4a\x9c\x83\xaa\x2d\x08\x70\xc1\x4a\x4d\x70\x29\xaf\x10\xac\xac\x30\x3d\xae\x3f\xa5\xa4\xe5\xd9\x3c\x44\xa4\xab\x25\xe9\x1b\xa1\xcd\xa5\x28\x67\xad\xb7\xf9\x08\x8d\x10\x32\x45\xda\xb9\xc5\x2c\x23\xdf\x8c\x93\x8c\x29\xe6\x8e\xa7\xcd\x34\x5c\x5e\x37\x33\x0d\xdf\xeb\x29\xd3\xf0\xc7\xa1\x4c\x43\xcc\x33\x99\xdf\xb8\xcb\x6b\x8e\x37\xdd\x62\xc4\xa2\xef\x9a\xb3\x5f\x10\xc1\x69\x4b\x45\xd9\x07\x91\xcc\x6f\xa8\xe3\xa5\xb8\xe5\xfa\xbb\x68\x16\xf8\xb9\x1f\xd1\x6e\xa5\x8d\xe7\x38\x4c\xdc\x4a\x27\x48\xee\xbd\xa5\x1e\x43\x3f\x86\xe3\xb7\x45\x6f\x2a\x1a\xeb\x35\xce\xec\x3e\xd5\x20\xe0\xbf\xef\x7e\x3d\x76\xcc\xd4\xb5\xf8\x17\x9d\x23\xbf\x68\xda\xe2\x04\xbc\xeb\x8c\x4d\xf8\x9f\x47\xa8\x73\xe8\xcc\x84\xb3\x5d\x33\xe4\x4f\x4a\x60\x76\x0e\x27\xa7\xe7\xb7\x96\xb3\x66\x94\x96\x0d\x65\x4e\xe6\x75\x98\xf1\x90\xcf\x83\xe6\x27\x44\x4c\x9b\x25\x1b\x45\x5b\x2a\x1e\xec\xce\xfc\x38\x96\x4a\xfb\xaf\x85\xd7\x21\x49\x28\x2a\x98\xef\x33\x4b\x69\x5b\x2d\x4c\xea\x5c\x85\xa6\x40\x41\xee\xde\x85\xc3\x63\xd1\x54\x0e\xd3\x2f\x1c\xed\x14\x36\x6a\x37\xa1\xbe\x46\xad\x65\x8e\xcd\x64\x2a\x5e\x4d\x07\x93\x8f\x47\x2a\xb2\x72\x96\x70\xc4\x8c\x67\xa0\x8e\x81\xec\x82\x4f\x6f\x21\x77\x9f\xcd\x59\xa2\x40\x8a\x82\x70\x50\xa3\xc8\x53\x9c\xe5\x71\xc1\xb8\xfd\x71\x6d\x31\xe3\xc0\x2d\xf2\x12\xc4\x6a\x85\x2a\x9f\x79\xc2\xbc\x6d\x35\xa3\xb0\x9e\x25\x89\x87\xc9\x0f\x68\x63\x03\xfc\x78\xf7\x39\x4d\x70\xb9\xa6\x31\xc2\xeb\xe0\xcd\x08\xc3\xe5\xc8\x90\xa3\xa0\x64\x9c\xab\x06\xad\xe9\xbd\x74\x9a\x34\x3f\xfd\x3b\xef\x1f\xc3\x33\xe9\xa7\x3f\xc7\x33\x76\xaa\x87\x49\x7c\x2a\xfc\xa0\xaa\x4e\x32\xe4\x8c\x66\xb8\x6e\xc9\x6b\x54\x70\xbe\x2e\x0a\xd4\x40\x39\xd0\x97\x83\x30\x92\xa6\xbc\xd6\x93\x30\x3b\x5f\x17\x3e\x89\xb9\x06\x93\x89\xf3\xb1\x54\xd6\x81\x81\x34\x6c\xc4\x39\x41\x73\x30\xdb\x81\x40\xad\x63\x87\x28\xa2\x50\xf7\xe5\x82\x58\xa2\xae\x36\xf5\x15\xdb\x0c\x74\xb6\x9b\xa2\x9d\xec\xa8\x5e\xc6\xe5\xb2\xc9\x77\xf4\xc9\xf8\x71\xb7\xad\xc3\xe8\x9c\xaf\x6d\x71\x7e\xf7\x80\xcd\x0c\x78\x58\x12\xe8\x27\xcd\x7e\x41\x20\xd8\x9c\x6e\x24\xbd\x13\x5f\x9d\x5c\xbb\x25\xba\x62\x88\xe4\x1c\xaa\x28\x64\x58\x65\xba\xb3\x88\xca\xb7\x42\xc3\xa5\xa2\xba\x69\xca\xc4\x74\x32\xf1\xb7\xdf\x58\x1b\x9f\x18\xab\x9b\xa4\x85\x7b\x00\xd9\x6e\xbf\xe6\x4e\x6f\xfc\x56\x45\x5e\xeb\xf4\x25\x85\x3f\x76\xde\x69\xd1\xbe\xd1\x89\xeb\x5d\xfc\xf9\xed\x2d\xa7\x1b\xcd\x6e\xdb\x80\x2a\x9f\xab\x0b\x29\xe3\x7a\xaa\x66\x54\xb9\x84\x17\xe1\x33\x4b\xa4\x74\xe2\xeb\xed\xc7\x39\x91\xfc\x97\x2c\x44\xb4\x9a\x9b\x93\x49\xf4\xcd\xc9\x02\xe4\xbc\x15\x1e\x9c\x35\x4a\x57\xbe\xdb\x01\x53\x04\x40\xc6\x8a\xc4\x53\x83\x3e\x56\x1c\x1e\x54\x1d\x48\xea\xb6\xfa\xf0\x0c\xda\x8f\xd6\x85\xc7\x14\x06\x3a\x80\xbf\x4a\x8c\xcd\xe0\xe2\xf0\xe4\x7e\xdf\xea\x4f\x47\x06\xed\xf9\x4b\xcf\x48\xf7\x9f\x59\xa1\x27\xf4\xc7\xa0\x86\xff\xe2\x33\xb6\xd5\x57\xa8\xa7\x34\x56\x16\xc0\x07\x75\x04\x99\xd4\x7f\x41\x1b\x59\xfa\xd6\xeb\xd3\x33\xf5\xb3\xed\x1a\x68\x0b\xab\x9b\x81\x96\x70\xb8\x27\xec\x16\x84\x6e\x35\xf0\x31\xcc\xe5\x80\xef\x9d\x0f\x28\x07\x9d\x16\x73\xb4\x1e\x8c\xa7\xe0\xcf\xae\x08\xc3\x09\x76\xbf\xfc\x3a\xee\x04\x4d\xf9\x1c\xcd\x9c\xe1\xf5\xd0\x9e\x5d\x09\x70\x03\xf3\x41\xec\xe2\x4e\x6d\x14\xba\xb1\x18\xfe\x4c\xe0\x86\x22\x74\xdf\x00\x6d\xe2\x93\x7d\xb3\xf1\xe1\x42\x94\x3c\x27\xbd\xdf\xdb\xe4\x4e\xd7\x38\x6a\xf3\x78\x30\xef\x6f\xf5\x60\xa8\xee\x17\xa9\xfb\x99\xd3\x0b\x37\xb5\x79\x5d\xa3\xc8\xcc\xd6\x5a\xcf\xa1\xbe\xe2\xce\x39\x0a\xdc\x13\xa1\x7c\x8f\x72\x4a\xda\x7e\x51\x5f\x79\x1d\x87\x37\x39\x9d\x55\x63\x67\xb0\xb1\x0a\xb2\xdd\x39\xa9\xc7\x27\x7d\x83\xfa\x02\x75\xf2\x0a\x76\xcb\xac\x78\xf3\x4c\xa8\xc4\x8f\x1c\xd8\x52\xe4\x89\xfc\x83\xec\xc4\x7d\xec\x1c\xdb\xf4\x18\x3b\xb7\xc8\x1c\xb3\xb3\x00\x1e\x2b\x3f\xc8\xd0\x62\x1f\x43\xc7\x36\x3d\xc6\xd0\x2d\x32\x77\x1b\xda\xde\x29\xda\x90\x73\xfa\x36\x73\x38\xf8\xfb\x6f\xf7\x74\xa4\x8a\x3a\x3d\x5e\x57\xa8\x65\xe6\xcb\x4c\x14\x13\x4e\x5f\xd5\x02\x11\x8f\xf0\xd2\x59\x51\xd6\xc2\x7e\xf7\x6d\xd2\x01\x62\xa0\xe0\xae\x15\xde\xac\x30\xb3\x98\xf7\x66\x93\x34\x16\x6d\x26\xa2\x0b\x1e\x89\xc6\x13\x51\xf3\x49\xda\xec\x12\x2c\x9f\x4e\xaa\xba\xee\xff\x15\xbd\x22\x61\x10\x2c\xfc\x67\x09\xf1\x8f\x62\xec\x3f\xe1\xc5\x0b\xb0\xf0\xef\x1e\xf9\xbb\x6f\x17\x84\x68\x6f\x08\xc9\x73\x56\x07\xe2\x90\xb8\x0f\x72\x58\xde\x07\x39\x2a\x70\xdd\x4a\x1c\xaa\xc9\x6d\x51\x84\x4f\x5a\xac\x4c\xfc\x73\x2a\x4f\x17\x2a\xe7\x5b\x50\x20\x54\x68\x2f\xeb\x1c\x3e\x49\x7b\x09\x1a\xb3\xfa\x9a\xaf\xbe\xa8\xcc\x5a\x23\xa8\x1a\x56\x42\xc9\xcc\x80\x54\xe0\xef\xa9\x52\x5d\xf8\x4a\x1e\x15\xe1\x22\x8f\x7e\x39\x02\x9e\x98\xc0\xc9\x69\xfb\x73\xa7\xfb\x04\x66\xbe\xde\x46\xe4\xfe\xe0\x2f\x47\x77\xf9\x76\xe2\xbd\xbf\xc8\x02\xae\xa9\xf4\xb0\x72\xee\x16\x7b\xdd\xa9\xbf\x34\x0b\xee\xb8\xc4\x97\xef\x83\x75\xac\x7c\xf3\x05\xcd\x1c\xae\xe9\x82\x53\x84\xda\x4b\x5e\x48\x2d\x8e\xbb\xe7\x05\xef\xca\xd3\x60\xc0\xbc\x87\x2e\x5f\x07\x36\xc0\x65\xf2\x63\xa1\x8c\x27\x60\x31\x9a\x4c\x0f\x60\xd2\x97\x9c\x0e\x4b\xbe\xa7\xb4\xc4\xe7\x40\xb2\x63\x5f\x07\x4c\x06\x12\xfd\xf5\x68\x10\xc7\x98\x79\x13\xca\x70\x2f\xd9\x00\x33\x2c\x3c\x16\xce\xee\x3c\x2e\x06\x34\xac\x04\x48\x79\x54\xef\x30\x0d\x77\xa7\x88\xfe\x8c\xb0\x06\x4b\x07\x80\x95\xcd\xad\x6d\x1b\xb4\x8d\x21\x7d\x70\x79\x4e\xb3\x01\x2d\x93\x1f\x0b\xec\xb6\xf9\xcd\x8c\xef\x3f\x8c\xdf\x9b\x76\x86\xf3\x2c\xf8\xb1\x39\x03\xe8\xb1\x12\xdb\xb1\x63\x2b\x36\x90\xe3\x7e\x76\x03\x39\x26\x3f\x16\xb9\x4e\xbb\x1e\x39\x24\xd3\x83\x3b\xba\x27\xf2\x46\xee\xb3\x5b\xe2\x33\x42\xc9\xf6\x0d\x40\x79\xe9\xfb\xfb\x6d\x50\x7a\xf5\xfb\x50\xfa\x46\x79\x03\x4b\x4f\x7f\x2c\x98\xdd\x8b\x40\x84\xa6\x5f\x48\xc8\x37\xfd\x61\x0e\x4e\xdf\xcc\xb7\xd4\x67\xc4\xd3\x1f\x3b\x00\xe8\x2a\x5c\x1f\xb6\x21\x1a\x4c\x98\x77\xee\x0e\xcd\xb0\xd2\x42\x3c\xae\x4c\x3a\x4f\x74\x59\xae\x35\xd8\xf4\x17\xa9\xf2\x59\x02\xcb\x65\xb3\xfe\xd6\x52\xa7\x36\xb1\xb0\x04\x9b\xbe\x2e\xb1\x9a\x75\x5a\x09\x3b\xbd\x9f\xfe\x3f\x00\x00\xff\xff\x01\x36\x8f\x15\xb7\x31\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 12727, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
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
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
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
