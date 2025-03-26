package input

import(
	"fmt"
	"log"
	"strconv"
	"bufio"
	"os"
	"strings"
)
/* 
	String
*/

func InputString(sentence string) string{
	fmt.Printf("============================")
	fmt.Printf("\n%s : ",sentence)
	reader  := bufio.NewReader(os.Stdin)
	str , errReading := reader.ReadString('\n')
	if errReading != nil{
		log.Fatal(errReading)
	}
	str = strings.TrimSpace(str)
	return str
}

/*
	Float64 & Float32
*/

func InputFloat64(sentence string) float64 {
	str := InputString(sentence)
	// Conversion
	float , errConversion := strconv.ParseFloat(str,64)
	if errConversion != nil{
		log.Fatal(errConversion)
	}
	return float
}


func InputFloat32(sentence string) float32 {
	str := InputString(sentence)
	// Conversion
	float , errConversion := strconv.ParseFloat(str,32)
	if errConversion != nil{
		log.Fatal(errConversion)
	}
	return float32(float)
}

/*
	Integers : Int32 Int64 Int
*/

func InputInt64(sentence string) int64 {
	str := InputString(sentence)
	integer , errConversion := strconv.ParseInt(str,10,64)
	if errConversion != nil{
		log.Fatal(errConversion)
	}
	return integer
}


func InputInt32(sentence string) int32 {
	str := InputString(sentence)
	integer , errConversion := strconv.ParseInt(str,10,32)
	if errConversion != nil{
		log.Fatal(errConversion)
	}
	return int32(integer)
}


func InputInt(sentence string) int {
	str := InputString(sentence)
	integer , errConversion := strconv.Atoi(str)
	if errConversion != nil{
		log.Fatal(errConversion)
	}
	return integer
}

/* 
 Rune
*/ 

func InputRune(sentence string) rune {
	str := InputString(sentence)
	return rune(str[0])
}

/* 
 Byte Slice
*/ 

func InputByteSlice(sentence string) []byte{
	str := InputString(sentence)
	return []byte(str)
}

/*
	Input Boolean
*/


func InputBool(sentence string)bool{
	str := InputString(sentence)
	if str == "y" || str == "yes" {
		return true
	} else if str == "n" || str == "no"{
		return false
	}else{
		return InputBool(sentence)
	}
}
