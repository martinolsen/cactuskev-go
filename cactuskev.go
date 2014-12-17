// Credit: http://suffecool.net/poker/evaluator.html
package cactuskev

import (
	"fmt"
	"log"
)

type Score int16

func (s Score) Less(other Score) bool {
	// CactusKevScore goes from 9999 towards zero, where 9999 is the lowest values
	return s > other
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

func (s Score) String() string {
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
