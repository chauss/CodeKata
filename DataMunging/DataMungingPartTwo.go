package main

type footballForAgainsData struct {
	teamName     string
	forGoals     int
	againstGoals int
}

func dataIsTeamData(possibleTeamData []int) bool {
	return len(possibleTeamData) > 0 && possibleTeamData[0] != -1
}

func getFootballForAgainsData() []footballForAgainsData {
	rawData := loadFile("football.txt")
	dataAsStrings := convertToStringMatrix(rawData)
	dataAsNumbers := convertToIntMatrix(rawData)

	var resultData []footballForAgainsData

	for idx, data := range dataAsNumbers {
		if dataIsTeamData(data) {
			resultData = append(resultData, footballForAgainsData{dataAsStrings[idx][1], data[6], data[8]})
		}
	}

	return resultData
}

func forAgainsDiffOfTeam(team footballForAgainsData) int {
	diff := team.forGoals - team.againstGoals
	if diff < 0 {
		return -diff
	}
	return diff
}

// GetForAgainstMinTeam returns the name of the team with the min diff in ForAgainst goals
func GetForAgainstMinTeam() string {
	footballData := getFootballForAgainsData()

	if len(footballData) == 0 {
		return "NoTeams"
	}

	var minDiffTeam = footballData[0]

	for _, team := range footballData {
		if forAgainsDiffOfTeam(team) < forAgainsDiffOfTeam(minDiffTeam) {
			minDiffTeam = team
		}
	}

	return minDiffTeam.teamName
}

// GetForAgainstMaxTeam returns the name of the team with the max diff in ForAgainst goals
func GetForAgainstMaxTeam() string {
	footballData := getFootballForAgainsData()

	if len(footballData) == 0 {
		return "NoTeams"
	}

	var minDiffTeam = footballData[0]

	for _, team := range footballData {
		if forAgainsDiffOfTeam(team) > forAgainsDiffOfTeam(minDiffTeam) {
			minDiffTeam = team
		}
	}

	return minDiffTeam.teamName
}
