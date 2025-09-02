package structs

const (
	twoByteLeftMask  uint8 = 0xF0
	twoByteRightMask uint8 = 0x0F
)

type TwoNumberByte struct {
	val byte // [HHHH][LLLL]
}

func (t *TwoNumberByte) GetLowValue() byte {
	return t.val & twoByteRightMask
}

func (t *TwoNumberByte) SetLowValue(val byte, maxVal byte) bool {
	if val > maxVal || val > 15 {
		return false
	}
	t.val = t.val&^twoByteRightMask | val

	return true
}

func (t *TwoNumberByte) GetHighValue() byte {
	return t.val >> 4
}

func (t *TwoNumberByte) SetHighValue(val byte) bool {
	if val > 15 {
		return false
	}
	t.val = t.val&^twoByteLeftMask | val<<4

	return true
}
