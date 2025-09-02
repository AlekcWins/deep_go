package structs

type GamePerson struct {
	x, y, z         int32           // 4 * 3
	gold            uint32          // 4 * 1
	healthManaFlags healthManaFlags // 4 * 1
	name            [42]byte        // 1 * 42
	stats           Stats           // 1 * 2
}

func NewGamePerson(options ...Option) GamePerson {
	person := GamePerson{}
	for _, opt := range options {
		opt(&person)
	}
	return person
}

func (p *GamePerson) Name() string {
	i := 0
	for ; i < len(p.name); i++ {
		if p.name[i] == 0 {
			break
		}
	}
	return string(p.name[:i])
}

func (p *GamePerson) X() int {
	return int(p.x)
}

func (p *GamePerson) Y() int {
	return int(p.y)
}

func (p *GamePerson) Z() int {
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int(p.healthManaFlags.GetMana())
}

func (p *GamePerson) Health() int {
	return int(p.healthManaFlags.GetHealth())
}

func (p *GamePerson) Respect() int {
	return int(p.stats.GetRespect())
}

func (p *GamePerson) Strength() int {
	return int(p.stats.GetStrength())
}

func (p *GamePerson) Experience() int {
	return int(p.stats.GetExperience())
}

func (p *GamePerson) Level() int {
	return int(p.stats.GetLevel())
}

func (p *GamePerson) HasHouse() bool {
	return p.healthManaFlags.GetHomeFlag()
}

func (p *GamePerson) HasGun() bool {
	return p.healthManaFlags.GetGunFlag()
}

func (p *GamePerson) HasFamily() bool {
	return p.healthManaFlags.GetFamilyFlag()
}

func (p *GamePerson) Type() int {
	return p.healthManaFlags.GetPlayerType()
}
