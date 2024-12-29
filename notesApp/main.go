package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error // it tells that a valid value must have a Save() method which takes no arguments and returns a error
}

// type displayer interface {
// 	Display() error
// }

// type outputable interface {
// 	Save()
// 	Display()
// }

type outputable interface {
	saver
	Display()
}

func main() {

	title, content := getNoteData()

	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
	}

	outputData(todo)

	usernote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(usernote)

}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("saving the data failed")
		return err

	}
	fmt.Println("saving the data succeeded")
	return nil
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
