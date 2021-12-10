package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic("Error: Problem opening the file\n")
	}
	//fmt.Println("Opened the file", filename, "successfully")
	return file
}

func readFelix(file *os.File) *Felix {
	F := Felix{}
	scanner := bufio.NewScanner(file)
	for lineno := 1; scanner.Scan(); lineno++ {
		layout := strings.Split(scanner.Text(), ",")
		for i, _ := range layout {
			layout[i] = strings.TrimSpace(layout[i])
		}
		//fmt.Println(layout)
		fillFelix(&F, lineno, layout)
	}
	file.Close()
	fmt.Println("\nWelcome to GS! I've successfully initialized Felix with file")
	return &F
}

func fillFelix(F *Felix, lineno int, layout []string) {
	if lineno == 1 {
		F.updateHead(layout)
	} else if lineno == 2 {
		F.updateAdapter(layout)
	} else if lineno == 3 {
		F.updateTips(layout)
	} else if lineno == 4 {
		F.wastePos, _ = strconv.Atoi(layout[0])
	} else {
		F.updatefelix(layout) //update [][]well
	}
}

func (F *Felix) updateHead(layout []string) {
	head := Head{name: layout[0], support: layout[1]}
	F.head = head
}

func (F *Felix) updateAdapter(layout []string) {
	if len(layout) == 2 {
		pos, _ := strconv.Atoi(layout[1])
		F.adapter.name, F.adapter.deckPos = layout[0], pos
		F.felix[pos-1] = &deck{deckPos: pos}
	}
}

func (F *Felix) updateTips(layout []string) {
	vol, _ := strconv.Atoi(layout[1])
	pos, _ := strconv.Atoi(layout[2])
	F.tips.name, F.tips.maxVol, F.tips.deckPos = layout[0], vol, pos
	F.felix[pos-1] = &deck{deckPos: pos}
}

func (F *Felix) updatefelix(layout []string) {
	for i := 1; i < len(layout); i++ {
		id, _ := strconv.Atoi(layout[0])
		pos, _ := strconv.Atoi(layout[i])
		deck := deck{deckPos: pos, labWareID: id}
		F.felix[pos-1] = &deck
		if deck.labWareID == 25 {
			plate := make([][]well, 1)
			plate[0] = make([]well, 1)
			F.felix[pos-1].plate = plate
		} else if deck.labWareID == 300 {
			plate := make([][]well, 8)
			for i := range plate {
				plate[i] = make([]well, 12)
			}
			F.felix[pos-1].plate = plate
		}
	}
}



func main() {
	filename := ""
	if len(os.Args) == 1{
		fmt.Println("Default setting: layout.txt for Felix layout file and mix.txt for command file")
		filename = "layout.txt"
	} else {
		filename = os.Args[1]
	}
	file := openFile(filename)
	F := readFelix(file)
	F.displayFelix()
	GenerateProtocol(F)
	//for testing, since there's no given file for layout, I will make it "layout.txt" as default
}
