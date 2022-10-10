package main

import (
	"bufio"
	"fmt"
	"os"
)

type Thought struct {
	title string
}

func main() {
	thoughts := []Thought{
		{
			title: "Go",
		},
		{
			title: "Java",
		},
		{
			title: "JavaScript",
		},
	}
	CreateThought(&thoughts)
	DeleteThought(&thoughts, 4)
	ReadAllThoughts(&thoughts)
}

func CreateThought(thoughts *[]Thought) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Add a thought: ")
	userInput, _ := inputReader.ReadString('\n')
	thought := Thought{title: userInput}
	*thoughts = append(*thoughts, thought)
	fmt.Println("Thought added successfully")
}

func DeleteThought(thoughts *[]Thought, thoughtId int) {
	fmt.Println(*thoughts)
	var tempThoughts []Thought
	tempThoughts = *thoughts
	for index, _ := range *thoughts {
		if index+1 == thoughtId {
			tempThoughts = append(tempThoughts[:index], tempThoughts[index+1:]...)
			*thoughts = tempThoughts
			fmt.Println(*thoughts)
		}
	}
}

func ReadAllThoughts(thoughts *[]Thought) {
	for _, thought := range *thoughts {
		fmt.Println(thought.title)
	}
}
