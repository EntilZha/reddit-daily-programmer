package main

import (
	"fmt"
	"strings"
)

// Challenge: Given a string of words and a string of letters.
// Find the largest string(s) that are in the 1st string of words that can be formed from the letters in the 2nd string.

var input1Words = "hello yyyyyyy yzyzyzyzyzyz mellow well yo kellow lellow abcdefhijkl hi is yellow just here to add strings fellow lellow llleow"
var input1Letters = "l e l o h m f y z a b w"
var input2Words = "sad das day mad den foot ball down touch pass play"
var input2Letters = "z a d f o n"
var input3Words = "a of lol"
var input3Letters = "o f"

func main() {
	answer := bruteForce(input1Words, input1Letters)
	fmt.Println("The answer is:")
	fmt.Println(answer)
	fmt.Println(len(answer))
}

func bruteForce(words string, letters string) []string {
	answers := make([]string, 0)
	wordsArray := strings.Split(words, " ")
	lettersArray := strings.Split(letters, " ")
	lettersDictionary := make(map[string]int64)
	minLength := 0
	for _, letter := range lettersArray {
		lettersDictionary[letter] += 1
	}
	for _, word := range wordsArray {
		if isValidWord(word, &lettersDictionary) {
			if len(word) > minLength {
				minLength = len(word)
				answers = nil
				answers = append(answers, word)
			} else if len(word) == minLength {
				answers = append(answers, word)
			}
		}
	}
	return answers
}

func isValidWord(word string, lettersDictionary *map[string]int64) bool {
	var tempLetters = make(map[string]int64)
	var condition = true
	copyMap(lettersDictionary, &tempLetters)
	for _, letter := range word {
		tempLetters[string(letter)] -= 1
		if tempLetters[string(letter)] < 0 {
			condition = false
			break
		}
	}
	return condition
}

func copyMap(src *map[string]int64, dest *map[string]int64) {
	for k, v := range *src {
		var m = *dest
		m[k] = v
	}
}
