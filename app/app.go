package app

import (
	"github.com/Jordation/jsonl/internal/distributor"
	"github.com/Jordation/jsonl/internal/merge"
	"github.com/Jordation/jsonl/internal/persistance"
	"github.com/Jordation/jsonl/internal/translate"
	"github.com/Jordation/jsonl/provider"
	"github.com/sirupsen/logrus"
)

type App struct {
	Feed               provider.Feed
	EventDistributor   *distributor.Distributor
	TranslationManager *translate.TranslationManager
	MergeManager       *merge.MergeManager
}

func New() (*App, error) {
	feed, err := provider.NewFeed()
	if err != nil {
		return nil, err
	}

	pers, err := persistance.New()
	if err != nil {
		return nil, err
	}

	distributor := distributor.New(feed)
	tm := translate.NewManager(pers, distributor)
	mm := merge.NewMerger(pers)

	return &App{
		Feed:               feed,
		EventDistributor:   distributor,
		TranslationManager: tm,
		MergeManager:       mm,
	}, nil
}

func (a *App) Start() {
	a.EventDistributor.BroadcastFeed()
	a.EventDistributor.BroadcastTranslatedEvents()
	logrus.Info("starting app")
}
