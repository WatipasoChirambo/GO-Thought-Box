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
			title: "Learn Go",
		},
		{
			title: "Buy new machine",
		},
		{
			title: "create a JavaScript TwitterBot",
		},
	}
	fmt.Println(GetThought(thoughts,2))
}

func CreateThought(thoughts *[]Thought) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Add a thought: ")
	userInput, _ := inputReader.ReadString('\n')
	thought := Thought{title: userInput}
	*thoughts = append(*thoughts, thought)
	fmt.Println("Thought added successfully")
}

func GetThought(thoughts []Thought, thoughtId int) (string) {
	for index, _ := range thoughts {
		if index+1 == thoughtId {
			return thoughts[index].title
		}
	}
	return thoughts[thoughtId-1].title
}

func DelThought(thoughts []Thought, thoughtId int)  {
	thoughts = append(thoughts[:thoughtId], thoughts[thoughtId+1:]...)
}


func UpdateThought(thoughts []Thought, thoughtId int, newThought string) {
	for index, _ := range thoughts {
		if index+1 == thoughtId {
			(thoughts)[index].title = newThought
			fmt.Println("Thought updated successfully")
			return
		}
	}
}

func ReadAllThoughts(thoughts []Thought) {
	for _, thought := range thoughts {
		fmt.Println(thought.title)
	}
}
