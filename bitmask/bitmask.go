// Package bitmask provides utilities for encoding and decoding sets of integer IDs
// using bitmasks. A bitmask is a compact way to store multiple boolean flags in a single int.
//
// For example, the integer 42 is 00101010 in binary — meaning bits 1, 3, and 5 are set.
//
// Key concept:
//
//	The bitmask for a single number n is: 1 << n == 2^n
//	- Bit 0 → 1 << 0 = 1      = 2⁰
//	- Bit 1 → 1 << 1 = 2      = 2¹
//	- Bit 2 → 1 << 2 = 4      = 2²
//	- Bit 3 → 1 << 3 = 8      = 2³
//	- ...
package bitmask

// Encode returns a bitmask with bits set for each ID in the input slice.
//
// Each ID corresponds to a bit position. For example, ID 3 sets bit 3 using `1 << 3`.
// This results in: 00001000 (decimal 8), which is 2³.
//
// Bitwise OR `mask |= 1 << id` ensures the bit at `id` is turned on without affecting others.
//
// Example:
//
//	Encode([]int{1, 3, 5}) => 42 (binary: 00101010)
//	  - 1 << 1 = 2  = 00000010
//	  - 1 << 3 = 8  = 00001000
//	  - 1 << 5 = 32 = 00100000
//	  - OR all of them:         00101010
func Encode(ids []int) int {
	var mask int

	for _, id := range ids {
		mask |= 1 << id
	}

	return mask
}

// Decode returns the list of IDs (bit positions) that are set in the given bitmask.
//
// It loops through the bits of the mask from right to left.
// At each step, it checks the lowest bit using `mask & 1`.
// If the result is 1, that bit is set, and its position is recorded.
// The mask is then shifted right (`mask >>= 1`) to check the next bit.
//
// Example:
//
//	mask = 42 = 00101010
//	→ bits 1, 3, and 5 are set
//	→ returns []int{1, 3, 5}
func Decode(mask int) []int {
	ids := make([]int, 0)

	for bit := 0; mask != 0; bit++ {
		if mask&1 == 1 {
			ids = append(ids, bit)
		}
		mask >>= 1
	}

	return ids
}

// HasBit tell if the bit at position `id` is set in the mask.
//
// It uses bitwise AND: `mask & (1 << id)`
//   - If the result is non-zero, then bit `id` is set.
//   - If the result is zero, then bit `id` is not set.
//
// Example:
//
//	mask = 00101010 (decimal 42)
//	id = 3 → 1 << 3 = 00001000
//	mask & (1 << 3) = 00101010 & 00001000 = 00001000 → bit 3 is set → true
func HasBit(mask int, id int) bool {
	return (mask & (1 << id)) != 0
}

// ToggleBit flips the bit at position `id` in the mask using XOR (`^`).
//
// If the bit is 1, it becomes 0. If it's 0, it becomes 1.
//   - 1 ^ 1 = 0 (turns off the bit)
//   - 0 ^ 1 = 1 (turns on the bit)
//
// Example:
//
//	mask = 00101010 (decimal 42)
//	id = 3 → 1 << 3 = 00001000
//	00101010 ^ 00001000 = 00100010 (decimal 34)
func ToggleBit(mask int, id int) int {
	return mask ^ (1 << id)
}
