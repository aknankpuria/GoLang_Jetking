package main

import "fmt"

func main() {

	stdGrades := make(map[string]int) // initializing an empty map

	stdGrades["harry"] = 100
	stdGrades["paddy"] = 70
	stdGrades["sandy"] = 90
	stdGrades["mangesh"] = 68
	stdGrades["sachin"] = 66
	stdGrades["urmal"] = 93

	// printing the map
	for key, value := range stdGrades {
		fmt.Println("initial Marks of", key, ":", value)
	}
	// accessing particular element
	fmt.Println("Marks of urmal is:", stdGrades["urmal"])

	// deleting an element
	delete(stdGrades, "mangesh")
	// printing the map
	for key, value := range stdGrades {
		fmt.Println("after deletion Marks of", key, ":", value)
	}
	// updating an element
	stdGrades["urmal"] = 100
	// printing the map
	for key, value := range stdGrades {
		fmt.Println("after updation Marks of", key, ":", value)
	}

	// cheking if the key exists
	if grade, exists := stdGrades["urmal"]; exists {
		fmt.Println("Marks of urmal is:", grade)
	} else {
		fmt.Println("Marks of urmal is not present")
	}

	if grade, exists := stdGrades["mangesh"]; exists {
		fmt.Println("Marks of mangesh is:", grade)
	} else {
		fmt.Println("Marks of mangesh is not present")
	}

	// cheking the length of the map
	fmt.Println("length of the map is:", len(stdGrades))

	// iterating over the map
	for key, value := range stdGrades {
		fmt.Println("Marks of", key, ":", value)
	}

}
