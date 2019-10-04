// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1561059285_add_whisper_keys.down.sql (25B)
// 1561059285_add_whisper_keys.up.sql (112B)
// doc.go (373B)

package sqlite

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __1561059285_add_whisper_keysDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\xcf\xc8\x2c\x2e\x48\x2d\x8a\xcf\x4e\xad\x2c\xb6\xe6\x02\x04\x00\x00\xff\xff\x42\x93\x8e\x79\x19\x00\x00\x00")

func _1561059285_add_whisper_keysDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1561059285_add_whisper_keysDownSql,
		"1561059285_add_whisper_keys.down.sql",
	)
}

func _1561059285_add_whisper_keysDownSql() (*asset, error) {
	bytes, err := _1561059285_add_whisper_keysDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1561059285_add_whisper_keys.down.sql", size: 25, mode: os.FileMode(0644), modTime: time.Unix(1564484687, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb9, 0x31, 0x3f, 0xce, 0xfa, 0x44, 0x36, 0x1b, 0xb0, 0xec, 0x5d, 0xb, 0x90, 0xb, 0x21, 0x4f, 0xd5, 0xe5, 0x50, 0xed, 0xc7, 0x43, 0xdf, 0x83, 0xb4, 0x3a, 0xc1, 0x55, 0x2e, 0x53, 0x7c, 0x67}}
	return a, nil
}

var __1561059285_add_whisper_keysUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xb1\x0a\xc2\x40\x0c\x06\xe0\xfd\x9e\xe2\x1f\x15\x7c\x03\xa7\xde\x19\x35\x18\x13\x09\x29\xb5\x53\x11\x3d\x68\xe9\x22\x56\x90\xbe\xbd\x5f\x71\x6a\x82\x10\x4d\x16\xc2\x6f\x9c\x96\x77\xfd\x0c\x73\x5d\x17\x6c\x12\xf0\x1c\x1f\xdf\x61\x7a\x21\xe8\x1e\xb8\x39\x5f\x1b\xef\x71\xa1\x1e\xa6\x28\xa6\x47\xe1\x12\xe0\x93\x9a\xd3\x2e\x01\x73\x5d\x91\xc5\x32\xd4\x02\xda\x8a\xa4\x2d\x3a\x8e\xb3\xb5\x01\xb7\x8e\x0f\xfb\xf4\x0f\x00\x00\xff\xff\x6e\x23\x28\x7d\x70\x00\x00\x00")

func _1561059285_add_whisper_keysUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1561059285_add_whisper_keysUpSql,
		"1561059285_add_whisper_keys.up.sql",
	)
}

func _1561059285_add_whisper_keysUpSql() (*asset, error) {
	bytes, err := _1561059285_add_whisper_keysUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1561059285_add_whisper_keys.up.sql", size: 112, mode: os.FileMode(0644), modTime: time.Unix(1564484687, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x25, 0x41, 0xc, 0x92, 0xdd, 0x9e, 0xff, 0x5d, 0xd0, 0x93, 0xe4, 0x24, 0x50, 0x29, 0xcf, 0xc6, 0xf7, 0x49, 0x3c, 0x73, 0xd9, 0x8c, 0xfa, 0xf2, 0xcf, 0xf6, 0x6f, 0xbc, 0x31, 0xe6, 0xf7, 0xe2}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x3d\x72\xeb\x30\x0c\x84\x7b\x9d\x62\xc7\x8d\x9b\x27\xb2\x79\x55\xba\x94\xe9\x73\x01\x98\x5a\x91\x18\x4b\xa4\x42\xc0\x7f\xb7\xcf\xc8\xe3\xc2\x5d\xda\x1d\x7c\x1f\x76\x63\xc4\x77\x51\xc3\xac\x0b\xa1\x86\xca\x44\x33\xe9\x0f\x9c\x98\xe4\x62\xc4\x21\xab\x97\xcb\x29\xa4\xb6\x46\x73\xf1\x8b\x8d\xba\xc6\x55\x73\x17\x67\xbc\xfe\x3f\x0c\x31\x22\x49\x3d\x3a\x8a\xd4\x69\xe1\xd3\x65\x30\x97\xee\x5a\x33\x6e\xea\x05\x82\xad\x73\xd6\x7b\xc0\xa7\x63\xa1\x98\xc3\x8b\xf8\xd1\xe0\x85\x48\x62\xdc\x35\x73\xeb\xc8\x6d\x3c\x69\x9d\xc4\x25\xec\xd1\xd7\xfc\x96\xec\x0d\x93\x2c\x0b\x27\xcc\xbd\xad\x4f\xd6\x64\x25\x26\xed\x4c\xde\xfa\xe3\x1f\xc4\x8c\x8e\x2a\x2b\x6d\xe7\x8b\x5c\x89\xda\x5e\xef\x21\x75\xfa\x7b\x11\x6e\xad\x9f\x0d\x62\xe0\x7d\x63\x72\x4e\x61\x18\x36\x49\x67\xc9\x84\xfd\x2c\xea\x1c\x86\x18\x73\xfb\xc8\xac\xdc\xa9\xf7\x8e\xe3\x76\xce\xaf\x2b\x8c\x0d\x21\xbc\xd4\xda\xaa\x85\xdc\x10\x86\xdf\x00\x00\x00\xff\xff\x21\xa5\x75\x05\x75\x01\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 373, mode: os.FileMode(0644), modTime: time.Unix(1564484687, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x23, 0x6a, 0xc1, 0xce, 0x94, 0xf6, 0xef, 0xf1, 0x97, 0x95, 0xb, 0x35, 0xaf, 0x5f, 0xe7, 0x5f, 0xac, 0x6e, 0xb8, 0xab, 0xba, 0xb5, 0x35, 0x97, 0x22, 0x36, 0x11, 0xce, 0x44, 0xfc, 0xfa, 0xac}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"1561059285_add_whisper_keys.down.sql": _1561059285_add_whisper_keysDownSql,

	"1561059285_add_whisper_keys.up.sql": _1561059285_add_whisper_keysUpSql,

	"doc.go": docGo,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"1561059285_add_whisper_keys.down.sql": &bintree{_1561059285_add_whisper_keysDownSql, map[string]*bintree{}},
	"1561059285_add_whisper_keys.up.sql":   &bintree{_1561059285_add_whisper_keysUpSql, map[string]*bintree{}},
	"doc.go":                               &bintree{docGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
