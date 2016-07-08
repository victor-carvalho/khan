// Code generated by go-bindata.
// sources:
// migrations/20160608133902_CreateGameTable.sql
// migrations/20160608150958_CreatePlayerTable.sql
// migrations/20160608174439_CreateClanTable.sql
// migrations/20160608182307_CreateMembershipTable.sql
// migrations/20160621161411_CreateHooksTable.sql
// migrations/20160627110742_LoadUUIDModule.sql
// migrations/20160627153918_CreateRetrieveClanIndexes.sql
// migrations/20160627155249_CreatePlayerMembershipAndOwnershipCount.sql
// migrations/20160628181530_CreateClanMembershipCount.sql
// DO NOT EDIT!

package db

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

var _migrations20160608133902_creategametableSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x53\xd1\x6e\xda\x30\x14\x7d\xcf\x57\xdc\x37\x40\x1b\x64\xeb\xd4\x3e\xd0\x69\x1a\x85\x74\x62\x4b\x43\x0b\xc9\x43\x9f\x2c\xe3\x5c\x12\xab\x89\x6d\xd9\x0e\xb4\x9a\xf6\x41\xfb\x8d\x7d\xd9\x9c\x10\x10\x42\xda\x1a\x6d\x79\x89\x7c\xcf\x39\x37\xe7\xfa\xe6\x0c\x87\xf0\x94\x53\xe1\xb9\x77\x6e\xad\x32\x63\xdf\xcf\xb8\xcd\xab\xf5\x88\xc9\xd2\xb7\x52\x6d\x34\x62\x46\x4b\x34\x7e\xcb\xab\xa9\x21\x67\x28\x0c\xa6\x50\x89\x14\x35\xd8\x1c\xe1\x6e\x1e\x43\xb1\x2f\x8f\x0f\xdd\x5c\xb3\xdd\x6e\x37\x92\xca\x55\x65\xa5\x19\x8e\xa4\xce\xfc\x96\x65\xfc\x92\xdb\x61\x7b\xa8\x15\x53\xa9\x5e\x34\xcf\x72\x0b\xbf\x7e\xc2\xc5\xbb\xf7\x57\x10\x4b\x05\xb7\xee\xfb\xf0\xa5\x36\x00\x1f\xd7\x94\x3d\xa1\x48\x3f\xdb\x4d\xc6\x64\x6d\xf0\x93\x57\x0b\xdf\x64\x52\x1a\x84\x44\xd5\x87\xd5\x43\x08\x5c\x80\x41\x66\xb9\x14\xd0\x4b\x54\x0f\xb8\x01\x7c\x46\x56\x59\xe7\x78\x97\xa3\x70\x86\x5d\xa9\xe4\x99\xa6\x0d\xc9\x1d\xa8\x52\x05\xc7\xd4\x9b\x2e\x83\x49\x1c\x40\x3c\xb9\x09\x03\x68\xe6\x86\xbe\x07\xee\xe1\xa9\xeb\xa9\x39\x2d\xe0\x7e\x39\xbf\x9b\x2c\x1f\xe1\x5b\xf0\xf8\xb6\x81\x54\xb5\x76\x63\x10\xc7\xd8\x52\xcd\x72\xaa\xfb\x1f\xae\x06\x10\x2d\x62\x88\x92\x30\xdc\x73\x84\x6b\x75\x84\x2f\x2e\x2f\xcf\xf1\x92\x0b\x52\x62\xb9\x46\x6d\x72\xae\x48\x81\x5b\x2c\xdc\x1c\x16\x33\x77\xbf\x67\x54\xfa\xdc\x99\xea\xba\x36\x38\xb1\x92\x50\xc6\x50\x59\xd2\x4c\xca\xda\xc1\x5f\x57\x31\x8d\xd4\x22\xe1\x62\xcb\x6d\x67\x91\xc6\x52\x6e\xb1\x75\xf9\xaa\x40\x6e\x36\x06\xed\xff\xe8\x94\x96\xa5\xb4\xff\x20\x4c\xb1\x83\xce\xdd\x37\x2b\xa8\x30\x44\xa1\x26\xaa\xa0\x2f\x7f\xa5\xb6\xab\xf9\x13\xe3\x6c\x71\x06\xbe\xae\x16\xd1\xcd\x91\x05\xb3\xe0\x76\x92\x84\x31\xf4\xbe\xff\xe8\x8d\xc7\x0d\x78\x50\x5a\x9a\x52\x4b\x3b\x0b\xf6\x9b\x4b\x09\xb5\xb0\xe6\x99\xb3\x73\x66\xa5\x52\xe9\x39\xde\x60\x0d\x38\x5d\x44\xab\x78\x39\x99\x47\xf1\xc9\xef\x9d\x44\xf3\x87\x24\xe8\x1f\x0b\x03\x6f\x70\x7d\x9a\xc0\x99\xdc\x89\x43\x06\x8f\x01\xac\x8b\x9d\x22\xa8\x65\x51\x38\xb4\x0e\xb9\x37\x5b\x2e\xee\x4f\x43\x78\xed\xfd\x0e\x00\x00\xff\xff\x61\x25\xb7\x2b\xa9\x04\x00\x00")

