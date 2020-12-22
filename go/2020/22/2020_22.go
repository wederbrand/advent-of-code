package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type deck struct {
	name  string
	cards []int
}

func (d *deck) addToBottom(card int) {
	d.cards = append(d.cards, card)
}

func (d *deck) getTopCard() int {
	card := d.cards[0]
	d.cards = d.cards[1:]

	return card
}

func (d *deck) getScore() int {
	score := 0

	for i, card := range d.cards {
		value := len(d.cards) - i
		score += value * card
	}

	return score
}

func newDeck(s string) *deck {
	deck := new(deck)
	split := strings.Split(s, "\n")

	for i, card := range split {
		if i == 0 {
			deck.name = card
		} else {
			atoi, _ := strconv.Atoi(card)
			deck.addToBottom(atoi)
		}
	}

	return deck
}

func main() {
	readFile, err := ioutil.ReadFile("22/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	decks := strings.Split(strings.TrimSpace(string(readFile)), "\n\n")

	p1 := newDeck(decks[0])
	p2 := newDeck(decks[1])

	winner := playGame(p1, p2)

	fmt.Println(winner.getScore())
}

func playGame(p1 *deck, p2 *deck) *deck {
	for {
		if len(p1.cards) == 0 {
			return p2
		}
		if len(p2.cards) == 0 {
			return p1
		}
		c1 := p1.getTopCard()
		c2 := p2.getTopCard()

		if c1 > c2 {
			p1.addToBottom(c1)
			p1.addToBottom(c2)
		} else {
			p2.addToBottom(c2)
			p2.addToBottom(c1)
		}
	}
}
