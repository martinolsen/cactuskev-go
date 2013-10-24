package cactuskev

import (
	"fmt"
	"log"
	"math/rand"
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
	switch n {
	case 5:
		return new(hand5ints)
	case 7:
		return new(hand7ints)
	default:
		return hand(make([]Card, n))
	}
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

func (h *hand5ints) String() string {
	return fmt.Sprintf("[%v %v %v %v %v]", h.one, h.two, h.three, h.four, h.five)
}

type hand7ints struct{ one, two, three, four, five, six, seven Card }

func (h *hand7ints) SetCard(n int, c Card) {
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
	case 5:
		h.six = c
	case 6:
		h.seven = c
	default:
		log.Panicf("index overflow: %d", n)
	}
}

func (h *hand7ints) Card(n int) Card {
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
	case 5:
		return h.six
	case 6:
		return h.seven
	default:
		log.Panicf("index overflow: %d", n)
		return 0
	}
}

func (h *hand7ints) Len() int { return 7 }

func (h *hand7ints) Cards() []Card {
	return []Card{h.one, h.two, h.three, h.four, h.five, h.six, h.seven}
}

func (h *hand7ints) Prime() int {
	return h.one.Prime() *
		h.two.Prime() *
		h.three.Prime() *
		h.four.Prime() *
		h.five.Prime() *
		h.six.Prime()
}

func (h *hand7ints) Bit() int {
	return h.one.Bit() |
		h.two.Bit() |
		h.three.Bit() |
		h.four.Bit() |
		h.five.Bit() |
		h.six.Bit() |
		h.seven.Bit()
}

func (h *hand7ints) String() string {
	return fmt.Sprintf("[%v %v %v %v %v %v %v]", h.one, h.two, h.three, h.four, h.five, h.six, h.seven)
}

type hand []Card

func (h hand) SetCard(n int, c Card) { h[n] = c }

func (h hand) Card(n int) Card { return h[n] }

func (h hand) Len() int { return len(h) }

func (h hand) Cards() []Card { return h }

// Product of each Card's Prime
func (h hand) Prime() int {
	product := 1

	for _, c := range h {
		product *= c.Prime()
	}

	return product
}

func (h hand) Bit() int {
	var bit int

	for _, c := range h {
		bit |= c.Bit()
	}

	return bit
}

func RandomHand(n int) Hand {
	h := NewHand(n)
	d := NewDeck()

	for i, j := range rand.Perm(d.Len()) {
		d.Swap(i, j)
	}

	for i := 0; i < n; i++ {
		c := d[i]
		d = d[1:]
		h.SetCard(i, c)
	}

	return h
}
