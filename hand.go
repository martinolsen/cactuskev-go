package cactuskev

import (
	"fmt"
	"log"
	"math/rand"
)

type Hand interface {
	Eval() Score
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
		return NewFiveCardHand()
	case 7:
		return NewSevenCardHand()
	default:
		panic(fmt.Errorf("hand of %d cards not supported", n))
	}
}

type FiveCardHand struct {
	A, B, C, D, E, F Card
}

func NewFiveCardHand() *FiveCardHand { return new(FiveCardHand) }

// IsSuited returns true if entire hand is of the same Suit.
func (h *FiveCardHand) IsSuited() bool {
	return h.A&h.B&h.C&h.D&h.E& 0xf000 != 0
}

func (h *FiveCardHand) Eval() Score {
	// Flushes and Straight Flushes
	if h.IsSuited() {
		return Flushes[h.Bit()]
	}

	// Straights and High Cards
	if s := Unique5[h.Bit()]; s != 0 {
		return s
	}

	// and others... [inlined `findit()`]
	var (
		k = int((h.A&0xff) * (h.B&0xff) * (h.C&0xff) * (h.D&0xff) * (h.E&0xff))
	)
	for low, mid, high := 0, 4887>>1, 4887;; mid = (high + low) >> 1 {
		if product := products[mid]; k < product {
			high = mid - 1
		} else if k > product {
			low = mid + 1
		} else {
			return values[mid]
		}
	}
}

func (h *FiveCardHand) SetCard(n int, c Card) {
	switch n {
	case 0:
		h.A = c
	case 1:
		h.B = c
	case 2:
		h.C = c
	case 3:
		h.D = c
	case 4:
		h.E = c
	default:
		log.Panicf("index overflow: %d", n)
	}
}

func (h *FiveCardHand) Card(n int) Card {
	switch n {
	case 0:
		return h.A
	case 1:
		return h.B
	case 2:
		return h.C
	case 3:
		return h.D
	case 4:
		return h.E
	default:
		log.Panicf("index overflow: %d", n)
		return 0
	}
}

func (h FiveCardHand) Len() int { return 5 }

func (h *FiveCardHand) Cards() []Card {
	return []Card{h.A, h.B, h.C, h.D, h.E}
}

func (h *FiveCardHand) Prime() int {
	return int((h.A&0xff) * (h.B&0xff) * (h.C&0xff) * (h.D&0xff) * (h.E&0xff))
}

func (h *FiveCardHand) Bit() int {
	return int(h.A|h.B|h.C|h.D|h.E) >> 16
}

func (h *FiveCardHand) String() string {
	return fmt.Sprintf("[%v %v %v %v %v]", h.A, h.B, h.C, h.D, h.E)
}

type SevenCardHand struct{ A, B, C, D, E, F, G Card }

func NewSevenCardHand() *SevenCardHand { return new(SevenCardHand) }

func (h *SevenCardHand) Eval() Score {
	var (
		sh   = NewFiveCardHand()
		best = Score(9999)
	)

	for i := 0; i < 21; i++ {
		sh.SetCard(0, h.Card(Perm7[i][0]))
		sh.SetCard(1, h.Card(Perm7[i][1]))
		sh.SetCard(2, h.Card(Perm7[i][2]))
		sh.SetCard(3, h.Card(Perm7[i][3]))
		sh.SetCard(4, h.Card(Perm7[i][4]))

		if q := sh.Eval(); best.Less(q) {
			best = q
		}
	}

	return best
}

func (h *SevenCardHand) SetCard(n int, c Card) {
	switch n {
	case 0:
		h.A = c
	case 1:
		h.B = c
	case 2:
		h.C = c
	case 3:
		h.D = c
	case 4:
		h.E = c
	case 5:
		h.F = c
	case 6:
		h.G = c
	default:
		log.Panicf("index overflow: %d", n)
	}
}

func (h *SevenCardHand) Card(n int) Card {
	switch n {
	case 0:
		return h.A
	case 1:
		return h.B
	case 2:
		return h.C
	case 3:
		return h.D
	case 4:
		return h.E
	case 5:
		return h.F
	case 6:
		return h.G
	default:
		log.Panicf("index overflow: %d", n)
		return 0
	}
}

func (h SevenCardHand) Len() int { return 7 }

func (h *SevenCardHand) Cards() []Card {
	return []Card{h.A, h.B, h.C, h.D, h.E, h.F, h.G}
}

func (h *SevenCardHand) Prime() int {
	return h.A.Prime() *
		h.B.Prime() *
		h.C.Prime() *
		h.D.Prime() *
		h.E.Prime() *
		h.F.Prime() *
		h.G.Prime()
}

func (h *SevenCardHand) Bit() int {
	return h.A.Bit() |
		h.B.Bit() |
		h.C.Bit() |
		h.D.Bit() |
		h.E.Bit() |
		h.F.Bit() |
		h.G.Bit()
}

func (h *SevenCardHand) String() string {
	return fmt.Sprintf("[%v %v %v %v %v %v %v]", h.A, h.B, h.C, h.D, h.E, h.F, h.G)
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
	RandomizeHand(h)
	return h
}

func RandomizeHand(h Hand) {
	d := NewDeck()

	for i, j := range rand.Perm(d.Len()) {
		d.Swap(i, j)
	}

	for i := 0; i < h.Len(); i++ {
		c := d[i]
		d = d[1:]
		h.SetCard(i, c)
	}
}
