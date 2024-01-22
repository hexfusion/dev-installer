// Code generated for package template_assets by go-bindata DO NOT EDIT. (@generated)
// sources:
// bindata/templates/installer/aws/install-config.yaml
// bindata/templates/installer/azure/install-config.yaml
// bindata/templates/installer/gcp/install-config.yaml
// bindata/templates/installer/libvirt/install-config.yaml
package template_assets

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

// Mode return file modify time
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

var _bindataTemplatesInstallerAwsInstallConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x4f\x6b\x1b\x31\x10\xc5\xef\xfa\x14\x73\xcb\xc9\x5b\xaf\x53\xd2\xa2\x6b\x1b\x28\x98\xa6\xa6\x0e\xe9\x79\xac\x7d\xb1\x44\xf5\x8f\x91\xd6\x89\x71\xf7\xbb\x17\xad\x83\x89\x4f\x85\x5e\x9f\xe6\xbd\xdf\xbc\x41\x9c\xdd\x13\xa4\xb8\x14\x35\x1d\x7a\xb5\xe3\x82\xaf\x29\xb0\x8b\x9a\x30\xec\xb1\x28\x92\x3a\xb1\x30\x09\x71\xdf\x99\x14\x94\x49\x21\x8f\x15\x5a\x2d\x88\xc5\x58\x57\x61\xea\x28\xd0\xc4\x61\xb8\xfb\xa8\x88\xec\x31\x43\xaa\x15\xf0\xe0\xe2\x5e\xd3\x7d\xe4\x9d\xc7\xa0\x88\x22\x07\x68\x7a\x49\xf2\x1b\xa2\x88\xb2\xe7\xfa\x9c\x24\x68\x3a\x4d\x8a\x48\x90\xbd\x33\x5c\x34\x9d\x4e\xd4\xfd\x9a\xa7\x7e\xbe\x69\x34\x4d\xca\xa4\x58\x25\xf9\x8d\xe7\x08\xad\xe8\x3f\xe9\x81\x4b\xfd\x37\xfd\xfb\x3c\xf5\x9e\x1e\x50\x79\xe0\xca\x8d\x6c\x04\x5c\x5d\x8a\x8f\x2e\xa0\x54\x0e\x59\x53\x1c\xbd\xbf\x30\x5a\xc2\x17\x3f\xb6\x88\x07\x0e\x68\xf6\x88\xda\x7a\xb7\x95\x5a\xc0\xdb\xe3\x59\x6c\xca\x82\x8c\x1b\x44\x53\xbf\xec\xfa\xd5\xe7\x6e\xd9\x2d\x3f\xf4\xad\x0e\x91\x4d\xa5\x6e\x04\xcf\xee\x55\xd3\xea\x56\x11\x05\x36\xd6\x45\xbc\x33\x5f\xd9\x97\x67\xf3\x5d\xdb\xe6\x3c\xf2\x78\xcc\xd0\xf4\xe3\xe9\x61\x3d\xee\x20\x11\x15\x45\x11\x15\xc8\xc1\x19\x5c\xed\xd0\x7f\x5a\x75\xb7\x97\x80\xcb\x85\xda\xb1\x5f\xca\x99\x24\xd8\xcf\x7f\xa5\x75\xdc\x48\x3a\xb8\xa1\xdd\xa9\x69\xad\x66\x1e\x77\xde\x15\xab\xe9\xfe\xb5\x42\x22\x7b\x95\x47\xef\xb7\x30\x82\xaa\xe9\x66\x36\x5d\x04\x9a\xa6\x1b\x55\x8a\x5d\xe3\xa8\xe9\x8f\xa2\x39\x73\xbb\xfd\xb6\xc6\xb1\x65\xfd\x0d\x00\x00\xff\xff\x03\x74\xff\x64\x9a\x02\x00\x00")

func bindataTemplatesInstallerAwsInstallConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesInstallerAwsInstallConfigYaml,
		"bindata/templates/installer/aws/install-config.yaml",
	)
}

