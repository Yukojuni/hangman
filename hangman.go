package hangman

import (
	"math/rand"
	"strings"
	"time"
)

// Structure used for play
type HangManData struct {
	Word          string
	ToFind        string
	Attempts      int
	TriedLetters  []rune
	HangmanStages []string
}

// Start a new game
func NewGame(word string, stages []string) *HangManData {
	game := &HangManData{
		Word:          strings.Repeat("_", len(word)),
		ToFind:        word,
		Attempts:      10,
		HangmanStages: stages,
	}
	revealInitialLetters(game)
	return game
}

// Reveals len(word)/ 2 -1 letters
func revealInitialLetters(game *HangManData) {
	rand.Seed(time.Now().UnixNano())
	lettersToReveal := len(game.ToFind)/2 - 1
	for i := 0; i < lettersToReveal; i++ {
		randomIndex := rand.Intn(len(game.ToFind))
		revealLetter(game, rune(game.ToFind[randomIndex]))
	}
}

func CheckEndGameCondition(game *HangManData) {
	if game.Attempts <= 0 {
		//perdu affiché nouvelle page
	} else if game.Word == game.ToFind {
		//gagné affiché nouvelle page
	}
}
