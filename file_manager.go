package hangman

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

// Randomly open a file then randomly choose a word
func LoadRandomWord(difficulty string) (string, error) {
	selectedFile := ""
	if difficulty == "Easy" {
		selectedFile = "word1.txt"
	} else if difficulty == "Medium" {
		selectedFile = "word2.txt"
	} else if difficulty == "Hard" {
		selectedFile = "word3.txt"
	}

	file, err := os.Open(selectedFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return words[rand.Intn(len(words))], nil
}

func LoadHangmanStages(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var stages []string
	var currentStage strings.Builder

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		currentStage.WriteString(scanner.Text() + "\n")
		lineCount++
		if lineCount == 8 {
			stages = append(stages, currentStage.String())
			currentStage.Reset()
			lineCount = 0
		}
	}
	if lineCount > 0 {
		stages = append(stages, currentStage.String())
	}

	return stages, scanner.Err()
}
