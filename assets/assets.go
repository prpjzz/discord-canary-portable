// Code generated by go-bindata. DO NOT EDIT.
// sources:
// DiscordCanary.lnk (1.923kB)
// pinned_update.json (16.178kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var _discordcanaryLnk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x95\x4d\x68\x2b\x55\x14\xc7\x7f\xf1\x3d\x79\xbe\x5d\xc4\xe0\x17\xef\xd9\x79\x48\xe4\x81\xcc\xe4\xa3\x79\x8f\x24\xa0\xe4\x35\x1f\xa6\x74\x30\x43\x3e\x68\x94\x11\x3b\x26\x03\x19\x69\x92\xe9\x4c\xd2\x0f\xbb\x10\x5d\xb8\x10\x0b\x82\x2e\x5c\x58\x1a\x44\xd1\x85\x2e\xc4\xba\x69\x45\xeb\xc2\xa5\xd2\x45\xab\xa5\x2b\x2b\xe8\x42\xa1\x5d\x08\xba\x51\x90\xb9\x4d\x4a\xd3\x46\xdd\xb8\x51\xfc\x5f\x72\x6f\xee\x3d\xe7\x9e\xff\xff\x9c\x19\xce\xa8\x80\x2f\x70\x1b\x1e\xb6\xc4\x4c\xee\x28\x05\x12\xf0\xfb\x03\x4f\x45\x57\x0b\xfb\xbe\xf5\x5f\x9e\x16\x2b\xfe\x47\x5f\xbd\x7a\x7d\xdf\x77\x6f\xf6\x1e\xe1\xe8\x63\x18\xbf\x12\x60\x4c\xfb\xb6\xb0\x2d\xfd\x98\xb4\xfc\x6f\xed\xdd\xc1\xc3\xe1\xf0\xea\x7d\x84\xd2\x49\x9d\xf3\x68\x10\x11\x6b\xa2\xba\xd5\xf5\x4b\x19\xcb\xad\xb5\x9d\x7a\xda\x68\x19\xce\x12\x2a\x97\xb9\xc8\xd1\xa7\x89\xea\xc7\x5d\xcf\xae\x00\x5f\xa6\x3c\xc2\xd4\x99\x28\xbd\x85\x4b\x64\xb0\x70\xa9\xd1\xc6\xa1\x4e\x1a\x83\x16\x06\x0e\x4b\xc0\x15\x66\x88\xe2\x49\x7e\x69\x9a\xb4\x44\xc5\xae\x1b\x1d\x53\x31\x17\x4d\xc8\x9d\xb0\x6c\x76\xbd\x9f\xc7\xf2\xd3\x9e\x57\x8c\xc0\x08\xbd\x15\x6c\xea\x18\x74\x30\x51\x30\x59\xc4\x04\xee\x17\x16\x8f\xe7\xb8\x20\xde\x1a\xef\xdf\x28\xf7\xf7\x17\x00\xed\xb3\x17\x9f\xf4\x03\x2d\xb3\x2c\xb9\x0b\xed\x7a\xcb\x9a\x26\x9d\xd4\x87\xd2\xd6\x4f\x8b\xbb\x76\x26\x2d\x09\x19\x89\x06\x1d\x3a\xd8\xb8\x24\x09\x11\xa2\x3e\xe4\x63\x60\x63\xa3\x88\x7d\x93\x10\xd7\x50\x50\xd0\x87\xe6\x3f\x2f\x96\x3e\x32\x45\x8f\x57\xc6\xc6\xa1\x4d\x0d\x13\x17\x97\x12\x1d\x71\xa7\x83\xf4\x17\xf1\x06\x11\x42\xa4\x49\x8a\xe8\x2e\x26\x0e\x2e\x3a\x35\x1c\x0c\x9e\x13\xac\xb7\x84\xea\x8c\xe0\x35\xd0\x51\x05\x93\xc1\xec\x39\xb5\x1a\x65\x26\xd0\x4f\xf2\xb4\x84\x25\x70\x01\x2e\x41\x2f\x58\x29\x65\x8b\x5a\xb1\x90\x9b\x54\xb3\x41\xfd\x96\x6d\x67\x8c\x8e\xa1\xab\xed\x9a\x31\x3b\xa8\xb3\x56\x9e\xd0\x0d\xdb\x56\xac\x5a\x7b\xc4\x13\xfe\x0f\x20\x48\x85\x12\x59\x8a\x68\x14\x29\x90\x63\x12\x95\x2c\xc1\x7f\xa0\xce\xff\xe3\xdf\x83\xc7\x7c\x70\x19\x7a\x2b\x40\xa4\xa4\x95\xbe\x7b\xb9\x9a\xfb\x44\x8d\xa7\x37\x7f\xbb\xeb\xb5\x87\xde\x68\x7e\x35\x07\x5c\xf4\x1c\xc7\x80\x10\x50\x42\x26\x82\xcc\x0d\x64\xa2\xe2\xdf\x38\xe3\x84\xb9\xc9\x38\x71\x22\x44\x89\x11\x45\x26\x46\x58\x8c\xb8\xb0\xc4\xb8\x29\x7c\x23\xc4\x88\x13\x25\x4e\x98\xa8\x18\x09\x71\xea\x79\x46\x4e\x44\x75\xfb\x62\x2a\xd7\xd5\xb5\xa5\xb5\xc4\xd4\x7b\xdb\x07\x3b\xf2\xc1\xee\xcf\x4f\x00\xb7\x0f\xc4\x3c\x08\xfd\xfe\xa5\xe0\x32\x47\x17\x0b\x07\x07\x93\x59\x94\x91\xef\xe9\xe8\xd3\x01\x12\x7d\xd2\x0f\xef\x6e\x66\x3e\x58\xb1\xf3\xef\xe6\x53\xca\xdb\x8f\x2c\xbe\x72\x55\x7c\x82\x80\xbc\xd7\xa1\x97\xa7\xf2\xc7\xee\xda\xd8\xe9\x32\xce\x1c\x77\xef\x5e\xb5\xbf\x5f\xb0\x5a\x72\x3b\x32\xdf\xb8\x31\x6f\x3c\xfb\x4c\xa3\x49\xde\x7a\xff\xfb\xec\xee\x95\xec\x9b\xcf\x7f\x71\xf8\x43\xa9\x37\xdf\xfb\xfc\xa3\xaf\x5f\xdf\x38\xbc\x73\xfd\xf1\x9d\x17\xbe\x59\x7e\x67\xe3\xef\xec\x5e\xcc\x3f\x02\x00\x00\xff\xff\x7c\xd1\xbe\x12\x83\x07\x00\x00")

