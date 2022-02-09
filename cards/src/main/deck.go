package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of Deck

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuites := deck{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardsValues := deck{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardsSuites {
		for _, value := range cardsValues {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}
	return cards
}

func (myDeck deck) print() {
	for i := range myDeck {
		fmt.Println(i, myDeck[i])
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (myDeck deck) saveToFile(filename string) error {
	bytes := []byte(myDeck.toString())
	return ioutil.WriteFile(filename, bytes, 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	anotherRand := rand.New(source)

	deckSize := len(d) - 1
	for i := range d {
		newPosition := anotherRand.Intn(deckSize)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
