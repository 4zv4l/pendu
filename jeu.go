package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func clearscreen() {
	if runtime.GOOS != "windows" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// count number of line
func nline() int {
	file, _ := os.Open("words.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	file.Close()
	return lineCount
}

// show menu
func menu() {
	print(`
Hi player !
Welcome in my game.
I hope you'll enjoy it :)
		
Here are the rules :
	- my coding skills aren't perfect so it mays have some bugs
	- take your time to focus and think which word it could be
	- your mouse is useless here
	- all words have to be in a "words.txt" file at the same place at this game
	- you have to write the all word to win
	- most important ENJOY
	- write "STOP" to get the answer if you wanna fed up
	
Ready to play ?

Press enter to continue...`)
	fmt.Scanln()
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// chose a random word to pay
func randomword(words []string, line int) string {
	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(line-1+1)]
	return word
}

func affiche(word string, letter string, lettre []string) []string {

	var find bool
	find = true

	// regarde si met plusieur fois la meme
	for J := 0; J < len(lettre); J++ {
		if letter == lettre[J] {
			find = false
		}
	}
	if find == true {
		lettre = append(lettre, letter)
	} else {
		println("already tried this one...")
	}

	for i := 0; i < len(word); i++ { // for every letter in the word
		for j := 0; j < len(lettre); j++ { // for every letter it the array
			if lettre[j] == string(word[i]) {
				print(lettre[j], " ")
				find = true
				break
			}

			find = false
		}
		if find == false {
			print("- ")
		}
	}

	return lettre
}

// main game
func showplateau(word string, line int, live int) bool {

	letter := "a"
	var tab = []string{}

	for i := 0; i < len(word); i++ {
		print("_ ")
	}
	for live != 0 {

		print("\n\n")
		print("try a letter or a word : ")
		fmt.Scan(&letter)
		if letter == word {
			return true
		} else if strings.ContainsAny(word, letter) == false {

			live--
			println("live :", live)
			print("\n")
		} else if strings.ContainsAny(word, letter) == true {
			println("great!")
			print("\n")
		}
		if live == 0 {
			break
		} else if letter == "STOP" {
			lose(word)
		}
		tab = affiche(word, letter, tab)
	}
	if live == 0 {
		return false
	}
	return false
}

func win() {
	println("\nCongratualion !")
	time.Sleep(time.Second * 3)
	os.Exit(2)
}

func lose(word string) {
	println("Sorry you'll do better next time !")
	println("The word was :", word)
	time.Sleep(time.Second * 5)
	os.Exit(1)
}

// call main function
func main() {
	clearscreen()

	file, err := os.Open("words.txt")
	if err != nil {
		print("error while opening file...")
		time.Sleep(2 * time.Second)
		os.Exit(1)
	}

	line := nline()               // count number of line
	words := make([]string, line) // create a array with a variable size
	scanner := bufio.NewScanner(file)
	for i := 0; i < line; i++ {
		scanner.Scan()
		words[i] = scanner.Text()
	}

	menu()
	print("\n")
	var live int
	println("Difficulty :")
	println("1. Easy -> 10 lives")
	println("2. Hard -> 5 lives")
	println("3. God -> infinity lives (easiest)")
	print("> ")
	fmt.Scanln(&live)
	if live == 1 {
		live = 10
	} else if live == 2 {
		live = 5
	} else if live == 3 {
		live = -1
	}
	word := randomword(words, line)
	score := showplateau(word, line, live)
	if score == true {
		win()
	} else {
		lose(word)
		println("the word was :", word)
	}
	file.Close()
}
