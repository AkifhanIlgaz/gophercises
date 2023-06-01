//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	// We use blank identifier to match each Rank with corresponding uint8 value
	_ Rank = iota
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
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss\n", c.Rank.String(), c.Suit.String())
}

func New() []Card {
	var cards []Card

	for _, suit := range suits {
		for i := minRank; i <= maxRank; i++ {
			card := Card{
				Suit: suit,
				Rank: Rank(i),
			}
			cards = append(cards, card)
		}
	}

	return cards
}
