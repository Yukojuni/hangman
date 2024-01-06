package github.com/Yukojuni/hangman

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// Randomly open a file then randomly choose a word
func LoadRandomWord(selectedFile) (string, error) {
	file, err := os.Open(selectedFile ) //remplacer par os qui lis l'input console
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
