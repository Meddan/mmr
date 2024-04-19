package model

import "fmt"

type ListMatchWrapperWrapper struct {
	Result ListMatchWrapper `json:"result"`
}

type ListMatchWrapper struct {
	Status int `json:"status"`
	NumResults int `json:"num_results"`
	TotalResults int `json:"total_results"`
	ResultsRemaining int `json:"results_remaining"`
    Matches []ListMatch `json:"matches"`
}

type ListMatch struct {
	MatchId int `json:"match_id"`
	StartTime int `json:"start_time"`
	LobbyType int `json:"lobby_type"`
	Players []Player `json:"players"`
	SeqNumber int `json:"match_seq_num"`
}

type OpenDotaListMatch struct {
	Id                int    `json:"match_id"`
 }

 func (m OpenDotaListMatch) String() string {
	p := fmt.Sprintf("Id: %d", m.Id)
	return p
}