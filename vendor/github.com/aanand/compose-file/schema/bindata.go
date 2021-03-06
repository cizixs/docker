// Code generated by go-bindata.
// sources:
// data/config_schema_v3.json
// DO NOT EDIT!

package schema

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

var _dataConfig_schema_v3Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x4b\x93\xdb\xa8\x13\xbf\xfb\x53\x4c\x29\xb9\xc5\x33\x93\xaa\x7f\xea\x5f\xb5\xb9\xed\x71\x4f\xbb\xe7\x75\x29\x2a\x2c\x61\x9b\x8c\x10\x04\x90\x13\x27\xe5\xef\xbe\xa0\x97\x01\xf1\xb2\xad\xec\xe4\xb0\x73\x98\x9a\x81\xee\xa6\x1f\x3f\x9a\xa6\xd1\x8f\xd5\xc3\x43\xf6\x96\x97\x07\x88\x41\xf6\xf1\x21\x3b\x08\x41\x3f\x3e\x3f\x7f\xe6\xa4\x79\xec\x47\x9f\x08\xdb\x3f\x57\x0c\xec\xc4\xe3\xfb\x0f\xcf\xfd\xd8\x9b\x6c\xad\xf8\x50\xa5\x58\x4a\xd2\xec\xd0\xbe\xe8\x67\x8a\xe3\xff\x9e\x14\x73\x4f\x20\x4e\x14\x2a\x12\xb2\xfd\x0c\x4b\xd1\x8f\x31\xf8\xa5\x45\x0c\x2a\xd6\x4d\x76\x84\x8c\x23\x49\x9d\xaf\x57\x6a\x8e\x32\x42\x21\x13\x08\x72\x39\xab\x54\x93\x63\x23\xc9\x38\xa0\x89\xe5\x82\xa1\x66\x9f\x75\xc3\xe7\x4e\x82\x9c\xe4\x90\x1d\x51\xa9\x49\x98\x14\x7d\xf3\x7c\x91\xff\x3c\x91\xad\x6d\xa9\x9a\xb2\xdd\x38\x05\x42\x40\xd6\xfc\x35\xd7\xad\x9b\xfe\xb4\x01\x8f\xdf\x7f\x7f\xfc\xfb\xfd\xe3\x6f\x4f\xc5\x63\xfe\xee\xad\x31\xad\xbc\xcb\xe0\xae\x5f\xbe\x82\x3b\xd4\x20\x21\xad\x99\xd6\xcf\x26\xca\xf3\xf0\xd7\x79\x5a\x18\x54\x55\x47\x0c\x6a\x63\xed\x1d\xa8\x39\x34\x6d\x6e\xa0\xf8\x4a\xd8\x4b\xcc\xe6\x89\xec\x95\x6c\x1e\xd6\x77\xd8\x6c\x9a\x73\x24\x75\x8b\xa3\x11\x1c\xa9\x5e\xc9\x98\x7e\xf9\xfb\xe2\xb7\x1a\x8d\x0e\xd2\xf6\x14\xda\xda\x9d\x82\x06\xda\x5d\xae\x72\xa1\xcd\xef\xab\xc9\x59\x1e\x2f\x55\x90\xd6\xe4\xa4\xc6\x3c\xfe\xe8\x09\x30\x6c\x44\x36\xb9\x40\xf2\x6d\x5b\x54\x57\xb6\x47\x49\x03\xff\x54\x22\x36\xda\xe0\x83\x94\x6c\x6d\x6c\x4d\x4e\x37\x6f\xfc\xe7\x0f\xf8\x34\xef\xb1\x65\x9a\x97\x99\x4b\xc0\x6f\xa2\x33\x2a\xbc\x74\xef\x02\x52\xbe\x40\xb6\x43\x35\x4c\xe5\x00\x6c\xcf\x03\x2e\xab\x11\x17\x05\x61\x45\x85\xa4\xf6\x67\x8b\x7d\x26\x2f\x8e\xa7\x89\x55\xfb\x2f\x5f\x39\x04\x66\x25\xa0\x85\x14\x67\xd8\x01\x18\x03\xa7\x6c\x2d\x01\x24\x20\xe6\x6e\x13\x1f\xb2\xb6\x41\x5f\x5a\xf8\xc7\x40\x22\x58\x0b\x6d\xb9\x95\x54\x6e\x79\xc1\x7b\x46\x5a\x5a\x50\xc0\x14\xc0\xc2\xee\x97\x71\xc5\x18\x34\x4b\xa1\xee\x1a\x3b\x12\x3c\x2f\x31\x07\x50\x03\x59\xd1\x00\x1c\x03\x92\xda\x75\xb0\xa9\x78\xd1\x9f\x7f\x41\x18\xed\x8a\x9e\x9f\x5b\x02\xa6\xc3\x70\xd1\x78\x54\x4d\x08\xd8\xbd\x18\x05\x6d\xa5\x5b\x66\x31\x16\x1c\x02\x56\x1e\x6e\xe4\x27\x58\xba\x2f\xc5\x77\x12\x28\xec\x44\x09\xea\xf1\xf2\xcb\x01\x01\x36\xc7\x62\xca\x25\x57\xbb\x41\x72\x23\x46\x1a\x3c\xee\x86\x94\x04\x33\x25\x79\xc5\xff\x8d\x12\x0e\x6d\xc7\x58\x06\xea\x53\x93\xa9\x86\x4f\x46\x8e\xcd\x68\xb8\x74\x4a\xd3\xe2\x2d\x64\xaa\xa4\x33\x28\x77\x84\x61\xa0\x94\x1d\xd7\xd6\xa6\x0d\x4f\x3b\x90\xa7\x3b\x50\xb7\x41\x1d\xeb\xa0\x96\xde\x69\x5e\x96\x87\xb8\x14\xcf\x40\x71\x20\x5c\xa4\xe7\x70\x8d\xfd\x00\x41\x2d\x0e\xb2\x28\x2e\x5f\x02\xec\x3a\x95\xc1\x2d\x97\x4d\x01\x39\xc2\x60\x1f\x27\xa2\x65\x8c\xa4\x06\x5b\x58\xdf\x64\xe7\xa2\xce\xd7\xc4\x92\xfd\x5e\x91\xfa\x10\x37\xab\x5c\x86\xe9\xd8\x99\x5f\x31\x24\x6f\x14\xa9\x07\x38\xa1\x97\x82\xcb\x9e\x8c\x17\x20\xbd\x42\xc1\xea\xd3\x20\xfd\xf4\xd4\x17\x9f\x81\x5d\xd5\xfd\x55\xd7\x59\x6e\x97\x0b\xea\x67\x3e\x66\x8e\x58\x16\xa6\x15\x14\x46\x54\x30\x28\x55\xdd\xc0\x20\xf7\xc4\xf5\x42\x3a\x14\xfb\x05\x26\x95\x0f\xa0\x33\x62\xdb\x37\xde\x4c\x7d\xf5\x41\xd8\xb1\x5d\x5d\x3f\x26\x85\x2e\x7a\x81\x88\x58\xe3\x53\x2f\x55\xcd\x8b\xba\x71\x88\x75\x74\xa0\x46\x80\xc3\xf8\x66\xf7\x3a\xd2\x90\x86\xe8\xf1\x43\x22\x26\x5c\xbc\xff\x0f\xf2\x7a\x58\xbd\x32\xd3\x6b\xe4\x88\xa8\x8b\x2a\xdd\x76\x73\x29\x92\x47\x76\xdb\x4f\x2e\xe1\x29\xaa\xfc\xb9\xa2\xcb\x10\xfa\x06\xa3\x84\x89\xd9\xee\xfa\x77\x8e\xfb\x7e\xe9\xbb\x4f\x7b\x2a\x13\xb7\x2c\x97\xf6\xd0\xbc\xb5\x6c\x09\xa9\x21\x68\x8c\xd4\xc3\x20\xa8\x64\xc9\x5c\x9f\x12\x28\xb9\x00\x2c\x7a\xa1\xe0\xb0\x6c\x19\x12\xa7\x42\x9e\x07\x8b\xd7\x19\xfc\x80\x0b\x8e\xbe\x43\x33\x9a\x97\x7c\x3f\x08\xca\x0d\x1e\x51\xa1\x46\x6a\x03\x9b\xa8\x89\x5c\x10\x2a\xe5\xef\x25\xe6\xa2\x66\x2a\xd2\x3d\x03\x25\x2c\x24\x36\x11\xa9\x5c\x0c\x6b\x3d\xb6\x55\xcb\x80\xc2\xb3\x21\x46\x60\xba\xbb\xf1\x76\x20\x44\x3c\x66\x6d\x8d\x30\xf2\x83\xd9\x91\x25\x13\x12\x79\x9f\xc4\xdd\xb9\x3b\x90\xb7\x2f\x9a\xca\x6b\x86\xc4\x26\x73\xa5\xbb\x40\xe9\x10\xae\x1c\x12\x4a\x86\x03\x60\x66\x94\x02\x7a\x74\x0c\x9c\xec\x84\x9b\xc1\x55\x50\x38\xf5\x32\x3a\xb8\x9d\xbc\xf5\xa0\x48\xee\xa4\xbf\x2a\x27\xdb\x6a\xe4\xde\xb4\x78\x76\xa6\xc5\x96\x47\xab\x3b\xbd\xbf\xb8\xe8\x4e\x56\x25\x8c\x42\x76\x85\xdc\x2a\xac\x2c\x75\xaf\xe8\xf0\x5a\xb7\x89\x51\x80\xab\xd7\xa7\x93\xda\xfd\xbe\xcd\x04\xb8\xf1\x94\xb8\x74\x49\x3d\x8d\x3f\x85\x0f\x76\x34\x92\x87\xcb\xa7\x02\x61\x48\x5a\x11\xa1\x62\x50\x8e\x59\x9e\x1f\x32\xdd\x2f\xdd\xc0\xb9\x29\x72\x1d\xa1\xd6\x16\x8d\xc4\x4d\xa3\x5c\x20\x6c\x81\xb2\x5b\x8b\x06\xad\x51\x09\x78\x2c\x81\xdc\x71\x3b\x6c\x69\x05\x04\x2c\xfa\x37\xa2\xab\x52\x76\x20\x57\x53\xc0\x40\x5d\x43\xb9\x28\x4e\xc9\x7d\x32\x06\x35\x38\xdd\x74\x96\x75\xec\x3b\x80\xea\x96\xc1\x02\x94\x62\x78\x88\x8a\x20\x4e\x3a\x5f\x3a\x86\x38\x93\x40\xda\x92\x18\x7c\x2b\xc6\x65\x3b\x12\xe7\x8e\xf1\xd6\x54\xa9\x17\x3b\x0d\x09\x9c\xb4\xac\x9c\x39\xfb\xe6\x10\x5d\xce\x68\x0f\x62\xc6\x15\x67\xa6\xcb\x09\x95\x6f\xa6\x7b\x77\x94\x3f\x7a\x24\x0c\x45\x5e\x41\x89\x44\xfb\x69\x29\x0b\x25\xa4\x7b\x27\xa7\x00\xe2\x4e\x04\x2a\x38\xa8\x12\x06\x53\x11\xdd\xac\x1d\xc3\x57\xd4\x54\xe4\xeb\x15\x0b\x2e\x07\x25\x5a\xcb\xfa\xd1\xca\x77\xf7\x3a\x5a\xea\x0e\xa4\xa9\x57\x9f\xd8\xf7\x9a\x75\xc7\x81\x3d\xe1\x33\x92\xf5\x27\xba\xf8\x33\xa6\x27\xd3\x97\xb4\x8d\x36\x63\x30\xc4\x84\x39\x01\x18\xb0\x31\xf1\xd5\x39\x66\xe1\x48\xb6\xc0\xa1\x96\xd4\xbc\x1b\xa8\xd4\x5d\x6d\xf1\x4b\x42\xbc\x41\x97\xc7\xf3\x11\xa2\x00\x2f\xb5\x39\x92\xdb\x99\x99\xf3\x08\x36\xd6\x9e\xb7\x01\x7a\x75\x9d\xad\x80\x98\xd6\x71\xdd\x07\x0a\xde\x6e\x25\x42\x42\xd0\xbc\xfc\x38\xdf\x58\xd3\x6f\x17\x67\xff\x5d\xe2\xbe\x9c\x37\xbe\x44\x78\xa2\xba\x99\xee\xb1\xeb\xc9\x57\x79\x72\x88\xbd\xcf\x00\xcb\xe9\x7f\x65\x7d\x77\x47\x5a\x1c\xbe\x9a\x88\xa4\x8c\x81\xea\xbf\x8c\x31\x48\x79\x7d\x7c\x05\xce\xc4\x1b\x2f\x07\x57\x80\xc6\x6a\x18\x69\xe0\x99\xdf\x08\x43\x71\x4e\x6e\x77\x0f\x1c\xb9\xa9\x86\x4d\xf6\x71\xfe\x45\x9a\x99\x42\x43\xbd\x84\x91\xc4\xd3\xfe\xb4\x16\x1d\x9c\x17\xb6\x7c\x41\xd8\x3e\xbd\x0b\x1c\x14\xa1\x67\xa9\x9f\x94\x61\x17\xe8\xd3\xb8\x63\x6a\x15\x97\xa3\x77\xe7\x9f\x55\x79\x32\x95\xc6\x3f\xfb\xc8\x4a\xd9\xd9\x9c\x66\x1d\x8b\x1f\x66\x03\xad\xff\x40\x2a\x37\xfc\x63\x91\xf4\x8f\xbc\x5a\x9e\xc8\xf5\x7a\xdb\x17\x46\xe7\xa7\x57\x76\xfb\x6e\xfc\x04\x2a\x0f\x6f\xf6\xd5\xf8\xfb\xbc\x3a\xaf\xfe\x09\x00\x00\xff\xff\xa9\x28\xf2\x96\x34\x2a\x00\x00")

func dataConfig_schema_v3JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v3Json,
		"data/config_schema_v3.json",
	)
}

func dataConfig_schema_v3Json() (*asset, error) {
	bytes, err := dataConfig_schema_v3JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/config_schema_v3.json", size: 10804, mode: os.FileMode(420), modTime: time.Unix(1478512187, 0)}
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
	"data/config_schema_v3.json": dataConfig_schema_v3Json,
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
	"data": &bintree{nil, map[string]*bintree{
		"config_schema_v3.json": &bintree{dataConfig_schema_v3Json, map[string]*bintree{}},
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

