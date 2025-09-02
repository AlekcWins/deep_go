package structs

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGamePersonJSON(t *testing.T) {
	// Создаем пример GamePerson
	p := NewGamePerson(
		func(p *GamePerson) { copy(p.name[:], "Alex") },
		func(p *GamePerson) { p.x = 10 },
		func(p *GamePerson) { p.y = 20 },
		func(p *GamePerson) { p.z = 5 },
		func(p *GamePerson) { p.gold = 100 },
	)
	p.stats.SetRespect(8)
	p.stats.SetStrength(7)
	p.stats.SetExperience(5)
	p.stats.SetLevel(3)
	p.healthManaFlags.SetFamilyFlag(true)
	p.healthManaFlags.SetFamilyFlag(false)
	p.healthManaFlags.SetHealth(100)
	p.healthManaFlags.SetMana(999)
	p.healthManaFlags.SetPlayerType(WarriorGamePersonType)

	// Маршалим в JSON
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	// Печатаем
	fmt.Println(string(jsonData))
}
