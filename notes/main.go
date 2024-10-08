package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

//type outputtable interface {
//	Save() error
//	Display()
//}

func main() {
	result := add(1, 2)
	fmt.Println(result)
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")
	title, content := getNoteData()
	todoText := getTodo()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(todo)

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Saving the todo succeeded")

	err = outputData(userNote)

	if err != nil {
		return
	}
}

// Generic add function with a dynamic type any (which is the same as interface{})
func add[T int | float64 | string](a, b T) T {

	return a + b

	// aInt, aIsInt := a.(int)
	// bInt, bIsInt := b.(int)

	// if aIsInt && bIsInt {
	// 	return aInt + bInt
	// }

	// aFloat, aIsFloat := a.(float64)
	// bFloat, bIsFloat := b.(float64)

	// if aIsFloat && bIsFloat {
	// 	return aFloat + bFloat
	// }

	// return nil
}

// any value is allowed type
func printSomething(value interface{}) {
	intVal, ok := value.(int) // if value of type int, true is returned

	if ok {
		fmt.Println("Integer:", intVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float", value)
	// case string:
	// 	fmt.Println("String")
	// }
	fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return SaveData(data)
}

func SaveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Saving the note succeeded")
	return nil
}

func getTodo() string {
	text := getUserInput("Todo text: ")
	return text
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
