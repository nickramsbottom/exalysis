// Code generated by go-bindata. DO NOT EDIT.
// sources:
// comment-format.md (1.36kB)
// generalize-name.md (243B)
// minimal-conditional.md (236B)
// missing-comment.md (121B)
// plus-used.md (224B)
// strings-join.md (214B)
// stub.md (463B)
// use-string-placeholder.md (331B)
// wrong-comment-format.md (121B)

package tpl

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

var _commentFormatMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x4d\x8f\xe4\x44\x0c\x3d\x93\x5f\x61\x89\xc3\xee\x48\xe9\x64\x00\x71\x60\xae\x2b\x18\x71\x40\x5a\xa1\xb9\x8d\x10\xa9\x54\x39\x29\xab\x2b\xe5\xa8\xec\xea\xd0\x20\xfe\x3b\x72\xfa\x43\x33\x8b\x84\xc4\x25\x87\x94\x9f\xed\xf7\xfc\xec\xe6\x70\x38\x34\xcf\x0c\xd1\x09\xbc\xce\x05\x9d\xc2\x5c\x29\x60\xa2\x8c\xf2\xdb\xc7\xa8\xba\xca\x53\xdf\xcf\x9c\x5c\x9e\x3b\x2e\x73\x1f\xd8\xf7\x38\x4d\xe8\x95\x4e\xf8\xfb\xcc\x5d\xd4\x25\x3d\x80\x1b\xb9\x2a\x44\xde\x40\x19\xb6\x42\x8a\xe0\x79\x59\x30\xab\x74\xf0\x12\xb1\xe0\x07\x01\x07\xa2\x2e\x07\x57\x02\x4c\x5c\x16\xa7\x70\xe6\x0a\xde\x65\xa8\x82\xf6\x0b\x02\xfb\x6a\x20\xa7\xc4\xf9\x9e\xa1\x85\x2d\x92\x8f\x7b\xa4\x46\xcc\x30\x22\x68\x2d\x19\x03\x50\x56\x86\x11\x5d\x55\x9a\x6a\xfa\x02\xef\xaa\xf2\xe2\x94\xbc\x4b\xe9\x0c\xe3\xd9\xc0\x30\xcc\x1c\xd8\x0f\xa0\xcc\xa9\x6b\x9a\x97\x88\xb0\x3a\x7f\x74\xf3\xbd\x63\x90\xc8\x35\x05\x6b\xb6\x28\x6c\xa4\x71\x07\x6e\x5c\x02\x0c\x9f\x2f\xb1\x03\x4c\x9c\x12\x6f\x18\x6e\x79\x6f\x49\xb2\x5b\xb0\x85\x44\x47\x04\x8d\x24\x4f\x4d\x33\x0c\xc3\xcc\x4d\xdf\xc3\xe7\x7b\x1d\x3e\xc2\x5a\xf8\x44\x01\x05\xa2\xcb\xe1\x0c\x9e\xf3\x09\x8b\x58\xdb\x0b\x6a\xe4\x20\xbb\x20\x35\x93\x0a\xe8\x79\xbd\x92\xa8\xb2\x93\x86\x82\x9e\x56\x94\xae\x59\xdf\xe4\xb4\x4a\x4d\xf3\x13\x17\xc0\x3f\x56\x2e\x8a\x01\xa6\x9a\xbd\x69\x21\xed\x2d\x6d\x6b\xa5\x6c\x0e\xa6\xab\xcb\xe1\xde\xf8\xc9\x15\x72\x63\x42\x69\x77\x3e\xff\x2d\xc6\x85\xa5\xc1\x3d\x67\xa5\x5c\xf1\xf2\xe6\xe0\x84\x65\x84\x35\x16\x27\xf8\x96\xfb\xcb\x28\x2f\xfc\x4b\xba\xf2\x34\x4e\x7b\xad\x95\x39\x8b\x79\x66\xa1\x94\x28\x91\x62\x91\xae\xb1\xae\x6f\x88\x8f\x3a\x8a\x8d\xf9\xc1\x3e\xf0\x57\xf3\x55\xdf\x43\xd7\x75\xcd\xdf\x17\xb6\x36\xbf\xd7\x4f\x1c\xf0\x57\x3c\x11\x6e\x9f\xae\x96\x79\xe3\x5d\xd2\x58\xc7\xce\xf3\x72\xb5\x71\x3f\x73\xbf\xd1\x91\xfa\x7f\xa3\xbe\xbe\x92\x3e\x08\x66\xc5\xec\x51\x1e\xc0\x42\x61\xaa\x45\x23\x16\x90\x15\x3d\x4d\x84\x02\x1a\x9d\x9a\xdd\xee\x26\xbd\x09\x35\x22\x4c\x35\x25\xb8\xa7\x68\x01\x73\xa0\x3c\xdf\x04\x5a\xb1\x10\x87\xae\x69\x7e\xce\xe0\x42\xa0\xdd\xaa\xca\xc0\xab\xd2\x42\x7f\x5a\xa4\x4d\xbe\xa0\x0b\x6e\x34\x49\xce\xed\xd5\x6c\xf6\x64\xe2\x17\xbc\x14\x0d\xbb\xcd\x05\x68\x31\x33\x99\x95\x78\xb3\x95\x2a\xe0\x39\x20\x24\xe6\xa3\x49\x07\xaf\xcf\xfc\x41\xbe\xd8\x0d\xb3\xff\x55\xa5\xa7\xbe\xdf\x22\xad\x2b\x16\x51\xe7\x8f\x58\x76\xb1\xbe\x7d\xfc\xe6\xfb\xfe\xf1\x87\xfe\xbb\xc7\x7e\xe6\xc3\x3b\xf0\x61\x5f\xa1\x77\x5f\x2e\xf3\xc1\xe5\x70\xb8\x84\xf6\x0f\x20\xd5\x47\xd8\xaf\x8a\xbd\xdb\xe9\xb8\x17\xbb\xff\x79\x68\x6f\xdb\x78\xb1\xd2\x30\xb3\x75\x39\xd8\x5a\xba\x23\x82\x83\x85\x77\x0f\x2a\xef\x72\xec\xe4\x05\x77\x43\x03\xe7\xf7\xf2\xd3\x7e\x1b\x80\xa7\x89\x3c\xb9\x04\xa2\xe7\x84\x97\x6b\xd6\xc2\xeb\x8f\xb7\x93\x05\xcf\xfc\x7f\xee\x9a\x32\x08\xe2\x55\x75\x27\x9c\x6d\x06\x23\x46\xca\x7b\x37\x82\xe0\x23\x93\xb7\x4d\xb4\x5b\xfa\x4f\x00\x00\x00\xff\xff\x87\x3e\x78\x45\x50\x05\x00\x00")

