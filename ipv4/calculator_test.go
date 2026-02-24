package ipv4

import (
	"testing"
)

// TestGetClass verifies that the tool correctly identifies IP classes
func TestGetClass(t *testing.T) {
	// 1. Define the table of scenarios
	tests := []struct {
		name     string  // Description of the test
		ip       Address // The input
		expected Class   // The expected output
	}{
		{"Class A Minimum", Init(1, 0, 0, 0), A},
		{"Class A Maximum", Init(127, 255, 255, 255), A},
		{"Class B Typical", Init(172, 16, 0, 1), B},
		{"Class C Typical", Init(192, 168, 1, 10), C},
		{"Class D Multicast", Init(224, 0, 0, 5), D},
		{"Class E Experimental", Init(240, 0, 0, 1), E},
	}

	// 2. Loop through the table and run each test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetClass(tt.ip); got != tt.expected {
				t.Errorf("GetClass() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestGetHostMaxNumber verifies the host calculation math
func TestGetHostMaxNumber(t *testing.T) {
	tests := []struct {
		name     string
		mask     Address
		expected int
	}{
		{"/24 Subnet", Init(255, 255, 255, 0), 254},
		{"/26 Subnet", Init(255, 255, 255, 192), 62},
		{"/30 Subnet", Init(255, 255, 255, 252), 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHostMaxNumber(tt.mask); got != tt.expected {
				t.Errorf("GetHostMaxNumber() = %v, want %v", got, tt.expected)
			}
		})
	}
}
