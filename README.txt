Descriptions and more updates can be viewed on github: https://github.com/geeeeeeeeew/GS
Cybio Composer file is on the box of Automation Lab 258, with folder named "AutomatedCyBio"
Changes: only generator is implemented. Debugger can be achieved by the cybio composer.
--------------------------------------------------------------------------------------
How to tell GS which deck that each labware locates at, and what labwares will be used?
Create a .txt file with below format (ignore the lines, and make sure to include comma)
---------------------------------------------------------------------------------------
head name, type
adapter name, adapter type, deck number which you want to put it to
tipbox name, tip type, deck number which you want to put it to
deck number of waste box
labware1 id, deck numbers which you want to put it to
labware1 id, deck numbers which you want to put it to
labware1 id, deck numbers which you want to put it to
----------------------------------------------------------------------------------------
you can put as many labwares as you want, as long as the total number of labwares is
smaller than or equal to 12, since there are only 12 decks available in Felix machine.

First line is the information for head.

Second line is for adapter. Some experiments may not use adapter, and you could type None.

Third line is for tipbox.

Forth line is for position to put waste box.

Remaining lines are for the decks' information of other labwares, like type of cell plate
applied. Plates that are generally used in liquid handling experiments are 300-falcon 96
PS tissue culture, which contains 96 wells, and 25-Nunc OmniTray, which contains only 1
big well. The number, for example 300 and 25, is the labware id, so you don't have to
bother typing the entire name of labware.

Example layout files are provided, naming layout.txt and layout1.txt.

------------------------------------------------------------------------------------------
Layout of 300 Falcon PS tissue culture (each square indicates 1 well, 8 * 12 wells)

    1   2   3   4   5   6   7   8   9  10  11  12
   -------------------------------------------------
A | A1|   |   |   |   |   |   |   |   |   |   |   |
   -------------------------------------------------
B |   |   |   |   |   |   |   |   |   |   |   |   |
   -------------------------------------------------
C |   |   |   |   |   |   |   |   |   |   |   |   |
   -------------------------------------------------
D |   |   |   |   |   |   |   |   |   |   |   |   |
   -------------------------------------------------
E |   |   |   |   |   |   |   |   |   |   |   |   |
   -------------------------------------------------
F |   |   |   |   |   |   |   |   |   |   |   |   |
   -------------------------------------------------
G |   |   |   |   |   |   |   |   |   |   |   |   |
   -------------------------------------------------
H |   |   |   |   |   |   |   |   |   |   |   |   |
   -------------------------------------------------