func commentFormatMdBytes() ([]byte, error) {
	return bindataRead(
		_commentFormatMd,
		"comment-format.md",
	)
}

func commentFormatMd() (*asset, error) {
	bytes, err := commentFormatMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "comment-format.md", size: 1360, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x97, 0x26, 0x5c, 0x8a, 0xbe, 0x4, 0x92, 0x9f, 0x4c, 0x2, 0xa5, 0x6a, 0x8d, 0x22, 0xdb, 0xba, 0x2a, 0xf0, 0x7d, 0x95, 0xfb, 0xa5, 0x93, 0xe7, 0x34, 0x84, 0x27, 0xaa, 0x5c, 0x7f, 0x8d, 0xd3}}
	return a, nil
}

var _generalizeNameMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xce\xb1\x4e\x04\x31\x0c\x04\xd0\x9e\xaf\x98\x0f\x80\xfd\x07\xe8\x10\x05\xc5\x21\xd1\x26\x9b\xcc\x5e\xac\xdb\xb3\x4f\x89\xa3\xd5\xfe\x3d\xf2\x35\xb4\x96\x67\xe6\xbd\xe1\x13\x83\xc4\x69\x13\x9d\x1b\x3b\xb5\xb0\xc2\x1b\xa1\xf9\xce\x81\xf4\xbe\x4b\x61\x42\xd6\x8a\xf4\x61\x6b\x42\x95\xce\xe2\xfb\x09\xd1\x88\x75\x14\xab\x5c\xf0\xd3\x64\x40\x06\xbe\xbf\x5e\xb1\x4e\x87\x38\x0e\x9b\x7b\xc5\x4a\xac\x74\x67\x87\x6c\xcf\xe2\x74\x69\xb9\xf3\x57\xbc\x25\x6c\x53\x8b\x8b\x29\x0e\xeb\x37\x56\x1c\xe2\x0d\x59\xcf\xe7\x3a\xc6\x7c\x3c\x76\x09\x8f\x41\x7c\xc1\x85\x8c\x92\xc0\x96\xac\xd8\xe4\x3a\x3b\x61\xd3\xd1\xec\x88\xa7\xd2\xb2\x5e\xf9\xcf\x8a\xdb\x3d\xdf\x08\x0f\x5d\x6c\x2c\x2f\x7f\x01\x00\x00\xff\xff\xbe\xd1\x58\x73\xf3\x00\x00\x00")

