package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"github.com/ouzrourextra/go_lvl_3/54_create_a_simple_command-line_quiz/input"
	"github.com/ouzrourextra/go_lvl_3/54_create_a_simple_command-line_quiz/quizGenerator"
)

func main(){
	quizGenerator.GenerateJSONQuiz("quiz.json")
	reader ,errReadFile := os.ReadFile("quiz.json")
	if errReadFile != nil {
		log.Fatal(reader)
	}
	var quiz []quizGenerator.Question
	errorUnmarshal := json.Unmarshal(reader,&quiz)
	if errorUnmarshal != nil{
		log.Fatal(errorUnmarshal)
	}
	score := 0
	for key , value := range quiz {
		fmt.Printf("\n======================")
		fmt.Printf("\n Question %d : ",key+1)
		fmt.Printf("%s",value.Question)
		fmt.Printf("\n======================")
		fmt.Printf("\n1) %s",value.Answers[0])
		fmt.Printf("\n2) %s",value.Answers[1])
		fmt.Printf("\n3) %s",value.Answers[2])
		fmt.Printf("\n4) %s",value.Answers[3])
		fmt.Printf("\n======================\n")
		answer := input.InputInt("Enter you asnwer")
		if answer == value.Correct {
			fmt.Printf("======================\n")
			fmt.Printf("==> ✅ Correct ! ")
			score++
		}else{
			fmt.Printf("======================\n")
			fmt.Printf("==> ❌ Incorrect ! ")
		}
		if key+1 == len(value.Answers){
			fmt.Printf("\n======================")
			fmt.Printf("\n☢️ Your SCORE is : %d/%d",score,len(value.Answers))
			fmt.Printf("\n======================")
		}
	}
}