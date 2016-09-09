package iputils

import (
	"net"
	"testing"
)

// TODO Add benchmarks

func TestIncrement(t *testing.T) {
	type testCase struct {
		in, out net.IP
		inc     uint32
	}

	testCases := []testCase{
		testCase{in: net.IPv4zero, inc: 0, out: net.IPv4zero},
		testCase{in: net.IPv4zero, inc: 1, out: net.ParseIP("0.0.0.1")},
		testCase{in: net.ParseIP("255.255.255.255"), inc: 1, out: net.IPv4zero},
	}

	for _, tc := range testCases {
		if ip := Increment(tc.in, tc.inc); !ip.Equal(tc.out) {
			t.Errorf("Failed to increment %s by %d. Expected %s, got %s.", tc.in, tc.inc, tc.out, ip)
		}
	}
}

func BenchmarkIncrement(b *testing.B) {
	ip := net.IPv4zero

	for i := 0; i < b.N; i++ {
		ip = Increment(ip, 1)
	}
}

func TestIPToUint32(t *testing.T) {
	type testCase struct {
		in  net.IP
		out uint32
	}

	testCases := []testCase{
		testCase{net.IPv4zero, 0},
		testCase{net.ParseIP("0.0.0.1"), 1},
		testCase{net.ParseIP("255.255.255.255"), 4294967295},
	}

	for _, tc := range testCases {
		if iip := IPToUint32(tc.in); iip != tc.out {
			t.Errorf("Failed to convert ip %s to uint32. Expected %d, got %d", tc.in, tc.out, iip)
		}
	}
}

func BenchmarkIPToUint32(b *testing.B) {
	ip := net.ParseIP("255.255.255.255")

	for i := 0; i < b.N; i++ {
		IPToUint32(ip)
	}
}

func TestUint32ToIP(t *testing.T) {
	type testCase struct {
		in  uint32
		out net.IP
	}

	testCases := []testCase{
		testCase{0, net.IPv4zero},
		testCase{1, net.ParseIP("0.0.0.1")},
		testCase{4294967295, net.ParseIP("255.255.255.255")},
	}

	for _, tc := range testCases {
		if ip := Uint32ToIP(tc.in); !ip.Equal(tc.out) {
			t.Errorf("Failed to convert uint32 %d to ip. Expected %s, got %s", tc.in, tc.out, ip)
		}
	}
}

func BenchmarkUint32ToIP(b *testing.B) {
	iip := IPToUint32(net.ParseIP("255.255.255.255"))

	for i := 0; i < b.N; i++ {
		Uint32ToIP(iip)
	}
}
