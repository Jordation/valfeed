package types

type AgentAbilityMapping struct {
	Name      string
	Abilities map[string]string
}

type APIAgent struct {
	UUID                      string        `json:"uuid,omitempty"`
	DisplayName               string        `json:"displayName,omitempty"`
	Description               string        `json:"description,omitempty"`
	DeveloperName             string        `json:"developerName,omitempty"`
	DisplayIcon               string        `json:"displayIcon,omitempty"`
	DisplayIconSmall          string        `json:"displayIconSmall,omitempty"`
	BustPortrait              string        `json:"bustPortrait,omitempty"`
	FullPortrait              string        `json:"fullPortrait,omitempty"`
	FullPortraitV2            string        `json:"fullPortraitV2,omitempty"`
	KillfeedPortrait          string        `json:"killfeedPortrait,omitempty"`
	Background                string        `json:"background,omitempty"`
	BackgroundGradientColors  []string      `json:"backgroundGradientColors,omitempty"`
	AssetPath                 string        `json:"assetPath,omitempty"`
	IsFullPortraitRightFacing bool          `json:"isFullPortraitRightFacing,omitempty"`
	IsPlayableCharacter       bool          `json:"isPlayableCharacter,omitempty"`
	IsAvailableForTest        bool          `json:"isAvailableForTest,omitempty"`
	IsBaseContent             bool          `json:"isBaseContent,omitempty"`
	Role                      *APIRole      `json:"role,omitempty"`
	Abilities                 []*APIAbility `json:"abilities,omitempty"`
	VoiceLine                 *APIVoiceLine `json:"voiceLine,omitempty"`
}
type APIRole struct {
	UUID        string `json:"uuid,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Description string `json:"description,omitempty"`
	DisplayIcon string `json:"displayIcon,omitempty"`
	AssetPath   string `json:"assetPath,omitempty"`
}
type APIAbility struct {
	Slot        string `json:"slot,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Description string `json:"description,omitempty"`
	DisplayIcon string `json:"displayIcon,omitempty"`
}
type APIMediaList struct {
	ID    int    `json:"id,omitempty"`
	Wwise string `json:"wwise,omitempty"`
	Wave  string `json:"wave,omitempty"`
}
type APIVoiceLine struct {
	MinDuration float64         `json:"minDuration,omitempty"`
	MaxDuration float64         `json:"maxDuration,omitempty"`
	MediaList   []*APIMediaList `json:"mediaList,omitempty"`
}
