package array

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Stack[T comparable] struct {
	element []T
}

func (s *Stack[T]) Push(element T) {
	s.element = append(s.element, element)
}

func (s *Stack[T]) Pop() T {
	if len(s.element) == 0 {
		var zero T
		return zero
	}
	element := s.element[len(s.element)-1]
	s.element = s.element[:len(s.element)-1]
	return element
}

func (s *Stack[T]) Splice(index int) {
	if index < 0 || index >= len(s.element) {
		return // Return the original slice if index is out of bounds
	}
	copyarr := append(s.element[:index], s.element[index+1:]...)
	s.element = copyarr
}

func (s *Stack[T]) FindArrayIndex(findone T) (index int, err error) {
	for i, val := range s.element {
		flg := Compare(val, findone)
		if flg {
			return i, nil
		}
	}
	return -1, errors.New("")
}

func (s *Stack[T]) FindArrayObjectIndex(Key string, value string, datatype string) (index int, err error) {
	for i, val := range s.element {
		d := ConvertObject(val)
		bytedata, _ := json.Marshal(d)
		var need map[string]interface{}
		json.Unmarshal(bytedata, &need)
		if need[Key] == nil {
			return -1, errors.New("key not found")
		} else {
			switch datatype {
			case "int":
				v, ok := need[Key].(float64)
				if !ok {
					return -1, errors.New("value is not an integer")
				}
				// Convert the float64 value to an int
				va := int(v)
				a, _ := strconv.Atoi(value)
				if va == a {
					return i, nil
				}
			case "string":
				if need[Key] == strings.TrimSpace(value) {
					fmt.Println("inside")
					return i, nil
				}
			case "float64":
				f, err1 := strconv.ParseFloat(value, 64)
				if err1 != nil {
					fmt.Println("Error:", err)
					return
				}
				if need[Key].(float64) == float64(f) {
					fmt.Println("inside")
					return i, nil
				}
			}

		}
	}
	return -1, errors.New("")
}

func Compare[T comparable](a, b T) bool {
	return a == b
}

func ConvertObject[T any](value T) interface{} {
	return value
}

func (s *Stack[T]) Concat(element ...[]T) {

	for _, val := range element {
		s.element = append(s.element, val...)
	}
}
