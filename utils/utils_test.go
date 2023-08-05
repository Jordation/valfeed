package utils

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetConfigs(t *testing.T) {
	wm := GetWeaponMappings()
	pm := GetPlayerMappings()
	aam := GetAgentAbilityMappings()
	logrus.Infof("WEAPONS %++v \n, PLAYERS %++v \n, AGENT:ABILITY %++v", wm, pm, aam)
}
