// bed is a bad editor
package main

import (
	"bufio"
	"fmt"
	"os"
)

// readInput reads input and makes sure we don't want to quit
func readInput(s *bufio.Scanner, cmd *string, cursor string) bool {
	fmt.Print(cursor)
	s.Scan()
	*cmd = s.Text()
	return *cmd != ""
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: no file specified")
		os.Exit(1)
	}

	fname := os.Args[1]

	f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error: could not open file: " + err.Error())
		os.Exit(1)
	}
	defer f.Close()

	cmd := ""
	text := ""

	scn := bufio.NewScanner(os.Stdin)

	for readInput(scn, &cmd, ": ") {
		switch {
		case cmd == "i":
			// insert mode
			for readInput(scn, &cmd, "- ") {
				text += cmd
				text += "\n"
				cmd = ""
			}
		case cmd == "w" && len(text) > 0:
			// write text if any
			bytes, err := f.Write([]byte(text))
			if bytes < len(text) || err != nil {
				fmt.Println("Error: failed to write all bytes to file")
				f.Close()
				os.Exit(1)
			}
			f.Sync()
			text = ""
		case cmd == "r":
			// reset current text value
			text = ""
		case cmd == "p":
			// print current text value
			fmt.Print(text)
		case cmd == "h":
			// print help message
			fmt.Println("how to use `bed`:\npress i for insert mode\npress w to write to file\npress r to reset current text buffer\npress p to print current text buffer\npress enter to quit current function")
		default:
			// anything we don't recognize
			fmt.Println("?")
		}
		cmd = ""
	}
}
