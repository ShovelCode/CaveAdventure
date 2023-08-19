package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Room represents a location in the game
type Room struct {
	description string
	exits       map[string]*Room
}

func main() {
	// Define rooms
	entrance := &Room{
		description: "You are at the entrance of a mysterious cave. There is a passage to the north.",
		exits:       make(map[string]*Room),
	}

	cave := &Room{
		description: "You are in a dark cave. There are passages to the south and to the east.",
		exits:       make(map[string]*Room),
	}

	treasure := &Room{
		description: "You have found the treasure room! There is a passage to the west.",
		exits:       make(map[string]*Room),
	}

	// Define exits
	entrance.exits["north"] = cave
	cave.exits["south"] = entrance
	cave.exits["east"] = treasure
	treasure.exits["west"] = cave

	// Game loop
	currentRoom := entrance
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(currentRoom.description)
		fmt.Print("What do you want to do? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		directions := []string{"north", "south", "east", "west"}
		validMove := false

		for _, direction := range directions {
			if input == direction {
				if nextRoom, exists := currentRoom.exits[direction]; exists {
					currentRoom = nextRoom
					validMove = true
					break
				}
			}
		}

		if !validMove {
			fmt.Println("You can't go that way.")
		}
	}
}
