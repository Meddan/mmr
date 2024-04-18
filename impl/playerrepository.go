package impl

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
	}

	return stackers
}

func GetNameForId(Id int) string {
	return GetStackers()[Id]
}