package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Heart})
	fmt.Println(Card{Rank: Nine, Suit: Heart})
	fmt.Println(Card{Suit: Joker})
	fmt.Println(Card{Rank: King, Suit: Spade})
	fmt.Println(Card{Rank: Queen, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Heart})

	//Ouput:
	//Ace of Hearts
	//Two of Hearts
	//Nine of Hearts
	//Joker
	//King of Spades
	//Queen of Diamonds
	//Jack of Hearts
}

func TestNew(t *testing.T) {
	cards := New()
	//13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Not enough cards in the deck")
	}
}

func index(vs []Card, t Card) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func include(vs []Card, t Card) bool {
	return index(vs, t) >= 0
}

func TestJokerInNewDeck(t *testing.T) {
	cards := New()
	exp := Card{Suit: Joker}
	if include(cards, exp) {
		t.Error("Joker in new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spaces as first card. Received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Diamond}
	if cards[13] != exp {
		t.Error("Expected Ace of Diamonds as first card. Received:", cards[13])
	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// First call to shuffleRand.Perm(52) should be:
	// rand.New(rand.NewSource(0)) = [40 35 50 0 ... 38 24]
	shuffleRand = rand.New(rand.NewSource(0))
	ogcards := New()
	first := ogcards[40]
	second := ogcards[35]
	last := ogcards[24]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected the first cards to be %s, received %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the first cards to be %s, received %s", second, cards[1])
	}
	if cards[51] != last {
		t.Errorf("Expected the first cards to be %s, received %s", last, cards[51])
	}

}

func TestJoker(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers, Recieved:", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Ace || card.Rank == Five
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Ace || c.Rank == Five {
			t.Error("Expected Aces and Fives to be filtered out")
		}
	}
}

func testDeck(t *testing.T) {
	cards := New(Deck(3))
	//13 ranks * 4 suits * 3 decks
	if len(cards) != 13*4*3 {
		t.Errorf("Expected %d cards, recieved %d cards.", 13*4*3, len(cards))
	}

}
