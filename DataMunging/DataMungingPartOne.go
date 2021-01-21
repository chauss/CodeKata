package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func loadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Can not open file \"%v\"", filename))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func convertStringsToFloats(strings []string, stringCleaner *regexp.Regexp) []float64 {
	var result []float64
	for _, value := range strings {
		cleanString := stringCleaner.ReplaceAllString(value, "")
		i, err := strconv.ParseFloat(cleanString, 64)
		if err != nil {
			result = append(result, -1)
		}
		result = append(result, i)
	}
	return result
}

func convertWeatherData(lines []string) [][]float64 {
	reg, err := regexp.Compile("[^0-9.]+")
	if err != nil {
		panic(err)
	}

	var result [][]float64
	for _, line := range lines {
		lineAsInts := convertStringsToFloats(strings.Fields(line), reg)
		result = append(result, lineAsInts)
	}
	return result
}

func weatherDataHasDayId(daysData []float64) bool {
	return len(daysData) > 0 && daysData[0] != -1
}

func filterDaysDataForTempValues(daysData []float64) []float64 {
	return []float64{daysData[0], daysData[1], daysData[2]}
}

func getWeatherTempData() [][]float64 {
	rawData := loadFile("weather.txt")
	allWeatherData := convertWeatherData(rawData)

	var filteredWeatherData [][]float64
	for _, daysData := range allWeatherData {
		if weatherDataHasDayId(daysData) {
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
