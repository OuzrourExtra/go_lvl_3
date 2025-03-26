package main

import (
	"fmt"
	"github.com/ouzrourextra/go_lvl_3/55_build_a_simple_to-do_list_with_generics/task"
	"github.com/ouzrourextra/go_lvl_3/55_build_a_simple_to-do_list_with_generics/input"
	"github.com/ouzrourextra/go_lvl_3/55_build_a_simple_to-do_list_with_generics/err"
	"github.com/ouzrourextra/go_lvl_3/55_build_a_simple_to-do_list_with_generics/clear"
	"encoding/json"
	"os"
)

func toJsonAndSaveIt(tasks []task.Task){
	jsonByte , jsonError := json.MarshalIndent(tasks,"","	")
	err.ErrorDetector(jsonError)
	errWriting := os.WriteFile("task.json",jsonByte,0644)
	err.ErrorDetector(errWriting)
}

func createTaskJson() []task.Task {
	length := input.InputInt("Number of tasks to enter ")
	tasks := task.CreateTasks(length)
	toJsonAndSaveIt(tasks)
	return tasks
}

func Application(tasks *[]task.Task){
	looped := true
	for{
		clear.ClearScreen()
		fmt.Printf("\n====================\n")
		fmt.Printf("TO-DO List App v1.0\n")
		fmt.Printf("By. Ilyas Ouzrour\n")
		fmt.Printf("====================\n")
		fmt.Printf("Choose an Option :\n")
		fmt.Printf("1) Added a Task\n")
		fmt.Printf("2) Delete a Task\n")
		fmt.Printf("3) Show All Tasks\n")
		fmt.Printf("4) Check/Uncheck Completed\n")
		choice := input.InputInt("Enter your choice ")
		switch choice {
		case 1 : 
			clear.ClearScreen()
			task.AppendTask(tasks)
			toJsonAndSaveIt(*tasks)
		case 2 :
			clear.ClearScreen()
			if(len(*tasks)!=0){task.ShowTasks(*tasks)}
			task.DeleteTask(tasks)
			if(len(*tasks)!=0){task.ShowTasks(*tasks)}

			toJsonAndSaveIt(*tasks)
		case 3:
			clear.ClearScreen()
			task.ShowTasks(*tasks)
		case 4:
			clear.ClearScreen()
			if(len(*tasks)!=0){task.ShowTasks(*tasks)}
			task.CheckUncheck(tasks)
			if(len(*tasks)!=0){task.ShowTasks(*tasks)}

			toJsonAndSaveIt(*tasks)
		default :
			clear.ClearScreen()
			fmt.Printf("Invalid Option , Reselect (1-3) !\n")
		}
		looped = input.InputBool("\nContinue The App (y/n)")
		if !looped {
			break
		}
	}
	
}

func UndoJson() []task.Task{
	dataJson , errReading := os.ReadFile("task.json")
	err.ErrorDetector(errReading)
	var tasks []task.Task
	json.Unmarshal(dataJson,&tasks)
	return tasks
}

func main(){
	_ , exist := os.Stat("task.json")
	if os.IsNotExist(exist){
		todo := createTaskJson()
		Application(&todo)
	}else{
		todo := UndoJson()
		Application(&todo)
	}
}
