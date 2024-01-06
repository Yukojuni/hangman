package github.com/Yukojuni/hangman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/rivo/tview"
)

// Update the game
func UpdateGame(game *HangManData, input string) {
	input = strings.ToLower(input)
	if len(input) == 1 {
		if strings.ContainsRune(string(game.TriedLetters), rune(input[0])) {
		}
		game.TriedLetters = append(game.TriedLetters, rune(input[0]))
		processSingleLetter(game, rune(input[0]))
	} else if len(input) > 1 {
		processWholeWord(game, input)
	}
}

// Checks if the entered letter is present in the word
func processSingleLetter(game *HangManData, letter rune) {
	if strings.ContainsRune(game.ToFind, letter) {
		if !strings.ContainsRune(game.Word, letter) {
			revealLetter(game, letter)
		}
	} else {
		game.Attempts -= 1
	}
}

// Checks if the word entered is equal to "STOP", the word found or if it is different
func processWholeWord(app *tview.Application, game *HangManData, word string) {
	if word == game.ToFind {
		game.Word = game.ToFind
	} else {
		game.Attempts -= 2
	}
}

// Reveal the letters found in the words
func revealLetter(game *HangManData, letter rune) {
	for i, l := range game.ToFind {
		if l == letter {
			game.Word = game.Word[:i] + string(letter) + game.Word[i+1:]
		}
	}
}
