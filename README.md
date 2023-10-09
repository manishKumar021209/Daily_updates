# Daily_updates

# Quiz Game in Go

This is a simple quiz game implemented in Go. The program reads questions and answers from a CSV file and allows the user to answer each question within a specified time limit. It then provides a score based on the number of correct answers.

## Features

- Loads questions and answers from a CSV file.
- Sets a time limit for answering questions.
- Accepts user input for answers.
- Calculates and displays the final score.

## Getting Started

To run the Quiz Game, follow these steps:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/quiz-game.git
2. Change into the project directory:cd quiz-game
3. Run the Go program with optional command-line arguments:go run main.go -csv questions.csv -Limit 10
   The -csv flag specifies the CSV file containing questions and answers (default is "problems.csv").
   The -Limit flag sets the time limit for the quiz in seconds (default is 10 seconds).

4. Answer the questions as they are presented, and the program will calculate your score at the end.


 CSV File Format
The CSV file should be in the format of 'question, answer' where each line represents a single question and its corresponding answer. For example:
5+5, 10
1+1, 2
8+3, 11
1+2, 3
10+99, 109
8+6, 14
3+1, 4
1+4, 5
2+4, 6
80-5, 75



Author
Manish
