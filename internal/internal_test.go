package internal_test

import (
	"testing"

	"github.com/Jordation/jsonl/internal/merge"
	"github.com/Jordation/jsonl/provider"
	"github.com/Jordation/jsonl/utils"
	"github.com/sirupsen/logrus"
)

var pump, _ = provider.NewPump(utils.GetRelativePath("../utils/test.jsonl"))

func TestGetMatchData(t *testing.T) {
	res, err := pump.GetDeltas(0, 9000)
	if err != nil {
		logrus.Fatal(err)
	}
	_, _ = merge.TranslateMatchData(res)
}
