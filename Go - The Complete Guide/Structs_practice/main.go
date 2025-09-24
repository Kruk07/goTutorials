package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type Saver interface {
	Save() error
}

type outputtable interface {
	Saver
	Display()
}

func main() {

	printSomething(1)
	printSomething(1.5)
	printSomething("Hello!")

	title, content := getNoteData()
	todoText := getUserInput("Todo text : ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	err = outputData(userNote)

}

func printSomething(value interface{}) {
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// }
	intVal, ok := value.(int)

	if !ok {
		fmt.Println("Integer:", intVal)
		return
	}

	float64Val, ok := value.(float64)

	if !ok {
		fmt.Println("Float:", float64Val)
		return
	}

	stringVal, ok := value.(string)

	if !ok {
		fmt.Println("String:", stringVal)
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	saveData(data)
}

func saveData(data Saver) error {

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
