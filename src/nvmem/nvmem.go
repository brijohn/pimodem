package nvmem

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

type Device struct {
	Device string
	Size   int
	Offset int
	Cells  map[string]Cell
}

type Cell struct {
	Offset int
	Len    int
	Type   string
}

type NVMEM struct {
	file   *os.File
	offset int
	size   int
	cells  map[string]Cell
}

var bus string = "/sys/bus/nvmem/devices"
var Config string

var Devices map[string]Device

func init() {
	if Config != "" {
		data, err := ioutil.ReadFile(Config)
		if err == nil {
			yaml.Unmarshal(data, &Devices)
		}
	}
}

func RegisterDevice(name string, device Device) error {
	if Devices == nil {
		Devices = make(map[string]Device)
	}
	if _, ok := Devices[name]; ok {
		return fmt.Errorf("Device %s is aleady registerd", name)
	}
	Devices[name] = device
	return nil
}

func Open(name string) (*NVMEM, error) {
	var nvmem NVMEM
	device, ok := Devices[name]
	if !ok {
		return nil, fmt.Errorf("Device %s has not been registerd", name)
	}
	file := path.Join(bus, device.Device, "nvmem")
	nvmem.cells = device.Cells
	nvmem.offset = device.Offset
	nvmem.size = device.Size
	f, err := os.OpenFile(file, os.O_RDWR, 0755)
	if err != nil {
		return nil, err
	}
	nvmem.file = f
	return &nvmem, nil
}

func (nv *NVMEM) ReadAt(p []byte, offset int64) (int, error) {
	length := int64(len(p))
	offset += int64(nv.offset)
	if offset >= int64(nv.size) {
		return 0, nil
	}
	if length+offset > int64(nv.size) {
		length = (length + offset) - int64(nv.size)
	}
	l, err := nv.file.ReadAt(p, offset)
	return l, err
}

func (nv *NVMEM) WriteAt(p []byte, offset int64) (int, error) {
	length := int64(len(p))
	offset += int64(nv.offset)
	if offset >= int64(nv.size) {
		return 0, nil
	}
	if length+offset > int64(nv.size) {
		length = (length + offset) - int64(nv.size)
	}
	l, err := nv.file.WriteAt(p, offset)
	return l, err
}

func (nv *NVMEM) Clear() error {
	zero := make([]byte, nv.size)
	_, err := nv.WriteAt(zero, 0)
	return err
}

func (nv *NVMEM) Close() error {
	return nv.file.Close()
}
