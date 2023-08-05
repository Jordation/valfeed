package valmap

import (
	"image"
	"math"

	"github.com/Jordation/jsonl/internal/pluck"
	riotTypes "github.com/Jordation/jsonl/provider/provider"
	"github.com/disintegration/imaging"
	"github.com/sirupsen/logrus"
)

const (
	x_divisor = 9
	y_divisor = 9
)

var log *logrus.Logger

func init() {
	log = logrus.New()
}

func GetImage(path string) (image.Image, error) {
	return imaging.Open(path)
}

func SaveImage(src image.Image, path string) error {
	return imaging.Save(src, path)
}

func GetPointsFromEvents(events []*riotTypes.Event) []image.Point {
	res := []image.Point{}

	for _, event := range events {
		players := event.GetSnapshot().GetPlayers()

		for _, player := range players {
			for _, ts := range player.GetTimeseries() {
				TransformPosition(ts.Position)
				res = append(res, image.Point{
					X: int(ts.Position.X),
					Y: int(ts.Position.Y),
				})
			}
		}
	}

	return res
}

func GetPointsFromAPIMap(m *riotTypes.APIMap) []image.Point {
	res := []image.Point{}
	posis := []*riotTypes.Position{}
	for _, callout := range m.Callouts {
		posis = append(posis, callout.Location)
	}

	pluck.TransformPlayerPositionWithMapData(posis, m)
	for _, pos := range posis {
		res = append(res, image.Point{
			X: int(pos.X),
			Y: int(pos.Y),
		})
	}

	return res
}

func GetPointsWithRoundsFromEvents(events []*riotTypes.Event, roundChanges []int) map[int][]image.Point {
	res := map[int][]image.Point{}

	for _, event := range events {
		eventRound := pluck.RoundNumberFromSequence(event.Metadata, roundChanges)
		players := event.GetSnapshot().GetPlayers()

		for _, player := range players {
			for _, ts := range player.GetTimeseries() {
				TransformPosition(ts.Position)
				res[eventRound] = append(res[eventRound], image.Point{
					X: int(ts.Position.X),
					Y: int(ts.Position.Y),
				})
			}
		}
	}

	return res
}

func TransformPosition(pos *riotTypes.Position) {
	pos.X = math.Floor(pos.X / x_divisor)
	pos.Y = math.Floor(pos.Y/y_divisor) + 575
}

func PasteOnPoint(dst, src image.Image, point image.Point) image.Image {
	return imaging.Paste(dst, src, point)
}
