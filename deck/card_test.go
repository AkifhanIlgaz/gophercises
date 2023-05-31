package deck

import "fmt"

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
