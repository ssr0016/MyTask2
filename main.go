package main

import (
	"errors"
	"fmt"
)

func FlattenList(input interface{}) ([]interface{}, error) {
	switch input.(type) {
	case nil:
		return nil, errors.New("input is nil")
	case []interface{}:
		fmt.Println("Current input list:", input)
		result := make([]interface{}, 0)
		for _, v := range input.([]interface{}) {
			if v != nil {
				switch v.(type) {
				case []interface{}:
					flattened, err := FlattenList(v)
					if err != nil {
						return nil, err
					}
					fmt.Println("Flattened list:", flattened)
					result = append(result, flattened...)
				default:
					result = append(result, v)
				}
			}
		}
		return result, nil
	default:
		return []interface{}{input}, nil
	}
}

func main() {
	input := []interface{}{1, []interface{}{2, 3, nil, 4}, []interface{}{nil}, 5}
	output, err := FlattenList(input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		finalOutput := make([]interface{}, 0)
		for _, val := range output {
			if val != nil {
				finalOutput = append(finalOutput, val)
			}
		}
		fmt.Println("Final output:", finalOutput)
	}
}

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// // FlattenList function to flatten a nested list
// func FlattenList(input interface{}) ([]interface{}, error) {
// 	result := make([]interface{}, 0)

// 	switch input.(type) {
// 	case nil:
// 		return result, errors.New("input is nil")
// 	case []interface{}:
// 		for _, v := range input.([]interface{}) {
// 			if v != nil {
// 				switch v.(type) {
// 				case []interface{}:
// 					flattened, err := FlattenList(v)
// 					if err != nil {
// 						return nil, err
// 					}
// 					result = append(result, flattened...)
// 				default:
// 					result = append(result, v)
// 				}
// 			}
// 		}
// 	default:
// 		return result, errors.New("input is not a list")
// 	}

// 	return result, nil
// }

// func main() {
// 	input := []interface{}{1, []interface{}{2, 3, nil, 4}, []interface{}{nil}, 5}
// 	output, err := FlattenList(input)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		finalOutput := make([]interface{}, 0)
// 		for _, val := range output {
// 			if val != nil {
// 				finalOutput = append(finalOutput, val)
// 			}
// 		}
// 		fmt.Println("Output:", finalOutput)
// 	}
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func flattenList(input interface{}) ([]interface{}, error) {
// 	var result []interface{}

// 	switch input.(type) {
// 	case []interface{}:
// 		for _, v := range input.([]interface{}) {
// 			if v != nil {
// 				flattened, err := flattenList(v)
// 				if err != nil {
// 					return nil, err
// 				}
// 				result = append(result, flattened...)
// 			}
// 		}
// 	default:
// 		if input != nil {
// 			result = append(result, input)
// 		} else {
// 			return nil, errors.New("Invalid input: nil value found")
// 		}
// 	}

// 	return result, nil
// }

// func main() {
// 	input := []interface{}{1, []interface{}{2, 3, nil, 4}, []interface{}{nil}, 5}
// 	flattenedList, err := flattenList(input)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println(flattenedList)
// 	}
// }
