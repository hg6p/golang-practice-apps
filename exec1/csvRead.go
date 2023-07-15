package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type Row struct {
	question string
	answere  string
}

var counter int = 0

func quizQuestions(data [][]string) []Row {
	var questions []Row
	//  i represent index and other variable is dataType[i]
	for _, line := range data {
		var quest Row
		for j, field := range line {

			if j == 0 {
				quest.question = field
			} else if j == 1 {
				quest.answere = field
			}
		}
		questions = append(questions, quest)
	}
	return questions
}

func main() {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	quizList := quizQuestions(data)

	duration := 5 * time.Second

	timer := time.NewTimer(duration)

	// Waiting for the timer to expire
	fmt.Printf("start timer press entrer")
	var start string
	fmt.Scanln(&start)
	fmt.Println("Timer started")
	for _, q := range quizList {
		answerCh := make(chan string)
		go func() {
			fmt.Printf("%s, Enter your answer: ", q.question)
			var answer string
			fmt.Scanln(&answer)
			if answer == q.answere {
				counter++
			}
			answerCh <- q.answere

		}()
		select {
		case <-timer.C:
			fmt.Println("Timer expired")
			fmt.Printf("score %d/%d", counter, len(quizList))
			return
		case answer := <-answerCh:
			fmt.Println("Received answer:", answer)
		}

	}

	fmt.Println("Timer expired")
	fmt.Printf("score %d/%d", counter, len(quizList))

}