func generalizeNameMdBytes() ([]byte, error) {
	return bindataRead(
		_generalizeNameMd,
		"generalize-name.md",
	)
}

func generalizeNameMd() (*asset, error) {
	bytes, err := generalizeNameMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "generalize-name.md", size: 243, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x67, 0xaf, 0x4d, 0x60, 0xd2, 0x7c, 0x68, 0xe6, 0x2, 0x7a, 0xcb, 0x21, 0x55, 0x20, 0xc1, 0x53, 0xd7, 0x66, 0x34, 0xb4, 0xff, 0x98, 0xd, 0xdc, 0x9, 0x7b, 0xa3, 0xcf, 0xc8, 0x7b, 0x2f, 0x9b}}
	return a, nil
}

var _minimalConditionalMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8f\x31\x52\xc3\x40\x0c\x45\x7b\x4e\xf1\x3b\x9a\x90\x2b\x50\xa4\xa0\xa0\xa0\xe1\x02\x8b\xad\x78\x77\xb0\x57\x9e\xaf\x2f\x3c\xbe\x3d\x63\x27\xb5\x9e\xde\x93\xde\x70\x4b\xd2\xba\xe6\x1d\xbb\x27\x0a\x0d\xb4\xd5\x8a\x5a\x9f\xb0\x78\x08\x7e\x87\xaa\xc1\x53\x6b\x0a\x21\xb6\x3e\x5d\xf1\x5d\x5b\xa0\x05\xbe\x3e\x2f\xf8\x49\xa1\x75\x4c\xd6\x8d\x65\xbe\xe0\xc3\xb1\xd2\x27\x96\x65\x31\x06\xc4\x1d\x72\x94\x3f\x6f\x23\xc6\x5c\xe7\x36\x3c\xf4\x83\x8f\x86\xad\x1a\xed\x28\xec\x18\x4a\x7f\x52\x4d\x57\xdc\x4a\x3f\x4f\xda\x9c\xbf\x47\x1d\xd5\xb7\xc3\x43\x0b\x31\x07\x25\xed\x98\xf3\xa1\x09\x87\x6a\xd1\xb9\x31\x7a\x7f\xd5\xf3\x8d\x13\x09\x9b\xef\xef\x2f\xff\x01\x00\x00\xff\xff\x5f\x23\x66\x73\xec\x00\x00\x00")

func minimalConditionalMdBytes() ([]byte, error) {
	return bindataRead(
		_minimalConditionalMd,
		"minimal-conditional.md",
	)
}

