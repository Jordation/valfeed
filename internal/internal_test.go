package internal_test

import (
	"testing"
	"time"

	"github.com/Jordation/jsonl/internal/core"
	"github.com/Jordation/jsonl/internal/translation"
	"github.com/Jordation/jsonl/provider"
	"github.com/Jordation/jsonl/utils"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

var pump, _ = provider.NewPump(utils.GetRelativePath("../utils/test.jsonl"))

func TestGetMatchData(t *testing.T) {
	res, err := pump.GetDeltas(0, 9000)
	if err != nil {
		logrus.Fatal(err)
	}
	for _, evt := range res {
		delta := translation.TranslatorPipeline(evt)
		spew.Dump(delta)
	}
}

func TestIngest(t *testing.T) {
	feed, _ := provider.NewFeed()
	c := feed.Connect()
	for event := range c {
		spew.Dump(event)
	}
}

func TestCore(t *testing.T) {
	c, err := core.New()
	if err != nil {
		logrus.Fatal(err)
	}
	c.Start()
	time.Sleep(time.Second * 10)
}
