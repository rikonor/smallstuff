package iputils

import (
	"net"
	"strconv"
	"strings"
)

// Increment increments an IP by `inc` amount
func Increment(ip net.IP, inc uint32) net.IP {
	iip := IPToUint32(ip)
	iip += inc
	return Uint32ToIP(iip)
}

// IPToUint32 converts an IPv4 to a uint32
func IPToUint32(ip net.IP) uint32 {
	bs := strings.Split(ip.String(), ".")

	b0, _ := strconv.Atoi(bs[0])
	b1, _ := strconv.Atoi(bs[1])
	b2, _ := strconv.Atoi(bs[2])
	b3, _ := strconv.Atoi(bs[3])

	var iip uint32

	iip += uint32(b0) << 24
	iip += uint32(b1) << 16
	iip += uint32(b2) << 8
	iip += uint32(b3)

	return iip
}

// Uint32ToIP converts a uint32 to IPv4
func Uint32ToIP(iip uint32) net.IP {
	d := uint8(iip)
	iip = iip >> 8
	c := uint8(iip)
	iip = iip >> 8
	b := uint8(iip)
	iip = iip >> 8
	a := uint8(iip)

	return net.IPv4(a, b, c, d)
}
