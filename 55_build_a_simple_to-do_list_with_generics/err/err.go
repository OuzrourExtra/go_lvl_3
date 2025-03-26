package err

import (
	"log"
)

func ErrorDetector(err error){
	if err != nil{
		log.Fatal(err)
	}
}