func migrations20160608133902_creategametableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160608133902_creategametableSql,
		"migrations/20160608133902_CreateGameTable.sql",
	)
}

func migrations20160608133902_creategametableSql() (*asset, error) {
	bytes, err := migrations20160608133902_creategametableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160608133902_CreateGameTable.sql", size: 1193, mode: os.FileMode(420), modTime: time.Unix(1466797802, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160608150958_createplayertableSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x52\xc1\x6e\xd3\x40\x10\xbd\xfb\x2b\xe6\x96\x44\x34\x31\x14\xb5\x87\x14\x21\xdc\x64\x83\x02\xae\xd3\x3a\xf6\xa1\xa7\x68\x63\x4f\xed\x55\xed\xdd\xd5\x7a\x8d\xa9\x10\x1f\xc4\x6f\xf0\x65\xcc\xda\x69\x53\x72\x40\xe4\x12\xcd\xbc\x37\x6f\xdf\xcc\xf3\x74\x0a\x8f\x25\x97\x1e\xfd\x97\xd6\xea\x66\xee\xfb\x85\xb0\x65\xbb\x9f\x65\xaa\xf6\xad\xd2\x0f\x06\xb1\xe0\x35\x36\xfe\x81\xe7\xa8\xa1\xc8\x50\x36\x98\x43\x2b\x73\x34\x60\x4b\x84\x9b\x75\x02\xd5\xd0\x9e\x3f\xab\x91\x58\xd7\x75\x33\xa5\xa9\xab\x5a\x93\xe1\x4c\x99\xc2\x3f\xb0\x1a\xbf\x16\x76\x7a\x28\xdc\xc4\x42\xe9\x27\x23\x8a\xd2\xc2\xef\x5f\x70\xfe\xf6\xdd\x25\x24\x4a\xc3\x8a\xde\x87\xcf\xce\x00\x7c\xd8\xf3\xec\x11\x65\xfe\xc9\x3e\x14\x99\x72\x06\x3f\x7a\x6e\xf0\x4d\xa1\x54\x83\x90\x6a\x57\x6c\xef\x42\x10\x12\x1a\xcc\xac\x50\x12\x46\xa9\x1e\x81\x68\x00\xbf\x63\xd6\x5a\x72\xdc\x95\x28\xc9\x30\xb5\x6a\x51\x18\xde\x93\xa8\xe0\x5a\x57\x02\x73\x6f\x11\xb3\x20\x61\x90\x04\xd7\x21\x03\x5d\xf1\x27\x34\x0d\x8c\x3d\xa0\x9f\xc8\x49\xd5\x08\x5e\xc1\x6d\xbc\xbe\x09\xe2\x7b\xf8\xca\xee\xcf\x7a\x48\xb7\x7b\x5a\x64\x47\x8c\x6f\xdc\x64\x25\x37\xe3\xf3\x8b\x8b\x09\x44\x9b\x04\xa2\x34\x0c\x07\x92\xbb\xe2\x6b\xca\xfb\xcb\x23\x03\x62\xb6\x62\x31\x8b\x16\x6c\xdb\xf3\xe8\xcd\x17\xcd\xc9\x30\x2e\xa9\xfd\x2f\xf9\x1a\x2d\xcf\xb9\xe5\xf0\x65\xbb\x89\xae\x8f\xca\x4b\xb6\x0a\xd2\x30\x81\xd1\x8f\x9f\xa3\xf9\xbc\x07\x87\x81\xcc\x20\xa7\x8b\xec\xb8\x85\xbd\x28\x84\xb4\x27\x82\xad\xce\x4f\xf1\x1e\xeb\xc1\xc5\x26\xda\x26\x71\xb0\x8e\x92\xde\xaf\xc8\x77\xc3\xb1\x68\xbf\x34\x5a\xdf\xa5\x6c\x7c\x58\xf7\xec\x78\x9c\x89\x37\xb9\x7a\x1d\xd8\x52\x75\xf2\x39\xb2\x97\xbc\x5c\xf3\xbf\x12\x33\xaa\xaa\x08\x75\xdf\x84\xb7\x8c\x37\xb7\x7f\x67\x76\xe5\xfd\x09\x00\x00\xff\xff\xa6\xd1\x46\x2b\xda\x02\x00\x00")

