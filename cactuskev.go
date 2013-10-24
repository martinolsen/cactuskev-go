// Credit: http://suffecool.net/poker/evaluator.html
package cactuskev

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

type Evaluator func(Hand) Score

func Eval(h Hand) Score {
	switch h.Len() {
	case 5:
		return CactusKevScore(eval_5hand(h))
	case 7:
		return CactusKevScore(eval_7hand(h))
	default:
		log.Panicf("no evaluator for %d Card Hands", len(h.Cards()))
		return CactusKevScore(9999)
	}
}

type Score interface {
	Category() Category
}

type CactusKevScore int16

func (s CactusKevScore) Category() Category {
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

func (s CactusKevScore) String() string {
	return fmt.Sprintf("%v(%d)", s.Category(), s)
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

func eval_7hand(h Hand) int16 {
	var (
		sh   = NewHand(5)
		best = int16(9999)
	)

	for i := 0; i < 21; i++ {
		for j := 0; j < 5; j++ {
			sh.SetCard(j, h.Card(Perm7[i][j]))
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

func AllFive(t *testing.T, eval Evaluator) {
	rand.Seed(time.Now().UnixNano())

	var (
		freq  = make([]int, 9)
		hand  = NewHand(5)
		deck  = NewDeck()
		count int
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
						count++
					}
				}
			}
		}
	}

	if t == nil {
		return
	}

	verifyAllHands(
		t,
		freq,
		count,
		2598960,
		map[Category]int{
			StraightFlush: 40,
			FourOfAKind:   624,
			FullHouse:     3744,
			Flush:         5108,
			Straight:      10200,
			ThreeOfAKind:  54912,
			TwoPair:       123552,
			OnePair:       1098240,
			HighCard:      1302540,
		},
	)
}

func AllSeven(t *testing.T, eval Evaluator) {
	rand.Seed(time.Now().UnixNano())

	var (
		deck  = NewDeck()
		freq  = make([]int, 9)
		hand  = NewHand(7)
		count = 0
	)

	for a := 0; a < 46; a++ {
		hand.SetCard(0, deck[a])
		for b := a + 1; b < 47; b++ {
			hand.SetCard(1, deck[b])
			for c := b + 1; c < 48; c++ {
				hand.SetCard(2, deck[c])
				for d := c + 1; d < 49; d++ {
					hand.SetCard(3, deck[d])
					for e := d + 1; e < 50; e++ {
						hand.SetCard(4, deck[e])
						for f := e + 1; f < 51; f++ {
							hand.SetCard(5, deck[f])
							for g := f + 1; g < 52; g++ {
								hand.SetCard(6, deck[g])
								freq[eval(hand).Category()]++
								count++
							}
						}
					}
				}
			}
		}
	}

	if t == nil {
		return
	}

	verifyAllHands(
		t,
		freq,
		count,
		133784560,
		map[Category]int{
			StraightFlush: 41584,
			FourOfAKind:   224848,
			FullHouse:     3473184,
			Flush:         4047644,
			Straight:      6180020,
			ThreeOfAKind:  6461620,
			TwoPair:       31433400,
			OnePair:       58627800,
			HighCard:      23294460,
		},
	)
}

func verifyAllHands(t *testing.T, fs []int, count, total int, ts map[Category]int) {
	if count != total {
		t.Errorf("unexpected hand count: %d", count)
	}

	var freqSum int

	for c, a := range ts {
		b := fs[int(c)]

		switch {
		case a != b:
			t.Errorf("unexpected number of %s hands: %d (expected %d)", c, b, a)
		default:
			freqSum += a
		}
	}

	if freqSum != count {
		t.Errorf("unexpected hand frequency count: %d", freqSum)
	}
}
