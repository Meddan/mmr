package impl

import (
	"encoding/json"
	. "example/hello/model"
	"fmt"
	"log"
	"net/http"
)

func GetMatchIds(playerId int, limit int, offset int) []OpenDotaListMatch{
	

	url := fmt.Sprintf("https://api.opendota.com/api/players/%d/matches?limit=%d&offset=%d", playerId, limit, offset)
	//log.Println(url)

	resp, err := http.Get(url)
 	//Handle Error
	if err != nil {
	   log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	var cResp []OpenDotaListMatch

	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal(err)
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	return cResp
}