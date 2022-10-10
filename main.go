package main

import (
	"bufio"
	"fmt"
	"os"
	"errors"
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
	fmt.Println(GetThought(&thoughts, 3))
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
	var tempThoughts []Thought
	tempThoughts = *thoughts
	for index, _ := range *thoughts {
		if index+1 == thoughtId {
			tempThoughts = append(tempThoughts[:index], tempThoughts[index+1:]...)
			*thoughts = tempThoughts
		}
	}
}

func GetThought(thoughts *[]Thought, thoughtId int) (Thought, error) {
	var tempThoughts []Thought
	tempThoughts = *thoughts

	for index, _ := range *thoughts {
		if index+1 == thoughtId {
			return tempThoughts[index], errors.New("")
		}
	}
	return tempThoughts[thoughtId-1], errors.New("does not exist")
}

func ReadAllThoughts(thoughts *[]Thought) {
	for _, thought := range *thoughts {
		fmt.Println(thought.title)
	}
}