func migrations20160608150958_createplayertableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160608150958_createplayertableSql,
		"migrations/20160608150958_CreatePlayerTable.sql",
	)
}

func migrations20160608150958_createplayertableSql() (*asset, error) {
	bytes, err := migrations20160608150958_createplayertableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160608150958_CreatePlayerTable.sql", size: 730, mode: os.FileMode(420), modTime: time.Unix(1466707165, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160608174439_createclantableSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x52\xc1\x72\xd3\x30\x10\xbd\xfb\x2b\xf6\x96\x64\x68\x62\x28\xd3\x1e\x52\x86\x21\x4d\x1c\x26\xe0\x3a\xad\x63\x1f\x7a\xca\x28\xf6\xc6\x16\x55\x24\x8d\x2c\x63\x3a\x0c\x1f\xc4\x6f\xf0\x65\xac\x1c\x37\xc9\x14\xe8\xe0\x8b\x47\xbb\xef\x3d\xad\xde\xbe\xe1\x10\x1e\x4a\x26\x3d\xfa\x97\xd6\xea\x6a\xec\xfb\x05\xb7\x65\xbd\x19\x65\x6a\xe7\x5b\xa5\xb7\x06\xb1\x60\x3b\xac\xfc\x0e\xe7\xa0\x21\xcf\x50\x56\x98\x43\x2d\x73\x34\x60\x4b\x84\x9b\x45\x02\x62\x5f\x1e\x3f\xa9\x91\x58\xd3\x34\x23\xa5\xa9\xaa\x6a\x93\xe1\x48\x99\xc2\xef\x50\x95\xbf\xe3\x76\xd8\x1d\x1c\x63\xaa\xf4\xa3\xe1\x45\x69\xe1\xd7\x4f\x38\x7f\xfd\xe6\x12\x12\xa5\x61\x4e\xf7\xc3\x47\x37\x00\xbc\xdb\xb0\xec\x01\x65\xfe\xc1\x6e\x8b\x4c\xb9\x01\xdf\x7b\x8e\xf8\xaa\x50\xaa\x42\x48\xb5\x3b\xac\xee\x42\xe0\x12\x2a\xcc\x2c\x57\x12\x7a\xa9\xee\x01\xaf\x00\xbf\x61\x56\x5b\x9a\xb8\x29\x51\xd2\xc0\x54\xda\xf1\xc2\xb0\x16\x44\x07\xa6\xb5\xe0\x98\x7b\xd3\x38\x98\x24\x01\x24\x93\xeb\x30\x80\x4c\x30\x59\x41\xdf\x03\xfa\x78\x4e\x9a\x86\x33\x01\xb7\xf1\xe2\x66\x12\xdf\xc3\xe7\xe0\xfe\xac\x6d\xe9\x7a\x43\xcf\x58\x13\xe2\x2b\x33\x59\xc9\x4c\xff\xfc\xe2\x62\x00\xd1\x32\x81\x28\x0d\xc3\x3d\xc8\x79\x78\x0a\x79\x7b\x79\x44\x40\x1c\xcc\x83\x38\x88\xa6\xc1\xaa\xc5\xd1\x9d\x07\xcd\xc1\x9e\x2e\xa9\xfc\x92\xfc\x0e\x2d\xcb\x99\x65\xf0\x69\xb5\x8c\xae\x8f\xca\xb3\x60\x3e\x49\xc3\x04\x7a\xdf\x7f\xf4\xc6\xe3\xb6\xb9\x27\x30\x21\x54\xb3\x6e\x9f\x9d\xed\x5d\xd8\x28\x25\x90\xc9\x3f\xb9\xd6\xd4\xd8\x91\x6a\xab\xd6\x5f\x14\x7f\x01\xbc\x65\xa2\xea\xd0\x99\x41\x46\x96\xaf\x99\x85\x0d\x2f\xb8\xb4\xcf\x66\xae\x75\xfe\xbc\x7f\xe8\xe5\x28\xf0\x5f\x3d\xd5\x48\x34\xce\x4b\x2a\x63\x41\xf9\xfb\x9b\x8f\x5a\xb0\x47\x34\xe4\x64\x6b\x61\xcb\x9b\x2e\xa3\x55\x12\x4f\x16\x51\xd2\xba\xcc\xf3\xb5\x5b\x30\xe9\xa4\xd1\xe2\x2e\x0d\xfa\xdd\x8a\xce\x8e\x0b\x1d\x78\x83\xab\xd3\x88\xcd\xe8\xea\xa7\x90\x1d\x12\xe6\x8a\xff\x95\x31\xa3\x84\xa0\xae\x4b\xb1\x37\x8b\x97\xb7\xa7\x29\xbb\xf2\x7e\x07\x00\x00\xff\xff\xb8\x66\x58\xaf\x8a\x03\x00\x00")

