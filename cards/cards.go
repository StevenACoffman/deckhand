package cards

import (
	"math/big"
	"math/rand"
	"strings"
)

// A Suit represents one of the four standard suites.
type Suit uint8

// The four standard suites in SHoCkeD order
const (
	Spade   Suit = 0
	Heart   Suit = 1
	Club    Suit = 2
	Diamond Suit = 3
)

// ♠	♥	♦	♣   ♤	♡	♢	♧
// SHoCkeD or "♠♡♣♢"
func (s Suit) String() string {
	const suites = "♠♡♣♢" // "SHCD" == "♠♡♣♢"
	return string([]rune(suites)[s : s+1])
}

/*
AS 2S 3S 4S 5S 6S 7S 8S 9S TS JS QS KS
AH 2H 3H 4H 5H 6H 7H 8H 9H TH JH QH KH
AC 2C 3C 4C 5C 6C 7C 8C 9C TC JC QC KC
AD 2D 3D 4D 5D 6D 7D 8D 9D TD JD QD KD
*/
// Rank is the rank or pip value of a card from Ace==1 to King==13.
type Rank uint8

// The ranks from Ace to King.
const (
	Ace   Rank = 1
	Two   Rank = 2
	Three Rank = 3
	Four  Rank = 4
	Five  Rank = 5
	Six   Rank = 6
	Seven Rank = 7
	Eight Rank = 8
	Nine  Rank = 9
	Ten   Rank = 10
	Jack  Rank = 11
	Queen Rank = 12
	King  Rank = 13
)

func (r Rank) String() string {
	const ranks = "A23456789TJQK"
	return string([]rune(ranks)[r-1 : r])
}

// A Card represents a specific playing card.
// It's an encoded representation of Rank and Suit
// with a valid range of [0,51].
type Card uint8

// NewCard returns the Card representation for the specified rank and suit.
func NewCard(r Rank, s Suit) Card {
	return Card(13*uint8(s) + uint8(r-1))
}

// RankSuit returns the rank and suit of the card.
func (c Card) RankSuit() (Rank, Suit) {
	return Rank(c%13 + 1), Suit(c / 13)
}

// Rank returns the rank of the card.
func (c Card) Rank() Rank {
	return Rank(c%13 + 1)
}

// Suit returns the suit of the card.
func (c Card) Suit() Suit {
	return Suit(c / 13)
}

func (c Card) String() string {
	return c.Rank().String() + c.Suit().String()
}

func (c Card) toSingleLetter() string {
	const base = 52
	//convertString := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	convertString := "🂡🂢🂣🂤🂥🂦🂧🂨🂩🂪🂫🂭🂮🂱🂲🂳🂴🂵🂶🂷🂸🂹🂺🂻🂽🂾🃁🃂🃃🃄🃅🃆🃇🃈🃉🃊🃋🃍🃎🃑🃒🃓🃔🃕🃖🃗🃘🃙🃚🃛🃝🃞"
	return string([]rune(convertString)[c%base])
}

func fromSingleLetter(s string) Card {
	const symbolIndex = "🂡🂢🂣🂤🂥🂦🂧🂨🂩🂪🂫🂭🂮🂱🂲🂳🂴🂵🂶🂷🂸🂹🂺🂻🂽🂾🃁🃂🃃🃄🃅🃆🃇🃈🃉🃊🃋🃍🃎🃑🃒🃓🃔🃕🃖🃗🃘🃙🃚🃛🃝🃞"
	const letterIndex = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	value := strings.Index(letterIndex, s)
	if value == -1 {
		for i, r := range []rune(symbolIndex) {
			if r == []rune(s)[0] {
				value=i
			}
		}
	}
	return Card(uint8(value))
}

// A Deck represents a set of zero or more cards in a specific order.
type Deck []Card

// NewDeck returns a regular 52 deck of cards in A-K order.
func NewDeck() Deck {
	d := make(Deck, 52)
	for i := range d {
		d[i] = Card(i)
	}
	return d
}

