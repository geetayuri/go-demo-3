package usecases

import (
	"github.com/nguitarpb/7-solutions/modules/searchlist/entities"
	"regexp"
	"strings"

	// "fmt"
)

type searchlistUse struct {
	SearchlistRepo entities.SearchlistRepository
}

// Constructor
func NewSearchlistUsecase(searchlistRepo entities.SearchlistRepository) entities.SearchlistUsecase {
	return &searchlistUse{
		SearchlistRepo: searchlistRepo,
	}
}

func (u *searchlistUse) Search() (*entities.BeefRes, error) {
	res, err := u.SearchlistRepo.SearchListDb()
	if err != nil {
		return nil, err
	}

	cleanRes := cleanString(res)

	words := strings.Fields(strings.ToLower(cleanRes))
	wordCount := make(map[string]int32)

	// Count occurrences of each word
	for _, word := range words {
		wordCount[word]++
	}

	// Print the JSON as a native Go object
	// fmt.Println(wordCount)

	newRes := new(entities.BeefRes)
	newRes.Beef = wordCount

	return newRes, nil
}

func cleanString(input *string) string {
	// Step 1: Remove punctuation like ',' '.' using regex
	rePunct := regexp.MustCompile(`[,.]`)
	cleaned := rePunct.ReplaceAllString(*input, "")

	// Step 2: Replace multiple spaces/newlines with a single space
	reSpace := regexp.MustCompile(`\s+`)
	cleaned = reSpace.ReplaceAllString(cleaned, " ")

	// Step 3: Trim leading/trailing spaces
	return strings.TrimSpace(cleaned)
}
