package cactuskev

import (
	"testing"
)

func TestDeck(t *testing.T) {
	d := NewDeck()
	if d == nil {
		t.Fatalf("could not create deck")
	}
}
