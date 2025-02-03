package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	// title, content := getNoteData()
	// todoText := getUserInput("Todo text:")

	// todo, err := todo.New(todoText)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//userNote, err := note.New(title, content)

	// if err != nil {
	// 	fmt.Println((err))
	// 	return
	// }

	// err = outputData(todo)

	// if err != nil {
	// 	return
	// }

	// err = outputData(userNote)
}

func getTodoData() string {
	text := getUserInput("Todo text:")
	return text
}

func printSomething(value any) {

	typedVal, isInt := value.(int)

	if isInt {
		fmt.Println("Integer:", typedVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)

	// case float64:
	// 	fmt.Println("Float:", value)

	// case string:
	// 	fmt.Println(value)

	// 	//default:
	// 	//...
	// }
}

func outputData(data outputtable) {
	data.Display()
	saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

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
