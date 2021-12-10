package main

import (
	"bufio"
	"fmt"
	os2 "os"
	"strconv"
	"strings"
)

func GenerateProtocol(F *Felix) {
	outFile := createFile()
	if len(os2.Args) > 2 {
		fileProvided, filename := getCommandFile()
		if fileProvided {
			commandFile := openFile(filename)
			generate(F, commandFile, outFile)
		} else {
			generateWithQuery(F, outFile)
		}
	} else if len(os2.Args) == 1 {
		commandFile := openFile("mix.txt")
		generate(F, commandFile, outFile)
	}
}

func generate(F *Felix, commandFile *os2.File, outFile *os2.File) {
	defer commandFile.Close()
	protocolArr := initialArr(F)
	generatedArr := parseInstruction(commandFile, F)
	protocolArr = append(protocolArr, generatedArr...)
	protocolArr = append(protocolArr,unload(F)...)
	for _, protocol := range protocolArr {
		fmt.Fprintln(outFile, protocol)
	}
	fmt.Println("done. Please check protocol.txt")
}

func initialArr(F *Felix) []string {
	Arr := make([]string, 0)
	//to let cyBio interface read the csv file, hardcoding the first line
	Arr = append(Arr, "command,var1,var2,var3,var4,var5,var6")
	Arr = append(Arr, "light,on")
	if F.adapter.name != "" {
		Arr = append(Arr,"load,"+F.head.name+","+getAdapterString(F)+","+F.head.support)
		Arr = append(Arr,"set tool,"+ F.adapter.name+",no")
		Arr = append(Arr, "load,"+F.adapter.name+","+strconv.Itoa(F.tips.maxVol)+","+
				strconv.Itoa(F.tips.deckPos)+","+getTipId(F))
	} else {
		Arr = append(Arr,"load,"+F.head.name+","+getTipString(F)+","+F.head.support)
		Arr = append(Arr,"set tool,"+F.head.name+","+getTipString(F))
	}
	return Arr
}

func askInput() []string{
	fmt.Println("Enter your command:")
	scanner := bufio.NewScanner(os2.Stdin)
	scanner.Scan()
	input := scanner.Text()
	commands := strings.Split(input,",")
	trimStringSpace(commands)
	fmt.Println(commands)
	return commands
}

func generateWithQuery(F *Felix, file *os2.File) {
	protocols := initialArr(F)
	commands := askInput()
	for commands[0] != "quit" {
		if len(commands) > 0 {
			if commands[0] == "mix" {
				protocol := generateMix(commands, F)
				protocols = append(protocols, protocol...)
			} else if commands[0] == "transfer"{ //transfer
				protocol := generateTrans(commands, F)
				protocols = append(protocols, protocol...)
			}
		}
		commands = askInput()
	}
	protocols = append(protocols,unload(F)...)
	for _, protocol := range protocols{
		fmt.Fprintln(file, protocol)
	}
	fmt.Println("done.Please check protocol.txt")
}
func createFile() *os2.File {
	file, err := os2.Create("protocol.txt")
	if err != nil {
		panic("Error: Problem Creating a file\n")
	}
	//fmt.Println("Created protocol.txt successfully")
	return file
}

/*
getCommandFile() returns a boolean value and a string
return true, filename when user provides a file that
contains commands, and filename is the name of this file.
return false, empty string when user does not provide
or wants to input one command at a time in terminal.
 */
func getCommandFile() (bool, string){
	if len(os2.Args) > 2 {
		fmt.Print("\nI've found a file named ", os2.Args[2])
		fmt.Print(". Does it contain the instructions for the protocol? ")
		var isFile string
		fmt.Scanln(&isFile)
		if !strings.Contains(strings.ToLower(isFile), "n") {
			fmt.Println("\nGot it. Protocols will be generated based on", os2.Args[2])
			fmt.Println("\nStart parsing and generating protocols")
			return true, os2.Args[2]
		} else {
			fmt.Println("\nGot it. file", os2.Args[2], "will be ignored")
		}
	}
	return false,""
}


