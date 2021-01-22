# What to do (simply said)

## Part One

Read the weather data from weather.txt and calculate the day with the lowest and the highest weather spread (difference between min and max temprature).

First column is the day, second column the max temp, third column the min temp.

Implement the following functions:

```go
// GetMinTempSpread returns the id of the day with the minimum temp spread
func GetMinTempSpread() int

// GetMaxTempSpread returns the id of the day with the maximum temp spread
func GetMaxTempSpread() int
```

## Part Two

Read the football data from football.txt and calculate the team with the smallest and biggest difference in "for" (Column F) and "against" (Column A)V goals.

Implement the following functions:

```go
// GetForAgainstMinTeam returns the name of the team with the min diff in ForAgainst goals
func GetForAgainstMinTeam() string

// GetForAgainstMaxTeam returns the name of the team with the max diff in ForAgainst goals
func GetForAgainstMaxTeam() string
```

## Part Three

Take a closer look at the parts that you were able to reuse from Part One in Part Two. Did you start Part One so you could easily use the common logic or did you have to refactor a lot?

# Kata Questions

- To what extent did the design decisions you made when writing the original programs make it easier or harder to factor out common code?

- Was the way you wrote the second program influenced by writing the first?

- Is factoring out as much common code as possible always a good thing? Did the readability of the programs suffer because of this requirement? How about the maintainability?

# Origin

This Code Cata comes from http://codekata.com/kata/kata04-data-munging/
