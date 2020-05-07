package game

type Mark string

const (
	FLAG     Mark = "FLAG"
	QUESTION Mark = "QUESTION"
	NONE     Mark = "NONE"
)

type Cells []Cell

type Cell struct {
	X           int  `json:"X"`
	Y           int  `json:"Y"`
	HasMine     bool `json:"has_mine"`
	WasRevealed bool `json:"was_revealed"`
	Mark        Mark `json:"Mark"`
}
