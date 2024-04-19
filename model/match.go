package model

import (
	"fmt"
)

type MatchWrapper struct {
	Result				Match		`json:"result"`
}


type Match struct {
	Id              	int    		`json:"match_id"`
	Players 		 	[]Player	`json:"players"`
	LobbyType			int			`json:"lobby_type"` //0 for unranked, 7 for ranked
	RadiantWin			bool		`json:"radiant_win"`
 }

 func (m Match) String() string {
	p := fmt.Sprintf("Match Id: %d, Players: %s", m.Id, m.Players)
	return p
}