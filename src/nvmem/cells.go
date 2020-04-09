package nvmem

import (
	"encoding/binary"
	"fmt"
	"net"
)

type FromBytes func([]byte) (interface{}, error)
type ToBytes func(interface{}) ([]byte, error)

type CellType struct {
	from FromBytes
	to   ToBytes
}

var CellTypes map[string]CellType

func init() {
	RegisterCellType("bytes", nil, nil)
	RegisterCellType("string", StringFromBytes, StringToBytes)
	RegisterCellType("uint16", UInt16FromBytes, UInt16ToBytes)
	RegisterCellType("ipv4", IPv4FromBytes, IPv4ToBytes)
}

func IPv4FromBytes(value []byte) (interface{}, error) {
	if len(value) != 4 {
		return nil, fmt.Errorf("length of []byte must be equal to 4")
	}
	return net.IPv4(value[0], value[1], value[2], value[3]), nil
}

func IPv4ToBytes(value interface{}) ([]byte, error) {
	ipv4, ok := value.(net.IP)
	if !ok {
		return nil, fmt.Errorf("parameter is of type %T, must be a net.IP", value)
	}
	ret := ipv4.To4()
	if ret == nil {
		return nil, fmt.Errorf("parameter is not an IPv4 address")
	}
	return ret, nil
}

func UInt16FromBytes(value []byte) (interface{}, error) {
	if len(value) != 2 {
		return nil, fmt.Errorf("length of []byte must be equal to 2")
	}
	return binary.BigEndian.Uint16(value), nil
}

func UInt16ToBytes(value interface{}) ([]byte, error) {
	ret := make([]byte, 2)
	i, ok := value.(uint16)
	if !ok {
		return nil, fmt.Errorf("parameter is of type %T, must be a uint16", value)
	}
	binary.BigEndian.PutUint16(ret, i)
	return ret, nil
}

func StringFromBytes(value []byte) (interface{}, error) {
	return string(value), nil
}

func StringToBytes(value interface{}) ([]byte, error) {
	s, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("parameter is of type %T, must be a string", value)
	}
	return []byte(s), nil
}

func RegisterCellType(name string, from FromBytes, to ToBytes) {
	if CellTypes == nil {
		CellTypes = make(map[string]CellType)
	}
	CellTypes[name] = CellType{from, to}
}

func (nv *NVMEM) ReadCell(name string) (interface{}, error) {
	var value interface{}
	cell, ok := nv.cells[name]
	if !ok {
		return nil, fmt.Errorf("No cell named '%s'", name)
	}
	handler, ok := CellTypes[cell.Type]
	if !ok {
		return nil, fmt.Errorf("No registered type named '%s'", cell.Type)
	}
	buffer := make([]byte, cell.Len)
	_, err := nv.ReadAt(buffer, int64(cell.Offset))
	if err != nil {
		return nil, err
	}
	if handler.from != nil {
		value, err = handler.from(buffer)
	} else {
		value = buffer
	}
	return value, err
}

func (nv *NVMEM) WriteCell(name string, value interface{}) error {
	cell, ok := nv.cells[name]
	var buffer []byte
	var err error
	_ = cell
	if !ok {
		return fmt.Errorf("No cell named '%s'", name)
	}
	handler, ok := CellTypes[cell.Type]
	if !ok {
		return fmt.Errorf("No registered type named '%s'", cell.Type)
	}
	if handler.to != nil {
		buffer, err = handler.to(value)
	} else {
		buffer, err = value.([]byte), nil
	}
	if err == nil {
		_, err = nv.WriteAt(buffer, int64(cell.Offset))
	}
	return err
}
