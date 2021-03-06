// Code generated by go-bindata.
// sources:
// temp/index.html
// DO NOT EDIT!

package something

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

var _tempIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\x5d\x6b\xf4\x44\x14\xbe\x5f\xd8\xff\x30\x9d\x9b\xa6\xa8\xc9\xad\xf8\x26\x7b\xe1\x6b\xe5\x15\x5e\x51\xb0\x22\x5e\x66\x33\x67\xbb\xa1\xd9\x49\x98\x99\x24\x96\xb2\xb0\x5d\x14\xab\xd8\x56\xe9\x8a\xb4\x15\x4b\xc1\x42\x05\x65\x85\x4a\xa1\x1f\xe0\x8f\xa9\xc9\xb6\x57\xfe\x05\x99\x49\xd2\x66\x3f\xdb\xa5\xe0\xdc\xf4\xcc\xcc\x39\xe7\x79\xce\x33\xa7\x67\x63\x2e\xbc\xf7\xd1\xcb\x95\xcf\x3f\x5e\x46\x4d\xd1\xf2\x6a\xd5\x8a\x29\xff\x22\xcf\xa6\xab\x16\x06\x8a\x6b\xd5\x8a\x3c\x03\x9b\xd4\xaa\x15\x84\x10\x32\x5b\x20\x6c\xe4\x34\x6d\xc6\x41\x58\xf8\xd3\x95\xf7\xdf\x7a\x1b\x17\x77\xc2\x15\x1e\xd4\x3e\x83\xfa\x27\xbe\xb3\x06\x02\x25\x3b\x47\x83\xde\xd1\x60\xe7\xb7\xf4\xec\x47\xd3\xc8\x6e\xab\x15\xd3\xc8\xf3\x55\x2b\x66\xdd\x27\xeb\x45\x34\x71\x23\xe4\x78\x36\xe7\x16\x2e\x32\xaa\xf3\x80\x41\x69\x2b\xd7\xa0\x7b\x91\x7c\x7d\x25\xad\xc5\xf4\x9b\xbd\xe4\xba\xb3\x88\x2c\x94\x6c\x1d\x26\x57\x97\xc9\xd6\x9f\xe9\xcf\xdb\xc9\xb7\xc7\xc9\xfe\xe9\xe0\xe0\xcb\xdb\xbf\x7f\x49\x77\x4e\xfe\xbd\xfe\x0e\x0d\x67\x28\xd6\x62\xb2\xfb\xc3\x5d\x67\x53\xc6\xa7\xe7\x5b\xe9\x66\x3f\xdb\x97\xb3\x4c\x0b\xfc\xea\xec\xee\xa7\x3f\x14\xb0\xb2\x32\xa4\x9b\x4e\x77\x04\x29\xed\x9e\x26\xbb\xfd\x7f\xae\x4e\xd2\xc3\xbf\xd2\xde\x45\x8e\x72\x71\x9e\x03\xfd\x7a\x90\xfe\x7e\x7c\xd3\xe9\x4e\xac\xd7\x34\x88\x1b\x3d\x26\x4f\xc3\x67\xad\x11\x7d\xcc\x7a\x28\x84\x4f\x91\x4b\x2c\xec\x07\xf2\x19\x33\x99\x4c\x23\xbb\x98\xe1\xee\x78\x3e\x07\x5c\xcb\x6a\x9a\xe6\xef\xd2\x20\x14\xca\x5d\x59\x18\x89\xf5\x00\x2c\x2c\xe0\x0b\x81\x51\x64\x7b\x21\x58\xf8\x15\x78\x9e\x8f\xee\x7b\x61\x01\xcf\x00\xe5\x40\x09\xae\x65\x8a\x4c\xc0\x34\x8d\x52\x8d\x43\x9a\xd4\x19\x32\x0a\xbb\xc9\x26\x9d\x4a\xd5\x94\x0c\xa1\x90\x4c\x47\x92\x98\x46\xde\x80\x26\x77\x98\x1b\x88\xfc\x3a\x76\x29\xf1\x63\xdd\x26\x64\x39\x02\x2a\x5e\xbb\x5c\x00\x05\xa6\x61\xcf\xb7\x09\x7e\x13\x35\x42\xea\x08\xd7\xa7\x1a\x44\x62\x09\x6d\xc8\x4e\x2e\xb8\x46\x36\x43\x19\x16\xb2\x10\xf1\x9d\xb0\x05\x54\xe8\xab\x20\x96\x3d\x90\xe6\xbb\xeb\x1f\x10\xad\x60\xb3\xf4\x62\x38\x30\xd3\x75\x46\x5c\x26\xf7\x68\x58\xcc\x5f\x8c\x52\x08\x98\x4b\x65\xa6\x7b\xa6\x2d\xe0\xdc\x5e\x05\xc5\xb6\xfc\x0c\xd2\x99\x94\x21\x1d\x06\xb6\x80\x1c\x55\xc3\xc4\x8d\x86\xf0\xe4\x22\xba\x4b\x29\xb0\x57\x2b\x1f\xbe\x46\x16\xca\x33\x8f\xf8\x64\x15\xea\x76\x10\x00\x25\x2f\x9b\xae\x47\x34\x52\xce\xd3\x1e\xa2\x3c\x5d\x28\xd9\xbd\x4b\xba\x4f\x1d\xcf\x75\xd6\xca\x05\x15\xd2\x97\x41\xdd\x06\xd2\x62\x3e\x76\x2c\x17\x03\x11\x32\x8a\x1a\xb6\xc7\x47\xa9\xb6\x87\xb7\x31\x47\x16\xa2\x10\x3f\xf4\xae\x86\x37\x36\xf4\x76\x7b\x4c\x87\x98\xeb\x3e\x95\x0c\x1f\xe3\x25\x97\x7a\x11\x0d\x27\xd7\x9d\xe4\xfb\x7e\x36\x5c\xc6\x12\x8e\x31\x51\x85\xfb\x1c\xe6\x01\x50\xff\xb9\x53\x00\x1e\xea\x0b\x3d\xef\x09\xe0\xf9\xdb\xce\x03\xbf\xb7\x9d\x5c\xf6\xde\x41\x18\xbd\x81\x20\x12\x3a\xb1\x85\xfd\x94\x32\x81\x31\x9f\xcd\x81\x73\xd7\xdb\xbf\xed\xf7\xe7\xc2\x99\xd2\x02\x4f\xec\x44\x35\xa3\xe6\xea\xc4\x85\x67\xb7\x62\xa1\xa9\x1a\x8c\x59\xad\x6a\x02\xe8\x6a\xc6\x4e\x68\x47\x49\x52\x9b\xe1\xf2\x3c\x09\xb2\xdf\x86\xff\x59\x83\x98\xeb\x0a\x57\x9b\xbb\x96\xb6\x8c\x30\x8d\xfb\xc9\xae\xbe\x3c\xe4\x57\xce\x7f\x01\x00\x00\xff\xff\x7f\xbe\x4f\xde\xf5\x08\x00\x00")

func tempIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tempIndexHtml,
		"temp/index.html",
	)
}

func tempIndexHtml() (*asset, error) {
	bytes, err := tempIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "temp/index.html", size: 2293, mode: os.FileMode(438), modTime: time.Unix(1472101561, 0)}
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
	"temp/index.html": tempIndexHtml,
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
	"temp": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{tempIndexHtml, map[string]*bintree{}},
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

