package main

import "fmt"

func main() {
	correctAnswer = 137
	fmt.Print(findAnswer(0, 1000))
}

type NullInt struct {
	int
	bool
}

var lastGuess NullInt
var correctAnswer int
var steps = 0

type Response int
const (
	correct Response = iota
	incorrect
	warmer
	colder
	same
)

// Check if this guess is warmer or colder than the last guess. It could (in theory) be the same.
// If there is no former guess, return incorrect.
func checkGuess(guess int) Response {
	if lastGuess.bool {
		return incorrect
	}
	if guess == correctAnswer {
		return correct
	}
	if Abs(correctAnswer-guess) > (correctAnswer - lastGuess.int) {
		return colder
	}
	if Abs(correctAnswer-guess) < (correctAnswer - lastGuess.int) {
		return warmer
	}
	return same
}

// Apparently there is no math.Abs for ints in Go :(
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findAnswer(min, max int) int {
	var guessA int = (min + max) / 2
	var guessB int = guessA + 1
	if(checkGuess(guessA) == correct) {
		return guessA
	}
	switch checkGuess(guessB) {
	case correct: { return guessB }
	case warmer:
		steps++
		return findAnswer(guessA, max)
	case colder:
		steps++
		return findAnswer(min, guessA)
	case same:
		var possibility int = (guessA + guessB) / 2;
		if checkGuess(possibility) == correct {
			return possibility
		} else {
			return possibility + 1
		}
	default:
		return -100000
	}
}