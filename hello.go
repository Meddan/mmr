package main

import (
	"example/hello/impl"
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


	
	//impl.GetMatchesBySequenceNumber(seq)
	//7321070778
	//testOneMatch()
	testMMR()
}

func testMMR() {
	matches := impl.GetMatchIds(28193363, 500, 0)

	standings := impl.InitialStandings()

	log.Printf("Matches: %d", len(matches))

	unrankedStacks := 0

	for _, match  := range matches {
		m := impl.GetMatch(match.Id)

		if(impl.CheckIfUnrankedStack(m)) {
			unrankedStacks +=1
			winners, losers := impl.FindWinnersAndLosers(m)
			/*
			if(len(winners) != 0 && len(losers) != 0) {
				log.Printf("Outhouse: %d", m.Id)
			}
			*/
			for _, w  := range winners {
				standings[w] = standings[w] + 25
			}

			for _, l  := range losers {
				standings[l] = standings[l] - 25
			}
		} else {
			if(m.LobbyType == 0) {
				log.Printf("Not stack: %d : %d", m.LobbyType, m.Id)
			}
		}
		
	}

	log.Printf("# of stacks: %d", unrankedStacks)
	impl.PrettyPrintStandings(standings)
	
}


func testOneMatch() {
	m := impl.GetMatch(7467669827)
	stack := impl.CheckIfUnrankedStack(m)

	log.Printf("p: %v", m.Players)

	log.Printf("stack: %t", stack)
}
