package main

import (
	"errors"
	"fmt"
	"os"
)

func handsomeMan() error {
	bytes, readError := os.ReadFile("banner.txt")
	if readError != nil {
		return readError
	}
	clooneyOfDogs := string(bytes)
	fmt.Println(clooneyOfDogs)
	return nil
}

func playFetch(args []string) error {
	if args == nil || len(args) < 2 {
		return errors.New("There's really only one thing to do here...\n")
	}
	switch args[1] {
	case "throw":
		throwBall()
	case "help":
		getHelp()
	default:
		return fmt.Errorf("Just throw the ball for him: %v\n", args[1])
	}
	return nil
}

func throwBall() {
	fmt.Println("You throw the ball as far as you can...")
	fmt.Println("...")
	fmt.Println("*sprints out toward the ball*")
	fmt.Println("...")
	fmt.Println("*trots back*")
	fmt.Println("...")
	fmt.Println("*drops the ball at your feet*")
}

func getHelp() {
	fmt.Println("throw")
}

func main() {
	tired := handsomeMan()
	if tired != nil {
		fmt.Printf("%v", tired)
		os.Exit(1)
	}
	unableToThrow := playFetch(os.Args)
	if unableToThrow != nil {
		fmt.Printf("%v", unableToThrow)
		os.Exit(1)
	}
}