func bindataTemplatesInstallerAwsInstallConfigYaml() (*asset, error) {
	bytes, err := bindataTemplatesInstallerAwsInstallConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/templates/installer/aws/install-config.yaml", size: 666, mode: os.FileMode(420), modTime: time.Unix(1705947791, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataTemplatesInstallerAzureInstallConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xb1\x6e\x1b\x31\x0c\xdd\xf5\x15\xdc\x32\xf9\x80\x02\x9d\xb4\xa6\x45\x03\x04\x35\x8c\x38\x68\x67\x46\xa2\x7d\x42\x25\x51\xa0\x74\x6e\xdd\xeb\xfd\x7b\x41\x9d\x63\x37\xe8\x90\xf5\xe9\x3d\xbe\xf7\x28\x62\x09\xdf\x48\x6a\xe0\x6c\xe1\xf4\xc1\xbc\x60\xa5\x4f\x9c\x30\x64\x0b\x0e\x9b\x1b\x31\xc6\x01\x7f\x4f\x42\x83\xa7\x93\x8b\x53\x6d\x24\x03\x17\xca\x75\x0c\x87\x36\x38\x4e\xc6\x71\x2a\x53\x23\x6b\x36\x30\x9e\x0b\x49\x1b\x85\xd0\x87\x7c\xb4\xf0\x39\xe3\x4b\x24\x6f\x00\x32\x26\xb2\xf0\x93\xe5\x07\x89\x01\x28\x11\xdb\x81\x25\x59\x98\x17\x03\x20\x54\x62\x70\x58\x2d\xcc\x33\x0c\xdf\x3b\xeb\xe9\x82\xc1\xb2\x18\xc7\xb9\x09\xc7\x5d\xc4\x4c\xd6\xc0\xbb\x3e\x09\x35\xe7\x7b\x3e\x5f\x3b\xeb\x5f\x9f\x44\x0d\x3d\x36\x54\x0f\x27\x84\x2d\x70\x7e\x0e\x89\x6a\xc3\x54\x2c\xe4\x29\xc6\xab\x87\x4e\xb8\x5f\x17\xb2\xc5\x44\x2a\xcf\xd4\xb4\xa1\x46\xd2\x01\x97\xc7\x15\x54\x64\x03\x2e\x78\x79\x23\xbd\x0f\x5e\x54\x0a\x00\x30\x72\x6d\x3b\xa1\x43\xf8\xf5\x1f\xe5\xe1\xfa\xb4\x92\x13\xba\x31\x64\x7a\x9d\x0d\xf3\x2c\x98\x8f\xa4\x9d\xfa\x83\x6a\x3a\xf1\x66\x39\x2c\xcb\x3c\x53\xf6\x1d\xbe\x24\x7d\x3e\x97\x4b\x93\xed\x0d\x58\x1d\x2a\xc9\x29\x38\x7a\x93\x5e\x89\xfb\x15\x7f\xcd\x7d\x5d\xb0\x01\xe8\x77\x62\x7b\x95\xdb\x1d\x3d\x51\xe5\x49\x1c\x7d\x11\x9e\xca\xb6\x6f\x8e\xeb\xc7\x8d\xe3\x94\x38\x77\xae\xd0\xb1\x9f\x9f\x4e\xdf\x09\x9f\x82\xd7\x4f\x51\xac\x1b\x4c\x31\xee\xc9\x09\x35\x0b\x77\x9d\x72\x05\x60\x59\xee\x4c\xad\xe3\x23\x9d\x2d\xfc\x31\xb0\xe6\xdb\x3f\x3c\xd2\x59\x95\x7f\x03\x00\x00\xff\xff\xd7\x51\x47\xfc\xdb\x02\x00\x00")

func bindataTemplatesInstallerAzureInstallConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesInstallerAzureInstallConfigYaml,
		"bindata/templates/installer/azure/install-config.yaml",
	)
}