func minimalConditionalMd() (*asset, error) {
	bytes, err := minimalConditionalMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "minimal-conditional.md", size: 236, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x31, 0xcd, 0x9e, 0x36, 0xfb, 0xd4, 0xe2, 0x6e, 0x63, 0x53, 0xb8, 0x56, 0x1e, 0x36, 0x22, 0x41, 0x9b, 0x3, 0xc0, 0xca, 0x31, 0x97, 0x14, 0xb0, 0x2e, 0xf9, 0x8b, 0x25, 0xeb, 0x83, 0x3b, 0x29}}
	return a, nil
}

var _missingCommentMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xca\xb1\x0d\xc2\x40\x0c\x05\xd0\x9e\x29\x7e\x43\x07\x59\x02\x09\xe6\x30\x89\xa5\x3b\x61\xdf\x97\xe2\x1f\xd0\x6d\x4f\x95\xfe\xdd\xf1\x60\xe6\x31\xba\x26\x4a\x33\x1c\x65\xb3\xa0\x66\xc2\xb5\xd0\x07\x5e\x44\x35\x1e\xb1\xa1\xd9\xd7\x61\x58\x99\xe9\x43\x0b\x9e\xdc\x91\xdc\x1d\x9b\xcb\x7a\xd4\xed\x14\x41\x7e\x60\x82\x9a\x9f\xba\x50\xbe\xaa\x73\xe0\xed\xc1\xdf\x72\xf9\x07\x00\x00\xff\xff\x57\xf7\xfa\xba\x79\x00\x00\x00")

func missingCommentMdBytes() ([]byte, error) {
	return bindataRead(
		_missingCommentMd,
		"missing-comment.md",
	)
}

func missingCommentMd() (*asset, error) {
	bytes, err := missingCommentMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "missing-comment.md", size: 121, mode: os.FileMode(420), modTime: time.Unix(1542628451, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4, 0x25, 0x36, 0xe4, 0xde, 0x35, 0xd6, 0x99, 0x90, 0xef, 0xb2, 0x2d, 0x89, 0x9e, 0x17, 0x2c, 0xa5, 0xdb, 0x26, 0x10, 0xe8, 0xb3, 0xee, 0x54, 0x94, 0xab, 0x7, 0xa1, 0xa5, 0x43, 0xfb, 0xa6}}
	return a, nil
}

var _plusUsedMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8e\x3b\x6e\xc3\x30\x10\x44\xfb\x9c\x62\x52\xa5\x88\xa3\x7b\xa8\xcf\x01\xb4\x21\x97\xd4\x42\xf4\xae\x41\x0e\x11\xf8\xf6\x81\xac\x22\xe5\x0c\xe6\xf3\xbe\xb0\x62\xa8\xe2\x19\x13\xd2\x15\x73\x98\x57\x6c\x9f\x1b\x18\x48\xe1\x49\xa8\x2e\x54\x70\x57\x0c\x76\xf3\xba\xe0\x7b\x17\xc2\x06\x1e\xda\x8b\x26\xb6\x27\x9a\x56\xe3\x3b\xd6\xf2\x1a\xca\x96\xfd\x83\x38\x3c\x7e\x21\x3f\x31\xf9\x6a\x6f\xe5\xce\x0d\x0f\x49\x87\x54\xbd\x21\xed\x9a\x0e\x18\x11\x93\x0b\x56\x22\x89\x23\x07\x5a\x70\x20\x0a\xb8\x9b\xd7\x81\xae\x4d\xa8\xf9\xc4\xb9\xee\x51\xa2\xdf\x85\x34\xaf\x37\x98\xa7\x36\xf3\xe9\xfe\xb3\x9e\xea\x8a\x8e\xe5\xed\x2f\x00\x00\xff\xff\xd8\x15\x3a\x01\xe0\x00\x00\x00")

func plusUsedMdBytes() ([]byte, error) {
	return bindataRead(
		_plusUsedMd,
		"plus-used.md",
	)
}

