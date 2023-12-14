package github.com/Yukojuni/hangman

import (
	"strings"
)

// ASCIIArtMap associe chaque caractère à sa représentation en art ASCII.
var ASCIIArtMap map[rune][]string

// InitializeASCIIArtMap initialise la carte d'art ASCII à partir d'une chaîne contenant l'art ASCII.
func InitializeASCIIArtMap(asciiArtContent string) {
	ASCIIArtMap = make(map[rune][]string)
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	hauteurLigne := 9 // Nombre de lignes par caractère dans l'art ASCII

	lignes := strings.Split(asciiArtContent, "\n")
	for i, char := range alphabet {
		ligneDebut := i * hauteurLigne
		ligneFin := ligneDebut + hauteurLigne
		ASCIIArtMap[rune(char)] = lignes[ligneDebut:ligneFin]
	}
}

// GetASCIIArt retourne la représentation en art ASCII d'un caractère donné.
func GetASCIIArt(char rune) []string {
	if art, ok := ASCIIArtMap[char]; ok {
		return art
	}
	// Retourne un art ASCII vide pour les caractères non reconnus
	return strings.Split(strings.Repeat(" ", 9), "\n")
}
