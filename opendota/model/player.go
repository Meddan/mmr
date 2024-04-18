package model

import (
	"fmt"
)


type Player struct {
	Id                	int   	`json:"account_id"`
	Name 			  	string	`json:"personaname"`
	PartyId				int		`json:"party_id"`
 }

 func (p Player) TextOutput() string {
	s := fmt.Sprintf("Id: %d Name: %s PartyId: %d", p.Id, p.Name, p.PartyId)
	return s
}

func (p Player) String() string {
    return p.TextOutput()
}