func migrations20160608174439_createclantableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160608174439_createclantableSql,
		"migrations/20160608174439_CreateClanTable.sql",
	)
}

func migrations20160608174439_createclantableSql() (*asset, error) {
	bytes, err := migrations20160608174439_createclantableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160608174439_CreateClanTable.sql", size: 906, mode: os.FileMode(420), modTime: time.Unix(1466707165, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160608182307_createmembershiptableSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\x4b\x6e\xdb\x30\x10\xdd\xeb\x14\xb3\x4b\x8c\x26\x56\x3f\x40\x16\x49\x51\xd4\xb5\xe5\xc2\xa8\x22\x27\xb2\xb4\xc8\xca\xa0\xa4\xb1\x44\x84\x22\x59\x92\xb2\xeb\x23\xf5\x1a\x3d\x59\x87\xfe\x06\x4e\x80\x1a\xf5\xc6\xe0\xbc\xf7\x66\xde\x7c\x74\x7d\x0d\xcf\x0d\x93\x01\xfd\x37\xce\x69\x7b\x1b\x86\x35\x77\x4d\x57\xf4\x4b\xd5\x86\x4e\xe9\x85\x41\xac\x59\x8b\x36\xdc\xf1\x3c\x35\xe6\x25\x4a\x8b\x15\x74\xb2\x42\x03\xae\x41\xb8\x9f\x64\x20\xb6\xe1\xdb\x7d\x36\x4a\xb6\x5a\xad\xfa\x4a\x53\x54\x75\xa6\xc4\xbe\x32\x75\xb8\x63\xd9\xb0\xe5\xee\x7a\xf7\xf0\x8a\xa1\xd2\x6b\xc3\xeb\xc6\xc1\x9f\xdf\xf0\xf1\xfd\x87\x1b\xc8\x94\x86\x31\xd5\x87\xef\xde\x00\x7c\x2e\x58\xf9\x8c\xb2\xfa\xea\x16\x75\xa9\xbc\xc1\x2f\x81\x17\xbe\xab\x95\xb2\x08\xb9\xf6\x8f\xd9\x63\x0c\x5c\x82\xc5\xd2\x71\x25\xe1\x22\xd7\x17\xc0\x2d\xe0\x2f\x2c\x3b\x47\x8e\x57\x0d\x4a\x32\x4c\xa1\x96\xd7\x86\x6d\x48\xf4\x60\x5a\x0b\x8e\x55\x30\x4c\xa3\x41\x16\x41\x36\xf8\x16\x47\xd0\x62\x5b\xa0\xb1\x0d\xd7\x16\x2e\x03\xa0\x1f\xaf\x28\xb3\xe1\x4c\xc0\x43\x3a\xb9\x1f\xa4\x4f\xf0\x23\x7a\xba\xda\x40\x7e\x48\x73\xc2\x97\xcc\x94\x0d\x33\x97\x9f\x6e\x7a\x90\x4c\x33\x48\xf2\x38\x86\x34\x1a\x47\x69\x94\x0c\xa3\xd9\x86\x47\xe9\x74\x57\x50\xef\x24\xe8\x6d\xe5\xa5\x60\xd2\xcb\xb9\x74\x58\xd3\x4c\xdf\x92\x7a\x0e\x49\x0f\x1a\x2d\xd8\x1a\xcd\xbf\x54\x5b\xd6\x4b\xdd\xb1\xaf\xb9\xc0\x25\x8a\x37\x3d\x6f\xa9\x34\x17\xa3\x96\x34\xb7\x42\x29\x81\x4c\x1e\x2b\x8c\xa2\xf1\x20\x8f\x33\x58\x30\x61\x71\x4b\xae\x50\xf2\x33\xa9\x05\x93\xf2\x4c\xaa\xc1\x9f\x1d\x5a\xa7\xfe\xa3\xd1\xd2\x20\xa3\xa5\xcf\x99\x83\x82\xd7\xa4\x3d\xe9\xae\xd3\xd5\x29\x7e\xc0\x2a\x14\xe8\xb1\x62\x7d\x2c\xfa\x0a\x3c\x15\x6e\xc0\xe1\x34\x99\x65\xe9\x60\x92\x64\x3b\x4b\xbc\x9a\xfb\xd5\x91\xfd\x3c\x99\x3c\xe6\xd1\xe5\x61\x71\x57\xfb\xbd\xf7\x82\xde\xdd\xcb\x63\x1e\xa9\x95\xdc\x9f\xf3\xe1\x96\x7d\xf0\xac\x6b\x36\x4a\x08\x3f\x5d\xfa\x5e\x82\x51\x3a\x7d\x78\x7d\xcf\x77\xc1\xdf\x00\x00\x00\xff\xff\xf2\xac\x4f\xcf\xfa\x03\x00\x00")

