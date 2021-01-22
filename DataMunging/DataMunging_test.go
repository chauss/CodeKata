package main

import (
	"testing"
)

func TestMinTempSpread(t *testing.T) {
	expected := 14
	result := GetMinTempSpread()
	if result != expected {
		t.Errorf("Expected day %v but got day %v for the min temp spread", expected, result)
	}
}

func TestMaxTempSpread(t *testing.T) {
	expected := 9
	result := GetMaxTempSpread()
	if result != expected {
		t.Errorf("Expected day %v but got day %v for the max temp spread", expected, result)
	}
}

func TestForAgainstMinTeam(t *testing.T) {
	expected := "TeamName"
	result := GetForAgainstMinTeam()
	if result != expected {
		t.Errorf("Expected team %v but got team %v for the min ForAgainst difference", expected, result)
	}
}

func TestForAgainstMaxTeam(t *testing.T) {
	expected := "TeamName"
	result := GetForAgainstMaxTeam()
	if result != expected {
		t.Errorf("Expected team %v but got team %v for the max ForAgainst difference", expected, result)
	}
}
