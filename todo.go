package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TodoItem struct {
	Text string
	Done bool
}

type TodoList []TodoItem

func (list TodoList) Print() {
	for i, item := range list {
		status := " "
		if item.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d. %s\n", status, i+1, item.Text)
	}
}

func main() {
	todos := TodoList{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Todo List")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		} else if input == "list" {
			todos.Print()
		} else if strings.HasPrefix(input, "add ") {
			itemText := strings.TrimPrefix(input, "add ")
			todos = append(todos, TodoItem{Text: itemText})
			fmt.Println("Added:", itemText)
		} else if strings.HasPrefix(input, "done ") {
			itemIndex := strings.TrimPrefix(input, "done ")
			index := parseIndex(itemIndex)
			if index >= 0 && index < len(todos) {
				todos[index].Done = true
				fmt.Printf("Marked item %d as done.\n", index+1)
			} else {
				fmt.Println("Invalid index.")
			}
		} else {
			fmt.Println("Invalid command.")
		}
	}
}

func parseIndex(indexStr string) int {
	var index int
	_, err := fmt.Sscanf(indexStr, "%d", &index)
	if err != nil {
		return -1
	}
	return index - 1
}