func migrations20160608182307_createmembershiptableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160608182307_createmembershiptableSql,
		"migrations/20160608182307_CreateMembershipTable.sql",
	)
}

func migrations20160608182307_createmembershiptableSql() (*asset, error) {
	bytes, err := migrations20160608182307_createmembershiptableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160608182307_CreateMembershipTable.sql", size: 1018, mode: os.FileMode(420), modTime: time.Unix(1466787920, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160621161411_createhookstableSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x51\xcb\x6e\xdb\x30\x10\xbc\xeb\x2b\xf6\x16\x1b\x8d\xad\x3e\x80\x1c\x92\xa2\xa8\xeb\x30\x85\x51\x45\x4e\x64\xe9\x90\x93\x41\x53\x1b\x8a\xb0\x4c\x12\x24\x15\x25\x9f\xd4\xdf\xe8\x97\x75\x29\x3b\xa9\x6b\xf4\x50\x5d\x84\xdd\x99\x59\xce\xee\x4c\x26\xb0\x6d\xb8\x4e\xe8\xdf\x84\x60\xfd\x65\x9a\x4a\x15\x9a\x6e\x33\x15\x66\x97\x06\x63\x1f\x1d\xa2\xe4\x3b\xf4\xe9\x81\x17\xa9\x99\x12\xa8\x3d\xd6\xd0\xe9\x1a\x1d\x84\x06\xe1\x76\x51\x42\xbb\x6f\x5f\xbe\x4e\xa3\x61\x7d\xdf\x4f\x8d\xa5\xae\xe9\x9c\xc0\xa9\x71\x32\x3d\xb0\x7c\xba\x53\x61\x72\x28\xa2\x62\x6e\xec\x8b\x53\xb2\x09\xf0\xeb\x27\x7c\x7c\xff\xe1\x02\x4a\x63\xe1\x86\xde\x87\xef\xd1\x00\x7c\xde\x70\xb1\x45\x5d\x7f\x0d\x8f\x52\x98\x68\xf0\x4b\x12\x85\xef\xa4\x31\x1e\xa1\xb2\xb1\x58\xdd\x67\xa0\x34\x78\x14\x41\x19\x0d\x67\x95\x3d\x03\xe5\x01\x9f\x51\x74\x81\x1c\xf7\x0d\x6a\x32\x4c\xad\x9d\x92\x8e\x0f\x24\x2a\xb8\xb5\xad\xc2\x3a\x99\x17\x6c\x56\x32\x28\x67\xdf\x32\x06\x8d\x31\x5b\x0f\xa3\x04\xe8\x53\x35\xcd\x74\x8a\xb7\x70\x57\x2c\x6e\x67\xc5\x03\xfc\x60\x0f\xe7\x03\x14\xcf\xb3\x26\xfc\x89\x3b\xd1\x70\x37\xfa\x74\x31\x86\x7c\x59\x42\x5e\x65\x19\x14\xec\x86\x15\x2c\x9f\xb3\xd5\xc0\xa3\x71\xb6\xdb\xd0\xd6\x24\x18\xef\xe5\x6f\xf5\x3f\x07\xec\x39\xf8\x84\x3a\xac\xc3\x8b\x45\x5a\x2e\xa0\xa4\xa3\xff\x4d\xe8\x5c\x0b\x01\x9f\xc3\x49\x5b\x38\xe4\xb4\xf5\x9a\x07\xd8\x28\x49\xd2\x53\x99\xad\x4f\xf1\x01\x1b\xc0\xf9\x32\x5f\x95\xc5\x6c\x91\x97\xc3\x25\x54\xbd\xde\x5b\x25\xa7\x55\xbe\xb8\xaf\xd8\xe8\xb0\xf9\xf9\x9f\x1d\xc6\xc9\xf8\xea\x38\x94\x6b\xd3\xeb\xd7\x58\xde\x32\x89\xcd\xff\x4a\xc5\x99\xb6\x25\x34\xe6\x9e\x5c\x17\xcb\xbb\xe3\x5c\xae\x92\xdf\x01\x00\x00\xff\xff\xe0\xb0\xeb\xc6\xbc\x02\x00\x00")

