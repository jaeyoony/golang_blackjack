package blackjack

import( 
	"fmt"
	"strconv"
	"github.com/jaeyoony/deck_of_cards/deck"
)
type BlackjackGame struct {
	HumanWins int
	DealerWins int
	Rounds int 
}

type Round struct {
	PlayDeck []deck.Card
	Index int  // the index in the deck we're currently drawing from, aka WHICH CARD is next to draw
	Dealer Player
	Human Player
}

type Player struct {
	Cards []deck.Card
	Score int
}

var suits = [...]string{"Spades", "Diamonds", "Clubs", "Hearts"}

// ********* BLACKJACK struct methods ********* // 
func StartGame(rounds int, addl_decks int) {
	newGame := BlackjackGame{0, 0, rounds}

	for i:=0; i < rounds; i++ {
		new_round := StartRound(addl_decks)
		new_round.PlayerTurn()
		new_round.DealerTurn()
		result := new_round.EndRound()

		if(result == 1) {
			newGame.HumanWins++
		} else {
			newGame.DealerWins++
		}
		newGame.Rounds--
	}

	fmt.Println("FINAL SCORE : ", newGame.HumanWins, " VS ", newGame.DealerWins)
	
	if(newGame.HumanWins > newGame.DealerWins) {
		fmt.Println(" YOU BEAT THE DEALER!")

	} else {
		fmt.Println(" YOU GOT YO ASS BEAT BY THE DEALER!")
	}
}


// ********* ROUND struct methods ********* //

// start a round with a deck containing n complete decks 
func StartRound(decks int) Round {
	temp_deck := deck.New(deck.MultiDeck(decks), deck.Shuffle)
	player := Player{}
	dealer := Player{}

	player.Cards = append(player.Cards, temp_deck[0])
	dealer.Cards = append(dealer.Cards, temp_deck[1])
	player.Cards = append(player.Cards, temp_deck[2])
	dealer.Cards = append(dealer.Cards, temp_deck[3])
	player.Score = player.getTotalScore(true)
	dealer.Score = dealer.getTotalScore(false)
	
	// reveal Dealer's first card 

	fmt.Println(" ** Dealers first card is : ", dealer.Cards[0], " ** ")
	return Round{temp_deck, 4, dealer, player}
}

// returns the next card from a PlayDeck
func (r *Round) Hit() deck.Card {
	temp := r.PlayDeck[r.Index]
	r.Index++
	return temp
}


// handle human player turn 
func (r *Round) PlayerTurn() {
	var response string 
	fmt.Println("              --- PLAYER TURN ---")

	// debug; change this back to 21 when done testing
	for(r.Human.Score < 21 && r.Human.Score > 0) {
		fmt.Println("*************************************************")
		fmt.Println("Your current score : ", r.Human.Score)
		fmt.Print("Your current hand : ")
		for _, i := range(r.Human.Cards){
			fmt.Print(i, " / ")
		}

		fmt.Println("\nDo you want to hit or stand?")
		fmt.Scanln(&response)

		if(response == "hit") {
			new_card := r.Hit()
			fmt.Println("You drew a :", new_card)
			r.Human.Cards = append(r.Human.Cards, new_card)
			r.Human.Score += getCardValue(new_card)

			if(r.Human.Score > 21){
				fmt.Println("BUST!")
				r.Human.Score *= -1
			} else if(r.Human.Score == 21) {
				fmt.Println("BLACKJACK!")
			}

		} else if(response == "stand") {
			break
		} else {
			fmt.Println(" ! Invalid response, please enter either \"hit\" or \"stand\" ! ")
		}
	}

	fmt.Println("Your round is over. Your final score : ", r.Human.Score)
	fmt.Println("*************************************************")
} 

// in this version, just reveal the hand of the dealer
func (r Round) DealerTurn() {
	return
}

func (r Round) EndRound() int {
	// print dealer info
	fmt.Println("Dealers hand : ")
	for _, i := range(r.Dealer.Cards) {
		fmt.Println(" . . . ", i, "  ")
	}
	fmt.Println("Dealers score : ", r.Dealer.Score)

	// print player info
	fmt.Println("Your cards : ")
	for _, i := range(r.Human.Cards) {
		fmt.Println(" . . . ", i, "  ")
	}
	fmt.Println("Your score : ", r.Human.Score)

	// determine the winner
	if(r.Human.Score > r.Dealer.Score) {
		fmt.Println(" ~~~~~~ Congratulations!!! you beat the dealer! ~~~~~~ ")
		return 1
	} else {
		fmt.Println(" ------- YOU LOSE -------")
		return 2
	}
}


func (p Player) getTotalScore(isHuman bool) int {
	temp_sum := 0
	for _, i := range(p.Cards) {
		if(i.Val == 1 && !isHuman) {
			temp_sum += 11
		} else {
			temp_sum += getCardValue(i)
		}
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