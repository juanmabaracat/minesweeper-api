package game

type Mark string

const (
	FLAG     Mark = "FLAG"
	QUESTION Mark = "QUESTION"
)

type Cells []Cell

type Cell struct {
	x           int  `json:"x"`
	y           int  `json:"y"`
	hasMine     bool `json:"has_mine"`
	wasRevealed bool `json:"was_revealed"`
	mark        Mark `json:"mark"`
}

type Board struct {
	Width  int   `json:"width"`
	Height int   `json:"height"`
	Mines  int   `json:"mines"`
	Cells  Cells `json:"cells"`
}
