package main

import "fmt"

func main() {
	correctAnswer = 137
	fmt.Print("Found answer: ", findAnswer(0, 1000000000))
	fmt.Print("\nIn: ", steps, " steps")
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
	var curEval int = Abs(correctAnswer - guess)
	var prevEval int = Abs(correctAnswer - lastGuess.int)
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

func findAnswer(min, max int) int {
	fmt.Print("Operating on range ", min, " ", max, "\n")
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
		return findAnswer(min, guessB)
	case same:
		var possibility int = (guessA + guessB) / 2;
		if checkGuess(possibility) == correct {
			return possibility
		} else {
			return possibility + 1
		}
	default:
		print("fail")
		return -100000
	}
}