func NewDeckFromString(s string) Deck {
	//trim all whitespace, and internal spaces
	s = strings.Replace(strings.TrimSpace(s), " ", "", -1)
	r := []rune(s)
	d := make(Deck, len(r))
	for i := range d {
		d[i] = fromSingleLetter(string(r[i]))
	}
	return d
}

// String returns a string representation of the cards in the deck with
// a newline ('\n') separating the cards into groups of thirteen.
func (d Deck) String() string {
	s := ""
	for i, c := range d {
		switch {
		case i == 0: // do nothing
		case i%13 == 0:
			s += "\n"
		default:
			s += " "
		}
		s += c.String()
	}
	return s
}

func (d Deck) SymbolString() string {
	s := ""
	for _, c := range d {
		s += c.toSingleLetter()
	}
	return s
}

// Shuffle randomises the order of the cards in the deck.
func (d Deck) Shuffle() {
	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

// Contains returns true if the specified card is withing the deck.
func (d Deck) Contains(tc Card) bool {
	for _, c := range d {
		if c == tc {
			return true
		}
	}
	return false
}

// AddDeck adds the specified deck(s) to this one at the end/bottom.
func (d *Deck) AddDeck(decks ...Deck) {
	for _, o := range decks {
		*d = append(*d, o...)
	}
}

// AddCard adds the specified card to this deck at the end/bottom.
func (d *Deck) AddCard(c Card) {
	*d = append(*d, c)
}

// Draw removes the selected number of cards from the top of the deck,
// returning them as a new deck.
func (d *Deck) Draw(n int) Deck {
	old := *d
	*d = old[n:]
	return old[:n:n]
}

// DrawCard draws a single card off the top of the deck,
// removing it from the deck.
// It returns false if there are no cards in the deck.
func (d *Deck) DrawCard() (Card, bool) {
	if len(*d) == 0 {
		return 0, false
	}
	old := *d
	*d = old[1:]
	return old[0], true
}

// Deal deals out cards from the deck one at a time to multiple players.
// The initial hands (decks) of each player are provided as arguments and the
// modified hands are returned. The initial hands can be empty or nil.
// E.g. Deal(7, nil, nil, nil) deals out seven cards to three players
// each starting with no cards.
// If there are insufficient cards in the deck the hands are partially dealt and
// the boolean return is set to false (true otherwise).
func (d *Deck) Deal(cards int, hands ...Deck) ([]Deck, bool) {
	for i := 0; i < cards; i++ {
		for j := range hands {
			if len(*d) == 0 {
				return hands, false
			}
			hands[j] = append(hands[j], (*d)[0])
			*d = (*d)[1:]
		}
	}
	return hands, true
}

// space inefficient encoding.
// each card = 1 byte = 128 max
// each card should be only be max 52
func (d *Deck) Encode() string {
	var bs []byte

	for _, c := range *d {
		bs = append(bs, byte(uint8(c)))
	}
	bigNum := big.NewInt(0)
	bigNum = bigNum.SetBytes(bs)
	return bigNum.Text(big.MaxBase)
}

func Decode(s string) Deck {
	bigNum := big.NewInt(0)
	bigNum.SetString(s, big.MaxBase)
	bs := bigNum.Bytes()

	//leading 0 got eaten
	if len(bs) == 51 {
		bs = append([]byte{0}, bs...)
	}
	d := make(Deck, 52)
	for i, b := range bs {
		d[i] = Card(b)
	}
	return d
}

func (d *Deck) ExpEncode() string {
	value := big.NewInt(0)
	base := big.NewInt(52)
	numBytes := len(*d)
	for i := numBytes-1; i > 0; i-- {
		card := big.NewInt(int64((*d)[i]))
		exp := big.NewInt(0).Exp(base, big.NewInt(int64(i)), nil)
		cardExp := exp.Mul(exp, card)
		value = big.NewInt(0).Add(value, cardExp)
	}

	return value.Text(big.MaxBase)
}

func ExpDecode(s string) Deck {
	value := big.NewInt(0)
	r :=  big.NewInt(0)
	value.SetString(s, big.MaxBase)
	base := big.NewInt(52)
	d := make(Deck, 52)

	for i := range d {
		value, r = value.QuoRem(value, base, r)
		d[i] = Card(uint8(r.Int64()))
	}
	return d
}