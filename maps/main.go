package main

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
		println("initial Marks of", key, ":", value)
	}

}
