package provider

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	riotType "github.com/Jordation/jsonl/provider/provider"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

type JsonlPump struct {
	Decoder *json.Decoder
	Config  *Configuration
}

type Configuration struct {
	Source *os.File
}

func NewPump(source string) (*JsonlPump, error) {
	fp, err := os.Open(filepath.Clean(source))
	if err != nil {
		return nil, err
	}

	d := json.NewDecoder(fp)
	cfg := &Configuration{
		Source: fp,
	}

	return &JsonlPump{
		Decoder: d,
		Config:  cfg,
	}, nil
}

func (p *JsonlPump) GetDeltas(from, to int) ([]*riotType.Event, error) {
	res := []*riotType.Event{}
	for ; from < to; from++ {
		temp := riotType.Event{}
		decErr := p.Decoder.Decode(&temp)

		if decErr == io.EOF {
			logrus.Info("EOF decoder returning early")
			return res, nil
		} else if decErr != nil {
			return nil, decErr
		}

		res = append(res, &temp)
	}

	return res, nil
}

func (p *JsonlPump) GetDelta() (*riotType.Event, error) {
	res := &riotType.Event{}
	if err := p.Decoder.Decode(res); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *JsonlPump) SpewDeltas(spewers []*riotType.Event) {
	for _, spewer := range spewers {
		spew.Dump(spewer)
	}
}
