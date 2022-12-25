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

func (d deck) copy(count int) *deck {
	deck := new(deck)
	deck.name = d.name
	for _, card := range d.cards[:count] {
		deck.addToBottom(card)
	}
	return deck
}

func (d deck) getKey() string {
	key := ""

	for _, card := range d.cards {
		key += strconv.Itoa(card) + " "
	}

	return key
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
	p1SeenDecks := make(map[string]bool)
	p2SeenDecks := make(map[string]bool)
	for {
		if p1SeenDecks[p1.getKey()] {
			return p1
		}
		if p2SeenDecks[p2.getKey()] {
			return p1
		}

		p1SeenDecks[p1.getKey()] = true
		p2SeenDecks[p2.getKey()] = true

		if len(p1.cards) == 0 {
			return p2
		}
		if len(p2.cards) == 0 {
			return p1
		}
		c1 := p1.getTopCard()
		c2 := p2.getTopCard()

		var winner *deck
		if c1 <= len(p1.cards) && c2 <= len(p2.cards) {
			// subgame
			p1sub := p1.copy(c1)
			p2sub := p2.copy(c2)
			winner = playGame(p1sub, p2sub)
		} else {
			if c1 > c2 {
				winner = p1
			} else {
				winner = p2
			}
		}

		if winner.name == p1.name {
			p1.addToBottom(c1)
			p1.addToBottom(c2)
		} else {
			p2.addToBottom(c2)
			p2.addToBottom(c1)
		}
	}
}
