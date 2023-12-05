# quiz_game
A simple math quiz game. 

v.1
Well this program reads csv file in the format of q,a. It uses a flag to get the input csv file. Its default file to read is problems.csv, but with the csv flag you can provide the path to any csv file.

It uses the os package to open & read the csv file & displays appropriate errors if any. The input records are stored into a slice of string slices. 

A parsing function is implemented to create a slice of struct {p,q}. It iterates over this struct using ranges to ask a q & wait for user to input the answer. It also trims whitespaces while paresing answer.

An int variable is used to keep track of all correct answers & finally a score is presented to user after the completetion of quiz.