package bitmask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		ids      []int
		expected int
	}{
		{
			name:     "empty slice",
			ids:      []int{},
			expected: 0,
		},
		{
			name:     "single bit zero",
			ids:      []int{0},
			expected: 1,
		},
		{
			name:     "single bit non zero",
			ids:      []int{3},
			expected: 8,
		},
		{
			name:     "multiple bits",
			ids:      []int{1, 3, 5},
			expected: 42,
		},
		{
			name:     "duplicate ids",
			ids:      []int{1, 1, 3},
			expected: 10,
		},
		{
			name:     "high bit",
			ids:      []int{31},
			expected: 1 << 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := Encode(tt.ids)
			assert.Equal(t, tt.expected, encoded, tt.name)
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		mask     int
		expected []int
	}{
		{
			name:     "zero mask",
			mask:     0,
			expected: []int{},
		},
		{
			name:     "single bit",
			mask:     1,
			expected: []int{0},
		},
		{
			name:     "multiple bits",
			mask:     42,
			expected: []int{1, 3, 5},
		},
		{
			name:     "all bits",
			mask:     15,
			expected: []int{0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoded := Decode(tt.mask)
			assert.Equal(t, tt.expected, decoded, tt.name)
		})
	}
}

func TestHasBit(t *testing.T) {
	tests := []struct {
		name     string
		mask     int
		id       int
		expected bool
	}{
		{
			name:     "bit set",
			mask:     42,
			id:       3,
			expected: true,
		},
		{
			name:     "bit not set",
			mask:     42,
			id:       2,
			expected: false,
		},
		{
			name:     "zero mask",
			mask:     0,
			id:       0,
			expected: false,
		},
		{
			name:     "high bit set",
			mask:     1 << 31,
			id:       31,
			expected: true,
		},
		{
			name: "high bit not set",
			mask: 1 << 30,
			id:   31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasBit := HasBit(tt.mask, tt.id)
			assert.Equal(t, tt.expected, hasBit, tt.name)
		})
	}
}

func TestToggleBit(t *testing.T) {
	tests := []struct {
		name         string
		existingMask int
		id           int
		expectedMask int
	}{
		{
			name:         "toggle on",
			existingMask: 0,
			id:           2,
			expectedMask: 4,
		},
		{
			name:         "toggle off",
			existingMask: 4,
			id:           2,
			expectedMask: 0,
		},
		{
			name:         "toggle unrelated",
			existingMask: 2,
			id:           3,
			expectedMask: 10,
		},
		{
			name:         "toggle high bit",
			existingMask: 0,
			id:           31,
			expectedMask: 1 << 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toggedMask := ToggleBit(tt.existingMask, tt.id)
			assert.Equal(t, tt.expectedMask, toggedMask, tt.name)
		})
	}
}
