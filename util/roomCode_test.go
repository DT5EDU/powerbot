package util

import (
	"testing"
)

func TestRoomCode(t *testing.T) {
	testCases := []struct {
		roomNumber string
		expected   string
	}{
		// 就用我们班的 11 个寝室测试吧~
		{"113401", "6X0VVY3XlmTqcnYUfy7CAA=="},
		{"113402", "rwN/OuzGgqw41gEJJzu+5A=="},
		{"113403", "yM5iS3mIyBDjZ8l6P1NNmg=="},
		{"113404", "i08WJz55Vrn3VQi4wieD2w=="},
		{"113405", "D20r+1BG747jBFdVvoT/0g=="},
		{"113406", "GQIkJP+4yrb8eZhX2cGEGA=="},
		{"113407", "Yd2l/9QQDSMhs/i6+qrQqg=="},
		{"113408", "yhMPbSrVnBS0p4sYCi6sTw=="},
		{"104129", "ocGXfr1ERAGgHLScZXmd3w=="},
		{"104131", "CoY0KmKDCPGwmrYfh4p4iQ=="},
		{"104133", "oaMrMPAoFl2ZmLaWjGaMrg=="},
	}

	for _, tc := range testCases {
		if result := RoomCode(tc.roomNumber); result != tc.expected {
			t.Errorf("RoomCode(%s) = %s; expected: %s", tc.roomNumber, result, tc.expected)
		}
	}
}
