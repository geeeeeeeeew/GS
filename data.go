package main

/*
1."lights"
	status: on, off

2."load"
	general1: "Head R 96"
	general2: "8-Channel Adapter; Head R 96"
	deckPos: 6
	labware.type: "Support; 37.0 mm height (OL3317-11-120)"
	labware.well: ""

	general1: "8-Channel Adapter"
	general2: "CyBi-Tip 250µ"
	deckPos: 7
	labware.toolType: "CyBi-TipBox 96"
	labware.well: "1A"

3."setTool"
	head: "8-Channel Adapter"
	adapter: "No add-on"

4."moveTo"
	tipLayout: "8-channel (column)" 一列
			   "96-channel" 全部一起
	deckPos: 2
	labware.type: "25 - Nunc OmniTray"
				  "300 - Falcon 96 PS tissue culture"
	labware.well: "1"
				  "1A"
	labware.reference: "well bottom"
					   "well top"
	offset.x: 10
	offset.y: 0
	offset.z: 2.5 只看这个就行了，代码里只有这个

5."pistonToZero"

6."aspirate"
	overstroke: 0
	volume: "200" 有"max"和"min"

7."dispense"
	blowout: 0
	volume: "25" 有"max"和"min"

8."unload"
	general1: "8-Channel Adapter"
			  "Head R 96"
	deckPos: 4
	labware.type: "Waste" 对于unload，labware里只看这个
	labware.well: ""
	labware.reference: ""

9."setPipetteSpeed"
	speed: "10" 有"max","min"和"default"

10."verticalDrive" //right before turn off the light
	zReference: "topmost" 有"well bottom"，"well top"
	absolute:""
 */

/*
Head R 96
Support; 37.0 mm height (OL3317-11-120) //head的support，在第一次load的时候
1)8-Channel Adapter; Head R 96
2)CyBi-RoboTipTray 96 / CyBi-Tip 250µl DW
25 - Nunc OmniTray
300 - Falcon 96 PS tissue culture
Labware Type: CyBi-TipBox 96
tips: CyBi-Tip 250µl
layout: 8-channel (column)
 */

type deck struct {
	deckPos int
	labWareID int //uniquely define which labware applied (25 -> Nunc OmniTray, 300 -> Falcon 96 PS tissue culture
	plate [][]well
}

type well struct {
	vol float64 //may not be used
	solution string //may not be used
}

type Felix struct {
	head Head
	adapter Adapter
	tips Tips
	wastePos int //position of waste box
	felix [12]*deck
}

type Head struct {
	name string //head96
	support string //37.0 mm height (OL3317-11-120), for example
}

type Adapter struct {
	name string //8chAdapter
	deckPos int
}

/*
maxVol: maximum volume of one tip. According to the setup of
		Cybio Felix machine, it should be integer only.
 */
type Tips struct {
	name string //cyBi96
	maxVol int
	deckPos int
}

type Move struct { //source
	deckPos int
	platePos int
	vol int
}