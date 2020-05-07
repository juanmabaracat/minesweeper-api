package game

type CreateRequest struct {
	PlayerId uint64 `json:"player_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Mines    int    `json:"mines"`
}
