// Code generated by go-bindata. DO NOT EDIT.
// sources:
// tkg/manifest/telemetry/config-aws.yaml
// tkg/manifest/telemetry/config-azure.yaml
// tkg/manifest/telemetry/config-docker.yaml
// tkg/manifest/telemetry/config-vsphere.yaml
// tkg/manifest/telemetry/zz_generated.bindata.go

package telemetry

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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataPkgV1TkgManifestTelemetryConfigawsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x5d\x6b\xdb\x4a\x10\x7d\xd7\xaf\x18\x0c\x21\x10\xae\xfc\xc1\x85\x70\x11\xe4\xe1\xde\x5c\x68\x08\x25\x09\x4e\x28\x2d\xa5\x94\xd1\x6a\x6c\x6f\xbc\xda\x15\x3b\xb3\x4a\xdd\x5f\x5f\x56\x96\x82\x3f\xe4\xd8\x29\x74\xfd\xa4\x9d\x99\x73\xce\xcc\x9e\x71\x9a\xa6\x09\x56\xfa\x13\x79\xd6\xce\x66\x50\x4f\x92\xa5\xb6\x45\x06\x77\x58\x12\x57\xa8\x28\x29\x49\xb0\x40\xc1\x2c\x01\xb0\x58\x52\x06\xb2\x9c\xa7\xbc\x62\xa1\x32\x15\x32\x54\x92\xf8\x55\x92\x1c\x84\x7a\x24\x5f\x6b\x45\xff\x2a\xe5\x82\x95\x03\x78\xaf\x40\x29\x63\x1b\x68\xe8\xdf\x62\x5b\xc3\x5f\x9b\xc0\x42\x7e\xea\x0c\x6d\xf1\xfb\x1c\xd5\x10\x83\x2c\x9c\xd7\x3f\x51\xb4\xb3\xc3\xe5\x3f\x3c\xd4\x6e\x54\x4f\x8e\x8a\x50\x6b\xd0\xd4\x47\x54\x1f\x0c\x71\xcc\x4c\x01\x2b\xfd\xc1\xbb\x50\x71\x06\x5f\x07\x83\x6f\x09\x00\x80\x27\x76\xc1\x2b\x6a\xee\x98\x94\x27\xe1\xc1\x5f\x30\x78\xed\xa1\xf9\x52\xce\xce\xf4\xbc\xc4\x8a\xdb\xb2\x9a\x7c\xde\x94\xcc\x49\x62\x82\xd1\x2c\x4d\x68\x87\xa5\x95\x32\xfc\x91\xae\xe5\xf7\xb1\xb6\x39\x0d\x51\x89\x6a\xa1\x2d\x15\x54\x19\xb7\x2a\xc9\xca\x7b\x09\xb5\x9d\x79\x64\xf1\x41\x49\xf0\x34\x3c\x85\x1f\x5f\x78\x53\x02\xbe\x70\xab\x42\xa8\xac\x0c\x0a\xbd\xbb\x69\x67\xc5\x3b\x53\x19\xb4\xa7\x29\x58\x86\x9c\xb0\x28\x37\xeb\x8e\x70\xf6\x5b\xe8\x3f\x6d\x0b\x6d\xe7\x7f\xc6\x49\x69\xde\xa2\x73\xc8\x9f\x49\x49\x6b\xaa\xde\x35\x89\xca\x0f\xae\xc7\xf1\x05\x89\x74\x53\x9a\x45\x82\x6e\xb0\x6f\x34\x92\x00\xec\x2f\xd3\x09\x9b\xb1\xb7\xf4\x39\x8a\x5a\x8c\xea\x49\x4e\x82\xdd\xfe\x5f\x7b\x67\x6f\x5d\x7e\x6c\x52\x27\x6c\x3d\x80\xc1\x9c\x4c\x33\x36\x80\x5d\x5b\x8c\x3a\x69\x6b\xf0\xf3\x33\x3e\x4f\xb8\x22\x15\xb3\x59\x2d\xa8\x08\x86\x32\x18\x8c\xe1\x62\x74\x09\x17\xf1\x37\x48\x00\x94\xb3\x2a\x78\x4f\x56\xad\x1e\x9c\xd1\x6a\x95\xc1\x94\x2a\x13\xff\xf9\x00\x66\xa8\x0d\x15\xb7\x2e\xe7\x1b\xcd\xe2\xfc\xea\xa3\x2e\xb5\x64\x30\x19\x27\x00\xcf\x2e\x7f\x6a\xed\xbd\x16\xd4\x91\xc5\x23\x5b\x91\xdd\x68\xf3\xbd\xf5\xde\x07\x9e\x79\x7d\x50\x89\xae\xe9\x7f\xc2\xc2\x68\x4b\x8f\xa4\x9c\x2d\x38\x83\xbf\x2f\xc7\xe3\x8d\xac\xe8\x7d\xd4\x96\x3c\x6f\xd2\x44\x83\xf5\x4f\x7b\xf3\xe8\x12\xe7\xdd\xcc\xb6\x23\xca\x95\x25\xda\x22\xdb\xb9\x8e\xb0\xa3\xb7\x10\x63\x42\x9a\xb2\xb3\x2e\x0f\x6e\x15\x9d\x9f\x56\x28\x8b\xab\x51\x77\xd5\x9b\x1f\xfd\x11\x07\x4d\x69\xf0\xe6\xea\x8c\x7b\x93\xd6\x26\xd8\x8f\x92\xad\x77\x65\x76\xbd\xdf\x3c\x3d\x3d\x7c\x7f\x98\xde\x7f\xfe\xb2\x87\x58\xa3\x09\xfd\xad\x6f\x16\x3f\xfe\x76\xf5\xdd\xfd\xbb\x4a\x3d\xb1\xa0\x97\xce\x8b\x77\x54\x93\x4f\x7e\x05\x00\x00\xff\xff\x9e\xe9\xbb\x80\xa7\x07\x00\x00")

func bindataPkgV1TkgManifestTelemetryConfigawsYamlBytes() ([]byte, error) {
	return bindataRead(
		_bindataPkgV1TkgManifestTelemetryConfigawsYaml,
		"tkg/manifest/telemetry/config-aws.yaml",
	)
}

