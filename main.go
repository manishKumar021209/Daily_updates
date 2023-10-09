package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

//new

func main() {
	//taking the csv file
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("Limit", 10, "the time limit for te quiz in seconds")

	flag.Parse() //parsing the command line arguments(in this case those arguments will be the answers to the questions in the csv file)

	file, err := os.Open(*csvFilename) //opening the csv file with os.Open method
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename))

	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the csv file.")
	}

	problems := parselines(lines)

	//created a timer to limit te duration of game
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	//printing out each problem in a formatted the manner
	// 	Problem #1: 5+5=
	// Problem #2: 1+1=
	// Problem #3: 8+3=
	// Problem #4: 1+2=
	// Problem #5: 10+99=
	// Problem #6: 8+6=
	// Problem #7: 3+1=
	// Problem #8: 1+4=
	// Problem #9: 2+4=
	// Problem #10: 80-5=

	correct := 0

problemloop:
	//a counter that keeps the count of correct answers
	for i, p := range problems {

		fmt.Printf("Problem #%d: %s= ", i+1, p.q)

		answerCh := make(chan string)

		go func() {
			//giving the answer to the question as user input
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}

			// if answer == p.q {
			// 	fmt.Println("Correct!")
			// } else {
			// 	fmt.Println("wrong!")
			// }
		}

	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

func parselines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],                    //at column zero a question is stored
			a: strings.TrimSpace(line[1]), //at colimn one its answer is stored
		}
	}
	return ret //returning the 2d slice which is storing qna
	//[{5+5 10} {1+1 2} {8+3 11} {1+2 3} {10+99 109} {8+6 14} {3+1 4} {1+4 5} {2+4 6} {80-5 75}]
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
