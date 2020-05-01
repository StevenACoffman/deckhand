# Deckhand - Covid Coding Challenge

Note: this is just for isolated computer science fun, not a real, life-saving project.

Two socially distant players wish to play cards with one another. They must each start with an identically sequenced decks of cards. How can they encode the sequence of cards and decode that back into a sequence of cards?

Assume you have a standard deck of playing cards (52 cards with the four french suits [ordered as SHoCkeD](https://www.deceptionary.com/aboutsuits.html): (spades, hearts, clubs, diamonds) and cards within suits ordered A, 2,.. Q, K).

### Challenge: How can they encode the sequence of cards as a single (large) integer and decode that back into a sequence of cards?

Put another way, a standard deck has 52! = 8x10^67 unique sequences of cards.

If it helps, 8x10^67 in hexidecimal is `2F7A53A390F4323B0F54BBBB472FA8C5DB448DF400000000000000000` which is
57 characters long. Try to get it under 52 characters!

### Resources
* Rosetta Code has a [Playing Cards](https://rosettacode.org/wiki/Playing_cards) sample in most languages. I chose Go.

* Golang has [math/big](https://golang.org/pkg/math/big/) for arbitrary precision numbers

* [Print all possible strings of length k that can be formed from a set of n characters](https://www.geeksforgeeks.org/print-all-combinations-of-given-length/)

* [Converting non-numeric string to integer](https://stackoverflow.com/questions/19724818/converting-non-numeric-string-to-integer)

* For added flavor, the [playing card symbols are available in unicode](https://en.wikipedia.org/wiki/Playing_cards_in_Unicode)
