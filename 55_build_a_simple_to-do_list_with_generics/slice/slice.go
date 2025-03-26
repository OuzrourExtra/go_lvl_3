package slice

import (
	"fmt"
	"reflect"

	"github.com/ouzrourextra/go_lvl_3/55_build_a_simple_to-do_list_with_generics/input"
)

func CreateSlice[T any]() []T {
	size := input.InputInt("The number of elements you want to add ")
	slice := make([]T, 0, size)

	for i := 0; i < size; i++ {
		val := readValue[T](fmt.Sprintf("Value of [%d]", i))
		slice = append(slice, val)
	}
	fmt.Printf("\n======= Created Successfuly ! =======")

	return slice
}


// Affiche les éléments du slice
func ReadSlice[T any](slice []T) {
	fmt.Printf("\n======= Content of the Slice =======")
	for i, v := range slice {
		fmt.Printf("\n[%d] %v\n", i, v)
	}
}

// Met à jour un élément du slice à l'index donné
func UpdateElement[T any](slice *[]T, index int) {
	if index < 0 || index >= len(*slice) {
		fmt.Println("\nInvalid Index.")
		return
	}
	fmt.Printf("\n======= Updating Element %d =======",index)
	newValue := readValue[T](fmt.Sprintf("New Value for [%d] : ", index))
	(*slice)[index] = newValue
	fmt.Printf("\n======= Done Successfuly ! =======")

}

func UpdateAll[T any](slice *[]T){
	fmt.Printf("\n======= Updating all Table =======")

	for i:=0 ; i<len(*slice) ; i++{
		newValue := readValue[T](fmt.Sprintf("\nNew Value for [%d] : ", i))
		(*slice)[i] = newValue
	}
	fmt.Printf("\n======= Done Successfuly ! =======")

}

// Supprime un élément du slice à l'index donné
func DeleteElement[T any](slice *[]T, index int) {
	fmt.Printf("\n======= Deleting Element from The Slice =======")

	if index < 0 || index >= len(*slice) {
		fmt.Println("\nIndex invalide.")
		return
	}
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
	fmt.Printf("\n======= Done Successfuly ! =======")

}

// Ajoute un élément à la fin du slice
func AppendElement[T any](slice *[]T) {
	fmt.Printf("\n======= Appending ELement to the Slice =======")

	newValue := readValue[T]("Value to be Added : ")
	*slice = append(*slice, newValue)
	fmt.Printf("\n======= Done Successfuly ! =======")

}

// Fonction privée pour lire une valeur selon le type (string, int, float64, etc.)
func readValue[T any](message string) T {
	var result any
	var zero T
	kind := reflect.TypeOf(zero).Kind()

	switch kind {
	case reflect.String:
		result = input.InputString(message)
	case reflect.Int:
		result = input.InputInt(message)
	case reflect.Int32:
		result = input.InputInt32(message)
	case reflect.Int64:
		result = input.InputInt64(message)
	case reflect.Float32:
		result = input.InputFloat32(message)
	case reflect.Float64:
		result = input.InputFloat64(message)
	default:
		panic("Type not supported in readValue")
	}

	return result.(T)
}
