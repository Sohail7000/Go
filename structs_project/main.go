package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// "strings"
	"example.com/structs-example/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")

	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	
	if err!=nil{
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r\n")  // trimming linebreak for windows

	// var value string
	// fmt.Scanln(&value)

	return text
}
