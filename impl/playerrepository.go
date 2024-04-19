package impl

import "log"

func getStackerIds() []int {
	stackers := GetStackers()
	keys := make([]int, 0, len(stackers))
	for k := range stackers {
		keys = append(keys, k)
	}
	return keys
}

func GetStackers() map[int]string {
	

	stackers := map[int]string{
		115787026: "KK",
		107068931: "Water",
		29517560: "Barry",
		86818489: "Kodo",
		28193363: "Nils Olav",
		67029306: "Crysen",
		43181484: "Glete",
		203843297: "Scotadog",
		217230027: "TammyB",
		66095501: "Myst",
		86839100: "Chillkroete",
		101385227: "Lokie",
		117156158: "Thunderdyne",
		4294967295: "Glimmer of Dawn",
		228440453: "Szajtek",
		75126809: "Karenlyl",
		152273237: "Cory",
		159789890: "TTTM",
		82358773: "Sky",
		317725979: "RCE",
		160061113: "Gimmick",
		110008130: "Belle",
		315792033: "Dekait",
		111326429: "Hazel",
		84852390: "Weather-Event",
		119109388: "Bambohitter",
		114314986: "Sun",
		168600255: "Sandfish",
		270293472: "TeaGuvnor",



		//Ros
		76561198093600151: "PTS",
		54151884: "Tanner",
	}

	return stackers
}

func GetNameForId(Id int) string {
	return GetStackers()[Id]
}

func InitialStandings() map[int]int {

	standings := map[int]int{}
	for _, v  := range getStackerIds() {
		standings[v] = 5000
	}

	return standings
}

func PrettyPrintStandings(standings map[int]int) {
	for id, mmr  := range standings {
		log.Printf("%s: %d", GetNameForId(id), mmr)
	}
}