package main

import (
	opendota "example/hello/opendota/impl"
	"log"
)


 
func main() {
	/*
	match := impl.GetMatch(3842445305)
	unranked, _ := impl.CheckIfUnrankedStack(match, 28193363)

	log.Println(unranked)

	playerMatches, err := impl.GetMatchesForPlayer(28193363, 1000)

	if(err != nil) {
		log.Fatal(err)
	}

	log.Println(len(playerMatches))

	for i := 0; i < len(playerMatches); i++ {
		log.Println(playerMatches[i].Id)
	}
	*/

	/*
	match := impl.GetMatch(9311668);
	log.Println(match)

	stack := impl.CheckIfUnrankedStack(match)
	log.Println(stack)

	winners, losers := impl.FindWinnersAndLosers(match)
	log.Printf("Winners %v", winners)
	log.Printf("Losers %v", losers)
	*/

	//matches := impl.GetMatchesForPlayer(28193363, 5, 0)
	//log.Printf("Matches: %v", matches)

	/*
	stackers := impl.GetStackers()

	for id, name  := range stackers {
		
		matches := impl.GetMatchesForPlayer(id, 100, 0)
		
		for len(matches) != 0 {
			lastIndex := matches[len(matches)-1].MatchId
			
			//log.Printf("Matches got: %d, first index %d, last index %d", len(matches), matches[0].MatchId, lastIndex)

			matches = impl.GetMatchesForPlayer(id, 100, lastIndex-1)
			if(len(matches) == 0) {
				log.Printf("%s : %d", name, lastIndex)
			}
		}

	}*/


	matches, _ := opendota.GetMatchesForPlayer(28193363, 5000);

	log.Printf("Matches: %d", len(matches))
	
	//impl.GetMatchesBySequenceNumber(seq)
	//7321070778
}

