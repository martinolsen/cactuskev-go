package cactuskev

import (
	"testing"
	"sync"
	"time"
	"math/rand"
)

func AllFive(t *testing.T, handfn func() Hand, zerofn func(Hand) Hand) {
	rand.Seed(time.Now().UnixNano())

	var (
		hand  = handfn()
		freq  = make([]int, 9)
		deck  = NewDeck()
		count int
	)

	for a := 0; a < 48; a++ {
		for b := a + 1; b < 49; b++ {
			for c := b + 1; c < 50; c++ {
				for d := c + 1; d < 51; d++ {
					for e := d + 1; e < 52; e++ {
						hand = zerofn(hand)
						hand.SetCard(0, deck[a])
						hand.SetCard(1, deck[b])
						hand.SetCard(2, deck[c])
						hand.SetCard(3, deck[d])
						hand.SetCard(4, deck[e])
						freq[hand.Eval().Category()]++
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

func AllSeven(t *testing.T, handfn func() Hand, zerofn func(Hand) Hand) {
	rand.Seed(time.Now().UnixNano())

	var (
		deck    = NewDeck()
		freqch  = make(chan Category, 1e6)
		freq    = make([]int, 9)
		count   = 0
	)
	
	go func() {
		for c := range freqch {
			freq[c]++
			count++
		}
	}()
	
	var wg sync.WaitGroup
	for i := 0; i < 46; i++ {
		wg.Add(1)
		go func(a int) {
			var hand = handfn()
			
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
									freqch <- hand.Eval().Category()
									hand = zerofn(hand)
								}
							}
						}
					}
				}
			}
			wg.Done()
		}(i)
	}
	
	wg.Wait()
	close(freqch)

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