func parseInstruction(commandFile *os2.File, F *Felix) []string{
	protocols := make([]string, 0)
	scanner := bufio.NewScanner(commandFile)
	for lineno := 1; scanner.Scan(); lineno++ {
		commandStr := strings.TrimSpace(scanner.Text())
		commands := strings.Split(commandStr, ",")
		trimStringSpace(commands)
		fmt.Println(commands)
		if len(commands) == 0 {
			continue
		}
		if commands[0] == "mix" {
			protocol := generateMix(commands, F)
			protocols = append(protocols, protocol...)
		} else { //transfer
			protocol := generateTrans(commands, F)
			protocols = append(protocols, protocol...)
		}
	}
	return protocols
}

func generateMix(commandArr []string, F *Felix) []string {
	//mix, src1 info, src2 info, dest info
	//mix,deck10 1A 20,deck12 2B 10,deck8 1A
	src1 := strings.Split(commandArr[1]," ")
	src2 := strings.Split(commandArr[2]," ")
	dest := strings.Split(commandArr[3]," ")
	if len(dest) != 0 {
		protocol := Mix(src1, src2, dest, F)
		return protocol
	} else {
		protocol := Trans(src1, src2, F)
		return protocol
	}
}

func generateTrans(commandArr []string, F *Felix) []string {
	if len(commandArr) <= 1{
		fmt.Println("below command not included\n",commandArr)
		return make([]string,0)
	}
	src1 := trimStringSpace(strings.Split(commandArr[1]," "))
	dest := trimStringSpace(strings.Split(commandArr[2]," "))
	protocol := Trans(src1, dest, F)
	return protocol
}

func Mix(src1, src2, dest []string, F *Felix) []string {
	if len(dest) == 0 {
		return Trans(src1, src2, F)
	}
	protocols := make([]string, 0)
	trans1 := Trans(src1, dest, F)
	protocols = append(protocols, trans1...)
	trans2 := Trans(src2, dest, F)
	protocols = append(protocols, trans2...)
	return protocols
}

func Trans(src, dest []string, F *Felix) []string {
	protocol := make([]string,0)
	srcInfo := getMoveInfo(src)
	labware := strconv.Itoa(F.felix[srcInfo.deckPos-1].labWareID)
	tool := "" //adapter or tipbox
	if F.adapter.name != "" {
		tool = getAdapterId(F)
	} else {
		tool = getTipId(F)
	}
	protocol = append(protocol, "piston")
	protocol = append(protocol,"move,"+strconv.Itoa(srcInfo.deckPos)+","+
		strconv.Itoa(srcInfo.platePos)+","+tool+","+labware+",bottom,2.5")
	protocol = append(protocol, "aspirate,"+strconv.Itoa(srcInfo.vol))

	destInfo := getMoveInfo(dest)
	protocol = append(protocol,"move,"+strconv.Itoa(destInfo.deckPos)+","+
		strconv.Itoa(destInfo.platePos)+","+tool+","+labware+",up,0")
	protocol = append(protocol, "dispense,"+strconv.Itoa(srcInfo.vol))

	return protocol
}

func getMoveInfo(moveCommand []string) *Move {
	deckPos, platePos := getPos(moveCommand)
	vol := getVol(moveCommand)
	return &Move{deckPos,platePos,vol}
}

func unload(F *Felix) []string {
	protocols := make([]string, 0)
	var pos string
	if F.adapter.name != "" {
		pos = strconv.Itoa(F.adapter.deckPos)
	} else {
		pos = strconv.Itoa(F.tips.deckPos)
	}
	protocols = append(protocols,"unload,"+F.head.name+","+pos +","+F.head.support)
	protocols = append(protocols,"light,off")
	protocols = append(protocols,"vertical")
	return protocols
}