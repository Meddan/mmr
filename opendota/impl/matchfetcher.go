package impl

import (
	"encoding/json"
	"errors"
	. "example/hello/opendota/model"
	"fmt"
	"log"
	"net/http"
)

func GetMatchIds(playerId int, limit int, offset int) []PlayerMatch{
	

	url := fmt.Sprintf("https://api.opendota.com/api/players/%d/matches?limit=%d&offset=%d", playerId, limit, offset)
	//log.Println(url)

	resp, err := http.Get(url)
 	//Handle Error
	if err != nil {
	   log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	var cResp []PlayerMatch

	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal(err)
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	return cResp
}

func GetMatch(matchId int) Match{
	resp, err := http.Get(fmt.Sprintf("https://api.opendota.com/api/matches/%d", matchId))

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()
	 //sb := string(body)
	 //log.Printf(sb)
 
 
	 var cResp Match
	 //Decode the data
	 if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		 log.Fatal(err)
		 log.Fatal("ooopsss! an error occurred, please try again")
	 }

	 return cResp;
}

func CheckIfUnrankedStack(match Match, playerId int) (bool, error){
	if(match.LobbyType == 0) {
		log.Println(match)
		partyId, err := findPartyId(match.Players, playerId)
		if(err != nil) {
			return false, fmt.Errorf("could not find player %d in match %d", playerId, match.Id)
		}

		partySize := countPlayersWithSameParty(match.Players, partyId);
		return partySize == 5, nil;
	}
	return false, nil
}

func findPartyId(players []Player, id int) (int, error) {
    for _, player := range players {
        if player.Id == id {
            return player.PartyId, nil
        }
    }
	log.Printf("Could not find player %d in game", id)
	return -1, errors.New("could not find player in match")
}

func countPlayersWithSameParty(players []Player, partyId int) int {
    count := 0
    for _, player := range players {
        if player.PartyId == partyId {
            count++
        }
    }
    return count
}

func GetMatchesForPlayer(playerId int, maxNumberOfMatches int) ([]Match, error){
	matches := make([]Match, 0, maxNumberOfMatches)

	matchIds := GetMatchIds(playerId, maxNumberOfMatches, 0);
	//data, _ := json.Marshal(matchIds);
	//log.Printf("Got matchIds: %s", data)
	log.Println(len(matchIds))

	for _, matchId := range matchIds {
		//log.Printf("matchid: %d", matchId)
		match := GetMatch(matchId.Id)
		//log.Printf("Match: %d", match.Id)
		unrankedStack, err := CheckIfUnrankedStack(match, playerId)
		if(err != nil) {
			return nil, err
		}
		if(unrankedStack) {
			matches = append(matches, match)
		}
	}

	return matches, nil;
}
