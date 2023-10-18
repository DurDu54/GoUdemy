package main

import "fmt"

// Create a new typr of 'deck'
// wich is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}