package input

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"log"
)

func InputInt(sentence string)(number int){
	buffer := bufio.NewReader(os.Stdin)
	fmt.Printf("%s : ",sentence)
	str , errInput := buffer.ReadString('\n')
	if errInput != nil{
		log.Fatal(errInput)
	}
	str = strings.TrimSpace(str)
	number , errNumber := strconv.Atoi(str)
	if errNumber != nil{
		log.Fatal(errNumber)
	}
	return number
}