func migrations20160621161411_createhookstableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160621161411_createhookstableSql,
		"migrations/20160621161411_CreateHooksTable.sql",
	)
}

func migrations20160621161411_createhookstableSql() (*asset, error) {
	bytes, err := migrations20160621161411_createhookstableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160621161411_CreateHooksTable.sql", size: 700, mode: os.FileMode(420), modTime: time.Unix(1466601889, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160627110742_loaduuidmoduleSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\xd2\xd5\x55\xd0\x4e\xcf\xcf\x2f\x4e\x55\x08\x2d\xe0\x72\x0e\x72\x75\x0c\x71\x55\x70\x8d\x08\x71\xf5\x0b\xf6\xf4\xf7\x53\xf0\x74\x53\xf0\xf3\x0f\x01\x0a\x78\x06\x87\x04\x2b\x28\x95\x96\x66\xa6\xe8\xe6\x17\x17\x17\x28\x59\x73\x71\x21\x69\x75\xc9\x2f\xcf\xe3\x02\x04\x00\x00\xff\xff\xd3\x7c\xf3\xf6\x4b\x00\x00\x00")

func migrations20160627110742_loaduuidmoduleSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160627110742_loaduuidmoduleSql,
		"migrations/20160627110742_LoadUUIDModule.sql",
	)
}

