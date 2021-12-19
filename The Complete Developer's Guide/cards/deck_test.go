package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)



func TestNewDeck(t *testing.T) {
	asserter := assert.New(t)

	d := newDeck()

	asserter.Equal(20, len(d))
	asserter.Equal("Ace of Hearts", d[0])
	asserter.Equal("Five of Spades", d[len(d)-1])
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	asserter := assert.New(t)
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	os.Remove("_decktesting")

	asserter.Equal(20, len(loadedDeck))


}

