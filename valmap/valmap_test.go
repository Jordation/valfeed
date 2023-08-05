package valmap

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/Jordation/jsonl/internal/pluck"
	"github.com/Jordation/jsonl/provider"
	riotTypes "github.com/Jordation/jsonl/provider/provider"
	"github.com/Jordation/jsonl/utils"
	"github.com/disintegration/imaging"
)

var pump, _ = provider.NewPump(utils.GetRelativePath("../utils/test.jsonl"))

func TestMapUse(t *testing.T) {
	events, err := pump.GetDeltas(30, 2000)
	if err != nil {
		log.Info(err)
	}
	dst, _ := GetImage("./img/blank.png")
	mark1, _ := GetImage("./img/col1.png")
	mark2, _ := GetImage("./img/col2.png")
	mark3, _ := GetImage("./img/col3.png")
	mark4, _ := GetImage("./img/col4.png")
	mark5, _ := GetImage("./img/col5.png")
	mark6, _ := GetImage("./img/col6.png")
	playerEvents := pluck.SplitDeltasByPlayer(events)
	targEvents := playerEvents[4]
	rounds := pluck.GetRoundTimings(events)
	points := GetPointsWithRoundsFromEvents(targEvents, rounds)
	for k, pointGroup := range points {
		for _, point := range pointGroup {
			switch k {
			case 1:
				dst = PasteOnPoint(dst, mark1, point)
			case 2:
				dst = PasteOnPoint(dst, mark2, point)
			case 3:
				dst = PasteOnPoint(dst, mark3, point)
			case 4:
				dst = PasteOnPoint(dst, mark4, point)
			case 5:
				dst = PasteOnPoint(dst, mark5, point)
			case 6:
				dst = PasteOnPoint(dst, mark6, point)
			default:
				log.Infof("%v round", k)
			}
		}
	}
	if err := SaveImage(dst, "./img/dotsonblank.png"); err != nil {
		log.Fatal(err)
	}
}

func TestResizeImage(t *testing.T) {
	im, _ := imaging.Open("./img/clean-lotus.png")
	im = imaging.Resize(im, 0, 1300, imaging.Linear)
	imaging.Save(im, "./img/resized-clean-lotus.png")
}

func TestPlotMapLocations(t *testing.T) {
	m := &riotTypes.APIMap{}
	dst, _ := GetImage("./img/blank.png")
	mark1, _ := GetImage("./img/col1.png")
	bytes, readErr := os.ReadFile("./lotusdata.json")
	if readErr != nil {
		log.Fatal(readErr)
	}
	if err := json.Unmarshal(bytes, m); err != nil {
		log.Fatal(err)
	}
	points := GetPointsFromAPIMap(m)
	for _, point := range points {
		dst = PasteOnPoint(dst, mark1, point)
	}

	if err := SaveImage(dst, "./img/dotsonblank.png"); err != nil {
		log.Fatal(err)
	}
}

func TestShowMeTheSeeqz(t *testing.T) {
	res, _ := pump.GetDeltas(0, 4000)
	for i, event := range res {
		if event.RoundStarted != nil {
			log.Info("round started on json #", i)
		}
	}
}
