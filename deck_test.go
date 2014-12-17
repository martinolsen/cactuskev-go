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

func TestRemove(t *testing.T) {
	d := NewDeck()
	d.Remove(NewCard(Club, Deuce))
	if d.Len() != 51 {
		t.Errorf("expected Deck length of 51, not %d", d.Len())
	}
	for {
		c, ok := d.Draw()
		if !ok {
			break
		}
		if c.Suit() == Club && c.Rank() == Deuce {
			t.Errorf("unexpected Card in Deck: %s", c)
		}
	}
}

func TestMustDraw(t *testing.T) {
	var i int
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("did not panic")
		} else if i < 52 {
			t.Errorf("panic at card no. %d", i)
		}
	}()
	deck := NewDeck()
	for ; i < 53; i++ {
		deck.MustDraw()
	}
}
