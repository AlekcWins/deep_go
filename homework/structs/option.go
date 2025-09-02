package structs

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		p := [42]byte{}
		copy(p[:], name)
		person.name = p
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.healthManaFlags.SetMana(uint16(mana))
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.healthManaFlags.SetHealth(uint16(health))
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats.SetRespect(int32(respect))
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats.SetStrength(int32(strength))
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats.SetExperience(int32(experience))
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats.SetLevel(int32(level))
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.healthManaFlags.SetHomeFlag(true)
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.healthManaFlags.SetGunFlag(true)
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.healthManaFlags.SetFamilyFlag(true)
	}
}

func WithType(personType uint8) func(*GamePerson) {
	return func(person *GamePerson) {
		person.healthManaFlags.SetPlayerType(personType)
	}
}
