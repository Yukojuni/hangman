package github.com/Yukojuni/hangman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/rivo/tview"
)

// Update the game
func UpdateGame(app *tview.Application, game *HangManData, input string) {
	input = strings.ToLower(input)
	if len(input) == 1 {
		if strings.ContainsRune(string(game.TriedLetters), rune(input[0])) {
		}
		game.TriedLetters = append(game.TriedLetters, rune(input[0]))
		processSingleLetter(game, rune(input[0]))
	} else if len(input) > 1 {
		processWholeWord(app, game, input)
	}
}

// Checks if the entered letter is present in the word
func processSingleLetter(game *HangManData, letter rune) {
	if strings.ContainsRune(game.ToFind, letter) {
		if !strings.ContainsRune(game.Word, letter) {
			revealLetter(game, letter)
		}
	} else {
		game.Attempts--
	}
}

// Checks if the word entered is equal to "STOP", the word found or if it is different
func processWholeWord(app *tview.Application, game *HangManData, word string) {
	if word == "stop" {
		err := game.SaveToFile()
		if err != nil {
			fmt.Println("Error saving game:", err)
			return
		}
		saveScreen := tview.NewModal().
			SetText("Your game has been saved.\n\nSee you soon !").
			AddButtons([]string{"Quit"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				app.Stop()
			})
		app.SetRoot(saveScreen, false)
	} else if word == game.ToFind {
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

// Save the current game to a save.txt file
func (game *HangManData) SaveToFile() error {
	filename := "/home/lucasmcn/dev/Hangman/save/save.txt"
	data, err := json.Marshal(game)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// Load game state from file
func LoadGameFromFile(filename string) (*HangManData, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var game HangManData
	err = json.Unmarshal(data, &game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}
