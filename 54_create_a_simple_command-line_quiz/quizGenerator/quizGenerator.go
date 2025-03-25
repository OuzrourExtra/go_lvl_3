package quizGenerator

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
)

type Question struct{
	Question string `json:"question"`
	Answers []string `json:"answers"`
	Correct int `json:"correct"`
}

func GenerateJSONQuiz(filename string){
	questions := []Question{
		{
			"What is the square of 4 ?",
			[]string{"4","8","16","2"},
			3},
		{
			"What is the square of 9 ?",
			[]string{"18","81","99","3"},
			2},
		{
			"what is the Square root of 64 ?",
			[]string{"4","32","16","8"},
			4},
		{
			"what is the square root of 25 ?",
			[]string{"2","10","25","5"},
			4},
	}
	
	jsonData , errJson := json.MarshalIndent(questions,"","	")
	if errJson != nil{
		log.Fatal(errJson)
	}
	
	errWriting := os.WriteFile(filename,jsonData,0644)
	if errWriting != nil{
		log.Fatal(errWriting)
	}
	fmt.Printf("File Generation Succefully !")
}