package main

import (
	"errors"
	"fmt"
)

func flattenList(input interface{}) ([]interface{}, error) {
	var result []interface{}

	switch input.(type) {
	case []interface{}:
		for _, v := range input.([]interface{}) {
			if v != nil {
				flattened, err := flattenList(v)
				if err != nil {
					return nil, err
				}
				result = append(result, flattened...)
			}
		}
	default:
		if input != nil {
			result = append(result, input)
		} else {
			return nil, errors.New("Invalid input: nil value found")
		}
	}

	return result, nil
}

func main() {
	input := []interface{}{1, []interface{}{2, 3, nil, 4}, []interface{}{nil}, 5}
	flattenedList, err := flattenList(input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(flattenedList)
	}
}