func migrations20160627110742_loaduuidmoduleSql() (*asset, error) {
	bytes, err := migrations20160627110742_loaduuidmoduleSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160627110742_LoadUUIDModule.sql", size: 75, mode: os.FileMode(420), modTime: time.Unix(1467060502, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160627153918_createretrieveclanindexesSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\xc1\x4e\xc3\x30\x10\x44\xef\xfe\x8a\xbd\x15\x44\xfb\x05\x39\x21\x12\xa4\x48\xa8\x85\xa6\x95\x7a\xab\x9c\x64\x95\xac\x70\xd6\x96\x1d\x28\x9f\x8f\x8d\x5b\x54\x51\xda\xf8\xb8\x2f\x33\x93\xf5\xac\x58\x2c\xe0\xa1\xd3\xda\x21\x6c\x4d\x18\xaa\xb7\x17\x20\x06\x87\xcd\x48\x9a\x61\xb6\x35\x33\x20\x07\xf8\x85\xcd\xc7\x88\x2d\x1c\x7a\x64\x18\x7b\x8f\x06\xea\xac\xfc\x11\xf9\x41\x1a\xa3\x08\x5b\xf1\xb4\x2e\x1e\x37\x05\x94\xcb\xbc\xd8\xc1\x80\x43\x8d\xd6\xf5\x64\xdc\xde\x0b\xac\xfe\xf4\x01\xab\xe5\x39\xbf\x3b\xf1\xfb\xec\xba\xb7\x96\xcc\x97\xce\x48\x6f\xf9\x5a\x64\xba\xf4\x45\x7a\xcb\x67\x90\x5b\xe2\xee\xda\xaa\x73\x88\xbf\x9e\xc3\x74\x54\xa3\x24\xff\xcd\x09\x6c\x4f\xc1\x26\xce\xea\xcf\xf5\x81\x4f\x07\xf8\x6d\x3f\xc0\xa4\xfe\xad\x56\xca\x7f\xad\x65\xf3\x2e\xf2\xf5\xea\xf5\xb8\x4a\xf9\x0c\xc5\xae\xac\x36\xd5\xbf\xb7\xc8\xa6\xa5\xf1\xa5\x09\xc2\x58\x45\x82\xf0\x58\x6e\x82\x32\xf4\x94\x89\xef\x00\x00\x00\xff\xff\xad\x5b\xe0\x06\xa2\x02\x00\x00")

func migrations20160627153918_createretrieveclanindexesSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160627153918_createretrieveclanindexesSql,
		"migrations/20160627153918_CreateRetrieveClanIndexes.sql",
	)
}

func migrations20160627153918_createretrieveclanindexesSql() (*asset, error) {
	bytes, err := migrations20160627153918_createretrieveclanindexesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160627153918_CreateRetrieveClanIndexes.sql", size: 674, mode: os.FileMode(420), modTime: time.Unix(1467060534, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160627155249_createplayermembershipandownershipcountSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x8f\x41\x4f\xc4\x20\x10\x85\xef\xfd\x15\xef\xb6\x07\xb3\x89\xf7\x3d\xa1\xd4\x13\xb6\xba\xc2\xd9\xb0\xec\xa4\x25\x52\x20\xc0\x66\xf5\xdf\x4b\x4d\x34\x8d\xa9\xb1\xc7\xf7\x66\xde\x9b\xf9\x9a\xfd\x1e\x37\x43\x08\x99\xa0\xe2\x2c\x5e\x9e\x05\xac\x47\x26\x53\x6c\xf0\xd8\xa9\xb8\x83\xcd\xa0\x77\x32\x97\x42\x67\x5c\x47\xf2\x28\x63\xb5\x26\x3b\x24\xfd\xb5\x54\x85\x8e\xd1\x59\x3a\x37\x4c\xc8\xf6\x08\xc9\xee\x44\x8b\xe8\xf4\x07\xa5\x0c\xc6\x39\xee\x7b\xa1\x1e\x3b\x4c\x34\x9d\xaa\x35\xda\xf8\x6a\xc2\xc5\x97\x7a\xaa\xd0\x40\x09\x5d\x2f\xd1\x29\x21\xc0\xdb\x07\xa6\x84\xc4\xed\xe1\xbf\xae\x70\xf5\x9b\xab\x16\x98\xbc\xe6\xbe\x41\x7f\x28\x67\x73\x13\x67\x0a\xce\xd5\xe9\x49\x9b\xb7\xd5\xff\xf8\xb1\x7f\xfa\x0b\x76\x9d\x68\x99\xf8\x85\x74\x68\x3e\x03\x00\x00\xff\xff\x47\x2a\x44\xfc\x9f\x01\x00\x00")

func migrations20160627155249_createplayermembershipandownershipcountSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160627155249_createplayermembershipandownershipcountSql,
		"migrations/20160627155249_CreatePlayerMembershipAndOwnershipCount.sql",
	)
}

