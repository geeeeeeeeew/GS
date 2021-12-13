# GS

To run the program, in the terminal, enter ./GS_finalProject layoutFileName.txt commandFileName.txt
The program will print out the layout of Felix according to the layout File you provided, and it will ask it the command file containing the commands. If answer is yes, then the program will generate protocols based on the command file, and write to file called protocol.txt. If answer is no, the program will ask you to enter a valid command at a time, and genrate it based on it until you input quit. 

Default setting(if user only input ./GS_finalProject): layoutFile will be layout.txt, and commandFile will be mix.txt

Valid Command: 

               1)mix deck+deck_number source1_col volume deck+deck_number source2_col volume deck+deck_number dest_col volume
                 mix deck10 1 20 deck12 2 10 deck2 10 30
               
               2)transfer deck+deck_number source_col volume deck+deck_number dest_col volume
                 transfer deck10 1 20 deck12 2 10
                 
Note that this progect does not implement debugger, so that users have to follow the rules when writing commands (one of the changes compared to the proposal). 
                 
