package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"

	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/sirupsen/logrus"
)

func GetRelativePath(target string) string {
	_, path, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(path), target)
}

func GetWeaponMappings() map[string]string {
	weps := []*riotTypes.APIWeapon{}
	res := map[string]string{}

	data, err := os.ReadFile(GetRelativePath("./weapons.json"))
	if err != nil {
		logrus.Fatal(err)
	}

	if err := json.Unmarshal(data, &weps); err != nil {
		logrus.Fatal(err)
	}

	for _, wep := range weps {
		res[wep.UUID] = wep.WeaponName
	}

	return res
}

func GetPlayerMappings() map[int]string {
	event := riotTypes.Event{}
	res := map[int]string{}

	data, err := os.ReadFile(GetRelativePath("./players.json"))
	if err != nil {
		logrus.Fatal(err)
	}

	if err := json.Unmarshal(data, &event); err != nil {
		logrus.Fatal(err)
	}

	for _, player := range event.Configuration.Players {
		res[player.PlayerID.Value] = player.DisplayName
	}

	return res
}

// abilities are mapped to agent UUID -> ability SLOT
// i.e. abilities[uuid][slot]
// abilities[1234][Ultimate] // abilities[1234][Ability2]
func GetAgentAbilityMappings() map[string]riotTypes.AgentAbilityMapping {
	res := map[string]riotTypes.AgentAbilityMapping{}
	mappings := []*riotTypes.APIAgent{}

	data, err := os.ReadFile(GetRelativePath("./agents.json"))
	if err != nil {
		logrus.Fatal(err)
	}

	if err := json.Unmarshal(data, &mappings); err != nil {
		logrus.Fatal(err)
	}

	for _, agent := range mappings {
		temp := riotTypes.AgentAbilityMapping{
			Name:      agent.DisplayName,
			Abilities: map[string]string{},
		}

		for _, ability := range agent.Abilities {
			temp.Abilities[ability.Slot] = ability.DisplayName
		}
		res[agent.UUID] = temp
	}

	return res
}
