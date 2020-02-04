package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type result struct {
	computerChoice string
	yourChoice     string
	status         bool
	message        string
}

func printToScreen(s string) {
	fmt.Println(s)
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	return input
}

func analyze(i string) result {
	var yourChoice string
	computerChoice := "rock"
	var message string
	var status bool

	if i == "paper" {
		yourChoice = i
		computerChoice = "rock"
		message = "You win"
		status = true
	}

	if i == "rock" {
		yourChoice = i
		computerChoice = "rock"
		message = "It's a draw"
		status = false
	}

	boo := result{
		computerChoice: computerChoice,
		yourChoice:     yourChoice,
		message:        message,
		status:         status,
	}

	return boo
}

func start() {
	printToScreen("What is your choice?\n")
	input := readInput()
	result := analyze(input)
	fmt.Printf("%+v", result)
}
