package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var score byte = 0
var tests = [12]string{
	"Do you use source control?",
	"Can you make a build in one step?",
	"Do you make daily builds?",
	"Do you have a bug database?",
	"Do you fix bugs before writing new code?",
	"Do you have an up-to-date schedule?",
	"Do you have a spec?",
	"Do programmers have quiet working conditions?",
	"Do you use the best tools money can buy?",
	"Do you have testers?",
	"Do new candidates write code during their interview?",
	"Do you do hallway usability testing?",
}

var (
	whiteText   = color.New(color.FgHiWhite).PrintfFunc()
	greenText   = color.New(color.FgHiGreen).PrintfFunc()
	redText     = color.New(color.FgHiRed).PrintlnFunc()
	greenTextLn = color.New(color.FgHiGreen).PrintlnFunc()
)

func scan(test string) bool {
	greenText("? ")
	whiteText(test + " (yes/no): ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		redText("Please type yes or no and then press enter")
		return scan(test)
	}

	input = strings.TrimSpace(input)

	switch strings.ToLower(input) {
	case "yes", "y":
		score++
		return true
	case "no", "n":
		return true
	default:
		redText("Please type yes or no and then press enter")
		return scan(test)
	}
}

func main() {
	greenTextLn("The Joel Test: 12 Steps to Better Code!")

	for _, s := range tests {
		scan(s)
	}

	var result string
	switch score {
	case 12:
		result = "Perfect!"
	case 11:
		result = "Tolerable!"
	default:
		result = "Unacceptable!"
	}

	greenTextLn("Your score is", score, "out of", len(tests), ", which is", result)
}
