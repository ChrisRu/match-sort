package main

import (
	"fmt"
	"sort"
	"strings"
)

// Match object after string comparison
type Match struct {
	Match       bool
	Value       string
	Start       int
	End         int
	MatchLength int
	TypeIndex   int
}

var rankingSystem = []func(s string, substr string) Match{
	CaseSensitiveEquals,
	Equals,
	StartsWith,
	WordStartsWith,
	CaseStartsWith,
	Contains,
	Acronym,
	SimpleMatch}

func main() {
	list := []string{"hey", "what", "is", "hanging"}
	fmt.Println(list)
	fmt.Printf("Search the array: ")

	var name string
	fmt.Scanln(&name)

	fmt.Println(Sort(list, name))
}

// Sort the list by match
func Sort(list []string, match string) []string {
	filteredList := []Match{}

	// Get a match for every string in the list
	for _, word := range list {
		for _, method := range rankingSystem {
			res := method(word, match)
			if res.Match {
				filteredList = append(filteredList, res)
				break
			}
		}
	}

	// Sort the filtered list
	sort.Slice(filteredList, func(i, j int) bool {
		return filteredList[i].TypeIndex < filteredList[j].TypeIndex
	})

	sortedList := []string{}
	for _, match := range filteredList {
		sortedList = append(sortedList, match.Value)
	}
	return sortedList
}

func CaseSensitiveEquals(s string, substr string) Match {
	result := Match{
		Match:     false,
		TypeIndex: 1,
		Value:     s,
	}

	if s == substr {
		length := len(substr)

		result.Match = true
		result.Start = 0
		result.MatchLength = length
		result.End = length
	}

	return result
}

func Equals(s string, substr string) Match {
	result := Match{
		Match:     false,
		TypeIndex: 2,
		Value:     s,
	}

	if strings.ToLower(s) == strings.ToLower(substr) {
		length := len(substr)

		result.Match = true
		result.Start = 0
		result.MatchLength = length
		result.End = length
	}

	return result
}

func StartsWith(s string, substr string) Match {
	result := Match{
		Match:     false,
		TypeIndex: 3,
		Value:     s,
	}

	if strings.HasPrefix(s, substr) {
		length := len(substr)

		result.Match = true
		result.Start = 0
		result.MatchLength = length
		result.End = length
	}

	return result
}

func WordStartsWith(s string, substr string) Match {
	result := Match{
		Match:     false,
		TypeIndex: 4,
		Value:     s,
	}

	// TODO: Implement

	return result
}

func CaseStartsWith(s string, substr string) Match {
	result := Match{
		Match:     false,
		TypeIndex: 5,
		Value:     s,
	}

	// TODO: Implement

	return result
}

func CaseAcronym(s string, substr string) Match {
	result := Match{
		Match:     false,
		TypeIndex: 6,
		Value:     s,
	}

	// TODO: Implement

	return result
}

func Contains(s string, substr string) Match {
	result := Match{
		Match:     false,
		TypeIndex: 7,
		Value:     s,
	}

	// TODO: Implement

	return result
}

func Acronym(s string, substr string) Match {
	result := Match{
		Match:     false,
		TypeIndex: 8,
		Value:     s,
	}

	// TODO: Implement

	return result
}

// Simple match where string contains substring anywhere
func SimpleMatch(s string, substr string) Match {
	result := Match{
		Match:     false,
		TypeIndex: 9,
		Value:     s,
	}

	startIndex := strings.Index(strings.ToLower(s), strings.ToLower(substr))
	if startIndex != -1 {
		length := len(substr)

		result.Match = true
		result.MatchLength = length
		result.Start = startIndex
		result.End = startIndex + length
	}

	return result
}
