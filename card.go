//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	// Suit of Cards, Standard 4 Suits. Joker for special cases.
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker //Special case only used sometimes
)

var suits = [...]Suit{Spade, Diamond, Club, Heart} //array of Suits

type Rank uint8

const (
	// Rank of Cards. Adding _ to index 0 so
	// Ace is numerically mapped to 1... King is numerically mapped to 13.
	_ Rank = iota //index 0
	Ace
	Two
	Three
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
)

const (
	//define what is the minimum Rank and max Rank of the cards (Ace: 1, King:13)
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	// creating a string join for Rank of Suit plural
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(opts ...func([]Card) []Card) []Card {
	//for each of the 4 suits, 13 ranks to give 52 cards. no joker initially
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func DefaultSort(cards []Card) []Card {
	//sort least to greatest
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	// 4 of Diamonds ==  1 * 13 * 4, Diamonds (1) always greater than Spade (0)..etc
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	perm := shuffleRand.Perm(len(cards))
	//perm = [0,2,4,5,6,7,1,3...etc] rand index to card index
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

func Jokers(n int) func([]Card) []Card {
	//add n jokers to the deck
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}
		return cards
	}
}

func Filter(f func(card Card) bool) func([]Card) []Card {
	//filter out certain cards
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

func Deck(n int) func([]Card) []Card {
	//Duplicate n number of decks
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
