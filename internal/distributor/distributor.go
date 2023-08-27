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
	TransEventListeners map[string]chan *types.TranslatedEvent

	Feed         provider.Feed
	IncDmgEvents chan *types.CombatEvent
	IncRndEvents chan *types.RoundEvent
	// incoming new type of event
}

func New(feed provider.Feed) *Distributor {
	return &Distributor{
		Feed:                feed,
		RiotEventListeners:  map[string]chan *riotTypes.Event{},
		TransEventListeners: map[string]chan *types.TranslatedEvent{},
		IncDmgEvents:        make(chan *types.CombatEvent),
		IncRndEvents:        make(chan *types.RoundEvent),
	}
}

func (d *Distributor) AddRiotEventListener(id string, listener chan *riotTypes.Event) {
	d.RiotEventListeners[id] = listener
}
func (d *Distributor) AddTranslatedEventListener(id string, listener chan *types.TranslatedEvent) {
	d.TransEventListeners[id] = listener
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
				d.sendTranslatedEvent(&types.TranslatedEvent{
					Event: event,
					ID:    event.ID,
				})

			case event := <-d.IncRndEvents:
				d.sendTranslatedEvent(&types.TranslatedEvent{
					Event: event,
					ID:    event.GameID,
				})
			}
		}
	}()

}

func (d *Distributor) sendTranslatedEvent(evt *types.TranslatedEvent) {
	for _, listener := range d.TransEventListeners {
		listener <- evt
	}
}
