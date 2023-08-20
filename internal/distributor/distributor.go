package distributor

import (
	"time"

	"github.com/Jordation/jsonl/internal/types"
	"github.com/Jordation/jsonl/provider"
	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/sirupsen/logrus"
)

type Distributor struct {
	RiotEventListeners  map[string]chan *riotTypes.Event
	TransEventListeners map[string]chan any

	Feed         provider.Feed
	IncDmgEvents chan *types.CombatEvent
	IncRndEvents chan *types.RoundEvent
}

func New(feed provider.Feed) *Distributor {
	return &Distributor{
		Feed:               feed,
		RiotEventListeners: map[string]chan *riotTypes.Event{},
		IncDmgEvents:       make(chan *types.CombatEvent),
		IncRndEvents:       make(chan *types.RoundEvent),
	}
}

func (d *Distributor) Start() {
	go func() {
		time.Sleep(time.Second * 5)
		close(d.Feed.Stream())
	}()

	for evt := range d.Feed.Stream() {
		// Send the message to each listener
		for lisID, listener := range d.RiotEventListeners {
			logrus.Tracef("sending a message to %v", lisID)
			listener <- evt
		}
	}

}

func (d *Distributor) AddListener(id string, listener chan *riotTypes.Event) {
	d.RiotEventListeners[id] = listener
}

func (d *Distributor) BroadcastFeed() {
	go func() {
		for {
			if len(d.RiotEventListeners) == 0 {
				logrus.Info("no riot event listeners, sleeping")
				time.Sleep(time.Second * 1)
			} else {
				break
			}
		}

		for event := range d.Feed.Stream() {
			for _, listener := range d.RiotEventListeners {
				listener <- event
			}
		}
	}()
}
func (d *Distributor) BroadcastTranslatedEvents() {
	go func() {
		for {
			if len(d.RiotEventListeners) == 0 {
				logrus.Info("no translated event listeners, sleeping")
				time.Sleep(time.Second * 1)
			} else {
				break
			}
		}

		for {
			select {
			case event := <-d.IncDmgEvents:
				d.SendTranslatedEvent(event)
			case event := <-d.IncRndEvents:
				d.SendTranslatedEvent(event)
			}
		}
	}()

}

func (d *Distributor) SendTranslatedEvent(evt any) {
	for _, listener := range d.TransEventListeners {
		listener <- evt
	}
}
