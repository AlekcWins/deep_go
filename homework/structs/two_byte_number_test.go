package structs

import (
	"fmt"
	"testing"
)

func TestTwoNumberByte_AllCases(t *testing.T) {
	for high := uint8(0); high <= 15; high++ { // 0–15
		for low := uint8(0); low <= 15; low++ { // 0–15
			var b TwoNumberByte
			b.SetHighValue(high)
			b.SetLowValue(low, 15)

			if got := b.GetHighValue(); got != high {
				t.Errorf("High mismatch: expected %X, got %X (val=%X)", high, got, b.val)
			}

			if got := b.GetLowValue(); got != low {
				t.Errorf("Low mismatch: expected %X, got %X (val=%X)", low, got, b.val)
			}

			expectedVal := (high << 4) | low
			if b.val != expectedVal {
				t.Errorf("Byte mismatch: expected %X, got %X (High=%X, Low=%X)", expectedVal, b.val, high, low)
			}
		}
	}
	fmt.Println("All 256 cases passed!")
}
