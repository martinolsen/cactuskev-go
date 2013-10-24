package cactuskev

import (
	"log"
)

type Deck []Card

func NewDeck() Deck {
	deck := make([]Card, 52)

	for i, suit := range []Suit{Club, Diamond, Heart, Spade} {
		for j, rank := range []Rank{Deuce, Trey, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace} {
			deck[(i*13)+j] = NewCard(suit, rank)
		}
	}

	return deck
}

func (d Deck) Len() int      { return len(d) }
func (d Deck) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

func (d Deck) Less(i, j int) bool {
	return d[i].Rank() < d[j].Rank() &&
		d[i].Suit() < d[j].Suit()
}

type Suit uint16

const (
	Club    Suit = 0x8000
	Diamond      = 0x4000
	Heart        = 0x2000
	Spade        = 0x1000
)

type Rank uint16

const (
	Deuce Rank = iota
	Trey
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// Bitmask: xxxbbbbb bbbbbbbb cdhsrrrr xxpppppp
//  p: prime number, r: rank, cdhs: suit, b: rank

type Card int32

func NewCard(s Suit, r Rank) Card {
	if int(r) >= len(Primes) {
		log.Panicf("unknown rank: %d", r)
	}

	return Card(Primes[r] | (int(r) << 8) | int(s) | (1 << uint(16+r)))
}

func (c Card) Suit() Suit {
	return Suit(c & 0xf000)
}

func (c Card) Rank() Rank {
	return Rank((c >> 8) & 0x0f)
}

func (c Card) Prime() int {
	return int(c) & 0xff
}

func (c Card) Bit() int {
	return int(c) >> 16
}
