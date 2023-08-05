package types

type APIMap struct {
	UUID         string        `json:"uuid,omitempty"`
	DisplayName  string        `json:"displayName,omitempty"`
	DisplayIcon  string        `json:"displayIcon,omitempty"`
	ListViewIcon string        `json:"listViewIcon,omitempty"`
	Splash       string        `json:"splash,omitempty"`
	AssetPath    string        `json:"assetPath,omitempty"`
	MapURL       string        `json:"mapUrl,omitempty"`
	XMultiplier  float64       `json:"xMultiplier,omitempty"`
	YMultiplier  float64       `json:"yMultiplier,omitempty"`
	XScalarToAdd float64       `json:"xScalarToAdd,omitempty"`
	YScalarToAdd float64       `json:"yScalarToAdd,omitempty"`
	Callouts     []*APICallout `json:"callouts,omitempty"`
}
type APILocation struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}
type APICallout struct {
	RegionName      string    `json:"regionName,omitempty"`
	SuperRegionName string    `json:"superRegionName,omitempty"`
	Location        *Position `json:"location,omitempty"`
}
