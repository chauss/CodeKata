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

func convertStringsToInt(strings []string, stringCleaner *regexp.Regexp) []int {
	var result []int
	for _, value := range strings {
		cleanString := stringCleaner.ReplaceAllString(value, "")
		i, err := strconv.Atoi(cleanString)
		if err != nil {
			result = append(result, -1)
		}
		result = append(result, i)
	}
	return result
}

func createNumberRegex() *regexp.Regexp {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		panic(err)
	}
	return reg
}

func convertToFloatMatrix(lines []string) [][]float64 {
	reg := createNumberRegex()

	var result [][]float64
	for _, line := range lines {
		lineAsFloats := convertStringsToFloats(strings.Fields(line), reg)
		result = append(result, lineAsFloats)
	}
	return result
}

func convertToIntMatrix(lines []string) [][]int {
	reg := createNumberRegex()

	var result [][]int
	for _, line := range lines {
		lineAsInts := convertStringsToInt(strings.Fields(line), reg)
		result = append(result, lineAsInts)
	}
	return result
}

func convertToStringMatrix(lines []string) [][]string {
	var result [][]string
	for _, line := range lines {
		lineAsStrings := strings.Fields(line)
		result = append(result, lineAsStrings)
	}
	return result
}