func plusUsedMd() (*asset, error) {
	bytes, err := plusUsedMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plus-used.md", size: 224, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x34, 0xb4, 0x90, 0x2b, 0x1e, 0x21, 0xac, 0xf3, 0x73, 0x84, 0xe9, 0xf7, 0x37, 0x2d, 0x47, 0xb3, 0x79, 0xe8, 0x7c, 0x33, 0x13, 0x5b, 0x93, 0x2e, 0x6c, 0x3d, 0xa5, 0xaf, 0x44, 0xb7, 0xae, 0xc8}}
	return a, nil
}

var _stringsJoinMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8e\x41\x52\xc4\x30\x0c\x04\xef\xbc\x62\x6e\x7b\x01\x7f\x80\x17\xc0\x15\x1e\x60\xef\xae\x9c\xa8\x62\x5b\x2e\x4b\x22\x95\xdf\x53\x4e\xf6\x3c\xd3\xd3\xf3\x81\x2f\x68\x17\x33\x7a\xc2\xd6\x64\x38\xc4\x6f\x83\xe0\xca\x6d\x41\x54\x1b\xdc\x16\x0d\xdf\xc2\x2d\xc2\x04\x77\xe7\x32\x9b\x84\x2b\x0a\xf8\x9d\xd4\x2e\x63\xd3\x77\xdc\xfd\x1c\x40\xe5\x65\x35\x14\xde\x68\x32\x45\x64\x43\xb2\x93\x8a\xb9\x5a\xf8\xe9\x83\x9b\xe5\x88\xec\xed\x61\x2c\x0d\xdc\xd4\x28\x3d\x3f\xc1\x76\x53\x24\xfc\xd1\x38\xd0\x65\xa7\x91\xbd\x60\x4f\x07\x24\x5f\xee\x79\xcb\xfb\xcb\xae\xc8\x43\x2a\xaa\x17\xe3\x5e\x08\x9d\xe9\x41\x1a\xde\xfe\x03\x00\x00\xff\xff\x48\xad\x17\xe3\xd6\x00\x00\x00")

func stringsJoinMdBytes() ([]byte, error) {
	return bindataRead(
		_stringsJoinMd,
		"strings-join.md",
	)
}

func stringsJoinMd() (*asset, error) {
	bytes, err := stringsJoinMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "strings-join.md", size: 214, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfc, 0xff, 0x6a, 0xf1, 0xe7, 0x47, 0xc5, 0xf6, 0x40, 0x60, 0x6e, 0x1a, 0x7b, 0xa4, 0x8f, 0xc9, 0x12, 0x5f, 0xef, 0xac, 0xad, 0xe1, 0xc8, 0x2d, 0x3c, 0xd4, 0x39, 0xdd, 0x16, 0xf, 0x7e, 0xc1}}
	return a, nil
}

