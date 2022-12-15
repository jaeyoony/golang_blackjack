package main 

import (
	"fmt"
	"github.com/jaeyoony/deck_of_cards/deck"
	"github.com/jaeyoony/blackjack/dealer"
)

func main() {
	fmt.Println("Hello from no AI blackjack!")	
	_ = deck.New()
	// for _, i := range(new_deck) {
	// 	fmt.Println(i)
	// }

	dealer.DealerPrint()
}