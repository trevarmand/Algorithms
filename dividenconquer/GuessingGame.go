package dividenconquer

/**
Essentially a modification of binary search, this models a game where a player is given a game in which
they attempt to guess a number in some range as quickly as possible.
 */

//func main() {
//	// Test for validity
//	// Based on how the algorithm is designed, 0 will have the longest execution time.
//	for i := 0; i < 592989111; i++ {
//		correctAnswer = i
//		if findAnswer(0, 1000000000000000) != correctAnswer {
//			fmt.Print("\nFailed to find correctly answer: ", correctAnswer, " in ", steps, " steps")
//		}
//		if i % 1000000 == 0 {
//			fmt.Print("\nVerfied up to ", i / 1000000, " million")
//		}
//	}
//}

type NullInt struct {
	int
	bool
}

var lastGuess NullInt
var correctAnswer int

//Essentially an enum
type Response int
const (
	correct Response = iota
	incorrect
	warmer
	colder
)

// Check if this guess is warmer or colder than the last guess. It could (in theory) be the same.
// If there is no former guess, return incorrect.
func checkGuess(guess int) Response {
	// if lastGuess has not been initialized, this is first iteration, response doesn't matter.
	if lastGuess.bool {
		return incorrect
	}
	if guess == correctAnswer {
		return correct
	}
	curEval := Abs(correctAnswer - guess)
	prevEval := Abs(correctAnswer - lastGuess.int)
	if curEval > prevEval {
		lastGuess.int = guess
		return colder
	} else {
		lastGuess.int = guess
		return warmer
	}
}

// Apparently there is no math.Abs for ints in Go :(
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//Finds the variable "correctAnswer" within the given range via a binary search.
func findAnswer(min, max int) int {
	guessA := (min + max) / 2
	guessB := guessA + 1
	if checkGuess(guessA) == correct {
		return guessA
	}
	switch checkGuess(guessB) {
	case correct:
		return guessB
	case warmer:
		return findAnswer(guessB, max)
	case colder:
		return findAnswer(min, guessA)
	default:
		print("Failed to find val bc checkGuess returned a bad value")
		return -123456789
	}
}
