package core

import (
	"github.com/Jordation/jsonl/internal/manager"
	"github.com/Jordation/jsonl/provider"
)

type CoreService struct {
	Upstreams *Upstreams
}
type Upstreams struct {
	RiotFeed *manager.RiotFeed
}

func New() (*CoreService, error) {
	feed, err := provider.NewFeed()
	if err != nil {
		return nil, err
	}

	return &CoreService{
		Upstreams: &Upstreams{
			RiotFeed: manager.NewRiotFeed(feed),
		},
	}, nil
}

func (c *CoreService) Start() error {
	go c.Upstreams.RiotFeed.Start()

	return nil
}
