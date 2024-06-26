// Code generated by go-bindata. DO NOT EDIT.
// sources:
// DiscordCanary.lnk (1.923kB)
// pinned_update.json (12.712kB)

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

	info := bindataFileInfo{name: "DiscordCanary.lnk", size: 1923, mode: os.FileMode(0666), modTime: time.Unix(1709215620, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbb, 0xc, 0x3a, 0x83, 0x1a, 0x8f, 0x80, 0xb, 0x96, 0x4d, 0xb4, 0x62, 0x81, 0xaf, 0x3c, 0xdb, 0xc4, 0xdf, 0xab, 0x6a, 0xc6, 0x10, 0xe1, 0xff, 0xd, 0x2a, 0x13, 0xd6, 0xf1, 0x2b, 0x7a, 0xa2}}
	return a, nil
}

var _pinned_updateJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x9a\xdd\x4e\x65\xb9\x8e\x80\xef\xeb\x29\xd0\xbe\x9e\x29\x62\x3b\x76\xec\x7e\x95\xa3\xa3\x96\xff\xd2\x85\x8a\x2e\x18\xa0\x7a\xa6\x75\xd4\xef\x3e\xda\x14\x45\x01\xea\x3e\x3a\x42\xbd\x25\xc4\x0d\x22\x2b\x90\xb5\xf2\xad\x6f\x59\x8e\x93\x7f\x7d\x38\x3b\x3b\xfc\x7a\x55\x5f\x2f\xfb\xf6\xf0\xd3\xd9\xb1\x79\x76\x76\xa8\x8b\xdb\xbc\xba\xa9\x9f\x7f\xed\xba\xf0\xc7\xcb\x67\x67\x87\xfd\xf5\xf2\xf2\x49\xfb\xec\xec\xf0\xe9\xea\xf6\xee\xe7\xdf\xfa\xe6\xf6\xe2\xea\xcb\xe1\xa7\xb3\x7f\x3c\xf6\x9c\x9d\xc1\x7f\x3d\x69\x8c\xa7\x0d\x1c\xf3\xb1\xf5\xcf\x1f\x1d\x0f\x0f\xf2\x64\x38\x7c\xd2\x79\xed\xf9\xd9\x7f\xe9\x9f\x6f\x3f\x39\xb2\x1c\x7e\x3a\x3b\x2c\xd3\x51\xe2\x69\x64\xb6\x8b\x7b\x7b\x61\x67\x60\x18\x8c\xd9\x84\xce\x35\x86\xc3\x34\xd7\x68\x1e\xc2\xcb\xd6\x06\x5e\xf7\x57\x8d\xed\xf0\x64\xf4\xaf\x37\xc7\x89\x1d\x3e\xdd\xdd\x5d\xdf\xfe\x74\x7e\x5e\x97\xff\x9d\xfe\xc5\x6f\x7e\xff\xf8\xc0\xc2\xaf\xaf\x3f\x7e\xe9\xbb\xf3\xba\xb8\xbd\xbb\xb9\x3a\xf7\xeb\xeb\xf3\x6f\x7f\x70\xfe\xbf\x17\x5f\xce\xff\x4f\xe5\x1c\x3e\x8e\x8f\x38\xe6\xf9\x33\x78\xe7\x78\x7e\x64\xf6\xf1\xdb\xbf\x1d\x1e\xee\xf7\xc7\xf7\x1b\x1f\xaa\x2f\xef\xfc\xf6\x19\xb8\x7f\x3d\x01\xf5\xef\xf0\xbe\x00\xfc\x02\xf1\x73\xc8\xcf\x30\xff\x19\xe8\x67\x43\xfd\x09\xea\x61\x23\x22\x08\x46\x2d\xa3\x18\x86\x95\xa5\x7b\x6b\x5b\x8f\x59\x83\xd6\x80\x95\x19\x9b\x78\x77\xb6\x9a\x88\x27\x4d\x9b\xd1\xb1\xe6\x54\x38\x3c\x1b\xff\x84\xb0\x6f\xae\x7e\x7d\xec\x84\xc3\xe3\x4d\xff\x78\xf8\xed\x9f\x1f\x9e\xf0\x7f\xf4\x3c\x2f\xaf\xbe\xd6\xed\xef\x5f\xf2\xed\xba\x8e\xea\x33\xaa\xa3\x45\x36\x4d\x45\x65\x8f\x30\xee\x3d\x6d\x75\xc5\x06\x81\x21\xc8\x46\x7b\x9b\x24\x8d\xa5\x5d\x94\x6b\x06\x67\x14\xef\x7d\x5a\xd7\x1f\x01\xbe\x17\xdf\x39\xda\x23\xf6\xc2\x01\x83\x88\xb7\xeb\x82\x04\x31\x82\x19\xdb\x91\x49\x9b\x60\xf7\x90\xc4\x85\x23\x2c\x60\x72\xea\x5c\xb5\x01\x6a\xe3\xc9\x7d\x7f\x06\xfc\x55\xce\x7f\xbe\xb9\xb8\xbd\x7e\xbb\xbe\x83\x18\x54\x60\xc2\x5e\x69\x3d\x02\xcd\x12\x58\x52\x49\x86\x57\xcd\x21\x5d\xa5\x24\x9a\x0b\x1d\x87\xcd\xe0\x11\xc6\x05\xbb\x16\x95\xf3\x69\x7d\xbf\x87\xf7\x5e\x5c\x47\x70\x82\x99\x24\x03\x41\x09\x55\x74\xed\x44\x70\xa7\x5c\x41\x82\x5e\x39\x73\x2e\x4d\x28\xe4\x86\x3d\x74\xc5\x9a\x80\xc7\xc8\xde\xcb\xf5\xd4\xae\x3f\xc2\x7e\x95\xe7\xb7\xfd\xf9\xe6\xe2\xee\x0d\x8b\x0e\x9d\x53\x58\x34\x24\xba\x24\x08\x65\x63\xaf\x01\xb4\x26\xf7\xa4\x4e\xf6\xe4\xe3\x0b\xe9\x36\x74\xd8\xa3\x47\xec\xd9\x6d\x53\x05\xfa\xc4\xa2\x7f\xa3\xf7\x5e\x4c\x9f\xbc\xac\x03\x16\x05\x7a\x56\x0c\x5e\x63\x38\x6a\xf0\x11\x6b\xe5\xd8\xd2\x11\x43\x08\x25\x6d\x2a\x5b\xc7\xa2\x61\x9e\x9b\x0c\x2a\x35\x4e\x6d\xfa\x0f\xda\xaf\x52\xfd\x65\x1e\xff\xf6\x5c\xc7\x0e\x81\x20\xed\x95\x14\x23\xf3\xf8\x73\x10\x81\x34\xa3\xb9\x83\xe6\xd8\x6b\xf8\x58\xc2\x90\x68\x09\x40\xbd\xd4\xb7\xc9\x96\x41\xf3\xc4\x09\xfb\x37\x7c\xef\x45\x76\xe2\x6a\xd9\x30\x82\xa3\x96\x07\x72\xc7\x90\xe1\xe6\x3e\xcd\xca\xd4\x06\xcf\xa2\x64\x10\x2e\x1c\x96\x30\xdd\xc2\x1c\x92\x15\x91\x08\x4f\x9e\xb2\xff\xc0\xfd\x2a\xdb\x7f\xbb\xba\xc8\x7e\xbb\xae\xd7\x8e\x54\x73\x95\xa9\xa6\xd1\xba\xb9\xc4\x98\x6c\x98\xfa\x6e\xc7\xe1\x64\x0e\xe5\x3c\xaa\x24\x7b\x02\x0f\x5f\x5a\x63\x28\x0f\x99\xe4\xa7\x75\xfd\x1e\xde\x7b\x31\x3d\x60\x6e\xaa\xe0\x20\x54\x5d\xb4\x1c\x65\xb1\x83\x67\x2d\xf2\x89\xb3\xaa\xdd\xd0\xf6\xe4\xe5\x0a\x33\x59\x07\x0a\xd6\x54\xd7\x4c\x83\x3a\xb5\xe9\x8f\xb0\x5f\xe5\xf9\xd7\xbb\x8b\xcb\x37\x1c\xd3\x4d\x47\x63\x1d\x73\x98\xcc\x25\x65\x96\x2a\x73\x4e\xcd\xe3\x52\x34\x47\xf5\xee\xb9\xe6\xf1\xd5\x4c\xe3\xb9\x56\x85\xc1\x32\x5e\xbe\x27\x11\x2f\x39\xad\xe7\xf7\xf0\xde\x8b\xe7\x63\x2c\x73\x19\xdd\x44\x3a\x06\xcd\xb5\x80\xac\x45\x28\x2c\x5c\x07\x24\x42\x6c\xdc\x40\xe9\xd8\x9b\x27\x6a\xb4\xcf\xdc\xd4\x63\x53\x95\x9f\xda\xf3\x47\xd8\xaf\x8b\xe7\x17\xbf\x5c\x5c\xfa\x97\xbb\x37\x1c\xd3\x39\xab\xa7\x6e\x8b\x74\x68\x97\xd0\xb4\xa1\x9a\xb3\x7a\xcf\xa1\x35\x12\xa5\xbd\xd0\xb7\xa4\x49\x49\xe6\x2e\x71\x1e\x4c\x0b\x60\x8d\xac\x13\xc7\xf4\xef\x00\xdf\x8b\xef\x38\x61\x56\x17\x30\x75\x2d\x8a\x4d\x3b\x96\x66\xf6\xde\x25\x2a\x7d\x5c\x1e\xad\x18\x00\xb5\x76\xae\x8d\xb3\xa6\x00\x5b\xb7\x99\xfb\xc6\x3a\x79\x06\xf3\x0c\xf8\xab\x9c\xbf\xfa\xad\x6f\x2e\xfd\x77\x7c\xc3\xca\x2b\x34\x80\x4e\xec\xc5\x33\x14\x69\x0a\x93\xb2\xd9\x68\x5b\xb0\x16\x37\xb3\xd2\xc8\x62\x9f\xec\x82\x33\x78\xc8\x2e\x5e\xa5\x1a\xac\xeb\xb4\xca\x7f\xe7\xf7\x5e\x8c\xaf\x14\xe8\x19\x06\xde\x0c\x93\xbd\x22\xba\x1d\x98\x10\x69\xc9\x68\x4a\x85\x20\x2c\xc4\xad\xc5\xbb\x68\xb8\x69\x97\xe1\x9a\xc7\x2f\xe5\xe4\x11\xfe\x29\xef\xd7\x55\x63\xae\xfb\xf2\x32\x3f\x75\x7e\x7e\xbb\xca\x27\xcd\x45\x43\xe6\xdc\x73\x81\x1b\x03\x53\x49\x57\xec\x12\x83\xa9\x06\x38\x19\xb8\x8c\x86\x35\x7c\x2b\xb5\xcb\x02\x85\x31\x6d\xe2\xf2\x13\x6f\x2b\xfd\x20\xf8\x5e\xa4\x27\xea\xb5\xac\xb9\xbb\x5c\x05\xb6\xef\xdc\x1b\x26\x14\xaa\x04\xd0\x71\xfd\xbf\x7b\x49\x53\x5b\x69\xed\x4e\x69\xcf\x28\xf2\x4d\xc8\xd4\xa7\xaf\xca\x3c\x23\xfe\x2a\xed\x3f\x5d\x5d\xbd\x61\xe1\x21\x5c\x04\xd5\x04\x81\x73\xbb\xeb\x32\x5f\x40\x1d\x3b\x68\x53\x04\xd7\x7d\x11\x21\x47\x2d\x33\x1b\xe6\x28\x4a\xd9\x3a\xb9\x62\xad\xcc\xd3\x0a\x7f\x64\xf7\x5e\x54\x87\x05\x5c\xc2\x23\x22\xad\x51\x29\x5a\x97\xb1\xf9\xec\x3d\x61\xa9\x3b\x96\xce\x35\x49\x74\x58\xeb\x9c\x3a\x69\xab\xee\x55\xd8\xb8\x37\xd3\xa9\x55\xff\xce\xfa\x55\x92\xd7\xc5\xed\xb5\xdf\xe5\xa7\xb7\x2b\xba\x91\xef\xe3\x22\x69\x09\xb5\x28\x33\x82\x92\x2f\xd9\x0a\x13\x34\x48\x73\x48\xd6\x30\x82\xc6\x36\x76\xa0\x7d\x4c\x23\x27\x83\x27\x03\x8d\x13\x47\xf6\xef\xfc\xde\x8b\xec\x53\xa4\xa3\x82\x6a\xcd\x24\xdc\x91\x63\x8d\x8c\x95\x45\x16\x69\x1b\xa6\x24\x56\x29\x9a\x09\x6d\x4a\x6d\x35\xcc\x05\x81\x8a\xa1\xbd\x4f\x1e\xd7\x9f\xf2\x7e\x95\xf0\x37\xd7\x6f\xf8\xc0\x40\x09\x0d\x1d\xe8\x9b\x09\x45\x62\x47\xa8\x16\xaf\x1e\xda\xc6\xc8\x4b\x79\x98\xb0\xf9\x98\x9b\x05\xa1\x31\xc8\xca\x06\xca\x71\xe1\x5a\xcf\x76\xf5\x4e\xc0\xfe\xe6\xfa\xdd\x1c\x15\xc0\xbd\x5d\x75\x57\x83\x44\x60\x6e\x1a\x85\x15\x28\x9e\x8e\x53\x6c\xda\x58\xdd\x09\x34\x7a\xae\x35\x71\x49\x86\xa7\x0c\x30\x67\x28\x81\x93\xc7\xf4\x07\xd4\xaf\x0b\xe9\x7d\xfb\xf9\xee\xea\xfa\xe7\xbc\xba\x79\xc3\x65\x19\xdc\x96\xd5\x36\xf5\x7e\xbd\xe4\x50\x8b\xcd\x76\x27\x6b\xaf\xce\x2d\xb1\x0d\x75\xe5\x82\xd9\xdb\x53\x5d\xd4\xf6\x98\x9a\x1e\xd8\x1b\xe1\xc4\x65\x99\xa7\x0c\xdf\x8b\xf3\x4c\x4c\x8b\x75\xed\xc6\xa9\xa0\x34\x61\x6b\x37\x0b\x45\x63\xec\x39\x92\xdd\x79\xcc\xa4\x95\x73\x8b\x8c\xd2\x2d\xb8\x36\x0d\x0d\x77\x86\x3e\x79\x68\x7f\xc1\xfc\x55\xf2\xff\xe2\xbf\xf6\x5b\xaf\xbe\xa3\x24\x17\xd0\x2e\x71\x40\x03\xf7\x2e\x58\x73\x14\x47\x44\x43\xda\xfd\x91\x81\x4c\xa8\x63\x10\xf2\x4e\x8a\xd2\x29\xea\xa4\x5b\x43\xa3\x4f\xab\xfe\x0f\x82\xef\x45\xfc\x9a\x8c\x44\xe1\x9d\x53\x7d\xac\xb9\xa4\x04\x63\x70\x20\x8e\xbd\xc8\xd3\x2a\x43\x1a\x71\x0b\xed\x6a\x2d\x41\x41\x65\x98\xdb\x45\xf6\xe9\x73\x9a\xe7\xc4\x5f\xa5\x7d\xdf\x5c\x1e\xe7\xfd\x76\x9d\x27\x25\xa1\x75\x4c\x1c\x77\x10\x78\xcc\x4e\xc1\x11\xa5\x51\x44\xc6\x2b\x70\xe7\xe6\x0e\xa9\x22\x81\x05\x82\x65\xec\xca\x02\x2d\x83\xfa\xc4\x59\xfc\x03\xbe\xf7\x22\x7c\xeb\xd2\x30\x8b\xae\x16\x8c\xb6\xb9\x57\x34\x30\xb7\xba\x66\xc8\xc4\x51\x01\xa3\xc7\xe6\x86\xcd\x6b\xcd\xad\x15\x12\x73\x8e\x6e\x52\xb4\x53\x0b\xff\x04\xf7\x7f\x66\xfb\x87\x87\xb7\xf0\x5c\xea\xbf\xe2\xfd\x88\xe7\x91\xf1\x77\xba\x0f\x5c\xff\x6c\x3b\xba\x75\x64\xa5\xec\x50\x34\x07\x23\xeb\x58\x19\x3d\x68\xa7\x4c\xc0\x48\x98\xac\x29\x63\x6c\xc2\xe2\xb1\x9c\x64\x6c\x06\x43\xa0\x59\xf5\xbd\x9c\xf2\x37\xc2\x7a\x61\xe2\xb7\xf9\xbf\x30\xf0\xf1\x63\xff\x4b\xf3\x9e\xa8\xf2\xc4\x38\x1c\xf4\xe1\x85\x69\x7f\x7a\x7e\x3c\x78\xc5\xb4\x34\x21\xd6\x48\x18\xd6\x01\x93\x7a\x74\x4d\xb1\xa1\x3c\x82\x55\x85\xf7\xa4\xdd\x34\x73\xd9\xf4\x3d\x99\x0a\x61\x1a\xfe\x38\xf9\xf3\x77\x52\xf9\x21\x0c\x1d\x1e\xe5\xb8\x9f\xc5\xe1\xa6\xff\xe7\xeb\xc5\x4d\x3f\x3d\x53\xf5\x8f\x7f\x93\x1a\xff\x45\x08\xfd\xeb\xe2\xf7\x8b\x9e\x6f\x59\xc6\xcb\x0d\xd1\xfb\x03\x2e\xc7\x47\xfa\xf0\xc7\xff\x07\x00\x00\xff\xff\x43\xdb\x4d\x07\xa8\x31\x00\x00")

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

	info := bindataFileInfo{name: "pinned_update.json", size: 12712, mode: os.FileMode(0666), modTime: time.Unix(1711103992, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xbe, 0x4a, 0x9, 0xbc, 0x67, 0xcb, 0x90, 0xc9, 0x64, 0xc3, 0x92, 0xf0, 0xa5, 0x39, 0xbc, 0x54, 0x33, 0xeb, 0xd9, 0xde, 0xa9, 0xa, 0x32, 0x48, 0xfd, 0xe8, 0x38, 0x3e, 0xe9, 0x99, 0xaf}}
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
	"DiscordCanary.lnk": {discordcanaryLnk, map[string]*bintree{}},
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
