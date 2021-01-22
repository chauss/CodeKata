package main

func weatherDataHasDayID(daysData []float64) bool {
	return len(daysData) > 0 && daysData[0] != -1
}

func filterDaysDataForTempValues(daysData []float64) []float64 {
	return []float64{daysData[0], daysData[1], daysData[2]}
}

func getWeatherTempData() [][]float64 {
	rawData := loadFile("weather.txt")
	allWeatherData := convertToFloatMatrix(rawData)

	var filteredWeatherData [][]float64
	for _, daysData := range allWeatherData {
		if weatherDataHasDayID(daysData) {
			filteredWeatherData = append(filteredWeatherData, filterDaysDataForTempValues(daysData))
		}
	}

	return filteredWeatherData
}

// GetMinTempSpread returns the id of the day with the minimum temp spread
func GetMinTempSpread() int {
	weatherData := getWeatherTempData()
	var result int
	var currentMinTempSpead float64 = 9999999
	for _, dayData := range weatherData {
		daysTempSpread := dayData[1] - dayData[2]
		if daysTempSpread < currentMinTempSpead {
			result = int(dayData[0])
			currentMinTempSpead = daysTempSpread
		}
	}

	return result
}

// GetMaxTempSpread returns the id of the day with the maximum temp spread
func GetMaxTempSpread() int {
	weatherData := getWeatherTempData()
	var result int
	var currentMaxTempSpead float64 = -273
	for _, dayData := range weatherData {
		daysTempSpread := dayData[1] - dayData[2]
		if daysTempSpread > currentMaxTempSpead {
			result = int(dayData[0])
			currentMaxTempSpead = daysTempSpread
		}
	}

	return result
}
