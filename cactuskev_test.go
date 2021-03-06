package cactuskev

import (
	"testing"
)

func BenchmarkFiveHand(b *testing.B) {
	bench(b, RandomHand(5))
}

func BenchmarkSevenHand(b *testing.B) {
	bench(b, RandomHand(7))
}

func bench(b *testing.B, hand Hand) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hand.Eval()
	}
}

func BenchmarkFiveCardHandPrime(b *testing.B) {
	h := RandomHand(5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Prime()
	}
}

func TestFive(t *testing.T) {
	RandomHand(5).Eval()
}

func TestSeven(t *testing.T) {
	RandomHand(7).Eval()
}

func TestAllFive(t *testing.T) {
	AllFive(t,
		func() Hand { return NewFiveCardHand() },
		func(h Hand) Hand { return h } /* no need to zero it */)
}

func TestAllSeven(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping, because.")
	}

	AllSeven(t,
		func() Hand { return NewSevenCardHand() },
		func(h Hand) Hand { return h } /* no need to zero it */)
}

func TestCard(t *testing.T) {
	tests := []struct {
		s       Suit
		r       Rank
		str     string
		v, p, b int
	}{
		{Diamond, King, "K♦", 0x8004B25, 37, 1 << 11},
		{Spade, Five, "5♠", 0x81307, 7, 1 << 3},
		{Club, Jack, "J♣", 0x200891D, 29, 1 << 9},
	}

	for _, test := range tests {
		c := NewCard(test.s, test.r)

		if test.s != c.Suit() {
			t.Errorf("expected Suit %v, got %v", test.s, c.Suit())
		}

		if test.r != c.Rank() {
			t.Errorf("expected Rank %v, got %v", test.r, c.Rank())
		}

		if str := c.String(); str != test.str {
			t.Errorf(`expected "%s", got "%s"`, test.str, str)
		}

		if v := int(c); v != test.v {
			t.Errorf("expected %032b, got %032b", test.v, v)
		}

		if p := c.Prime(); p != test.p {
			t.Errorf("expected prime %d, got %d", test.p, p)
		}

		if b := c.Bit(); b != test.b {
			t.Errorf("expected bit %012b, got %012b", test.b, b)
		}
	}
}

func TestHand(t *testing.T) {
	tests := []struct {
		cards [5]Card
		c     Category
		s     Score
	}{ // AKQJ9
		{
			[...]Card{NewCard(Heart, Ace), NewCard(Heart, King), NewCard(Heart, Queen), NewCard(Heart, Jack), NewCard(Heart, Nine)},
			Flush,
			323,
		},
	}

	for _, test := range tests {
		h := NewHand(len(test.cards))
		for i, c := range test.cards {
			h.SetCard(i, c)
		}

		t.Logf("%v", h)

		if s := h.Eval(); s != test.s {
			t.Errorf("expected Score %d, got %d", test.s, s)
		} else if c := s.Category(); c != test.c {
			t.Errorf("expected Category %v, got %v", test.c, c)
		}
	}
}
