package main

import "fmt"

func (F *Felix) displayFelix() {
	fmt.Println("--Felix--")
	fmt.Println("	Head name:", F.head.name, ", Head support:", F.head.support)
	if F.adapter.name != "" {
		fmt.Println("	Adapter:", F.adapter.name, ", deck", F.adapter.deckPos)
	}
	fmt.Println("	Tips:", F.tips.name, ", deck", F.tips.deckPos)
	fmt.Println("	Waste box: deck", F.wastePos)
	fmt.Println("--Layout--")
	F.displayfelix()
}

func (F *Felix) displayfelix() {
	felix := F.felix
	for i, deck := range felix {
		if deck != nil {
			fmt.Println("	deck", deck.deckPos)
			if i == F.adapter.deckPos-1 {
				fmt.Println("		Adapter:",F.adapter.name)
			} else if i == F.tips.deckPos-1 {
				fmt.Println("		tipBox:",F.tips.name)
			} else if deck.labWareID == 25 {
				fmt.Println("		25 - Nunc OmniTray")
			} else if deck.labWareID == 300 {
				fmt.Println("		300 - Falcon 96 PS tissue culture")
			}
			//deck.displayPlate()
		} else {
			fmt.Println("	deck",i+1)
		}
	}
}

//when advanced track mode is activated, volume indicates the actual volume
//of each well, while the default mode, standard track mode, only tracks the
//changes of volume.
//Therefore, volume can never be negative under advanced track mode, but could
//be if under standard track mode.
func (deck *deck) displayPlate() {
	fmt.Println("		plate volume:")
	for i, _ := range deck.plate {
		fmt.Print("	        ")
		for j, _ := range deck.plate[i] {
			fmt.Print(deck.plate[i][j].vol," ")
		}
		fmt.Println()
	}
}