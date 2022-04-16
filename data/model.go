package data

type Option struct {
	Choice string `json:"choice"`
	Next   int    `json:"next"`
	Key    string `json:"key"`
}

type Stats struct {
	Stamina int `json:"stamina"`
	Skill   int `json:"skill"`
}

type Character struct {
	Name          string `json:"name"`
	CharacterType int    `json:"type"`
	Stats         Stats  `json:"stats"`
}

type Stage struct {
	Page       int         `json:"page"`
	Narrative  string      `json:"narrative"`
	Action     string      `json:"action"`
	Options    []Option    `json:"options"`
	Characters []Character `json:"characters"`
}

type DungeonData struct {
	Stages []Stage `json:"stages"`
}
