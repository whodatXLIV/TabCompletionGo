package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var cList []string

func main() {

	var port int

	cList = append(cList, "apples", "banana", "coffee", "celery", "milk", "macaroni")

	// How to use flags
	// os.Args[0] is always the file name of the program

	//Can add in own flag usage called by main.exe -help
	flag.Usage = func() {
		fmt.Println("Usage of ", os.Args[0], ":")
		fmt.Printf("\t main.exe file1 file 2 ... \n")
		flag.PrintDefaults()
	}

	//Set up flags to call
	flag.IntVar(&port, "p", 8000, "Specify Port to use. Default is 8000.")
	flag.Parse() //Parse the flag inputs

	// Look into flags, https://thenewstack.io/cli-command-line-programming-with-go/

	// The following block of code is for getting single word inputs from the cmd line
	/*
			var	guessColor string

			const favColor = "blue"
		for {
				fmt.Println("Guess my favorite color:")
				if _, err := fmt.Scanf("%s", &guessColor); err != nil {
					fmt.Printf("%s\n", err)
					return
				}
				if favColor == guessColor {
					fmt.Printf("%q is my favorite color!\n", favColor)
					return
				}
				fmt.Printf("Sorry, %q is not my favorite color. Guess again.\n", guessColor)
			}*/

	// The following code uses bufio which allows more complex inputs
	// It is an echo program
	/*
		scanner := bufio.NewScanner(os.Stdin)
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}

		for scanner.Scan() {
			line := scanner.Text()
			if line == "exit" {
				os.Exit(0)
			}
			fmt.Println(line)
		}*/

	readerTab := bufio.NewReader(os.Stdin)
	readerLine := bufio.NewReader(os.Stdin)
	tabComplete, _ := readerTab.ReadString('\t')
	lineComplete, _ := readerLine.ReadString('\n')
	fmt.Println(tabComplete)
	fmt.Println("---------------")
	fmt.Println(lineComplete)

}
