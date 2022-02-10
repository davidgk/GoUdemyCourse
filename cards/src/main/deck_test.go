package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	fmt.Println("Starting Test of Deck")
	msg := testDeckCreation(t)
	msgs := []string{"when want to create a deck" + msg}

	msg = testFirstCard(t)
	msgs = append(msgs, "when want to create a deck and first card"+msg)

	msg = testLastCard(t)
	msgs = append(msgs, "when want to create a deck and last card"+msg)

	msg = testSaveFile(t)
	msgs = append(msgs, "when want to save a deck into a file"+msg)

	msg = testReadFile(t)
	msgs = append(msgs, "when want to read a deck from a file"+msg)

	for _, m := range msgs {
		fmt.Println(m)
	}
}

func testReadFile(t *testing.T) string {
	deck := NewDeck()
	msg := " should be read"
	filename := "testFile.txt"
	err := deck.saveToFile(filename)
	msg = checkErrorOnSave(t, err, msg)
	result := newDeckFromFile(filename)
	if len(result) != len(deck) {
		t.Errorf("error expected on loading from file and get %v", err)
		msg = " and FAIL"
	}
	os.Remove(filename)
	return msg
}

func checkErrorOnSave(t *testing.T, err error, msg string) string {
	if err != nil {
		t.Errorf("error expected on save file and get %v", err)
		msg = " and FAIL"
	}
	return msg
}

func testSaveFile(t *testing.T) string {
	deck := NewDeck()
	msg := " should be saved"
	filename := "testFile.txt"
	err := deck.saveToFile(filename)
	msg = checkErrorOnSave(t, err, msg)
	os.Remove(filename)
	return msg
}

func testDeckCreation(t *testing.T) string {
	deck := NewDeck()
	msg := " should be created with 52"
	if len(deck) != 52 {
		t.Errorf("error expectedFirst 52 and got %d", len(deck))
		msg = " and FAIL"
	}
	return msg
}

func testLastCard(t *testing.T) string {
	deck := NewDeck()
	expectedLast := "King of Clubs"
	msg := " should be " + expectedLast
	lastCard := deck[len(deck)-1]
	if lastCard != expectedLast {
		t.Errorf("error expected Last %v and got %v", expectedLast, lastCard)
		msg = " and FAIL"
	}
	return msg
}

func testFirstCard(t *testing.T) string {
	deck := NewDeck()
	expectedFirst := "Ace of Spades"
	first := deck[0]
	msg := " should be " + expectedFirst
	if first != expectedFirst {
		t.Errorf("error expectedFirst %v and got %v", expectedFirst, first)
		msg = " and FAIL"
	}
	return msg
}
