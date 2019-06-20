package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"net"
	"os"
	"strconv"
	"strings"
)

type InterfaceRoutes struct {
	Interface *net.Interface
	Routes    []Route
}

type Route struct {
	Destination net.IPNet
	Gateway     net.IP
	Metric      uint32
	Flags       uint32
	Default     bool
}

func GetRoutes() map[string]*InterfaceRoutes {
	var routes map[string]*InterfaceRoutes = make(map[string]*InterfaceRoutes)
	parseIPv4Routes(routes)
	parseIPv6Routes(routes)
	return routes
}

func swapBytes(bytes []byte) []byte {
	value := binary.BigEndian.Uint32(bytes)
	binary.LittleEndian.PutUint32(bytes, value)
	return bytes
}

func parseIPv6Routes(routes map[string]*InterfaceRoutes) {
	file, err := os.Open("/proc/net/ipv6_route")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		intf_name := tokens[len(tokens)-1]
		route_intf, ok := routes[intf_name]
		if !ok {
			intf, err := net.InterfaceByName(intf_name)
			if err != nil {
				continue
			}
			route_intf = &InterfaceRoutes{intf, nil}
			routes[intf_name] = route_intf
		}
		dest, err := hex.DecodeString(tokens[0])
		if err != nil {
			continue
		}
		prefixlen, err := strconv.ParseUint(tokens[1], 16, 8)
		if err != nil {
			continue
		}
		nexthop, err := hex.DecodeString(tokens[4])
		if err != nil {
			continue
		}
		metric, err := strconv.ParseUint(tokens[5], 16, 32)
		if err != nil {
			continue
		}
		flags, err := strconv.ParseUint(tokens[8], 16, 32)
		if err != nil {
			continue
		}
		deflt := ((tokens[0] == "00000000000000000000000000000000") && (flags&0x02 == 0x02))
		route := Route{net.IPNet{dest, net.CIDRMask(int(prefixlen), 128)}, net.IP(nexthop), uint32(metric), uint32(flags), deflt}
		route_intf.Routes = append(route_intf.Routes, route)
	}
}

func parseIPv4Routes(routes map[string]*InterfaceRoutes) {
	file, err := os.Open("/proc/net/route")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan() /* Skip first line - header */
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "\t")
		intf_name := tokens[0]
		route_intf, ok := routes[intf_name]
		if !ok {
			intf, err := net.InterfaceByName(intf_name)
			if err != nil {
				continue
			}
			route_intf = &InterfaceRoutes{intf, nil}
			routes[intf_name] = route_intf
		}
		dest, err := hex.DecodeString(tokens[1])
		if err != nil {
			continue
		}
		gateway, err := hex.DecodeString(tokens[2])
		if err != nil {
			continue
		}
		mask, err := hex.DecodeString(tokens[7])
		if err != nil {
			continue
		}
		flags, err := strconv.ParseUint(tokens[3], 16, 16)
		if err != nil {
			continue
		}
		metric, err := strconv.ParseUint(tokens[6], 10, 16)
		if err != nil {
			continue
		}
		deflt := ((tokens[1] == "00000000") && (flags&0x02 == 0x02))
		route := Route{net.IPNet{swapBytes(dest), swapBytes(mask)}, net.IP(swapBytes(gateway)), uint32(metric), uint32(flags), deflt}
		route_intf.Routes = append(route_intf.Routes, route)
	}
}