func migrations20160627155249_createplayermembershipandownershipcountSql() (*asset, error) {
	bytes, err := migrations20160627155249_createplayermembershipandownershipcountSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160627155249_CreatePlayerMembershipAndOwnershipCount.sql", size: 415, mode: os.FileMode(420), modTime: time.Unix(1467060707, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations20160628181530_createclanmembershipcountSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\xce\x31\x4f\xc5\x20\x14\x05\xe0\x9d\x5f\x71\xb6\x37\x98\x37\x38\xbf\x09\xa5\x4e\xd8\xea\x13\x66\x43\xe9\x4d\x4b\xa4\x40\x80\xa6\xfe\x7c\xa9\x89\xc6\x41\x13\xc7\x73\xee\x4d\xce\xc7\xce\x67\xdc\xcc\x31\x16\x82\x4e\x47\x78\x79\x96\x70\x01\x85\x6c\x75\x31\xe0\xa4\xd3\x09\xae\x80\xde\xc9\x6e\x95\x26\xec\x0b\x05\xd4\xa5\x55\xab\x9b\xb3\xf9\x7c\x6a\xc1\xa4\xe4\x1d\x4d\x8c\x4b\xd5\x5d\xa1\xf8\x9d\xec\x60\xbd\x09\x05\x5c\x08\xdc\x0f\x52\x3f\xf6\x58\x69\x1d\x29\x97\xc5\xa5\x57\x1b\xb7\x50\xdb\x50\xa5\x99\x32\xfa\x41\xa1\xd7\x52\x42\x74\x0f\x5c\x4b\x85\xdb\x0b\xfb\x21\x13\x71\x0f\x5f\xb6\x6f\xd8\x51\xfe\x8b\x96\xa3\xf7\xed\x3a\x1a\xfb\xf6\x0b\x4f\x5c\x87\xa7\xbf\x7c\x17\xf6\x11\x00\x00\xff\xff\x41\x07\xd0\x68\x1f\x01\x00\x00")

func migrations20160628181530_createclanmembershipcountSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations20160628181530_createclanmembershipcountSql,
		"migrations/20160628181530_CreateClanMembershipCount.sql",
	)
}

func migrations20160628181530_createclanmembershipcountSql() (*asset, error) {
	bytes, err := migrations20160628181530_createclanmembershipcountSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/20160628181530_CreateClanMembershipCount.sql", size: 287, mode: os.FileMode(420), modTime: time.Unix(1467154167, 0)}
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
	"migrations/20160608133902_CreateGameTable.sql": migrations20160608133902_creategametableSql,
	"migrations/20160608150958_CreatePlayerTable.sql": migrations20160608150958_createplayertableSql,
	"migrations/20160608174439_CreateClanTable.sql": migrations20160608174439_createclantableSql,
	"migrations/20160608182307_CreateMembershipTable.sql": migrations20160608182307_createmembershiptableSql,
	"migrations/20160621161411_CreateHooksTable.sql": migrations20160621161411_createhookstableSql,
	"migrations/20160627110742_LoadUUIDModule.sql": migrations20160627110742_loaduuidmoduleSql,
	"migrations/20160627153918_CreateRetrieveClanIndexes.sql": migrations20160627153918_createretrieveclanindexesSql,
	"migrations/20160627155249_CreatePlayerMembershipAndOwnershipCount.sql": migrations20160627155249_createplayermembershipandownershipcountSql,
	"migrations/20160628181530_CreateClanMembershipCount.sql": migrations20160628181530_createclanmembershipcountSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"20160608133902_CreateGameTable.sql": &bintree{migrations20160608133902_creategametableSql, map[string]*bintree{}},
		"20160608150958_CreatePlayerTable.sql": &bintree{migrations20160608150958_createplayertableSql, map[string]*bintree{}},
		"20160608174439_CreateClanTable.sql": &bintree{migrations20160608174439_createclantableSql, map[string]*bintree{}},
		"20160608182307_CreateMembershipTable.sql": &bintree{migrations20160608182307_createmembershiptableSql, map[string]*bintree{}},
		"20160621161411_CreateHooksTable.sql": &bintree{migrations20160621161411_createhookstableSql, map[string]*bintree{}},
		"20160627110742_LoadUUIDModule.sql": &bintree{migrations20160627110742_loaduuidmoduleSql, map[string]*bintree{}},
		"20160627153918_CreateRetrieveClanIndexes.sql": &bintree{migrations20160627153918_createretrieveclanindexesSql, map[string]*bintree{}},
		"20160627155249_CreatePlayerMembershipAndOwnershipCount.sql": &bintree{migrations20160627155249_createplayermembershipandownershipcountSql, map[string]*bintree{}},
		"20160628181530_CreateClanMembershipCount.sql": &bintree{migrations20160628181530_createclanmembershipcountSql, map[string]*bintree{}},
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

