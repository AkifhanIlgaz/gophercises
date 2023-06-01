package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Print(Card{Rank: Ace, Suit: Heart})
	fmt.Print(Card{Rank: Two, Suit: Spade})
	fmt.Print(Card{Rank: Nine, Suit: Diamond})
	fmt.Print(Card{Rank: Jack, Suit: Club})
	fmt.Print(Card{Suit: Joker})
	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)

	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received: ", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))

	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received: ", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))

	for _, card := range cards[len(cards)-3:] {
		if card.Suit != Joker {
			t.Error("Not joker")
		}
	}
}
