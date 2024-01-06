// ecrire une fonction qui ouvre le fichier puis qui parmi ce fichier séléctionne une lettre

// cache et reveal len /2 -1 lettre

// attemps = x
// 	alors afficher fichier lignes multiplier par x

// Structure used for play
type HangManData struct {
	Word          string
	ToFind        string
	Attempts      int
	TriedLetters  []rune
	HangmanStages []string
}

	word, err := hangman.LoadRandomWord(os.Args[4])
	if err != nil {
		log.Fatalf("Error loading random word: %v", err)
	}
	NewGame(word)

	func NewGame(word string) {
		game := &HangManData{
			Word:          strings.Repeat("_", len(word)),
			ToFind:        word,
			Attempts:      10,
			HangmanStages: stages,
		}
		revealInitialLetters(game)
		return game
	}

	func revealInitialLetters(game *HangManData) {
		rand.Seed(time.Now().UnixNano())
		lettersToReveal := len(game.ToFind)/2 - 1
		for i := 0; i < lettersToReveal; i++ {
			randomIndex := rand.Intn(len(game.ToFind))
			revealLetter(game, rune(game.ToFind[randomIndex]))
		}
	}

	func revealLetter(game *HangManData, letter rune) {
		for i, l := range game.ToFind {
			if l == letter {
				game.Word = game.Word[:i] + string(letter) + game.Word[i+1:]
			}
		}
	}


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
	func processWholeWord(game *HangManData, word string) {
		if word == game.ToFind {
			game.Word = game.ToFind
		} else {
			game.Attempts -= 2
		}
	} 

	// regarde si mot ou lettre

	// 	si lettre compare avec chacune des lettres du mot,

	// 		si présente dévoiler chaque occurence
	// 		si non attemps +1

	// 	si mot compare input avec mot du jeu
	// 		si égale alors gagné 
	// 		sinon attemps +2

	func CheckEndGameCondition(game *hangman.HangManData) {
		if game.Attempts <= 0 {
			//perdu
		} else if game.Word == game.ToFind {
			//gagné
		}
	}