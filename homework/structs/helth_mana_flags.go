package structs

// healthManaFlags packs multiple fields into a single 32-bit value.
//
// A bit numbering: 0 = least significant bit (rightmost), 31 = most significant bit (leftmost)
//
// |31|30 29 28  |27........16|15| 14 13 12|11........0|
// |X |PlayerType|   Health   |X | Flags   |    Mana   |
//
// X (bit 31): unused/reserved
// PlayerType (bits 28–30): 3-bit player type
//
//	C = Constructor, F = Forger, W = Warrior
//
// Health (bits 16–27): 12-bit health value (0…1000)
// Mana   (bits 0–11): 12-bit mana value (0…1000)
//
// Flags  (bits 12–14): 4-bit status flags
//
//	bit 12 = Family
//	bit 13 = Gun
//	bit 14 = Home
//
// X (bit 15): unused/reserved
//
// This layout counts bits from right (LSB = 0) to left (MSB = 31) and allows compact storage of stats and flags.
type healthManaFlags struct {
	val uint32
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

const (
	manaMask       uint32 = 0xFFF
	familyFlagMask uint32 = 1 << 12
	gunFlagMask    uint32 = 1 << 13
	homeFlagMask   uint32 = 1 << 14
	healthMask     uint32 = 0xFFF << 16
	playerTypeMask uint32 = 0x7 << 28
)

// GetBits извлекает значение из val по маске и сдвигу
func GetBits(val uint32, mask uint32, shift uint) uint32 {
	return (val & mask) >> shift
}

// SetBits записывает value в val по маске и сдвигу
func SetBits(val *uint32, mask uint32, shift uint, value uint32) {
	*val = (*val & ^mask) | ((value << shift) & mask)
}
func (h *healthManaFlags) GetHealth() uint16 {
	return uint16(GetBits(h.val, healthMask, 16))
}

func (h *healthManaFlags) SetHealth(health uint16) {
	SetBits(&h.val, healthMask, 16, uint32(health))
}

func (h *healthManaFlags) GetMana() uint16 {
	return uint16(GetBits(h.val, manaMask, 0))
}

func (h *healthManaFlags) SetMana(mana uint16) {
	SetBits(&h.val, manaMask, 0, uint32(mana))
}

func (h *healthManaFlags) GetFamilyFlag() bool {
	return GetBits(h.val, familyFlagMask, 12) != 0
}

func (h *healthManaFlags) SetFamilyFlag(hasFamily bool) {
	if hasFamily {
		h.val |= familyFlagMask
	} else {
		h.val &^= familyFlagMask
	}
}

func (h *healthManaFlags) GetGunFlag() bool {
	return GetBits(h.val, gunFlagMask, 13) != 0
}

func (h *healthManaFlags) SetGunFlag(hasGun bool) {
	if hasGun {
		h.val |= gunFlagMask
	} else {
		h.val &^= gunFlagMask
	}
}

func (h *healthManaFlags) GetHomeFlag() bool {
	return GetBits(h.val, homeFlagMask, 14) != 0
}

func (h *healthManaFlags) SetHomeFlag(hasHome bool) {
	if hasHome {
		h.val |= homeFlagMask
	} else {
		h.val &^= homeFlagMask
	}
}

func (h *healthManaFlags) GetPlayerType() int {
	return int(GetBits(h.val, playerTypeMask, 28))
}

func (h *healthManaFlags) SetPlayerType(pt uint8) {
	SetBits(&h.val, playerTypeMask, 28, uint32(pt))
}
