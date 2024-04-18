package model

import (
	"fmt"
)


type PlayerMatch struct {
	Id                int    `json:"match_id"`
	PartySize 		  int	 `json:"party_size"`
 }

 func (m PlayerMatch) TextOutput() string {
	p := fmt.Sprintf("Match Id: %d, PartySize: %d", m.Id, m.PartySize)
	return p
}