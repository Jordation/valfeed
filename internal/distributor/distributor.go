package distributor

import (
	"time"

	"github.com/Jordation/jsonl/provider"
	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/sirupsen/logrus"
)

type Distributor struct {
	Feed           provider.Feed
	EventListeners map[string]chan *riotTypes.Event
}

func New(feed provider.Feed) *Distributor {
	return &Distributor{
		Feed: feed,
	}
}

func (d *Distributor) Start() {
	go func() {
		time.Sleep(time.Second * 5)
		close(d.Feed.Stream())
	}()

	go func() {
		for evt := range d.Feed.Stream() {
			for id, listener := range d.EventListeners {
				logrus.Info("sending message to ", id)
				listener <- evt
			}
		}
	}()

}

func (d *Distributor) AddListener(id string, listener chan *riotTypes.Event) {
	d.EventListeners[id] = listener
}
