package godeck

import (
	"fmt"
	"math/rand"
	"testing"
)

// ExampleDeck is an Example function.
// Adding an Output comment turns it into a test function.
// Print the result of the tested logic and comment the expected output as assertion.
func ExampleDeck() {
	fmt.Println(Card{Heart, King})
	fmt.Println(Card{Spade, Three})
	fmt.Println(Card{Club, Nine})
	fmt.Println(Card{Diamond, Ace})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// King of Heart
	// Three of Spade
	// Nine of Club
	// Ace of Diamond
	// Joker
}

func TestNewDeck(t *testing.T) {
	cards := New(DefaultSort)
	actual := len(cards)
	wanted := 13 * 4
	if actual != wanted {
		t.Errorf("number of cards was %d, wanted %d", actual, wanted)
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Errorf("wrong number of jokers found, actual: %d, wanted: %d", count, 3)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Errorf("unexpected card found in godeck: %v", c)
		}
	}
}

func TestDeck(t *testing.T) {
	n := 3
	wanted := 13 * 4 * n
	cards := New(Deck(n))
	if len(cards) != wanted {
		t.Errorf("wrong number of cards, got %d but wanted %d", len(cards), wanted)
	}
}

func TestShuffle(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0))
	orig := New()
	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("unexpected card, got %s but wanted %s", cards[0], first)
	}
	if cards[1] != second {
		t.Errorf("unexpected card, got %s but wanted %s", cards[1], second)
	}
}
