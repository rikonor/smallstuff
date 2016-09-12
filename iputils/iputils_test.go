package iputils

import (
	"net"
	"testing"
)

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

func TestCIDRToIPRange(t *testing.T) {
	type testCase struct {
		in  string
		out [2]net.IP
	}

	testCases := []testCase{
		testCase{
			"193.254.30.0/24",
			[2]net.IP{
				net.ParseIP("193.254.30.0"),
				net.ParseIP("193.254.31.0"),
			},
		},
	}

	for _, tc := range testCases {
		if ip1, ip2 := CIDRToIPRange(tc.in); !ip1.Equal(tc.out[0]) || !ip2.Equal(tc.out[1]) {
			t.Errorf("Failed to convert CIDR string %s to IP range. Expected (%s-%s) but got (%s-%s)", tc.in, tc.out[0], tc.out[1], ip1, ip2)
		}
	}
}

func BenchmarkCIDRToIPRange(b *testing.B) {
	cidr := "193.254.30.0/24"

	for i := 0; i < b.N; i++ {
		CIDRToIPRange(cidr)
	}
}

func TestIPRangeSize(t *testing.T) {
	type testCase struct {
		in  [2]net.IP
		out uint32
	}

	testCases := []testCase{
		testCase{
			[2]net.IP{
				net.IPv4zero,
				net.IPv4zero,
			},
			0,
		},
		testCase{
			[2]net.IP{
				net.ParseIP("0.0.0.1"),
				net.ParseIP("0.0.0.2"),
			},
			1,
		},
		testCase{
			[2]net.IP{
				net.ParseIP("193.254.30.0"),
				net.ParseIP("193.254.30.255"),
			},
			256,
		},
	}

	for _, tc := range testCases {
		if sz := IPRangeSize(tc.in[0], tc.in[1]); sz != tc.out {
			t.Errorf("Failed to extract range size from IP range (%s-%s). Expected %d but got %d", tc.in[0], tc.in[1], tc.out, sz)
		}
	}
}

func BenchmarkIPRangeSize(b *testing.B) {
	ip1 := net.IPv4zero
	ip2 := net.ParseIP("255.255.255.255")

	for i := 0; i < b.N; i++ {
		IPRangeSize(ip1, ip2)
	}
}
