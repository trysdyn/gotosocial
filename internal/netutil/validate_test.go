package netutil

import (
	"net/netip"
	"testing"
)

func TestValidateIP(t *testing.T) {
	tests := []struct {
		name     string
		ip       netip.Addr
		expected bool
	}{
		// IPv4 tests
		{
			name:     "IPv4 this host on this network",
			ip:       netip.MustParseAddr("0.0.0.0"),
			expected: false,
		},
		{
			name:     "IPv4 dummy address",
			ip:       netip.MustParseAddr("192.0.0.8"),
			expected: false,
		},
		{
			name:     "IPv4 Port Control Protocol Anycast",
			ip:       netip.MustParseAddr("192.0.0.9"),
			expected: false,
		},
		{
			name:     "IPv4 Traversal Using Relays around NAT Anycast",
			ip:       netip.MustParseAddr("192.0.0.10"),
			expected: false,
		},
		{
			name:     "IPv4 NAT64/DNS64 Discovery 1",
			ip:       netip.MustParseAddr("192.0.0.17"),
			expected: false,
		},
		{
			name:     "IPv4 NAT64/DNS64 Discovery 2",
			ip:       netip.MustParseAddr("192.0.0.171"),
			expected: false,
		},
		{
			name:     "Leah's SOCKS proxy",
			ip:       netip.MustParseAddr("172.17.0.1"),
			expected: true,
		},
		// IPv6 tests
		{
			name: "IPv4-mapped address",
			ip:   netip.MustParseAddr("::ffff:169.254.169.254"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if valid := ValidateIP(tc.ip); valid != tc.expected {
				t.Fatalf("Expected IP %s to be: %t, got: %t", tc.ip, tc.expected, valid)
			}
		})
	}
}
