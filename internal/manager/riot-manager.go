package manager

import (
	"github.com/Jordation/jsonl/internal/merge"
	"github.com/Jordation/jsonl/internal/translation"
	"github.com/Jordation/jsonl/internal/types"
	"github.com/Jordation/jsonl/provider"
	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/sirupsen/logrus"
)

type RiotFeed struct {
	Feed        provider.Feed
	EventStream chan *riotTypes.Event
	Manager     *MapManager
}

type MapManager struct {
	LastEventsProcessed map[string]int                // The game id and sequence number
	HeldEvents          map[string]map[int]*types.Map // translated events out of seq or waiting to be processed are stored here
	ActiveMaps          map[string]*types.Map
}

func NewRiotFeed(f provider.Feed) *RiotFeed {
	return &RiotFeed{
		Feed:        f,
		EventStream: f.Connect(),
		Manager:     NewMapManager(),
	}
}

func NewMapManager() *MapManager {
	return &MapManager{
		LastEventsProcessed: make(map[string]int),
		HeldEvents:          make(map[string]map[int]*types.Map),
		ActiveMaps:          make(map[string]*types.Map),
	}
}

func (f *RiotFeed) Start() {
	for event := range f.EventStream {
		gameID := event.Metadata.GameID.Value
		eventSeq := event.Metadata.SequenceNumber
		/*
			we can use this once we're receving messages one at at ime rather than just looping over the chan
			f mapData, exists := f.Manager.HasNextEvent(gameID); exists {
			merge.MergeRiotMap(f.Manager.ActiveMaps[gameID], mapData)
			} else
		*/
		delta := translation.TranslatorPipeline(event)
		if f.Manager.IsNextEvent(eventSeq, gameID) {
			merge.MergeRiotMap(f.Manager.ActiveMaps[gameID], delta)
		} else {
			logrus.Errorf("SEQ: %v, EVT: %++v", eventSeq, event)
			panic("somehow got an event out of sequence")
		}
	}
}

func (m *MapManager) IsNextEvent(seq int, gameID string) bool {
	return m.LastEventsProcessed[gameID] == seq-1
}

func (m *MapManager) HasNextEvent(gameID string) (*types.Map, bool) {
	lastSeq := m.LastEventsProcessed[gameID]
	if e, ok := m.HeldEvents[gameID][lastSeq+1]; ok {
		return e, true
	}
	return nil, false
}
