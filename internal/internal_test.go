package internal_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Jordation/jsonl/internal/distributor"
	"github.com/Jordation/jsonl/internal/translate"
	"github.com/Jordation/jsonl/provider"
	"github.com/Jordation/jsonl/utils"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

var pump, _ = provider.NewPump(utils.GetRelativePath("../utils/test.jsonl"))

func TestIngest(t *testing.T) {
	feed, _ := provider.NewFeed()
	c := feed.Stream()
	for event := range c {
		spew.Dump(event)
	}
}

func TestPlayerEventManager(t *testing.T) {
	res, err := pump.GetDeltas(0, 2000)
	if err != nil {
		logrus.Fatal(err)
	}
	pm := utils.GetPlayerMappings()

	player := translate.NewDamageTranslator(pm, 1)

	ctx := context.Background()
	go player.Start(ctx)

	for _, evt := range res {
		player.Ingest(evt)
	}

	for _, event := range player.CombatEvents {
		fmt.Println(event)
	}
}

func TestDistributor(t *testing.T) {
	feed, err := provider.NewFeed()
	if err != nil {
		logrus.Fatal(err)
	}
	d := distributor.New(feed)
	d.Start()

}
