package array

import (
	"fmt"
	"log"
	"testing"
)

func TestArray(t *testing.T) {

	Arrayofinteger := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	Intarray := Stack[int]{}

	for _, val := range Arrayofinteger {
		//push array
		Intarray.Push(val)
	}
	//test pop array
	Intarray.Pop()

	//test find index of the array
	index, err := Intarray.FindArrayIndex(2)

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(index)

	//remove specific index array
	Intarray.Splice(index)

	fmt.Println(Intarray.element)
}

func TestArrayOfObject(t *testing.T) {

	type User struct {
		Id    int
		Name  string
		Email string
	}

	ArrayofObject := []User{
		{Id: 1, Name: "Remo", Email: "remo@gmail.com"},
		{Id: 2, Name: "Ambi", Email: "ambi@gmail.com"},
		{Id: 3, Name: "Anniyan", Email: "anniyan@gmail.com"},
	}

	arrayObj := Stack[User]{}

	for _, val := range ArrayofObject {
		//push array
		arrayObj.Push(val)
	}
	//test pop array
	arrayObj.Pop()

	//test find index of the array
	index, err := arrayObj.FindArrayObjectIndex("Email", "remo@gmail.com", "string")

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(index)

	//remove specific index array
	arrayObj.Splice(index)

	fmt.Println(arrayObj.element)
}
