package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Thought struct {
	title string
}

func main() {
	thoughts := []Thought{
		{
			title: "Learn Go",
		},
		{
			title: "Buy new machine",
		},
		{
			title: "create a JavaScript TwitterBot, should I?",
		},
	}
	CreateThought(&thoughts)
	UpdateThought(&thoughts,1,"Terminator")
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
	//We are not handling errors

	var tempThoughts []Thought
	tempThoughts = *thoughts

	for index, _ := range *thoughts {
		if index+1 == thoughtId {
			return tempThoughts[index], errors.New("")
		}
	}
	return tempThoughts[thoughtId-1], errors.New("does not exist")
}

func UpdateThought(thoughts *[]Thought, thoughtId int, newThought string){
	for index, _ :=  range *thoughts{
		if index+1 == thoughtId{
			(*thoughts)[index].title = newThought
		}
	}
}

func ReadAllThoughts(thoughts *[]Thought) {
	for _, thought := range *thoughts {
		fmt.Println(thought.title)
	}
}
