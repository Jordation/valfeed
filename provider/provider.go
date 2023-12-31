package provider

import (
	"time"

	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/Jordation/jsonl/utils"
	"github.com/sirupsen/logrus"
)

type Feed interface {
	Stream() chan *riotTypes.Event
}

type ValFeed struct {
	Pump *JsonlPump
}

func NewFeed() (Feed, error) {
	p, err := NewPump(utils.GetRelativePath("../utils/test.jsonl"))
	if err != nil {
		return nil, err
	}

	return &ValFeed{
		Pump: p,
	}, nil
}

func (f *ValFeed) Stream() chan *riotTypes.Event {
	c := make(chan *riotTypes.Event)
	go f.streamEvents(c)
	return c
}

func (f *ValFeed) streamEvents(c chan *riotTypes.Event) {
	for {
		evt, err := f.Pump.GetDelta()
		if err != nil {
			logrus.Info("pump error, probzbly closed", err.Error())
			time.Sleep(time.Second * 10)
			return
		}
		c <- evt
	}
}