var _stubMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x91\xc1\x8e\x14\x31\x0c\x44\xef\x7c\x45\x49\x1c\x60\xa5\xa5\xfb\x3e\x57\x40\x7b\x07\x2e\x68\x85\x58\x4f\xe2\x4e\xa2\x75\xda\x91\xe3\x4c\x33\x7f\x8f\xba\x81\x61\xae\x65\x57\xe9\xb9\xfc\x01\x9f\x74\x7d\xe7\x58\xd4\x12\x3b\x5c\x11\x59\xd8\x19\xb4\x46\x18\x37\xa1\xc0\x20\x11\x78\x66\x74\x1f\x67\x04\xad\x95\x57\xef\xd8\x8a\x67\x18\x93\xdc\xa4\x09\xdf\x75\xa0\x96\x94\x1d\x67\x46\x1f\xd6\xac\x74\x8e\x20\x47\xd6\x0d\xa5\x36\x35\xa7\xd5\x41\x21\x0c\x23\xe7\x47\x64\x96\xb6\x8c\xff\x19\x20\xe3\x1d\xe3\x69\x27\xb9\xb0\x68\x63\xeb\x13\xbe\xe5\xd2\x51\x3a\x08\xc9\x98\x1c\x69\x94\x78\xec\x6d\x56\xbc\xac\x09\x49\x35\xde\x42\x4e\x78\xfe\xbc\x2c\x1c\xbc\x5c\x18\x4f\x7a\xc2\xc7\x3f\x03\xb2\xeb\x8f\xf7\xd9\xbd\xf5\xd3\x3c\x27\x15\x5a\xd3\xa4\x96\xe6\xa8\x61\xe6\x7f\x86\x9f\x49\xa7\xec\x55\xde\x86\x9b\xe9\x61\xc2\x17\xae\x5c\xcf\x6c\x8f\x47\x13\x2f\x49\xa3\x86\x17\xb8\xaa\x60\x2b\x22\x68\xa6\x71\x04\x46\xd4\x30\x0e\x97\x17\x5d\xb1\x98\x56\x5c\x75\x18\x82\x46\xbe\x3b\x72\xb8\x56\xf2\x12\x48\xe4\x3a\xe1\x2b\x33\x9e\x8f\xc8\x1d\xe7\x1e\xf1\xaf\xf4\xb0\x3f\x08\xfc\x8b\x6a\x13\xee\xd0\x05\x5b\x26\x87\xef\xad\x88\xea\x6b\x87\x94\x57\x9e\xde\xfc\x0e\x00\x00\xff\xff\xda\xc3\xad\xf2\xcf\x01\x00\x00")

func stubMdBytes() ([]byte, error) {
	return bindataRead(
		_stubMd,
		"stub.md",
	)
}

func stubMd() (*asset, error) {
	bytes, err := stubMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stub.md", size: 463, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x88, 0xbb, 0xa4, 0x28, 0x23, 0x7d, 0xfc, 0x4f, 0xb4, 0xaa, 0xcb, 0x7d, 0xcd, 0x2d, 0x77, 0x96, 0xc7, 0x9c, 0x5f, 0x72, 0x20, 0x34, 0x1e, 0x52, 0x55, 0x2a, 0x8a, 0x58, 0xc0, 0x5c, 0x2c, 0x31}}
	return a, nil
}

var _useStringPlaceholderMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x90\x41\x4e\xc3\x30\x10\x45\xf7\x9c\xe2\x6f\x22\x40\x2a\x5e\xb0\xec\x01\x80\x03\xf4\x00\x1e\xdc\x71\x33\x52\xe2\xb1\xc6\xd3\x40\x6e\x8f\x6a\x17\xa9\x4b\xcb\xff\xbf\x67\xff\x37\x9c\xa4\x1e\x41\x65\x87\xef\x95\xa1\x19\x1b\x2d\x57\x86\x14\x7c\x2a\x12\x15\x7c\x33\xaa\x49\x71\x3e\xe3\xda\xa4\x5c\xe0\x33\x23\xab\xad\xe4\x48\x33\x19\x25\x67\x43\x9c\xb6\x18\xf0\xa5\x3f\xbc\xb1\x1d\x6e\x19\xe3\xe7\x06\x5a\x9a\x82\x70\x96\x9c\xd9\xb8\xf8\x7f\x31\xab\x81\x29\xcd\xc3\xfa\xd2\x8f\xbf\xb4\xd6\x85\x11\xa7\x16\xfb\x7d\x73\x93\x72\x69\xaf\x01\xa7\x99\x47\xb0\x55\x4e\x92\x25\xdd\x31\x77\xbe\xde\xe0\x58\xd5\x18\x5a\x5d\xb4\xb4\xde\x1f\x19\x1f\x4f\x26\xef\x84\x80\x8f\x07\xd5\xc3\x4f\xe2\x94\xe3\xa1\xd7\x62\x5e\x94\x3c\x8e\x1d\xda\xa1\x6f\xd0\x35\x43\xbe\xf7\x56\x35\x4e\xd2\x44\xcb\x11\x71\x0a\xef\x39\x86\xa7\xbf\x00\x00\x00\xff\xff\x11\x1e\xb5\xad\x4b\x01\x00\x00")

