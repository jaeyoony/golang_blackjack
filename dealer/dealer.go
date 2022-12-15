package dealer

import( 
	"fmt"
	"github.com/jaeyoony/deck_of_cards/deck"
)

type Player struct {
	Card1 deck.Card
	Card2 deck.Card
}

func DealerPrint() {
	fmt.Println("Dealer says hello")
	newdeck := deck.New()

	fmt.Println("First card : ", newdeck[0])
}

