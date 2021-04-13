//go:generate stringer -type=Suit,Rank
package godeck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit is the number associated to the card
type Suit uint8

const (
	// Spade is the value of a type of cards
	Spade Suit = iota
	// Diamond is the value of a type of cards
	Diamond
	// Club is the value of a type of cards
	Club
	// Heart is the value of a type of cards
	Heart
	// Joker is the value of a type of cards
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank is the value of the number of a card
type Rank uint8

const (
	_ Rank = iota
	// Ace is the numeric value of the number of a card
	Ace
	// Two is the numeric value of the number of a card
	Two
	// Three is the numeric value of the number of a card
	Three
	// Four is the numeric value of the number of a card
	Four
	// Five is the numeric value of the number of a card
	Five
	// Six is the numeric value of the number of a card
	Six
	// Seven is the numeric value of the number of a card
	Seven
	// Eight is the numeric value of the number of a card
	Eight
	// Nine is the numeric value of the number of a card
	Nine
	// Ten is the numeric value of the number of a card
	Ten
	// Jack is the numeric value of the number of a card
	Jack
	// Queen is the numeric value of the number of a card
	Queen
	// King is the numeric value of the number of a card
	King
)

const (
	minRank = Ace
	maxRank = King
)

// Card is the type of card object
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %s", c.Rank.String(), c.Suit.String())
}

// New creates a new godeck of cards
// It accepts option functions to modify the godeck
func New(options ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{
				Suit: suit,
				Rank: rank,
			})
		}
	}
	for _, option := range options {
		cards = option(cards)
	}
	return cards
}

// DefaultSort is the default sorting function of cards
func DefaultSort(cards []Card) []Card {
	return Sort(Less)(cards)
}

// Sort is the function used to sort cards in a godeck
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Less is the function used to compare cards while sorting
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

// Shuffle shuffles cards randomly in a godeck
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	permutations := shuffleRand.Perm(len(cards))
	for i, j := range permutations {
		ret[i] = cards[j]
	}
	return ret
}

// Jokers adds n number of jokers to the godeck
func Jokers(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Suit(Joker),
				Rank: Rank(i),
			})
		}
		return cards
	}
}

// Deck creates n number of decks
func Deck(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}

// Filter returns a godeck filtering out cards according
// to the criteria of the passed filtering function
func Filter(f func(card Card) bool) func([]Card) []Card {
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
