package model

import (
	"fmt"
)


type Player struct {
	Id                	int   	`json:"account_id"`
	TeamNumber			int		`json:"team_number"` //0 for radiant, 1 for dire
}

func (p Player) String() string {
    return fmt.Sprintf("{Id: %d, Team number: %d}", p.Id,  p.TeamNumber)
}