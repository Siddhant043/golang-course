package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
)

func main() {

	title, content := getNoteData()

	usernote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	usernote.Display()

	err = usernote.Save()

	if err != nil {
		fmt.Println("saving the notes failed")
		return
	}
	fmt.Println("saving the notes succeeded")
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")

	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text

}
