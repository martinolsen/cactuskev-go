package cactuskev

import (
	"log"
)

type Hand interface {
	SetCard(int, Card)
	Card(int) Card
	Len() int
	Cards() []Card
	Prime() int
	Bit() int
}

func NewHand(n int) Hand {
	return new(hand5ints)
}

type hand5ints struct {
	one, two, three, four, five Card
}

func (h *hand5ints) SetCard(n int, c Card) {
	switch n {
	case 0:
		h.one = c
	case 1:
		h.two = c
	case 2:
		h.three = c
	case 3:
		h.four = c
	case 4:
		h.five = c
	default:
		log.Panicf("index overflow: %d", n)
	}
}

func (h *hand5ints) Card(n int) Card {
	switch n {
	case 0:
		return h.one
	case 1:
		return h.two
	case 2:
		return h.three
	case 3:
		return h.four
	case 4:
		return h.five
	default:
		log.Panicf("index overflow: %d", n)
		return 0
	}
}

func (h *hand5ints) Len() int { return 5 }

func (h *hand5ints) Cards() []Card {
	panic("Hold, who goes there?")
	return []Card{h.one, h.two, h.three, h.four, h.five}
}

func (h *hand5ints) Prime() int {
	return h.one.Prime() * h.two.Prime() * h.three.Prime() * h.four.Prime() * h.five.Prime()
}
func (h *hand5ints) Bit() int {
	return h.one.Bit() | h.two.Bit() | h.three.Bit() | h.four.Bit() | h.five.Bit()
}

type hand5 [5]Card

func (h *hand5) SetCard(n int, c Card) {
	h[n] = c
}

func (h *hand5) Card(n int) Card {
	return h[n]
}

func (h *hand5) Cards() []Card {
	return h[:]
}

// Product of each Card's Prime
func (h *hand5) Prime() int {
	return h.Card(0).Prime() * h.Card(1).Prime() * h.Card(2).Prime() * h.Card(3).Prime() * h.Card(4).Prime()
}

func (h *hand5) Bit() int {
	return h.Card(0).Bit() | h.Card(1).Bit() | h.Card(2).Bit() | h.Card(3).Bit() | h.Card(4).Bit()
}
