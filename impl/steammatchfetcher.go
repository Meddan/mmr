package impl

import (
	"encoding/json"
	. "example/hello/model"
	"fmt"
	"log"
	"net/http"
)


func GetMatch(matchId int) Match{
	
	key := GetValveApiKey()
	resp, err := http.Get(fmt.Sprintf("https://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/v1?match_id=%d&key=%s", matchId, key));

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()
 
 
	 var cResp MatchWrapper
	 //Decode the data
	 if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		 log.Fatal(err)
		 log.Fatal("ooopsss! an error occurred, please try again")
	 }

	 return cResp.Result;
}

func FindWinnersAndLosers(match Match) ([]int, []int) {
	if(match.LobbyType != 0) {	//Unranked game
		return []int{}, []int{}
	}
	radiantStackers, direStackers := getStackersOnTeams(match)
	if(match.RadiantWin) {
		return radiantStackers, direStackers
	} else {
		return direStackers, radiantStackers
	}

}

//Find the on each side of the match, returns a list of playerIds
func getStackersOnTeams(match Match) ([]int, []int) {
	stackerIds := getStackerIds()
	radiant := []int{}
	dire := []int{}
	for _, p := range match.Players {
		if(hasPlayerId(stackerIds, p.Id)) {
			if(p.TeamNumber == 0) {
				radiant = append(radiant, p.Id)
			} else {
				dire = append(dire, p.Id)
			}
		}
	}

	return radiant, dire
}

func CheckIfUnrankedStack(match Match) bool {
	if(match.LobbyType == 0) {
		//TODO: Handle outhouses
		partySize := 0

		stackerIds := getStackerIds()

		for _, p  := range match.Players {
			if(hasPlayerId(stackerIds, p.Id)) {
				partySize += 1;
			}
		}
		return partySize >= 5
	}
	return false
}

func hasPlayerId(playerIds []int, playerId int) bool {
    for _, v := range playerIds {
        if v == playerId {
            return true
        }
    }
    return false
}


func GetMatchesForPlayer(Id int, Amount int, StartId int) []ListMatch {
	key := GetValveApiKey()

	url := fmt.Sprintf("https://api.steampowered.com/IDOTA2Match_570/GetMatchHistory/v1?key=%s&account_id=%d&matches_requested=%d&start_at_match_id=%d", key, Id, Amount, StartId)

	resp, err := http.Get(url);

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()
 
 
	 var cResp ListMatchWrapperWrapper
	 //Decode the data
	 if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		 log.Fatal(err)
		 log.Fatal("ooopsss! an error occurred, please try again")
	 }

	 return cResp.Result.Matches;
	
}

func GetMatchesBySequenceNumber(seq int) {
	key := GetValveApiKey()

	url := fmt.Sprintf("https://api.steampowered.com/IDOTA2Match_570/GetMatchHistoryBySequenceNum/v1?key=%s&matches_requested=100&start_at_match_seq_num=%d", key, seq)

	log.Println(url)
}