func useStringPlaceholderMdBytes() ([]byte, error) {
	return bindataRead(
		_useStringPlaceholderMd,
		"use-string-placeholder.md",
	)
}

func useStringPlaceholderMd() (*asset, error) {
	bytes, err := useStringPlaceholderMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "use-string-placeholder.md", size: 331, mode: os.FileMode(420), modTime: time.Unix(1542628451, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x62, 0x87, 0x67, 0x67, 0x0, 0xf3, 0x60, 0xfa, 0x61, 0x1b, 0x26, 0x92, 0x70, 0x9b, 0x6e, 0x8b, 0xd8, 0x81, 0x73, 0x77, 0x46, 0x4d, 0x9e, 0x1f, 0x35, 0x12, 0xeb, 0x71, 0x6f, 0x39, 0xa2, 0x28}}
	return a, nil
}

var _wrongCommentFormatMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\x31\x0e\xc2\x30\x0c\x05\xd0\x9d\x53\xfc\x05\x31\xd1\xd3\x70\x81\x34\x71\x89\xa5\xda\x86\xf8\x57\xd0\xdb\x23\x45\x62\x7d\xc3\xbb\xe3\xd1\x05\x35\xcc\xc4\x89\x26\x59\x87\xae\xea\x4f\xb0\x0b\xae\x09\x4d\xbf\x11\xef\x43\x29\x50\x9f\x9a\x2c\xde\xca\x68\xd8\x62\x58\xe1\x32\x87\x94\x4a\x0d\x47\xf8\x3f\x4b\xac\xb2\xc7\x07\x56\x4e\x74\xd9\x5f\x38\xe3\xc0\xa6\x5f\xb0\x6b\x2e\x97\x5f\x00\x00\x00\xff\xff\x63\xfa\x54\x9b\x79\x00\x00\x00")

func wrongCommentFormatMdBytes() ([]byte, error) {
	return bindataRead(
		_wrongCommentFormatMd,
		"wrong-comment-format.md",
	)
}

func wrongCommentFormatMd() (*asset, error) {
	bytes, err := wrongCommentFormatMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wrong-comment-format.md", size: 121, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe0, 0xae, 0x32, 0xff, 0x58, 0x57, 0xf7, 0xbf, 0xb0, 0xef, 0xfd, 0x9e, 0xfa, 0x84, 0xc0, 0xf, 0xb6, 0x8b, 0x1b, 0x84, 0x2a, 0x2e, 0x56, 0x15, 0x4e, 0xaf, 0x8e, 0x4d, 0x18, 0x4b, 0x1f, 0xfe}}
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
	"comment-format.md": commentFormatMd,

	"generalize-name.md": generalizeNameMd,

	"minimal-conditional.md": minimalConditionalMd,

	"missing-comment.md": missingCommentMd,

	"plus-used.md": plusUsedMd,

	"strings-join.md": stringsJoinMd,

	"stub.md": stubMd,

	"use-string-placeholder.md": useStringPlaceholderMd,

	"wrong-comment-format.md": wrongCommentFormatMd,
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
	"comment-format.md":         &bintree{commentFormatMd, map[string]*bintree{}},
	"generalize-name.md":        &bintree{generalizeNameMd, map[string]*bintree{}},
	"minimal-conditional.md":    &bintree{minimalConditionalMd, map[string]*bintree{}},
	"missing-comment.md":        &bintree{missingCommentMd, map[string]*bintree{}},
	"plus-used.md":              &bintree{plusUsedMd, map[string]*bintree{}},
	"strings-join.md":           &bintree{stringsJoinMd, map[string]*bintree{}},
	"stub.md":                   &bintree{stubMd, map[string]*bintree{}},
	"use-string-placeholder.md": &bintree{useStringPlaceholderMd, map[string]*bintree{}},
	"wrong-comment-format.md":   &bintree{wrongCommentFormatMd, map[string]*bintree{}},
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
