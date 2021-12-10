package main

import (
	"strconv"
	"strings"
)

func trimStringSpace(Arr []string) []string{
	for i := range Arr {
		Arr[i] = strings.TrimSpace(Arr[i])
	}
	return Arr
}

//return deckpos, platepos of src
func getPos(subCommand []string) (int,int) { //src command or dest command
	deckPos := 0
	platePos := 0
	trimStringSpace(subCommand)
	posStr := subCommand[0]
	for i := 1; i <= 12; i++ { //from col1 to col12 of the plate
		if strings.Contains(posStr, strconv.Itoa(i)) {
			platePos, _ = strconv.Atoi(subCommand[1])
			deckPos = i
		}
	}
	return deckPos, platePos
}

func getVol(src []string) int{
	vol, _ := strconv.Atoi(src[2])
	return vol
}

func getAdapterString(F *Felix) string {
	//two options: 8 channel adapter or cyBi 96
	if F.adapter.name == "8chAdapter" {
		return "8,"+ strconv.Itoa(F.adapter.deckPos)
	} else {
		return "96,"+strconv.Itoa(F.adapter.deckPos)
	}
}

func getTipString(F *Felix) string {
	return F.tips.name+","+strconv.Itoa(F.tips.maxVol)+","+strconv.Itoa(F.tips.deckPos)
}

func getTipId(F *Felix) string{
	return "96" //I only used cyBi96 this semester, but it needs to be updated when there are other kinds of tips applied
}

func getAdapterId(F *Felix) string{
	return "8" //need to be extended when more adapters options are applied
}