func bindataPkgV1TkgManifestTelemetryConfigawsYaml() (*asset, error) {
	bytes, err := bindataPkgV1TkgManifestTelemetryConfigawsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "tkg/manifest/telemetry/config-aws.yaml",
		size:        1959,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPkgV1TkgManifestTelemetryConfigazureYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\xdb\x6a\x1b\x31\x10\x7d\xdf\xaf\x18\x0c\x21\x10\x22\x5f\x28\x84\xb2\x90\x87\x36\x85\x86\x50\x92\xe0\x84\xd2\x52\x4a\x99\xd5\x8e\x6d\xc5\xba\x2c\xd2\x68\xa9\xf3\xf5\x45\xeb\xdd\xe0\x6b\xec\x14\xa2\x7d\x5a\xcd\xcc\x39\x47\xa3\x33\x12\x42\x64\x58\xa9\xef\xe4\x83\x72\x36\x87\x7a\x94\xcd\x95\x2d\x73\xb8\x45\x43\xa1\x42\x49\x99\x21\xc6\x12\x19\xf3\x0c\xc0\xa2\xa1\x1c\x78\x3e\x15\x61\x11\x98\x8c\x60\xd2\x64\x88\xfd\x22\xcb\xf6\x42\x3d\x90\xaf\x95\xa4\x4f\x52\xba\x68\x79\x0f\xde\x0b\x90\x08\xd8\x06\x1a\xfa\xd7\xd8\x96\xf0\x57\x3a\x06\x26\x3f\x76\x9a\xd6\xf8\x7d\x81\xb2\x8f\x91\x67\xce\xab\x67\x64\xe5\x6c\x7f\xfe\x31\xf4\x95\x1b\xd4\xa3\x83\x22\xe4\x12\x54\xf8\x84\xea\xa3\xa6\x90\x32\x05\x60\xa5\xbe\x7a\x17\xab\x90\xc3\xaf\x5e\xef\x77\x06\x00\xe0\x29\xb8\xe8\x25\x35\x7b\x81\xa4\x27\x0e\xbd\x73\xe8\xbd\x9c\xa1\xf9\x93\xce\x4e\xd4\xd4\x60\x15\xda\xb2\x9a\x7c\xd1\x94\x4c\x89\x53\x82\x56\x81\x9b\xd0\x06\x4b\x2b\xa5\xff\x57\x2c\xe5\xef\x62\x6d\x73\xde\x15\xfa\x1c\x7a\x06\xe5\x4c\x59\x2a\xa9\xd2\x6e\x61\xc8\xf2\x5b\x09\x95\x9d\x78\x0c\xec\xa3\xe4\xe8\xa9\x7f\x0c\x3f\x3e\x47\x4f\x2d\x71\x23\xa2\xd9\x58\x55\xb5\x9a\xc1\x64\x2a\x8d\x4c\x6f\xee\x84\xb3\xec\x9d\xae\x34\xda\xe3\x64\xcd\x63\x41\x58\x9a\xd5\xba\x03\x9c\xbb\x2d\xfb\x59\xd9\x52\xd9\xe9\xfb\x38\x57\x14\x2d\x7a\x88\xc5\x13\x49\x6e\x4d\xbc\x73\x2c\x93\xf2\xbd\xe3\x78\x78\x20\x13\xdd\x98\x26\x89\xa0\x6b\xec\x2b\x07\xc9\x00\xb6\x87\xf7\x88\x49\xdc\x7a\x64\x0a\x64\x39\x1b\xd4\xa3\x82\x18\xbb\xf7\xe6\xca\x3b\x7b\xe3\x8a\x43\x9d\x3a\xe2\x95\x01\xd0\x58\x90\x6e\xda\x06\xb0\x69\x8b\x41\x27\x6d\x09\x7e\x7a\x12\x4e\xb3\x50\x91\x4c\xd9\x41\xce\xa8\x8c\x9a\x72\xe8\x0d\xe1\x6c\x70\x01\x67\xe9\xeb\x65\x00\xd2\x59\x19\xbd\x27\x2b\x17\xf7\x4e\x2b\xb9\xc8\x61\x4c\x95\x4e\x2f\x2d\xc0\x04\x95\xa6\xf2\xc6\x15\xe1\x5a\x05\x76\x7e\xf1\x4d\x19\xc5\x39\x8c\x86\x19\xc0\x93\x2b\x1e\x5b\x7b\x2f\x05\x75\x64\x69\xf1\x5a\x64\x33\xda\xfc\xaf\xdd\xf7\x9e\x6b\x5e\x2e\x94\xac\x6a\xfa\x42\x58\x6a\x65\xe9\x81\xa4\xb3\x65\xc8\xe1\xc3\xc5\x70\xb8\x92\x95\xbc\x8f\xca\x92\x0f\xab\x34\xc9\x60\xbb\xbb\xbd\xba\x94\xc1\x69\xd7\xb3\xf5\x88\x74\xc6\xa0\x2d\xf3\x8d\xed\x04\x3b\x78\x0d\x31\x25\x08\x11\x9c\x75\x45\x74\x8b\xe4\x7c\x51\x21\xcf\x2e\x07\xdd\xd6\xce\xfc\xe4\x8f\xd4\x68\x12\xd1\xeb\xcb\x93\xb0\x33\x69\x69\x82\xed\x28\xd9\x7a\x53\x66\x77\xf6\xeb\xc7\xc7\xfb\x3f\xf7\xe3\xbb\x1f\x3f\xb7\x10\x6b\xd4\x71\xf7\xd1\x57\x8b\x1f\xfe\xbb\xfa\xf6\xee\x4d\xa5\x9e\x02\xa3\xe7\xce\x8b\xb7\x54\x93\xcf\xfe\x05\x00\x00\xff\xff\xa4\x0c\x47\x2e\x17\x08\x00\x00")

func bindataPkgV1TkgManifestTelemetryConfigazureYamlBytes() ([]byte, error) {
	return bindataRead(
		_bindataPkgV1TkgManifestTelemetryConfigazureYaml,
		"tkg/manifest/telemetry/config-azure.yaml",
	)
}

func bindataPkgV1TkgManifestTelemetryConfigazureYaml() (*asset, error) {
	bytes, err := bindataPkgV1TkgManifestTelemetryConfigazureYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "tkg/manifest/telemetry/config-azure.yaml",
		size:        2071,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPkgV1TkgManifestTelemetryConfigdockerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x6b\x6b\x1b\x3b\x10\xfd\xbe\xbf\x62\x30\x84\x40\x88\xfc\xe0\x42\xb8\x2c\xe4\xc3\xbd\xb9\x70\x43\x28\x49\x70\x42\x69\x29\xa5\xcc\x6a\xc7\xb6\x62\x3d\x16\x69\xb4\xd4\xfd\xf5\x45\xfb\x08\x7e\xc6\x4e\x21\xda\x4f\xab\x99\x39\xe7\x68\x74\x46\x42\x88\x0c\x2b\xf5\x99\x7c\x50\xce\xe6\x50\x4f\xb2\xa5\xb2\x65\x0e\xf7\x68\x28\x54\x28\x29\x33\xc4\x58\x22\x63\x9e\x01\x58\x34\x94\x03\x2f\xe7\x22\xac\x02\x93\x11\x4c\x9a\x0c\xb1\x5f\x65\xd9\x41\xa8\x27\xf2\xb5\x92\xf4\x8f\x94\x2e\x5a\x3e\x80\xf7\x0a\x24\x02\x76\x81\x86\xfe\x2d\xb6\x16\xfe\x46\xc7\xc0\xe4\xa7\x4e\xd3\x06\xbf\x2f\x50\x0e\x31\xf2\xc2\x79\xf5\x0b\x59\x39\x3b\x5c\xfe\x1d\x86\xca\x8d\xea\xc9\x51\x11\xb2\x05\x15\x3e\xa1\xfa\xa8\x29\xa4\x4c\x01\x58\xa9\xff\xbd\x8b\x55\xc8\xe1\xdb\x60\xf0\x3d\x03\x00\xf0\x14\x5c\xf4\x92\x9a\xbd\x40\xd2\x13\x87\xc1\x25\x0c\x5e\xcf\xd0\xfc\x49\x67\x67\x6a\x6e\xb0\x0a\x5d\x59\x4d\xbe\x68\x4a\xe6\xc4\x29\x41\xab\xc0\x4d\x68\x8b\xa5\x93\x32\xfc\x29\x5a\xf9\xfb\x58\xbb\x9c\x0f\x85\xbe\x84\x81\x41\xb9\x50\x96\x4a\xaa\xb4\x5b\x19\xb2\xfc\x5e\x42\x65\x67\x1e\x03\xfb\x28\x39\x7a\x1a\x9e\xc2\x5f\x3a\xb9\x24\xdf\x31\x37\x2a\xda\x9d\x75\x5d\x1b\x39\x4c\xa6\xd2\xc8\xf4\xee\x66\x38\xcb\xde\xe9\x4a\xa3\x3d\x4d\xd9\x32\x16\x84\xa5\x59\xaf\x3b\xc2\xb9\xdf\xb5\xff\x2a\x5b\x2a\x3b\xff\x18\xf3\x8a\xa2\x43\x0f\xb1\x78\x21\xc9\x9d\x8f\xf7\x4e\x66\x52\x7e\x70\x22\x8f\xcf\x64\xa2\x9b\xd2\x2c\x11\xf4\x8d\x7d\xe3\x20\x19\xc0\xee\xfc\x9e\x30\x8c\x3b\xef\x4c\x81\x2c\x17\xa3\x7a\x52\x10\x63\xff\xe4\xdc\x78\x67\xef\x5c\x71\xac\x53\x27\x3c\x34\x00\x1a\x0b\xd2\x4d\xdb\x00\xb6\x6d\x31\xea\xa5\xb5\xe0\xe7\x67\xe1\x3c\x0b\x15\xc9\x94\x1d\xe4\x82\xca\xa8\x29\x87\xc1\x18\x2e\x46\x57\x70\x91\xbe\x41\x06\x20\x9d\x95\xd1\x7b\xb2\x72\xf5\xe8\xb4\x92\xab\x1c\xa6\x54\xe9\xf4\xd8\x02\xcc\x50\x69\x2a\xef\x5c\x11\x6e\x55\x60\xe7\x57\x9f\x94\x51\x9c\xc3\x64\x9c\x01\xbc\xb8\xe2\xb9\xb3\x77\x2b\xa8\x27\x4b\x8b\x37\x22\xdb\xd1\xe6\x7f\xe3\xbe\x0f\x5c\x73\xbb\x50\xb2\xaa\xe9\x3f\xc2\x52\x2b\x4b\x4f\x24\x9d\x2d\x43\x0e\x7f\x5d\x8d\xc7\x6b\x59\xc9\xfb\xa8\x2c\xf9\xb0\x4e\x93\x0c\xb6\xbf\xdb\xeb\x4b\x19\x9c\xf7\x3d\xdb\x8c\x48\x67\x0c\xda\x32\xdf\xda\x4e\xb0\xa3\xb7\x10\x53\x82\x10\xc1\x59\x57\x44\xb7\x4a\xce\x17\x15\xf2\xe2\x7a\xd4\x6f\xed\xcd\x4f\xfe\x48\x8d\x26\x11\xbd\xbe\x3e\x0b\x7b\x93\x5a\x13\xec\x46\xc9\xd6\xdb\x32\xfb\xb3\xdf\x3e\x3f\x3f\xfe\x78\x9c\x3e\x7c\xf9\xba\x83\x58\xa3\x8e\xfb\x8f\xbe\x5e\xfc\xf4\xc7\xd5\xf7\x0f\xef\x2a\xf5\x14\x18\x3d\xf7\x5e\xbc\xa7\x9a\x7c\xf6\x3b\x00\x00\xff\xff\xf8\x9f\xf7\xc2\x1a\x08\x00\x00")

func bindataPkgV1TkgManifestTelemetryConfigdockerYamlBytes() ([]byte, error) {
	return bindataRead(
		_bindataPkgV1TkgManifestTelemetryConfigdockerYaml,
		"tkg/manifest/telemetry/config-docker.yaml",
	)
}

func bindataPkgV1TkgManifestTelemetryConfigdockerYaml() (*asset, error) {
	bytes, err := bindataPkgV1TkgManifestTelemetryConfigdockerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "tkg/manifest/telemetry/config-docker.yaml",
		size:        2074,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPkgV1TkgManifestTelemetryConfigvsphereYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x5d\x6b\xe3\x3a\x10\x7d\xf7\xaf\x18\x02\xa5\x50\xaa\x7c\x70\xa1\x5c\x0c\x7d\xb8\xb7\x17\x6e\x29\x4b\x5b\xd2\xb2\xec\xb2\x2c\xcb\x58\x9e\x24\x6a\xf4\x61\xa4\x91\xd9\xec\xaf\x5f\xe4\xd8\x25\x1f\x4e\x93\x2e\x54\x7e\xb2\x66\xe6\x9c\xa3\xd1\x19\x09\x21\x32\xac\xd4\x67\xf2\x41\x39\x9b\x43\x3d\xc9\x96\xca\x96\x39\xdc\xa3\xa1\x50\xa1\xa4\xcc\x10\x63\x89\x8c\x79\x06\x60\xd1\x50\x0e\xbc\x9c\x8b\xb0\x0a\x4c\x46\x30\x69\x32\xc4\x7e\x95\x65\x07\xa1\x9e\xc8\xd7\x4a\xd2\x3f\x52\xba\x68\xf9\x00\xde\x2b\x90\x08\xd8\x06\x1a\xfa\xb7\xd8\xd6\xf0\x37\x3a\x06\x26\x3f\x75\x9a\xb6\xf8\x7d\x81\x72\x88\x91\x17\xce\xab\x5f\xc8\xca\xd9\xe1\xf2\xef\x30\x54\x6e\x54\x4f\x8e\x8a\x90\x6b\x50\xe1\x13\xaa\x8f\x9a\x42\xca\x14\x80\x95\xfa\xdf\xbb\x58\x85\x1c\xbe\x0d\x06\xdf\x33\x00\x00\x4f\xc1\x45\x2f\xa9\xd9\x0b\x24\x3d\x71\x18\x5c\xc2\xe0\xf5\x0c\xcd\x9f\x74\x76\xa6\xe6\x06\xab\xd0\x96\xd5\xe4\x8b\xa6\x64\x4e\x9c\x12\xb4\x0a\xdc\x84\x76\x58\x5a\x29\xc3\x9f\x62\x2d\xbf\x8f\xb5\xcd\xf9\x50\xe8\x4b\x18\x18\x94\x0b\x65\xa9\xa4\x4a\xbb\x95\x21\xcb\xef\x25\x54\x76\xe6\x31\xb0\x8f\x92\xa3\xa7\xe1\x29\xfc\x75\xa8\x16\xe4\xa9\xa5\x6e\x64\xb4\x5b\x9b\xca\xb6\xb3\x98\x4c\xa5\x91\xe9\xdd\xfd\x70\x96\xbd\xd3\x95\x46\x7b\x9a\xb8\x65\x2c\x08\x4b\xb3\x59\x77\x84\xb3\xdf\xb8\xff\x2a\x5b\x2a\x3b\xff\x18\xff\x8a\xa2\x45\x0f\xb1\x78\x21\xc9\xad\x95\x7b\x87\x33\x29\x3f\x38\x94\xc7\xc7\x32\xd1\x4d\x69\x96\x08\xba\xc6\xbe\x71\x90\x0c\x60\x7f\x84\x4f\x98\xc7\xbd\xa7\xa6\x40\x96\x8b\x51\x3d\x29\x88\xb1\x7b\x75\x6e\xbc\xb3\x77\xae\x38\xd6\xa9\x13\xde\x1a\x00\x8d\x05\xe9\xa6\x6d\x00\xbb\xb6\x18\x75\xd2\xd6\xe0\xe7\x67\xe1\x3c\x0b\x15\xc9\x94\x1d\xe4\x82\xca\xa8\x29\x87\xc1\x18\x2e\x46\x57\x70\x91\xbe\x41\x06\x20\x9d\x95\xd1\x7b\xb2\x72\xf5\xe8\xb4\x92\xab\x1c\xa6\x54\xe9\xf4\xde\x02\xcc\x50\x69\x2a\xef\x5c\x11\x6e\x55\x60\xe7\x57\x9f\x94\x51\x9c\xc3\x64\x9c\x01\xbc\xb8\xe2\xb9\xb5\xf7\x5a\x50\x47\x96\x16\x6f\x45\x76\xa3\xcd\xff\xd6\x7d\x1f\xb8\xe6\xf5\x42\xc9\xaa\xa6\xff\x08\x4b\xad\x2c\x3d\x91\x74\xb6\x0c\x39\xfc\x75\x35\x1e\x6f\x64\x25\xef\xa3\xb2\xe4\xc3\x26\x4d\x32\x58\x7f\xb7\x37\x97\x32\x38\xef\x7a\xb6\x1d\x91\xce\x18\xb4\x65\xbe\xb3\x9d\x60\x47\x6f\x21\xa6\x04\x21\x82\xb3\xae\x88\x6e\x95\x9c\x2f\x2a\xe4\xc5\xf5\xa8\xdb\xea\xcd\x4f\xfe\x48\x8d\x26\x11\xbd\xbe\x3e\x0b\xbd\x49\x6b\x13\xec\x47\xc9\xd6\xbb\x32\xbb\xb3\xdf\x3e\x3f\x3f\xfe\x78\x9c\x3e\x7c\xf9\xba\x87\x58\xa3\x8e\xfd\x47\xdf\x2c\x7e\xfa\xe3\xea\xfb\x87\x77\x95\x7a\x0a\x8c\x9e\x3b\x2f\xde\x53\x4d\x3e\xfb\x1d\x00\x00\xff\xff\xbc\xed\x18\x04\x1d\x08\x00\x00")

func bindataPkgV1TkgManifestTelemetryConfigvsphereYamlBytes() ([]byte, error) {
	return bindataRead(
		_bindataPkgV1TkgManifestTelemetryConfigvsphereYaml,
		"tkg/manifest/telemetry/config-vsphere.yaml",
	)
}

func bindataPkgV1TkgManifestTelemetryConfigvsphereYaml() (*asset, error) {
	bytes, err := bindataPkgV1TkgManifestTelemetryConfigvsphereYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "tkg/manifest/telemetry/config-vsphere.yaml",
		size:        2077,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPkgV1TkgManifestTelemetryZzgeneratedBindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x9a\xdd\x6f\xdc\x46\x96\xc5\x9f\xd5\x7f\x45\xaf\x80\x19\x48\x0b\x8f\xc4\xaf\x26\x9b\x06\xf2\x32\x49\x16\xc8\x43\x32\x8b\x5d\xef\x02\x8b\xad\x45\x50\x24\x8b\x4a\xc3\x92\xda\xdb\x92\x92\xb2\x83\xfc\xef\x83\xdf\xbd\x87\xee\x8e\xe3\x7c\xd8\x93\x19\x60\x02\x74\xac\x66\x93\x55\xb7\xee\xc7\xb9\xe7\x54\xf1\xfa\x7a\xfd\xe9\x7e\x4a\xeb\x9b\x74\x9f\x0e\xf1\x31\x4d\xeb\xe1\xf5\xfa\x66\xff\xa7\x61\x77\x3f\xc5\xc7\x78\xb5\xfe\xec\x2f\xeb\xaf\xfe\xf2\x62\xfd\xf9\x67\x5f\xbc\xb8\x5a\x5d\x5f\xaf\x1f\xf6\x4f\x87\x31\x3d\x3c\xe7\xef\x57\x2f\x6f\xae\xbf\x2d\xaf\x1f\x5f\xde\x5c\xdf\xc5\xfb\xdd\x9c\x1e\x1e\xaf\x1f\xd3\x6d\xba\x4b\x8f\x87\xd7\xd7\xe3\xfe\x7e\xde\xdd\xfc\x29\x7e\xf7\x70\xf5\x3a\xde\xdd\xfe\xf6\x07\xde\x3c\x1d\xd2\x87\x3d\x32\xed\xc7\x97\xe9\xf0\x61\xcf\x7c\xfb\xf0\xea\x9b\xf4\x9b\x27\x7a\xf3\xe6\xeb\xb7\x2e\xba\x5a\x9c\x73\xb3\x5f\xad\x5e\xc5\xf1\x65\xbc\x49\xeb\xb7\xb7\xae\x56\xab\xdd\xdd\xab\xfd\xe1\x71\x7d\xb1\x3a\x3b\x1f\x5e\x3f\xa6\x87\xf3\xd5\xd9\xf9\xb8\xbf\x7b\x75\x48\x0f\x0f\xd7\x37\x6f\x76\xaf\xb8\x30\xdf\x3d\xf2\xcf\x6e\xef\xff\xbf\xde\xed\x9f\x1e\x77\xb7\x7c\xd9\xdb\x03\xaf\xe2\xe3\x37\xd7\xf3\xee\x36\xf1\x07\x17\x1e\x1e\x0f\xbb\xfb\x1b\xfb\xed\x71\x77\x97\xce\x57\x97\xab\xd5\xfc\x74\x3f\xae\x65\xce\x7f\xa4\x38\x5d\xf0\xc7\xfa\x7f\xff\x8f\x69\x9f\xad\xef\xe3\x5d\x5a\xfb\x63\x97\xeb\x8b\xe5\x6a\x3a\x1c\xf6\x87\xcb\xf5\xf7\xab\xb3\x9b\x37\xf6\x6d\xfd\xfc\x93\x35\x56\x5d\x7d\x95\xbe\x63\x90\x74\xb8\x30\xb3\xf9\xfe\xe7\xa7\x79\x4e\x07\x1b\xf6\xf2\x72\x75\xb6\x9b\xed\x81\x7f\xf9\x64\x7d\xbf\xbb\x65\x88\xb3\x43\x7a\x7c\x3a\xdc\xf3\xf5\xd9\x7a\xbe\x7b\xbc\xfa\x9c\xd1\xe7\x8b\x73\x06\x5a\xff\xe1\xff\x9f\xaf\xff\xf0\xed\xb9\x5b\x62\x73\x5d\xae\xce\x7e\x58\xad\xce\xbe\x8d\x87\xf5\xf0\x34\xaf\x7d\x1e\x9f\x64\x75\xf6\xb5\x9b\xf3\xc9\x7a\xb7\xbf\xfa\x74\xff\xea\xf5\xc5\x1f\x87\xa7\xf9\xd9\xfa\xe6\xcd\xe5\xea\x6c\xbc\xfd\x7c\xb1\xf4\xea\xd3\xdb\xfd\x43\xba\xb8\x5c\xfd\x5e\xf6\x30\x8c\x8f\xff\x33\x03\xa5\xc3\xc1\xed\xd6\xc5\xe1\x69\xbe\xfa\x33\xa6\x5f\x5c\x3e\xe3\x8e\xd5\x0f\xab\xd5\xea\xf1\xf5\xab\xb4\x8e\x0f\x0f\xe9\x11\x9f\x3f\x8d\x8f\x0c\x63\x0b\x54\x40\x56\x67\xbb\xfb\x79\xbf\x5e\x13\xd4\x2f\xee\xe7\xfd\xe7\x99\xe7\xec\xb1\xe3\xa5\xf5\xee\xfe\x31\x1d\xe6\x38\x26\x1e\xdf\x3f\x5c\xfd\x9b\x7e\x5a\x9d\x7d\xf9\xd9\xe6\xd3\x6f\xd2\xf8\xf2\xe1\xe9\xee\xe2\x52\x71\x7d\x3b\x82\x92\x60\xb9\xfb\xc4\x04\xcb\x02\xfd\xa7\x87\xce\x1e\x76\x6f\xde\x5e\xdb\xdd\x3f\xb6\xcd\xea\xec\x0e\x18\xd0\x7f\x9a\xf6\xcb\xfd\x94\xec\x87\x17\x3b\x0d\x41\xe2\x5d\xf1\x6d\x75\x76\x37\x6d\x46\x59\x73\x62\x8b\x25\xe4\xc5\xbc\x7b\xd7\x9e\xcb\xf5\x57\xf1\x2e\xbd\x35\x1b\xbb\xe4\xcb\x79\x77\x85\x85\xab\x1f\x7e\xe1\xd9\xff\xdc\xbd\xe1\x59\xb3\xf4\xc7\x8f\xb2\x90\x5f\x7c\x94\x35\x5c\x5c\x9e\xae\xe8\xc7\x03\xb0\xec\x5f\x1b\x80\x05\x5f\x5c\x1e\x17\xff\x93\x11\xcc\x23\xbf\x38\xc8\x7b\x42\xf7\xce\x28\x47\x77\xfe\xe2\x48\x5f\x3c\x7c\xb6\x3b\x5c\x5c\xae\x87\xfd\xfe\xf6\x74\x84\x78\xfb\xf0\x2b\x3e\x7c\xfd\xe0\x2e\xf4\xec\xfa\xfe\x87\x93\xa7\x95\xc2\x54\xe5\xd7\x7a\xf0\xdf\x5f\xde\xfc\x77\xf9\xe2\xe5\xcd\x97\xc2\xc2\x17\x0b\xbe\x7d\x6a\xf0\x19\xbf\x7b\xf8\x9f\x78\x77\xbb\xfe\x44\xb9\x7d\x71\x1e\x72\x39\x87\xbc\x1d\x42\x2e\xb6\x21\x17\xc5\xfb\x3f\xf3\x1c\xf2\xd0\x84\xbc\xe1\x33\x85\xdc\x0e\x21\x4f\x43\xc8\x4d\x0c\xb9\x2c\x42\xee\xa6\x90\xa7\x2e\xe4\x38\x87\x5c\x32\xce\x18\x72\x55\xfa\x6f\x31\x85\x3c\x8f\x21\x8f\x65\xc8\xdb\x4d\xc8\x5d\x11\x72\x59\x86\x9c\x9a\x90\x53\x19\xf2\x94\x42\xde\x8c\x21\xb7\x5b\xb7\xa1\xda\x84\x5c\xf4\x21\x37\x29\xe4\x8a\xef\x53\xc8\x71\x13\x72\xdf\x84\x3c\x95\x21\xb7\x31\xe4\x96\xfb\xb1\x69\x0c\x79\xc2\x86\x4d\xc8\xf5\x10\xf2\x50\xbb\x4d\xd3\x14\xf2\x66\xd6\xa7\x0d\xb9\x6f\x43\xde\x56\x21\xd7\xb3\xcf\x3b\x31\x6e\x1f\x72\xd7\x84\x3c\x33\x7e\x13\x72\x3f\x85\xdc\x73\xad\x0e\x79\x4c\x21\x8f\x63\xc8\x7d\x0a\xb9\x2b\x43\xee\x63\xc8\xb1\x75\xbb\x18\x6f\x66\xce\x3a\xe4\xae\x0f\x79\x6a\xfd\xfe\xb6\x0d\x79\x53\x84\xdc\xcc\x21\xf7\x95\xdb\x3c\xb4\x21\x37\xac\xa7\x0d\xb9\xeb\x42\xde\x6c\x43\x2e\xab\x90\x37\xf8\x6a\xab\xf5\xb1\xd6\x3e\xe4\x81\x67\x0b\xf7\x53\x35\x86\x5c\x94\x7e\x8d\x67\x36\x95\x8f\x31\x54\x21\xf7\x63\xc8\xb1\xf3\xb5\xb7\xcc\x53\x86\x5c\x57\xf2\x41\xe5\x31\x62\xfe\x99\xe7\x36\xfe\x77\x39\x86\xbc\x6d\x42\xee\x62\xc8\x55\xe3\x7e\x21\x86\xd8\x46\x6c\xab\x18\x72\xda\xb8\x8f\x7a\xec\x65\x6d\x5b\x8f\x27\x36\x61\x63\x5b\x87\x5c\x0e\x1e\xa7\xb4\x0d\x79\x9a\x7d\xfe\xcd\x10\xf2\x58\xfb\x98\x3d\x7f\xb3\x8e\x2a\xe4\x2e\x85\x9c\xa2\xe7\x42\x3b\x85\x3c\x97\x21\xcf\x83\xdb\x32\x6d\x3c\x37\xb6\xb5\xaf\x95\x35\x91\x3f\xc4\x67\x53\x7a\xae\x11\x4b\xc6\xc5\x2e\xae\xd7\xdc\xd3\x84\xdc\x26\x5f\x23\x7e\xde\xa6\x90\xb7\xd1\xfd\x6e\x39\x59\xf8\x3d\x4d\xe7\xf7\x90\xdb\xcc\x4f\x0c\x89\x67\x23\xbf\x56\x83\xc7\x1d\x5f\x0f\x83\x3f\xcf\x3d\xf8\x98\x75\x14\x9d\xdb\x46\x8e\x11\x73\x6a\x60\x8b\x1f\xba\x90\xbb\x56\x31\x1a\x3c\x17\x53\x0a\xb9\xdf\x7a\x7e\xf2\x3b\x39\xcd\xef\x8c\x53\x17\xee\xc7\x34\x79\x9c\xe6\xde\x63\x69\xf9\x45\x5d\x34\x1e\xc7\xb1\xf1\x75\x50\x53\x16\xe3\x18\x72\x37\x7a\x6d\x30\xae\xd9\xde\xbb\xbf\x89\xd1\xa4\x1c\x69\x26\x8f\xc7\xb6\x53\x8d\x28\x07\xa8\xb7\xbe\xf4\xda\xc3\x7f\x33\xf9\x52\x7a\x3e\xd6\xac\x67\x72\x3f\x6f\x7b\x5f\x03\xb6\x51\x67\xe4\x90\xf9\x8d\x7c\xec\x3c\x2f\xc6\x3e\xe4\x11\x7b\x06\xf7\xc5\xb0\xf1\x5a\x21\xf7\xdb\xc6\xe3\xda\x94\x6e\xdf\x5c\x84\x5c\x37\x1e\x2b\x6a\xbd\x54\x0d\xe2\x83\x4d\x1d\x72\x51\x85\x3c\x24\xc7\x80\x38\x78\x9e\x50\x07\x71\xf4\x38\x15\x31\xe4\x61\xf2\x98\xb4\x8c\xdf\x85\x3c\xaa\xb6\x79\xde\xec\x6a\x3c\xff\xc9\x49\x6a\x63\xdb\x7a\xfe\x50\x2b\xd4\x53\x4d\x0c\x06\xcf\x0d\xe6\x99\x5b\xcf\x51\x72\x96\x3c\x22\x17\x8b\xc1\xff\x4d\xc4\x73\x0e\x79\x6c\xe5\xdf\xde\x7d\x1e\x2b\xe1\x00\x71\x9d\xfd\x37\xfb\x7d\x08\x39\x46\xbf\xa7\xde\xb8\xff\xf0\x4d\x31\xbb\xaf\xc7\xce\xfd\x51\xe8\x37\xf0\x05\xbb\x89\x67\xa3\xfa\x22\x7f\xb1\x95\xdc\xaa\x55\xab\xe4\x57\x02\x33\xe7\x90\xa7\xda\xfd\x4c\x2d\x59\xcc\x3a\xcf\x13\xe2\x66\xb9\x1e\x43\xde\x8e\x5e\x5b\x8d\x70\xd2\x72\x82\x1a\x4b\xf2\xf7\xe0\xbe\x63\x0e\xf0\x9a\x5a\x04\x33\xc9\x3d\x72\x94\xbf\xeb\xc9\x73\x9a\x5c\xed\x55\x1b\x75\x1d\x72\xdd\xfa\x75\x9e\xc3\x6f\xcc\x55\x47\xc5\x63\xf0\x18\xcf\xb5\xee\x4d\x5e\x87\x03\x38\x26\x0c\xa2\x6e\x37\xbd\x70\x5b\x98\x4e\x8c\x0c\xa3\xb6\xb2\xbd\xf4\x5c\xc6\x3f\xac\xb5\x94\xdf\x88\x5d\xa3\x3e\xd4\x0b\x03\xc8\x41\xf2\xa7\x11\x96\x47\xf2\x34\x09\x33\xd4\x0f\xc8\x7b\x5b\x47\xe3\xf7\x83\xcd\xe0\xda\xac\xda\x05\x9b\xcd\x16\xe5\xb8\xe5\x76\xf2\xf8\x83\x69\x23\xb9\xa0\xba\xc1\xaf\xb1\xf6\xdf\xc9\x75\x72\x9c\x3a\x26\xf7\xf0\x1d\xe3\x8d\xb3\xff\x4e\xee\x18\x06\x4c\x9e\xb7\x65\xeb\xf8\x43\xed\x91\xb7\x55\xed\x3e\xe4\xfb\xa4\x5a\xc5\x77\xdc\x3b\x6f\x3c\x8f\xc0\x1a\xeb\x33\xbd\xf7\x51\xb0\x88\xdc\xb1\xda\xed\x95\x9f\x95\x7f\x78\xbe\x21\x26\xb3\x7a\xd6\xe8\x7e\xae\xb7\x8e\x57\xc3\xec\x39\x44\x1e\xa6\x3a\xe4\xed\xe4\xb5\x45\xee\x5b\x8f\x6e\xbc\x17\xd6\xea\x9f\x53\x15\x72\xd3\xf8\xb3\xd8\xd3\x28\x77\xb0\x31\xa9\x57\xe1\x23\x6c\x07\x83\xc0\x0f\xd6\x34\xca\xe7\xf4\x6a\xfa\x2f\xb1\x37\xec\xea\xbd\xee\xc8\x6f\x7c\x88\x9f\xc1\x79\x7a\x47\x23\x5c\x03\xdf\x06\xf9\x97\xba\xa1\x1e\x4a\x61\x4a\xaf\x18\x53\xc7\x70\x12\x30\x02\x9c\xa2\x77\xd2\x33\xa8\x09\x30\x05\xec\x86\x43\xe0\x3b\xc3\x40\x61\x4f\x6a\x3d\x7e\xe4\x18\xeb\xe8\x6b\xef\xcf\x55\xeb\x3d\x7f\xbb\xf4\xe6\xc9\xc7\x6c\x37\x8e\xe7\xc4\x63\x52\x8f\xa1\x2e\xf8\x80\xbf\xf8\x1e\x6c\x9f\x14\x6b\xea\x0b\xfb\xb7\xea\xe5\x51\xb5\x6a\x1c\x62\x72\x9f\x92\x03\xb5\xee\x65\x6c\x8b\xf3\xe4\xf7\x0c\xe2\x4c\x8c\x4f\x6d\x47\xf2\xbc\xf3\x5e\xc0\x7a\x3b\xf5\x2d\xec\xb6\x58\x24\xc7\x92\x53\x1e\xc6\xa7\x57\x6e\x30\x1e\xb5\x0a\x1e\x16\xdd\x72\xdf\xf9\x3b\x12\xf3\xb7\x91\x42\x29\xa3\xf7\x29\xce\x45\x3f\x9d\x28\xd6\xd5\xd9\xd9\x87\x71\xce\x67\xab\xb3\xb3\xf3\x0f\xda\x7a\x38\x7f\xb6\x3a\xbb\x34\x8d\xf6\x11\xab\x61\x21\xff\x6a\xb2\xee\x74\x21\xa6\xeb\xde\xaa\xe7\x8f\x71\xcf\xaf\x89\xd7\xb7\x9a\xd3\x44\xe3\x71\x92\x85\xd1\x73\x3f\xc2\xe9\xf9\xfa\xc3\x9d\x61\xfa\xef\xf9\xba\xec\x37\x3d\xdf\x4e\xb4\xc7\xf3\xf5\xb9\xdd\x80\x2c\x7a\x7e\xaa\x9a\x2e\x9a\xaa\xb8\xd4\x2f\xc8\x9d\xe7\x2e\x87\xfe\xeb\x7e\x97\x2f\xca\x67\x6b\xfb\x0d\x6b\x23\xa6\xfe\xd1\xfc\xf5\xbd\x39\xe9\xf9\x5a\xbe\x62\x1d\xcf\xed\xff\x27\x4a\x3a\x3e\xfb\x18\xf1\xf1\xe6\xe9\x90\xfe\x66\xf9\x31\x89\x76\x01\xdb\x75\x79\x22\x3f\xe6\xf7\xcb\x0f\xa0\x89\x96\x68\x90\xd2\x88\x5e\x15\x5e\xc2\xd6\xea\x36\xde\x9e\x80\x5d\xa0\x21\x15\x7e\x1f\xd0\x48\xbb\xa4\x5c\x91\x05\xc0\x0f\xd0\x06\x6d\x00\xda\x87\xe8\x14\x6c\x12\x2c\xc7\xde\x61\xd3\x60\x74\xe3\x14\x84\xb2\x4f\xa2\xa6\xb4\xf7\xb2\x11\xb5\x98\xbc\x8d\x8f\x82\x55\x5a\x3b\x94\x35\xaa\xad\x02\xbb\xc0\x25\xad\x67\xa3\xb1\x93\xa4\x0b\x50\xdd\x55\xb2\xbd\x73\x4a\x6f\x14\x71\x12\x24\x4f\x4e\x5d\x68\xc9\xd8\xd1\xd4\x4e\xef\x18\x0f\xba\xc0\x5a\xf0\x0d\x94\x87\x79\x4a\xd9\x8d\xcf\xa0\xee\x51\x92\x02\x9a\x4e\xdb\xad\x05\xcf\x6d\xe9\x90\xdb\xab\x85\xb7\x85\xd6\x8e\x4d\x5b\x87\x37\x7c\x3d\x0e\x4e\xb9\x98\x0f\x98\xc3\xd7\xc4\x05\xdb\xa2\x68\x3b\x7e\xc5\x7f\xf8\x8d\xb6\x47\x1b\xb0\x56\x38\xf8\xfd\xe4\x02\x70\x0d\x1d\x2f\xa3\xcb\x2d\xe4\x01\xd7\x68\x5f\x40\x3d\x12\x0a\x5a\x52\x08\xde\x81\x5c\xe0\x96\x16\x03\xe4\xcf\x6a\x79\xdb\x52\x90\x3d\x7b\x1b\x6e\x05\xb7\x71\xa1\x48\x8d\x53\xe4\x76\x3c\x42\x3c\x79\x05\xe5\xc2\x66\xf2\x8c\xb6\x49\x6c\x4c\xae\x54\x1e\x87\x52\x52\x6b\x56\x5e\x2d\xb4\x01\x99\xc8\x1a\xa0\x3b\x40\x37\x7e\x32\x29\x98\x3c\x56\x50\x9d\xa8\x75\x22\x6f\xf0\xcd\x26\x79\x7c\x81\x7f\xd6\x63\x74\x50\xb2\x6b\x90\xdc\xa2\xd5\x42\x39\x69\x6b\xac\x29\x4a\x6e\x14\xc9\xdb\x57\xa7\x96\xcd\xdf\x26\x11\x25\x37\x7a\xc9\x59\x68\x22\xeb\x30\x09\x55\x78\x4c\xa7\xf8\x8e\xdc\x10\x85\x82\x0a\x42\x41\xa0\xae\x50\x4b\x5a\x32\x73\xf6\x92\x58\xd8\x08\x0d\x41\x02\x45\x49\x48\xe2\x41\x3d\xe0\xcb\x24\x69\xd8\xca\xae\x49\x39\x64\xb2\x79\x74\x3f\x30\x06\xf7\x32\x97\xd1\xc5\xe8\x79\x3e\x89\x22\x93\x6f\xd0\xef\x59\x12\xd3\x28\xf1\x22\x3d\x54\xe3\xd4\x14\x3e\xb7\xad\x04\xd5\x49\x27\xc9\x67\x92\x31\x79\x5d\x91\x73\x56\x17\xa5\xd3\x17\x72\x7c\x92\x7c\x40\x22\x9b\xa4\x89\x4e\x21\x5b\x51\x48\xa8\x53\x23\x9a\x84\xc4\xdb\x8a\x8e\x32\x27\x36\x40\xcf\xa0\xef\xf8\x10\xaa\xd5\xaa\x06\x8d\xea\x94\x3e\xb6\xe5\xa2\xb0\x02\x49\x6d\x32\x65\x70\x5b\x90\xc2\xac\x7b\x92\x24\xe4\x5f\x5a\x3a\xd7\xa1\x73\xf8\x13\xbc\x30\xbb\xa3\xd7\x18\xb8\x03\xb5\x35\x29\x92\x9c\xfa\x20\x1f\xbb\x45\x46\xd7\xee\x2f\x72\x20\x49\xc2\xf3\xa9\x0a\x7f\x0e\x3c\xa2\x7e\xa9\x17\xa3\x98\x27\xd8\x44\x2c\xa1\x10\xa3\x28\xf3\x20\x7c\x41\x3a\x32\x37\x73\x4e\x8a\x1d\x6b\x2d\xf4\x2f\xd8\xf5\xa3\x1a\x11\xcd\x67\x9d\x3c\x0b\x25\x83\x2e\x22\xe1\xc8\x5d\xa3\xa5\xd1\x71\x09\x3f\x52\xc3\x69\xd9\x76\x28\x3d\xbf\xcc\x77\x85\xaf\x71\x52\xae\x98\x64\xea\xbd\x57\xe0\x5f\x93\x31\x93\xe7\x45\xb7\xd4\x56\xe5\x79\x6e\x12\xa4\xf0\x7a\x27\x86\xf8\x96\x67\xac\xee\x85\x5f\x96\xf3\x95\xe7\x0d\xb2\xc0\x24\xc9\xe8\x92\xb1\x10\x4e\x10\x0f\x70\xbc\x15\x05\xb3\x9e\xd0\x4b\x86\x4e\x6e\x03\x7e\x2b\xe5\x27\xdb\xb2\x12\x0d\x26\x77\x5b\xcd\xcf\xda\xec\x37\xe5\x21\x38\x5b\x08\x5f\xc0\x0b\x62\x66\x78\xbc\x71\x3b\xe8\x0d\x85\xa8\x29\x6b\x33\x89\x3d\xf9\xbc\x26\x87\x7b\x5f\x03\x7d\xcb\xae\x15\x1e\x7b\xb3\x25\x79\xbd\xe0\x6b\xe6\xb3\x71\x24\x63\x8c\x26\x6b\x7c\x8b\x5d\xed\xfe\x45\xf6\x80\x43\x46\x53\x25\x47\xc8\x0b\x70\xa9\xd2\x16\x04\x32\xa2\x15\x16\x83\xed\xc4\xc8\xb6\xe2\x36\xee\xa7\x05\x93\xf0\x3d\xd7\x58\xbf\x6d\x85\xd4\x6e\x33\x7e\x05\xcb\x7a\xc5\x9a\xbc\x05\x93\x4d\x4a\x54\x92\xd0\x83\xd7\x14\x39\x08\x46\x91\x43\xf6\x99\xbc\x86\xac\xcf\x95\x92\xe9\x9d\xcf\x45\xbf\x25\x3f\x91\x92\xf8\x90\x9c\xb6\xfa\x92\xfc\x41\x9a\xe1\x6b\xcb\x3f\xc5\x1a\x7c\xb7\x6d\xc6\xad\x7a\xda\xe8\x3e\xa2\x5f\x32\x0e\x39\x06\xa6\x80\x89\xc8\x0f\xeb\x97\x92\x7c\x85\xe4\xd0\xb2\xed\x67\x92\xb1\xf1\xb5\x18\xb6\x4a\x36\x81\xa5\xe4\xce\x28\x19\x42\xfe\xe2\xef\x56\xdb\x0c\xd4\x09\xb6\xc1\x4b\xc8\x4f\xf2\x17\xbc\x6d\xb5\x55\x48\x0f\xb5\xe7\x3b\x6d\x2b\x0c\x92\x15\x92\x76\x83\xa4\x10\xd7\xde\xca\x2c\xd5\x21\x38\xc7\x7d\xd8\x00\x8e\x6c\x17\xd9\x94\x54\x97\x51\xe3\x0b\x43\x6c\xeb\x28\xf9\xf6\x8c\xe1\xe9\x89\x24\x19\x4e\x24\xc9\xa8\x31\xde\x27\x49\x88\x3d\xb8\xdb\x48\x5a\x53\x13\x47\x4e\xf7\xa1\x92\x64\xa1\x8a\x7f\x37\x51\xb2\x4c\xf0\x01\xb2\xe4\xed\x01\xe7\x47\x0a\x93\x65\xca\xdf\x51\x9a\xbc\xe3\xa6\x7f\xb8\x38\xf9\x91\x4b\x24\x4f\xaa\xa2\x2b\xff\x19\xe5\x89\x9f\x46\xff\xcd\xfa\x84\x1a\xe4\x63\xfa\x64\xf0\xba\x34\xe9\x9f\x7c\x7b\x06\x5e\x54\x4b\x63\x34\x85\xb8\xf3\xe8\x7c\x0f\x5c\xb5\xad\x34\x61\xca\xa8\xad\x22\xfa\x09\xf8\x49\xcf\xac\xd4\x8b\x3b\xdd\x0f\x96\xc0\x13\xa9\xe1\x51\x18\x0d\x5e\xd0\xbf\x5b\x71\xf0\xb2\xf5\xfb\x06\xf1\xbf\x79\x3a\xea\x94\x59\x6b\xa3\xcf\xd8\x36\x68\x72\xbc\xa3\x2f\x34\xc2\x64\x34\x03\x3d\x07\x5c\x68\x85\xb5\x8d\x38\xfe\x76\xd1\x5b\x83\x8f\x69\x47\x28\xa3\xf7\x5f\x30\x24\x9d\x1c\x8b\x0c\xcb\xb1\x48\xa5\x9e\x9a\x1c\xa7\x5a\x1d\x85\xe0\xbf\xe5\x48\x04\xac\x1c\xa5\x7f\xe0\xdc\xf0\x2a\xc3\x28\x71\x2b\xdb\x3a\x6d\x9c\x57\xd1\xe7\xb0\x8d\xfb\x6c\x2b\xb9\x72\xbc\x2a\x97\xfe\xab\x63\x88\x41\xdb\xf2\xcc\xbd\x70\x02\x3b\x86\xe9\xc4\x9b\x74\x5c\x32\x0b\xcf\x6d\x3b\x31\x79\xbf\x84\x4b\x6e\x0b\xb7\xb7\xd0\x31\x0a\xf3\x74\xe2\xc6\xf4\x62\xb0\x91\xde\x83\xef\x4b\x7d\xc7\x4f\x86\xe1\x85\xc7\x34\x8a\x27\xd8\x36\x6f\x7b\xdc\xb2\xc3\x47\xc5\xa2\x3d\xb4\xdd\x87\x36\xc2\xa7\x85\xb8\x22\x31\xb0\xad\xf0\xd1\xc7\x2e\x27\x9f\x77\x14\x5f\x4d\xda\xae\xdf\x48\xab\x0d\xd2\x20\x76\x34\x94\xa4\x09\xa2\xd6\xdc\x4b\x8f\x8d\x47\x9f\x9a\x86\x9d\x3d\xe7\xc8\xf9\xcd\x49\x1f\x37\xdd\x5d\x68\x8b\x7f\x72\xbf\x17\xfa\x0e\x4f\x35\x9b\xb4\x9d\xdd\x6b\x1b\xbf\x11\x67\x82\x7b\xc2\x25\xc7\xa5\xef\x8e\xce\x2d\xe8\x27\xf0\xbd\x79\x39\xde\xe8\xf4\x89\xc7\x6d\x72\x7c\x6d\x5b\xd1\xbd\xf8\x6e\xd2\x36\x68\x14\xb7\x6c\x74\x84\x56\x7a\xae\xda\xf1\xd3\x72\x2c\x22\x7d\x42\x2e\x2c\xc7\x22\x76\x6c\x10\x5d\x1b\x19\x07\x17\x57\x85\x03\xc0\x41\xb6\xd2\x6a\xd4\xe7\x56\x47\x1b\xd4\x9f\xf1\xc9\xd6\xe7\x8d\x51\xf5\x3e\x88\x0b\x6a\x6c\xd6\xd4\x28\xb7\x59\x07\x5c\x0c\x3f\x31\x8f\x1d\x4d\xa8\xcf\xa3\x69\xc9\xe5\x51\x7b\x13\x70\x60\xe6\x20\xa7\x6c\x1b\x52\x5b\xe1\xe8\x10\xea\xd0\x34\xcf\xc6\xfd\x60\xb8\xd2\x7b\xdc\xe0\x8e\xf8\xd7\x74\x52\xe3\xfd\x1c\xfe\x80\x7e\x80\x77\x4c\xda\x26\x6c\xa4\x4d\x89\x1b\xbe\x21\xf7\x6d\x3b\xbe\xf5\x75\xe2\x23\xfc\x0b\x07\x24\xbf\xb0\x0b\x7b\xb0\x65\x38\x39\x22\x64\xbd\x95\xb8\xfc\xac\xbd\x96\x59\x7c\x1e\x4d\x05\x3e\xe1\x43\xea\xcc\x74\x47\xe9\xf9\x88\xfd\xdb\xf9\xc8\x91\xed\xa8\x54\x47\x84\x9d\xf6\x68\x8c\x0b\xce\xce\x4b\x19\xcb\xb4\x66\xe5\x58\x58\x4a\x8f\x93\xe3\x9d\xb4\x4d\x55\x1c\xb9\x6f\x12\x0f\xc3\xa7\x86\xad\xa3\xdb\x88\x0e\x35\x2d\xab\xe3\xb2\x52\x7b\x2c\xc6\xb5\xaa\xe3\xd6\xf4\xbc\x1c\x0b\x6b\xcf\xc3\xb8\xff\xec\x39\x0d\x97\x02\x53\xe0\x66\xe0\x84\x6d\x11\x37\x9e\xfb\xc4\xb1\x90\x9d\xe6\xd7\xe1\xa8\x9d\x5a\xf1\x4c\xdb\x17\x90\xc6\x04\xe3\x93\x72\x9c\x9a\x62\x2c\x9e\xa7\x3e\xcd\x4f\x8d\xd7\xb4\xf1\xc9\xca\xeb\x7a\xa9\xd5\xa5\x7e\xf0\x79\x25\x5c\x81\xdb\x75\xd2\x39\xe4\x6c\xa3\x23\x4c\x7c\xd3\x6b\x3f\x08\x0e\xd8\xe9\x58\x11\xbd\x03\x9e\xd8\x3e\x81\x6a\x35\xea\x48\x96\x5c\x02\x7f\x4d\x83\xea\x88\x1c\x7c\x1b\x85\xe5\xad\x34\xfa\xe2\x17\x62\x62\x47\xb1\xda\xe3\x60\x5d\xd8\xd3\x0b\x37\x37\xe2\xd5\xa3\xb4\x71\xbf\x1c\x69\x8f\xda\xd3\x50\x8e\xc1\xe1\x3b\xe9\xa8\x05\x8b\xed\x88\x63\x50\x0d\xe8\x68\xc1\xf8\x74\xeb\xfa\xde\xf8\xaf\x8e\x0f\x36\xd2\xbf\xcb\x31\xae\xf5\x98\xe8\xeb\xa8\x75\x74\x62\xc7\x5e\x83\x8f\x6b\x7b\x01\x3a\xee\x07\xab\x2d\x67\xb5\xbd\x4f\x7d\x91\x63\xe0\x75\x2b\xcc\x1d\x85\x13\x51\x7d\x79\x5e\x8e\xea\x93\xe3\x4a\xa1\x63\x11\x70\x1c\x3b\x93\xe6\x21\x37\x59\x33\x39\x5f\x2b\x37\xc8\xcb\x46\xf9\x55\xe8\xd5\x05\xea\x90\x3a\xb3\x7e\xfb\x9e\xa3\x93\x65\x0f\xa6\xd4\x51\x1d\x36\xe1\xe7\xa6\xf1\x7b\xa8\xed\x51\x7c\xdb\xf6\x58\x2a\xf7\x2d\x39\x6e\x7a\x64\x76\x9d\x31\x4e\xbe\x76\xeb\x49\x5b\x5f\x67\xd5\xaa\x4f\x08\x57\x6b\xed\xc1\x2d\x47\x27\x1b\x1d\xa5\xa5\x51\x7d\x4c\xaf\x3b\xb0\x16\xc3\x53\xbd\x72\x31\x69\xaf\xa5\x9e\x4f\x8e\x3e\xb6\xbe\x7e\xae\x77\x8b\xdf\xa2\xef\xa1\x2d\x47\x27\xdb\xe4\x63\x50\xfb\xcc\xbf\x51\x1d\x93\x7f\xa3\x8e\x7c\xac\xdf\xce\x9e\x77\xe4\x1f\xf9\x4b\xbf\xa1\x0e\x6a\xe9\x55\x9e\x1b\xf4\x5a\xc4\xb2\x97\x85\xbe\xa9\x87\x9f\xea\x94\x59\x47\xba\x3c\x47\x6e\x58\x0d\x7f\xac\x4e\x39\x72\xc6\xbf\x97\x50\x39\xce\xf0\xdb\x95\xca\xc9\x7b\x95\x1f\x27\x55\x8e\x93\xfe\x7e\x5a\xe5\x27\xae\xfa\x47\x8b\x95\x77\xbc\x72\x54\x2b\xcd\x3f\xa3\x5a\xd1\x7b\xb0\xbf\xdb\xdb\x5c\x40\x6a\x7d\xf2\x36\xd7\x7c\xfa\x36\x97\x68\x3b\xf0\x1a\x55\x5e\xf6\x06\x47\xe9\xf0\x6b\x6f\xd8\x4c\x0e\xeb\x06\x5b\x9d\xb7\x06\xa3\x65\x83\xd3\x06\x5a\x24\xd0\x0c\x94\x0c\x6a\xeb\xa3\xe8\x17\x6d\x1b\x48\xb6\xed\xa5\xc6\xb7\xcd\xa2\x4e\xd1\x6d\x5b\x54\xdb\x53\x9b\xd3\xb7\xb7\xf4\x76\xc9\x02\xc3\x95\xde\x0a\xea\x74\x42\x0b\xc5\x48\x7a\x33\x04\xa8\xa1\x7d\x00\xc9\x85\x8e\x3f\x6a\x49\x05\x83\x48\xd1\x7f\xda\x67\xad\x37\x2f\x1a\xd1\x25\x6b\xcb\xa2\xd8\xbd\xde\x9e\x99\x34\xa6\xad\x5f\x7e\x88\x7a\x4b\x67\xa1\x14\x5b\x6d\xd1\xd9\xf1\x46\xab\xe3\x92\xf2\x48\xf3\x07\xbd\xf5\x64\x5b\x2c\xa2\xb9\x8d\x60\xb6\x56\xbb\xc5\xc6\x51\x6b\xa2\x95\x9b\x64\xea\x7c\x3e\xa3\x8f\x5b\xdf\x06\xb3\x37\xde\xf4\x76\xdb\xe9\xd1\xca\xf2\xf6\xc3\xa4\x6d\xc9\x9f\x3b\x5a\x19\x74\xb4\x32\xe9\x68\xa5\x7e\xdf\xd1\x8a\xb6\xc8\x87\x77\x8e\x56\x3a\x49\x40\x62\x55\x9e\x1c\xad\x10\x2b\xec\xc3\x6f\x16\x43\xc9\x9a\xa4\xb7\xa2\x36\x3f\xf3\x26\x57\xaf\x6d\x2c\xfc\xc4\x33\xb4\x1f\x72\xce\xe4\x98\xb6\x1e\x19\x93\x16\x80\x1c\xeb\x75\xac\x60\x54\x55\xf9\x64\x6f\x1e\x8d\xa2\x08\x93\xaf\xc3\x5a\x54\x7d\xac\x01\xe6\xa4\x9d\x22\x81\x91\x5f\x51\x5b\xf1\xb4\x3f\x72\xb1\xd6\x71\x01\xf9\x4f\x8e\x8f\x3a\x1a\x9c\x25\x5f\x46\xbd\xad\x80\x14\x46\x02\x46\xe5\x1b\x73\x45\x1d\x93\x90\x2f\xb4\xb0\xa4\x2d\x45\xee\x81\x0a\x35\x7a\x7b\xa5\xac\xbc\x0e\x89\xfd\xb4\xbc\xa1\xd7\xea\x58\x22\x7a\xfd\x40\x9f\x7a\x6d\x1b\x98\xb4\xc4\x57\xf3\x5f\x03\x00\x00\xff\xff\x5c\x1a\xe7\xdc\x00\x30\x00\x00")

func bindataPkgV1TkgManifestTelemetryZzgeneratedBindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataPkgV1TkgManifestTelemetryZzgeneratedBindataGo,
		"tkg/manifest/telemetry/zz_generated.bindata.go",
	)
}

func bindataPkgV1TkgManifestTelemetryZzgeneratedBindataGo() (*asset, error) {
	bytes, err := bindataPkgV1TkgManifestTelemetryZzgeneratedBindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "tkg/manifest/telemetry/zz_generated.bindata.go",
		size:        28672,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

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
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

// AssetNames returns the names of the assets.
// nolint: deadcode
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"tkg/manifest/telemetry/config-aws.yaml":         bindataPkgV1TkgManifestTelemetryConfigawsYaml,
	"tkg/manifest/telemetry/config-azure.yaml":       bindataPkgV1TkgManifestTelemetryConfigazureYaml,
	"tkg/manifest/telemetry/config-docker.yaml":      bindataPkgV1TkgManifestTelemetryConfigdockerYaml,
	"tkg/manifest/telemetry/config-vsphere.yaml":     bindataPkgV1TkgManifestTelemetryConfigvsphereYaml,
	"tkg/manifest/telemetry/zz_generated.bindata.go": bindataPkgV1TkgManifestTelemetryZzgeneratedBindataGo,
}

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
				return nil, &os.PathError{
					Op:   "open",
					Path: name,
					Err:  os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"pkg": {Func: nil, Children: map[string]*bintree{
		"v1": {Func: nil, Children: map[string]*bintree{
			"tkg": {Func: nil, Children: map[string]*bintree{
				"manifest": {Func: nil, Children: map[string]*bintree{
					"telemetry": {Func: nil, Children: map[string]*bintree{
						"config-aws.yaml":         {Func: bindataPkgV1TkgManifestTelemetryConfigawsYaml, Children: map[string]*bintree{}},
						"config-azure.yaml":       {Func: bindataPkgV1TkgManifestTelemetryConfigazureYaml, Children: map[string]*bintree{}},
						"config-docker.yaml":      {Func: bindataPkgV1TkgManifestTelemetryConfigdockerYaml, Children: map[string]*bintree{}},
						"config-vsphere.yaml":     {Func: bindataPkgV1TkgManifestTelemetryConfigvsphereYaml, Children: map[string]*bintree{}},
						"zz_generated.bindata.go": {Func: bindataPkgV1TkgManifestTelemetryZzgeneratedBindataGo, Children: map[string]*bintree{}},
					}},
				}},
			}},
		}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