func bindataTemplatesInstallerAzureInstallConfigYaml() (*asset, error) {
	bytes, err := bindataTemplatesInstallerAzureInstallConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/templates/installer/azure/install-config.yaml", size: 731, mode: os.FileMode(420), modTime: time.Unix(1704128341, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataTemplatesInstallerGcpInstallConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\x4f\x6f\x13\x31\x10\xc5\xef\xfe\x14\xef\xd6\x53\x96\x6c\x8a\x00\xf9\xda\x20\x51\x55\x94\xa8\x5b\xc1\x79\xea\x9d\x64\x0d\xfe\xa7\xb1\xb3\x10\x85\xfd\xee\xc8\x9b\x6a\x45\x4f\xb9\x3e\xbf\x79\xbf\xf9\x63\x4a\xf6\x3b\x4b\xb6\x31\x68\x8c\xad\x7a\xa1\xcc\xdb\xe8\xc9\x06\x8d\x83\x49\x4d\xcf\xa3\x71\xc7\x5c\x58\x9a\x98\x38\xe4\xc1\xee\x4b\x63\xa2\x57\x26\xfa\x74\x2c\xac\xd5\x0a\xc3\x29\xb1\x94\x41\x98\x7a\x1b\x0e\x1a\x9f\x03\xbd\x38\xee\x15\x10\xc8\xb3\xc6\xef\x28\xbf\x58\x14\x90\x1c\x95\x7d\x14\xaf\x71\x9e\x14\x20\x9c\x9c\x35\x94\x35\xce\x67\x34\x3f\x66\xd7\xd3\xab\x86\x69\x52\x26\x86\x22\xd1\xed\x1c\x05\xd6\x0a\x57\x39\x9e\x6a\x9f\xd7\x38\x5f\x67\xd7\xff\x1c\xcf\x85\x7a\x2a\x54\x19\x46\x98\x8a\x8d\xe1\xd9\x7a\xce\x85\x7c\xd2\x08\x47\xe7\x16\x46\x4d\xb8\xbb\x2c\xe4\x91\x3c\xd7\xf2\xc0\xa5\x4e\x58\x5b\xaa\x01\xaf\x8f\x17\xb1\x2a\x2b\x18\xdb\x8b\x46\xbb\x6e\xda\xcd\xa7\x66\xdd\xac\xdf\xb5\xef\x15\x00\x0c\x31\x97\x9d\xf0\xde\xfe\xd1\xd8\xdc\x2a\xc0\x93\x19\x6c\xe0\xbb\xfb\xed\xd3\xec\x5f\x5f\xdc\x1f\x2a\xfe\x12\xf8\x7c\x4a\xac\xf1\x2d\x71\xe8\xea\x29\xba\xed\xa3\x02\x32\xcb\x68\x0d\xbf\x61\xb6\x1f\x37\xcd\xed\x52\xbf\x6c\x44\xa1\x9e\x55\xcf\xf8\x24\xf1\x27\x9b\x72\xbf\xd5\x58\x6e\xbb\x3a\x18\x5e\xf5\x3c\xb2\x9b\x2d\xc2\x87\xf9\x67\xd4\xb1\x77\x12\x47\xdb\xd7\xd5\x55\xad\x4e\x9e\x8e\xce\x75\x6c\x84\x8b\xc6\xcd\x6c\x59\x04\x4c\xd3\x8d\xca\x79\x78\xe0\x93\xc6\x5f\x85\x39\xa1\xeb\xbe\x3c\xf0\xa9\x56\xfe\x0b\x00\x00\xff\xff\x41\xdb\x33\x30\x76\x02\x00\x00")

func bindataTemplatesInstallerGcpInstallConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesInstallerGcpInstallConfigYaml,
		"bindata/templates/installer/gcp/install-config.yaml",
	)
}

