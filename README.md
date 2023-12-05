# quiz_game
A simple math quiz game. 

v.1
Well this program reads csv file in the format of q,a. It uses a flag to get the input csv file. Its default file to read is problems.csv, but with the csv flag you can provide the path to any csv file.

It uses the os package to open & read the csv file & displays appropriate errors if any. The input records are stored into a slice of string slices. 

A parsing function is implemented to create a slice of struct {p,q}. It iterates over this struct using ranges to ask a q & wait for user to input the answer. It also trims whitespaces while paresing answer.

An int variable is used to keep track of all correct answers & finally a score is presented to user after the completetion of quiz.

v.2

v.2.1
Another flag is implemented to have timelimit to the quiz. The idea is to use channels & go routines of golang to keep a timer running in the background & end the quiz once the timer stops.

NewTimer method of Time package is used create a new timer. In the loop iterations, select is implemented. A case is used to check for timer expiration & the default case asks the question & gets the answer from user. 

Limitation: If the program is waiting for user's input for answer then it continues to wait even after timer expiration. Although the answer after timer expiration is not considered in final score irrespective of its correctness, its not a great user experience.

v.2.2
Moved the question at the start of the loop. Implemented a go-routine(closure to be specific) that scans for an answer from user & sends via a channel.

Modify the select to have another case that checks if an answer is received via the answer channel & checks its correctness. There is no default case any more.
This way the quiz ends even rightaway after timer expiration & doesn't wait for user input.

v.2.3
Implement lable to break the loop. Kinda like goto's from C. But it is helpful to break the loop, instead of return statements, also allows to just use println to enter new line so that program exits & avoid code duplications in printing final score.
