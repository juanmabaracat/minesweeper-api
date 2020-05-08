package game

type Mark string

const (
	FLAG     Mark = "flag"
	QUESTION Mark = "question"
	NONE     Mark = "none"
)

type Cells []Cell

type Cell struct {
	X           int  `json:"X"`
	Y           int  `json:"Y"`
	HasMine     bool `json:"has_mine"`
	WasRevealed bool `json:"was_revealed"`
	Mark        Mark `json:"mark"`
}