func bindataTemplatesInstallerGcpInstallConfigYaml() (*asset, error) {
	bytes, err := bindataTemplatesInstallerGcpInstallConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/templates/installer/gcp/install-config.yaml", size: 630, mode: os.FileMode(420), modTime: time.Unix(1704128341, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataTemplatesInstallerLibvirtInstallConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\xcd\x6e\xdb\x30\x10\x84\xef\x7c\x8a\xb9\xe5\x14\x45\x52\x0a\x37\xe5\x35\x2e\x50\x23\x6d\x6a\xd8\x49\x7b\x5e\xcb\xab\x8a\x28\xff\x40\x52\x6e\x0c\x57\xef\x5e\x90\x32\x8c\x34\x40\xae\xb3\xbb\xdf\xa7\x11\xc9\xab\x1f\x1c\xa2\x72\x56\xe2\xd0\x88\x1d\x45\x5e\x3a\x43\xca\x4a\x0c\xfc\xd2\x8f\x79\x52\x69\xd7\x91\x16\x9d\x33\x7e\x4c\x2c\x05\x70\x8d\xe1\xe8\x39\xa4\x21\x30\xed\x95\xfd\x25\xf1\xd9\xd2\x4e\xf3\x5e\x00\x80\x25\xc3\x12\x7f\x5c\xf8\xcd\xa1\x04\x5e\x53\xea\x5d\x30\x12\xa7\xa9\x04\x81\xbd\x56\x1d\x45\x89\xd3\x09\xd5\xcf\xb2\xb9\x39\x67\x98\x26\xd1\x39\x9b\x82\xd3\x6b\x4d\xb6\xf8\xde\xb7\xcd\x2e\x43\x31\x15\xd7\x1b\xd3\xff\x9e\x6f\x65\xeb\xb5\xc7\x70\xa2\x3d\x25\xca\x8e\x2e\x30\x25\xe5\xec\x93\x32\x1c\x13\x19\x2f\x61\x47\xad\x2f\x8e\x4c\xb8\xd7\x63\x46\x3c\x92\xe1\x7c\x6e\x39\xe5\x96\xf9\x93\x32\xe0\x3c\x9c\x43\x59\x8a\x5e\xa3\x53\xfb\x20\xd1\xd4\x55\xd3\xde\x55\x75\x55\xdf\x34\x1f\xca\x04\x18\x5c\x4c\xeb\xc0\xbd\x7a\x91\x68\x6f\x05\x60\xa8\x1b\x94\xe5\xfb\xd5\x72\x23\xd1\x7c\x6a\xab\x66\x71\x57\x35\xed\xa2\xaa\x6f\xda\x7c\x74\xd6\x3d\x1d\x3d\x4b\x7c\xf7\x6c\xb7\x83\xea\xd3\x76\xf9\x28\x80\xc8\xe1\xa0\x3a\x7e\x23\x6f\x3e\xb6\xd5\x6d\x3d\x6b\x17\xe2\xf2\x73\x04\xa0\xd5\xee\xa0\x42\x9a\x17\x9f\x37\xab\xb9\xdf\xd7\x39\x7d\xde\xac\x72\xbd\xf2\x96\xaf\x81\x80\xea\x25\x52\xaa\x85\x1f\xb5\xde\x72\x17\x38\x49\x5c\xe5\xcb\xf5\x25\xc0\x34\x5d\x89\x18\x87\x07\x3e\x4a\xfc\x15\x28\xe0\xed\xf6\xcb\x03\x1f\x33\xf4\x5f\x00\x00\x00\xff\xff\x1b\xca\x10\xe2\x71\x02\x00\x00")

func bindataTemplatesInstallerLibvirtInstallConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesInstallerLibvirtInstallConfigYaml,
		"bindata/templates/installer/libvirt/install-config.yaml",
	)
}

func bindataTemplatesInstallerLibvirtInstallConfigYaml() (*asset, error) {
	bytes, err := bindataTemplatesInstallerLibvirtInstallConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/templates/installer/libvirt/install-config.yaml", size: 625, mode: os.FileMode(420), modTime: time.Unix(1704128341, 0)}
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
	"bindata/templates/installer/aws/install-config.yaml":     bindataTemplatesInstallerAwsInstallConfigYaml,
	"bindata/templates/installer/azure/install-config.yaml":   bindataTemplatesInstallerAzureInstallConfigYaml,
	"bindata/templates/installer/gcp/install-config.yaml":     bindataTemplatesInstallerGcpInstallConfigYaml,
	"bindata/templates/installer/libvirt/install-config.yaml": bindataTemplatesInstallerLibvirtInstallConfigYaml,
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
	"bindata": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"installer": &bintree{nil, map[string]*bintree{
				"aws": &bintree{nil, map[string]*bintree{
					"install-config.yaml": &bintree{bindataTemplatesInstallerAwsInstallConfigYaml, map[string]*bintree{}},
				}},
				"azure": &bintree{nil, map[string]*bintree{
					"install-config.yaml": &bintree{bindataTemplatesInstallerAzureInstallConfigYaml, map[string]*bintree{}},
				}},
				"gcp": &bintree{nil, map[string]*bintree{
					"install-config.yaml": &bintree{bindataTemplatesInstallerGcpInstallConfigYaml, map[string]*bintree{}},
				}},
				"libvirt": &bintree{nil, map[string]*bintree{
					"install-config.yaml": &bintree{bindataTemplatesInstallerLibvirtInstallConfigYaml, map[string]*bintree{}},
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
