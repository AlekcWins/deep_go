package structs

import "encoding/json"

type StatsJSON struct {
	Respect    int `json:"respect"`
	Strength   int `json:"strength"`
	Experience int `json:"experience"`
	Level      int `json:"level"`
}

type FlagsJSON struct {
	HasHouse  bool `json:"has_house"`
	HasGun    bool `json:"has_gun"`
	HasFamily bool `json:"has_family"`
}

type ResourcesJSON struct {
	Gold   int `json:"gold"`
	Health int `json:"health"`
	Mana   int `json:"mana"`
}
type GamePersonJSON struct {
	X          int           `json:"x"`
	Y          int           `json:"y"`
	Z          int           `json:"z"`
	Name       string        `json:"name"`
	Resources  ResourcesJSON `json:"resources"`
	PlayerType string        `json:"player_type"`
	Stats      StatsJSON     `json:"stats"`
	Flags      FlagsJSON     `json:"flags"`
}

func getTypeString(p GamePerson) string {
	switch p.Type() {
	case BuilderGamePersonType:
		return "Builder"
	case BlacksmithGamePersonType:
		return "Blacksmith"
	case WarriorGamePersonType:
		return "Warrior"
	default:
		return "Unknown"
	}
}

func (p GamePerson) MarshalJSON() ([]byte, error) {
	data := GamePersonJSON{
		X: p.X(),
		Y: p.Y(),
		Z: p.Z(),

		Name:       p.Name(),
		PlayerType: getTypeString(p),

		Resources: ResourcesJSON{
			Gold:   p.Gold(),
			Health: p.Health(),
			Mana:   p.Mana(),
		},

		Flags: FlagsJSON{
			HasHouse:  p.HasHouse(),
			HasGun:    p.HasGun(),
			HasFamily: p.HasFamily(),
		},

		Stats: StatsJSON{
			Respect:    p.Respect(),
			Strength:   p.Strength(),
			Experience: p.Experience(),
			Level:      p.Level(),
		},
	}

	return json.Marshal(data)
}