func discordcanaryLnkBytes() ([]byte, error) {
	return bindataRead(
		_discordcanaryLnk,
		"DiscordCanary.lnk",
	)
}

func discordcanaryLnk() (*asset, error) {
	bytes, err := discordcanaryLnkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "DiscordCanary.lnk", size: 1923, mode: os.FileMode(0666), modTime: time.Unix(1747043000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbb, 0xc, 0x3a, 0x83, 0x1a, 0x8f, 0x80, 0xb, 0x96, 0x4d, 0xb4, 0x62, 0x81, 0xaf, 0x3c, 0xdb, 0xc4, 0xdf, 0xab, 0x6a, 0xc6, 0x10, 0xe1, 0xff, 0xd, 0x2a, 0x13, 0xd6, 0xf1, 0x2b, 0x7a, 0xa2}}
	return a, nil
}

var _pinned_updateJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9b\xdb\x6e\x64\xc7\x6e\x86\xef\xf7\x53\x08\xba\x4e\x2c\x92\xc5\x22\x8b\x7e\x95\x8d\x0d\x83\xa7\xf2\x08\x96\x47\x8a\xa4\x71\xe2\x04\x7e\xf7\xa0\xc7\x73\xd0\x4c\xb7\x77\x9c\x40\xd3\x80\x15\xf5\x85\xa0\x5e\xbd\x84\x5e\xfa\xbf\xfa\x16\xeb\xb4\xfe\xeb\x6f\x17\x17\x97\x3f\xdf\xd6\xbb\x9b\x7e\xb8\xfc\xfe\xe2\xf0\xf6\xf0\xba\xac\xeb\x87\xbc\xbd\xaf\x1f\xf2\xe6\xf6\x5d\x3d\xfc\xfa\x36\x9f\x7c\xf8\xfe\x84\xfd\xee\xe6\xe6\xab\x63\xef\x8f\xbf\xb9\x7d\x78\xfc\xe1\x97\xbe\x7f\xb8\xbe\x7d\x7b\xf9\xfd\xc5\xdf\xbf\xfa\xfc\xf0\xc2\x7f\x39\x71\x10\x4e\x1d\x1c\x34\xbe\x3a\xfa\x8f\xaf\x4f\xfb\x70\xf1\x4f\xbe\x72\x1e\x9d\x72\xe7\xf9\x93\xff\xd8\x3f\x3c\xbc\x71\x9a\x72\xf9\xfd\xc5\x25\x37\x15\xe3\xca\x8c\x76\xaf\x88\x1c\xe1\xd4\x90\x95\xd8\xa8\xe1\x0b\x17\x06\x01\x0e\x45\xea\xe2\x51\x25\xda\x5b\xd7\xce\x1a\xa5\x1a\x97\x47\xdf\xf1\xee\xfe\x10\xc7\xe5\x9b\xc7\xc7\xbb\x87\xef\xaf\xae\xea\xe6\x5f\xd3\xdf\xfa\xfd\xaf\xdf\x7d\x48\xd2\xef\xee\xbe\x7b\xdb\x8f\x57\x75\xfd\xf0\x78\x7f\x7b\xe5\x77\x77\x57\xbf\x9f\x70\xf5\xef\xd7\x6f\xaf\xfe\x63\xc9\x15\x7e\x07\xdf\x0d\x1a\x57\x47\xd1\x5f\xcd\xab\x43\xda\xdf\xfd\xfe\xa7\x97\x4f\xbe\xf9\xb7\xa7\x97\x71\x59\x7d\xf3\xe8\x0f\x27\x42\xff\x1a\xd2\xc5\x9f\x04\x75\xf1\x47\xb0\x2e\xfe\x08\xd8\xc5\x49\x68\x17\xa7\xc0\x5d\x9c\x84\xc7\x27\x4f\x3b\x06\x58\x46\xae\x22\xdb\xb5\x71\xf8\xdc\xd6\x29\x2d\xdd\x40\x12\xba\x42\x87\x33\x22\xa3\x24\x0a\xf4\x70\x48\x9f\x6b\x97\xea\x76\x0d\x2b\xc3\x23\x80\x17\x67\x80\x78\x7f\xfb\xf3\xa7\x13\xf8\xf2\xab\x0b\xf8\xed\xc9\xfb\x7f\x7c\xf8\xfd\x13\xdf\x4f\x3e\x3e\xf4\x4f\xf7\xd7\x8f\x2f\x4b\x46\xc7\x11\xa6\x4e\xb0\x16\x4f\x2a\x02\x0a\x1e\x2e\x34\xe6\xd0\x46\xe3\x44\x27\xa4\x41\x6b\xef\x5d\x1e\x95\x40\x93\x90\x0a\x88\x5d\x54\xf3\x1c\x32\xfe\x9e\xfb\xab\x89\x27\x4c\x64\x38\xc0\xe2\xb1\xb8\x2b\x96\x07\xec\xe5\xc9\xd1\xec\xc4\x48\x3e\xc3\x2d\x10\x5c\xba\xdc\x8a\x72\x00\x41\x0c\x01\xc3\xa5\x83\x42\xce\x65\xe2\x67\x82\xcf\xa0\x61\xdf\xdf\x1c\x92\x78\x59\x1e\x2e\x2a\x75\x91\xa2\x6d\x3c\xcc\x3d\x57\xf7\x0c\x1c\x1c\x12\x3e\x28\x66\x96\x57\xcd\x5e\xb5\x33\x73\x5b\xf6\xd4\x06\x9f\xb6\xd7\xa0\xc0\x63\x92\xdf\x80\xe2\x87\xe0\x5f\x45\x3c\xd1\xa7\x29\x9e\xcb\x67\xc9\xde\x8b\xc1\x63\x12\x1d\x7e\xe0\x6e\xee\x9a\x93\x27\x87\x6c\x6f\x41\x07\xef\x69\x0c\x90\x39\xc5\xd5\x46\x34\x4e\x39\xee\xd3\x5c\x7c\x73\x84\xcf\x60\x62\xf5\xc3\x4f\x8f\xb7\x77\x3f\xe4\xed\x7d\xbf\x2c\x1d\xdd\x74\x0d\x53\x84\xc0\x32\x84\x9a\x04\x99\x63\x6b\xf8\x18\xa3\x29\x43\xf7\xc4\xca\x02\x30\x9f\xe6\x20\x6d\x90\xca\xad\xd3\x0c\xe6\xb0\x73\xe8\xf8\x34\xfd\x57\x27\x8f\x19\xc2\x26\x64\x92\x05\x3c\x17\x32\x7b\xcc\xe9\xb9\xd3\x16\xe4\xfb\x6a\xb9\x8c\x92\x56\x21\xa8\x2c\xe6\x1a\x3d\xaa\x23\xcc\x02\x46\xaa\x9d\xe8\xda\x5c\x9c\x87\xe3\x33\x88\xf9\x9f\x0f\x8f\xf5\xb2\x84\x6c\xdb\x1c\xd3\xb9\xbb\x61\x0d\x18\x53\x23\x39\xa7\xe3\xa1\x53\xb3\x16\xec\x49\x33\xa8\xc9\xbd\x78\x0e\x59\x38\x78\x05\xe9\xf0\x42\x4c\x60\x3e\x87\x90\x87\xd4\x5f\x45\x3c\xc1\x8e\x65\x90\x4d\x08\x8e\x31\x7d\xec\x05\x5b\xe7\x5e\x5c\x0b\xa4\x6a\xe0\x42\x26\xeb\x9a\x6e\x6a\x8e\x98\xb6\xaa\x2b\x79\x45\xd9\x46\x9c\xe3\x5c\x22\x7e\xe4\xf7\x0c\x02\xde\xfe\xd2\xf7\x37\xfe\x2b\xbd\x2c\x09\xb1\x05\x7b\x18\xc6\x16\xe1\x65\x43\x40\xdb\x70\x75\x39\x4c\x1a\xa9\xd1\xc1\xba\x05\xb8\xcc\x99\x09\x3b\xe7\x5e\x08\x2d\x2d\x0a\x53\xf7\x39\x24\xfc\x98\xfc\xab\x88\x27\x2a\xe2\x64\xdb\xe8\x31\xa8\xe4\x30\x32\x34\xde\x26\x94\xd8\xec\x53\x55\x8a\xda\x67\x2f\x27\xe7\x01\x12\x8c\xc9\x63\xac\x41\xc4\x56\xa8\x05\x67\xab\x88\x4f\x19\x3e\x83\x8c\x3f\xdd\x5f\x3f\xdc\xbd\x2c\x13\x47\xf7\x5e\xc5\x59\x01\xc4\x53\x10\x36\xed\x9a\xe5\x02\xd1\x56\x89\xbe\x61\x8a\xd8\x6a\x21\x3a\x94\xc6\x22\x33\x4c\xf5\xd5\x22\xde\x7a\x7c\x4b\xfd\x06\x14\xdf\xc7\xfe\xaa\xe1\x89\xc1\x45\x7b\x55\xc8\x4e\x92\x1e\x1b\xd4\xb8\x85\xa1\x60\x58\x1e\xee\xa6\x83\x26\xce\x00\xac\x18\x69\x2e\x4b\x07\x14\x83\xc8\x5c\x72\x80\x7b\xbe\xf9\xd3\x4f\x00\x9f\xc1\xc1\xe3\x65\x8e\x8b\xbf\xbe\x85\x3d\x58\xf6\x92\xae\x8d\x1b\xcb\xb9\x38\x79\x86\x82\x77\x66\x30\x54\x0d\x5d\xb2\x87\x8f\xae\x2d\x04\xa4\x39\x3b\xa5\x28\x3c\x08\xba\x8e\x41\x7e\x03\x88\x1f\x82\x7f\xf5\xf0\x44\x77\x66\x8d\x64\x61\x17\x6c\x24\x06\x36\xc0\xf0\x19\x85\xa1\x73\x4c\xd6\xd5\x73\xa3\xc3\x92\xd9\xa9\xb6\x25\x4d\x02\x44\xb6\x8b\x50\xcf\xf3\x79\xf8\x04\xe1\x73\x4c\xda\x5c\x3f\xdc\xf9\x63\xbe\x79\x59\x2a\xc2\x1e\x6d\x12\x1b\x00\x16\xf7\x3a\xd8\xb5\x11\xd6\x16\xb2\xd6\x98\x36\xb4\x07\x63\x28\x04\x99\x92\x21\x06\xaf\x98\x52\x0c\xc3\x56\xc2\x59\xba\xa6\x1f\x93\x7f\x75\xf1\x98\x9f\x86\xd1\x46\x16\xc3\x28\xe0\xe4\xca\x18\x01\x85\x6b\x22\xed\x01\x58\xe6\xbe\x9b\x26\xb0\xce\x65\x30\x47\xfa\x56\xdb\xb0\x19\x6a\x3a\xcf\xb3\x4d\xd6\x3c\x61\xf8\x0c\x32\xbe\x7b\xbc\xbe\x79\x61\x45\x71\xaf\x96\xde\x4a\xaa\xa3\x58\x6d\x3a\x18\x99\x54\xd6\xd2\xe2\x60\x45\x6f\xdb\x13\x44\xd5\x36\x12\xda\x3a\xd4\xc8\xa9\xaa\xe1\x54\x9e\x7d\x0e\x13\xdf\xc7\xfe\xaa\xe1\x31\x3c\x91\x64\xf6\x3d\x3a\x2d\x7b\xa2\xd2\x2e\x5b\xe4\x30\xcd\xa4\x59\x31\x94\x62\x75\x84\xb6\x6d\x01\x84\x8e\xad\x48\x58\x2e\x63\x30\xf2\xd9\x46\x88\x9f\x00\x3e\x83\x83\xbf\xdc\x5e\xe7\x5f\x68\xf9\x02\xff\x8c\x84\x26\x30\x3a\xa7\x09\x98\xac\x2a\xdc\xb5\x2c\x7d\x09\x21\xd3\x72\x85\x35\x11\x97\x8d\x95\x03\x52\xd7\x4e\x6f\x28\x65\xc1\x25\xbe\x76\xd1\x59\xa6\x4b\xdf\xe7\x7e\x85\xff\x8f\x2c\xc4\x3f\xad\x61\x12\x69\x72\xc8\x5e\x23\x42\xc0\x66\xe2\xc8\xa1\x56\x39\xda\x34\x16\x79\x8a\xf3\x18\x7b\x1e\xc6\x1f\x87\xde\xaa\x70\xcc\xb4\xa4\x0c\xdf\x67\xd3\xf0\x33\xc2\xa7\x1e\xe2\xff\x75\x7f\xcd\x5d\xdf\xdc\xe4\x9b\x7e\x69\x6b\xfb\xdd\x4b\xd1\x78\x2c\x1a\x73\xa4\x91\x77\xfa\xde\x5a\x3d\x15\xa1\x8c\xb6\x69\xa5\x2b\xcc\x6e\x06\x1b\xe1\xec\x73\xd1\x08\x0f\xee\x4a\xa3\xb3\xec\xb1\xf9\x94\xfd\x6b\x59\x3c\x26\x98\x88\x26\xd3\x14\x57\xb6\x78\x47\x81\x0f\x96\x00\x4d\x45\xe6\xc2\x5c\x21\xbe\x95\x7a\x0f\xee\x5c\xd4\x73\xc7\x1a\x39\xf3\x70\xe3\x15\xd4\xb3\xed\xb3\xf9\x82\xe2\x73\xd4\xc6\xeb\x1f\xaf\x6f\xfc\xed\xe3\x5f\xa8\x3e\xfe\xa9\xd1\x22\x82\x43\x97\x47\x4a\x63\xd5\x0c\xd9\x43\x27\x56\x46\xdb\x88\x1a\xe5\x13\xdb\xc5\xbd\xd1\xda\x72\x28\x76\x2f\xe1\x29\x21\x8b\x6c\x9c\xc5\xc8\x4f\xd1\xbf\x0a\x79\x0c\x90\xcc\x70\x6c\x2b\xa4\x5e\x9c\x3e\x8b\x0a\xbd\x7c\x2a\x56\x6d\x65\x98\xb1\x7c\x2c\x00\xa3\xf6\xc9\x0e\xce\xde\x3b\x59\xc2\x45\x19\xf2\x6c\xc3\xc5\x2f\x20\x3e\x83\x8f\x6f\x6e\x6f\x5f\x58\x71\x84\xb6\x8d\x31\x49\x68\xe8\xa6\x5d\x24\x23\x16\x43\x73\xe3\x82\xa5\x60\x32\xe7\xde\x6e\x2d\xba\x51\x23\x2b\x97\x2c\x14\x9f\x7b\x2d\x27\x5f\xe7\x50\xf1\x90\xfa\xab\x85\x27\xb6\x9f\x3a\x6f\x75\xe7\x39\x78\xf7\xe0\xda\x89\x1b\x06\x0a\xa2\xe2\x8e\x09\x8e\xe4\x39\x94\xd0\x33\x07\xed\x81\xbc\xb4\x03\x6c\xa9\x65\x52\x9c\x6d\xd7\xdb\x47\x7e\xcf\xb1\x8e\xd1\x75\xed\x2f\xcb\x40\x5a\x80\x06\xd4\x14\xad\x3a\x07\xe9\xd2\xb0\xa1\xc5\xdd\xee\x83\xc8\xd1\x30\x97\x44\xc5\xd2\xd9\x1d\xb8\x10\x28\xbd\x27\xf0\xc8\xb5\x8e\x3b\x37\xdf\x62\x0a\xfc\x10\xfb\xab\x82\x27\xf6\x0d\xef\x81\x88\x98\xba\x32\x25\xb2\xcd\x78\xee\xa6\x8c\xa1\x7b\x69\xd3\x5e\x7b\x9b\xd3\x5c\x44\x5e\x1a\x2a\x63\xe7\x24\xe7\x55\x2e\x00\x7d\x7c\xfb\xbc\xf8\xc6\x00\x9f\xc1\xc1\xfb\xbb\x17\xf6\x44\x14\xed\xcd\x4c\x3d\x46\xcd\x06\x59\xef\x67\x49\xf7\xc4\xd0\x15\xe5\xb0\xc9\xd9\x44\xd3\x60\xc6\x60\xcd\x1a\x63\xc7\xa4\x4a\x84\x26\x77\x9f\x75\x0e\x03\xef\xef\x5e\x9f\x85\x3a\x55\x02\xa1\xb1\x63\x2f\xdf\x56\x23\x07\x91\x0d\xe0\x35\x5c\x62\x96\x63\x6b\xd4\x5e\x82\xd3\x3d\x19\x4a\x54\x73\x37\x83\x9a\x2d\x52\x01\xaa\xe3\x8d\xc2\x17\xdf\x14\xdf\x33\xd8\xf7\xa3\xff\xdc\x2f\x71\xdd\x22\x63\x6d\xcb\x34\xdd\xa9\x04\x1a\xae\xbc\x71\xc0\x4a\x55\x31\xe2\x76\x56\x2d\xb4\xc4\x01\xca\xb2\x29\x16\x47\xec\x34\x54\xd6\x51\xe7\xe9\x88\x7e\xce\xfe\xd5\xc5\x13\xeb\xf9\x9e\x54\xb2\x20\x65\x74\xa6\x2e\x53\x0e\x0a\xef\xb4\xb2\x50\x0c\x8e\x6d\x82\x5b\x1b\xb6\xc1\x14\x8c\x4a\xdc\x7b\xb7\x44\xf4\x6a\xf2\x73\xb9\xf8\x25\xc5\xff\xbd\x92\x7f\xfb\x00\xf9\x6b\xe1\xfe\x19\xbe\x2f\x90\x7d\x81\xe9\x33\x9a\x4f\x38\x4e\xf5\x33\x2a\x66\x07\x89\x8f\xc1\x73\x43\xb5\x90\x49\xae\x10\x65\x5b\x80\x00\x32\x9b\x27\x13\xad\xe0\xea\xd9\x3e\x52\x80\xb4\x73\x59\x2c\xa3\xcf\x33\xd2\xcf\x98\xe7\x57\x02\xfc\x9e\xc9\x51\x83\xff\xe2\x16\xf5\x3f\x34\xf0\xa3\x86\x7d\xd4\xa0\x07\xd1\x53\x24\x5f\xb8\x76\xa2\x3a\x68\x7a\x9b\x34\xcb\xea\x88\xd9\x3b\xc8\x67\xa8\x17\x03\x8b\xd2\xdc\xb2\x64\x2a\x1b\xf1\xe2\xb1\xc3\x91\xa8\x6c\x33\xd8\x20\x1c\x93\xbf\xac\x0e\xcf\x99\xdc\xe7\x46\x47\x97\x4f\x9a\xd5\xfb\xff\xe7\xf2\xbe\xff\xed\xdd\xf5\x7d\x3f\xdd\xb4\xf5\xf7\xaf\xab\xc0\x17\x8f\xfe\x1c\xd5\x88\x8f\x8f\xe8\xfd\xb3\x29\xfe\x3f\x58\x0b\xff\x83\xe5\xb9\xc3\xc5\xfd\xed\xb7\xff\x0e\x00\x00\xff\xff\x08\x19\xc3\x80\x32\x3f\x00\x00")

func pinned_updateJsonBytes() ([]byte, error) {
	return bindataRead(
		_pinned_updateJson,
		"pinned_update.json",
	)
}

func pinned_updateJson() (*asset, error) {
	bytes, err := pinned_updateJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pinned_update.json", size: 16178, mode: os.FileMode(0666), modTime: time.Unix(1747043000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0x8c, 0x45, 0xef, 0xde, 0xb1, 0xd3, 0x67, 0x1d, 0xd8, 0x51, 0xd5, 0x5e, 0x35, 0x2f, 0x6d, 0x2, 0xca, 0x9d, 0xf0, 0x84, 0x4, 0xb1, 0x32, 0x49, 0xb3, 0x9b, 0x8e, 0x8b, 0x6, 0x90, 0xbb}}
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
	"DiscordCanary.lnk":  discordcanaryLnk,
	"pinned_update.json": pinned_updateJson,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"DiscordCanary.lnk":  {discordcanaryLnk, map[string]*bintree{}},
	"pinned_update.json": {pinned_updateJson, map[string]*bintree{}},
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
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
