package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spar", "Ruder", "Hjerter", "Klør"}
	cardValues := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Knægt", "Dame", "Konge"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" "+value)
		}
	}

	return cards
}

func (d deck) shuffleDeck() {
	// Sleep for a random amount of time to ensure that the random seed is different
	randomNumber := rand.Intn(20) + 10
	time.Sleep(time.Duration(randomNumber))

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Fisher-Yates shuffle algorithm
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}

}

// func dealCards(d deck, handSize int) (deck, deck) {
// 	return d[:handSize], d[handSize:]
// }

func (d deck) saveToFile(filename string) {
	deckStringed := deck.toString(d)
	deckBytes := []byte(deckStringed)
	err := os.WriteFile(filename, deckBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func loadFromFile(filename string) deck {
	deckBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	deckStringed := string(deckBytes)
	return strings.Split(deckStringed, "|")
}

// ############## HELPER FUNCTIONS ##############

// Convert a slice into a string
func (d deck) toString() string {
	return strings.Join([]string(d), "|")
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
