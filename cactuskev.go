// Credit: http://suffecool.net/poker/evaluator.html
package cactuskev

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Evaluator func(Hand) Score

type Score int16

func Eval(h Hand) Score {
	switch h.Len() {
	case 5:
		return Score(eval_5hand(h))
	case 7:
		return Score(eval_7hand(h))
	default:
		log.Panicf("no evaluator for %d Card Hands", len(h.Cards()))
	}

	return 0
}

func (s Score) Category() Category {
	switch {
	case s > 6185:
		return HighCard
	case s > 3325:
		return OnePair
	case s > 2467:
		return TwoPair
	case s > 1609:
		return ThreeOfAKind
	case s > 1599:
		return Straight
	case s > 322:
		return Flush
	case s > 166:
		return FullHouse
	case s > 10:
		return FourOfAKind
	default:
		return StraightFlush
	}
}

type Category int

const (
	StraightFlush Category = iota
	FourOfAKind
	FullHouse
	Flush
	Straight
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func (c Category) String() string {
	switch c {
	case StraightFlush:
		return "Straight Flush"
	case FourOfAKind:
		return "Four of a Kind"
	case FullHouse:
		return "Full House"
	case Flush:
		return "Flush"
	case Straight:
		return "Straight"
	case ThreeOfAKind:
		return "Three of a Kind"
	case TwoPair:
		return "Two Pair"
	case OnePair:
		return "One Pair"
	case HighCard:
		return "High Card"
	default:
		log.Panicf("unknown Category %d", c)
	}

	return ""
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

type Card int

func NewCard(s Suit, r Rank) Card {
	return Card(primes[r] | (int(r) << 8) | int(s) | (1 << uint(16+r)))
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

func eval_7hand(h Hand) int16 {
	var (
		sh   = NewHand(5)
		best = int16(9999)
	)

	for i := 0; i < 21; i++ {
		for j := 0; j < 5; j++ {
			sh.SetCard(j, h.Card(perm7[i][j]))
		}

		if q := eval_5hand(sh); q < best {
			best = q
		}
	}

	return best
}

func eval_5hand(h Hand) int16 {
	var q = h.Bit()

	// Flushes and Straight Flushes
	if h.Card(0).Suit()&h.Card(1).Suit()&h.Card(2).Suit()&h.Card(3).Suit()&h.Card(4).Suit() != 0 {
		return Flushes[q]
	}

	// Straights and High Cards
	if s := Unique5[q]; s != 0 {
		return s
	}

	// and others...
	return values[findit(h.Prime())]
}

func findit(k int) int {
	var low, mid, high = 0, 0, 4887

	for low <= high {
		mid = (high + low) >> 1 // Divide by two

		if k < products[mid] {
			high = mid - 1
		} else if k > products[mid] {
			low = mid + 1
		} else {
			return mid
		}

	}

	panic(fmt.Errorf("no match found for key %d", k))
}

func (s Suit) String() string {
	switch s {
	case Diamond:
		return "♦"
	case Club:
		return "♣"
	case Heart:
		return "♥"
	case Spade:
		return "♠"
	default:
		log.Panicf("unknown suit: %d", s)
	}

	return ""
}

func (r Rank) String() string {
	switch r {
	case Deuce:
		return "2"
	case Trey:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "T"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		log.Panicf("unknown rank: %d", r)
	}

	return ""
}

func (c Card) String() string {
	return fmt.Sprintf("%v%v", c.Rank(), c.Suit())
}

func AllFive(eval Evaluator) {
	rand.Seed(time.Now().UnixNano())

	deck := NewDeck()

	var (
		freq = make([]int, 9)
		hand = NewHand(5)
	)

	for a := 0; a < 48; a++ {
		hand.SetCard(0, deck[a])

		for b := a + 1; b < 49; b++ {
			hand.SetCard(1, deck[b])

			for c := b + 1; c < 50; c++ {
				hand.SetCard(2, deck[c])

				for d := c + 1; d < 51; d++ {
					hand.SetCard(3, deck[d])

					for e := d + 1; e < 52; e++ {
						hand.SetCard(4, deck[e])

						freq[eval(hand).Category()]++
					}
				}
			}
		}
	}

	for i, n := range freq {
		log.Printf("% 15v: % 8d", Category(i).String(), n)
	}
}
