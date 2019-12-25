// Package migrations Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// 1_add_table_gopkg.down.sql
// 1_add_table_gopkg.up.sql
package migrations

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

var __1_add_table_gopkgDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\xcf\x2f\xc8\x4e\x8f\x2f\x48\x4c\xce\x4e\x4c\x4f\xb5\x06\x04\x00\x00\xff\xff\xca\xcc\x4f\x81\x23\x00\x00\x00")

func _1_add_table_gopkgDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_add_table_gopkgDownSql,
		"1_add_table_gopkg.down.sql",
	)
}

func _1_add_table_gopkgDownSql() (*asset, error) {
	bytes, err := _1_add_table_gopkgDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_add_table_gopkg.down.sql", size: 35, mode: os.FileMode(420), modTime: time.Unix(1576313531, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_add_table_gopkgUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcf\x8b\xd3\x40\x14\xc7\xef\xfd\x2b\xde\xad\x09\x78\xa8\x8b\x05\x41\xf6\x30\x4d\x9f\x75\x30\x9d\x48\x76\x22\xee\x29\x13\x93\x31\x0c\xbb\x4d\x42\x9a\x8a\x47\x57\x10\x2a\x28\xdb\x83\xba\xc8\xae\xb0\x88\xb0\x17\x0f\x3d\x14\xc1\x55\xf0\x9f\x71\x92\xed\x7f\x21\x89\xfb\xa3\x2e\x15\x3c\xce\xf0\x3e\xdf\xcf\xfb\xf2\x2c\x17\x09\x47\xe0\xa4\x67\x23\xc4\x69\xb6\x13\xfb\x59\x10\xee\x04\xb1\x04\xa3\x05\xa0\x22\xe8\xd1\x01\x65\xdc\xd8\xe8\x98\xc0\x1c\x0e\xcc\xb3\x6d\x20\x1e\x77\x7c\xca\x2c\x17\x87\xc8\xf8\x8d\x16\x80\x48\x82\x91\x14\x60\xdd\x23\xae\xd1\x5d\x1d\xb5\x9c\x61\x3d\x03\x6d\xfd\xfa\xa5\x9e\xbd\xa9\x4e\xe6\xed\x7a\x7e\x9c\x4e\xf2\x50\xc2\x43\xe2\x36\xc8\xcd\xce\xbf\x98\xf2\x74\x56\x1d\xef\xe9\xa3\xb9\xfe\xf8\xbc\x21\x45\x24\xc7\x61\xae\xb2\x42\xa5\x89\xb8\x0c\xd8\xe8\x76\x4d\xe8\xe3\x5d\xe2\xd9\xd7\x43\xca\xfd\xfd\xb3\x9f\x7f\xac\x93\x2c\x0a\x0a\x19\xf9\x41\x01\x7d\xc2\x91\xd3\x21\xae\xb1\x96\x87\x8b\xf2\xfd\xbc\x3c\xf8\xba\x3c\x58\x34\x58\x24\x77\xe5\x39\x56\xf3\x85\x1a\x49\xa3\xde\xb7\xa6\xd6\x3b\xf5\xf4\x78\xf9\xe1\xf3\x4a\x44\x98\xcb\xff\x30\xeb\xe9\xa1\xfe\x7e\xba\x82\x89\xa7\x32\x1f\x37\x45\x55\x52\xfc\x7d\x83\x0b\x71\xe7\x0a\xaf\x5e\x4d\xcb\xa3\x2f\xf0\xeb\xdb\xec\xec\xe4\xc5\xf2\xed\x5e\x13\xf1\xc0\xa5\x43\xe2\x6e\xc3\x7d\xdc\x06\x43\x45\x66\xfd\x47\x59\x1f\x1f\x81\x50\xd1\x33\xff\xbc\x1b\x29\x84\x21\xae\x7a\x0a\x13\xbc\x2d\xca\x06\xd0\xe3\x2e\xe2\xf5\x5e\xd5\xe2\x93\xfe\xf1\xae\xdd\x32\x91\x0d\x28\xc3\x4d\x9a\x24\x69\xbf\x77\xb9\x51\x7d\x90\x2d\xe4\x9b\x93\xe2\xc9\xed\xd1\xe3\x5b\x60\x39\xb6\x4d\x38\x5e\xbc\xfd\x58\x26\x32\x0f\x76\xfd\x50\xdd\xf9\x1d\x00\x00\xff\xff\x40\x20\xfb\x05\x7f\x02\x00\x00")

func _1_add_table_gopkgUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_add_table_gopkgUpSql,
		"1_add_table_gopkg.up.sql",
	)
}

func _1_add_table_gopkgUpSql() (*asset, error) {
	bytes, err := _1_add_table_gopkgUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_add_table_gopkg.up.sql", size: 639, mode: os.FileMode(420), modTime: time.Unix(1576314788, 0)}
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
	"1_add_table_gopkg.down.sql": _1_add_table_gopkgDownSql,
	"1_add_table_gopkg.up.sql":   _1_add_table_gopkgUpSql,
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
	"1_add_table_gopkg.down.sql": &bintree{_1_add_table_gopkgDownSql, map[string]*bintree{}},
	"1_add_table_gopkg.up.sql":   &bintree{_1_add_table_gopkgUpSql, map[string]*bintree{}},
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
