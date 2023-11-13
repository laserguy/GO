/*
	For the test file, the suffix _test is important
	And to run test command "go test" is used

	But before doing that, ".mod" file in the directory is needed
	Therefore run this command "go mod init <directory_name>"
*/

package main

import (
	"os"
	"testing"
)

/*
	The convention is to write the test function in CamleCase starting with Capital letter
	Whereas the normal functions (which are not test function), is camle case but start with small letter
*/
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16{
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first caard of Ace od Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first caard of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	/*
		Golang doesn't do the cleanup of anything created by the test suites
		So it is the responsibility of the developer to do the cleanup
	*/

	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

