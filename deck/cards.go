package deck

type Suit int

const (
	Diamonds Suit = iota
	Clubs
	Hearts
	Spades
)

type Rank int

const (
	Two Rank = iota
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
	Ace
)

type card struct {
	Suit Suit
	Rank Rank
}

var Cards []card
