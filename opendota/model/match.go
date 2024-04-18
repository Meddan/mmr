package model

import (
	"fmt"
)


type Match struct {
	Id              	int    		`json:"match_id"`
	Players 		 	[]Player	`json:"players"`
	LobbyType			int			`json:"lobby_type"`
 }

 func (m Match) TextOutput() string {
	p := fmt.Sprintf("Match Id: %d, Players: %s", m.Id, m.Players)
	return p
}