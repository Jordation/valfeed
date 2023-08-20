package persistance

import (
	"encoding/json"
	"os"
	"testing"

	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_Pers(t *testing.T) {
	input := &riotTypes.Event{}
	data, _ := os.ReadFile("./test_cfg.json")

	err := json.Unmarshal(data, input)
	assert.NoError(t, err)

	p, err := New()
	assert.NoError(t, err)

	err = p.StoreGameConfig(input.Configuration, "test3")
	assert.NoError(t, err)

	cfg, err := p.GetGameConfig("test")
	assert.NoError(t, err)
	logrus.Info(cmp.Diff(input.Configuration, cfg))
}
