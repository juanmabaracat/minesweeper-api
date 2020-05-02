package game

type Status string

const (
	WON      Status = "WON"
	PLAYING  Status = "PLAYING"
	DEFEATED Status = "DEFEATED"
)

type Game struct {
	Id       int64  `json:"id"`
	PlayerId int64  `json:"player_id"`
	Status   Status `json:"status"`
	Board    Board  `json:"board"`
}
