package dealer

import( 
	"fmt"
	"strconv"
	"github.com/jaeyoony/deck_of_cards/deck"
)

type Player struct {
	Card1 deck.Card
	Card2 deck.Card
	HitCards []deck.Card
	Score int
}

type Round struct {
	PlayDeck []deck.Card
	Index int 
	Dealer Player
	Human Player
}

var suits = [...]string{"Spades", "Diamonds", "Clubs", "Hearts"}


// ********* ROUND struct methods ********* //

// start a round with a deck containing n complete decks 
func StartRound(decks int) Round {
	temp_deck := deck.New(deck.MultiDeck(decks), deck.Shuffle)
	player := Player{}
	dealer := Player{}

	player.Card1 = temp_deck[0]
	dealer.Card1 = temp_deck[1]
	player.Card2 = temp_deck[2]
	dealer.Card2 = temp_deck[3]

	player.Score = player.getTotalScore()
	
	// reveal Dealer's first card 

	fmt.Println(" ** Dealers first card is : ", dealer.Card1, " ** ")
	return Round{temp_deck, 4, dealer, player}
}

// returns the next card from a PlayDeck
func (r Round) Hit() deck.Card {
	temp := r.PlayDeck[r.Index]
	r.Index ++ 
	return temp
}


// handle the players turn
/*
func (r Round) PlayerTurn() {

} 
*/

// gets the value of a given card, in blackjack rules - 
//	aces = 1 or 11, 
func getCardValue(c deck.Card) int {
	if(c.Val >= 10) {
		return 10
	} else if(c.Val == 1) {
		return HandleAce(c)
	} else {
		return c.Val
	}
}

func (p Player) getTotalScore() int {
	temp_sum := getCardValue(p.Card1) + getCardValue(p.Card2)
	for _, i := range(p.HitCards) {
		temp_sum += getCardValue(i)
	}
	return temp_sum
}

// promts the user to pick either a 1 or 11 value for their ace card
func HandleAce(c deck.Card) int {
	var response string
	fmt.Println("You drew and Ace! The ace of", suits[c.Suit])
	fmt.Println(" . . Do you want it to have a value of 1 or 11?")
	fmt.Scanln(&response)

	for(response != "1" && response != "11") {
		fmt.Println(" . . Invalid value! Please enter either \"1\" or \"11\"")
		fmt.Scanln(&response)		
	}

	i, err := strconv.Atoi(response)
	if err != nil {
		panic(err)
	}
	return i
}

