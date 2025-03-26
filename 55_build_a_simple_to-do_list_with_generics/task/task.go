package task

import (
	"fmt"
	"time"
	"github.com/ouzrourextra/go_lvl_3/55_build_a_simple_to-do_list_with_generics/input"
)

type Task struct{
	Name string
	Creation_Date time.Time
	Done bool
}
func InputTask()(task Task){
	name := input.InputString("Enter the task ") 
	return Task{name,time.Now(),false}
}

func CreateTasks(length int) []Task{
	tasks := make([]Task,length)
	for i:=0 ; i<length ; i++{
		tasks[i].Name = input.InputString("Enter the task ")
		tasks[i].Creation_Date = time.Now()
		tasks[i].Done = false
	}
	return tasks
}

func AppendTask (tasks *[]Task){
	newTask := InputTask()
	*tasks = append(*tasks,newTask)
}

func ShowTasks ( tasks []Task){
	if len(tasks) == 0 {
		fmt.Printf("=====================\n")
		fmt.Printf("Empty List !\n")
	}else{
	for key , value := range tasks{
		fmt.Printf("=====================\n")
		fmt.Printf(" Task %d :\n",key+1)
		fmt.Printf("=====================\n")
		fmt.Printf(" To-do : %s\n",value.Name)
		fmt.Printf(" Created : %s\n",value.Creation_Date.Format("02 Jan 2006 - 15:04"))
		fmt.Printf(" Status : %t\n",value.Done)
	}
	}
	
}

func DeleteTask(tasks *[]Task){
	
	if len(*tasks) == 0 {
		fmt.Printf("=====================\n")
		fmt.Printf("\n==Sorry ! No tasks !\n")
	}else{
		index := input.InputInt(fmt.Sprintf("Element You want to Delete ( between : 1-%d)",len(*tasks)))
		if index <1 || index >len(*tasks) {
			fmt.Printf("\n OUT OF INDEX ! \n")
		}else{
			*tasks = append((*tasks)[:index-1],(*tasks)[index:]...)
		}
	}
}

func CheckUncheck(tasks *[]Task){
	if len(*tasks) == 0 {
		fmt.Printf("=====================\n")
		fmt.Printf("Sorry ! No tasks !\n")
	}else{
		index := input.InputInt(fmt.Sprintf("Element You want to Check/Uncheck ( between : 1-%d)",len(*tasks)))
		if index <1 || index >len(*tasks) {
			fmt.Printf("\n OUT OF INDEX ! \n")
		}else{
			(*tasks)[index-1].Done = !(*tasks)[index-1].Done
		}
	}
	
}