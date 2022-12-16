package main 

import (
	"fmt"
	//"github.com/jaeyoony/deck_of_cards/deck"
	"github.com/jaeyoony/blackjack/dealer"
)

func main() {
	fmt.Println("Hello from no AI blackjack!")	
	
	// test roundStart
	new_round := dealer.StartRound(0)
	fmt.Println("TEST ROUND started")
	fmt.Println(". . New Round deck size : ", len(new_round.PlayDeck))
	fmt.Println(". . new round dealer : ", new_round.Dealer)
	fmt.Println(". . new round player : ", new_round.Human)

	next_card := new_round.Hit()
	fmt.Println(". . Next hit card : ", next_card)

}