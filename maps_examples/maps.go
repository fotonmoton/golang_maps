package maps_examples

import "fmt"

/*
1. Fast search and fast insertion
2. Unordered
3. Only 1Some types are allowed as keys
4. Keys must be distinct
5. Create with make or literal, indicate size of the map beforehand
6. Delete entry with delete(), nil to clear memory
7. Pointer type (can be nil)
8. Two values assignment, if statement
9. Memory leaks with whole map in memory
10. Length with len()
11. Iterate with for range, unordered data
12. Map of maps
13. Maps package
*/

type testScore struct {
	studentName string
	score       uint8
}

type position struct {
	x int
	y int
}

func Instantiation() {
	// allocates a lot of memory
	withMake := make(map[string]bool, 1000)
	withLiteral := map[string]bool{}
	withValues := map[string]bool{"this is it": true}

	withMake["notFound"] = withLiteral["found?"]
	withValues["key"] = withMake["notFound"]
}

func FastSearch() {
	results := []testScore{
		{"John Doe", 20},
		{"Patrick Indexing", 15},
		{"Bob Ferring", 7},
		{"Claire Novalingua", 8},
	}
	fmt.Println(results)

	for _, result := range results {
		if result.studentName == "Claire Novalingua" {
			fmt.Println("Score Found:", result.score)
		}
	}

	mapResults := map[string]uint8{
		"John Doe":          20,
		"Patrick Indexing":  15,
		"Bob Ferring":       7,
		"Claire Novalingua": 8,
	}

	fmt.Println("Score Found:", mapResults["Claire Novalingua"])
}

func Unordered() {

	// Do not try this at home
	for iteration := range make([]bool, 100) {
		fmt.Println("Iteration #", iteration)

		capitals := map[string]string{
			"Ukraine": "Kyiv",
			"USA":     "Washington",
			"Brazil":  "Brasilia",
		}

		for country, capital := range capitals {
			fmt.Println(capital, " is the capital of ", country)
		}
	}
}

func OnlySomeTypesAllowedAsKeys() {
	// numbers
	ints := map[int]string{}
	ints[0] = "0"
	ints[1] = "1"

	// strings
	strings := map[string]int{}
	strings["email@gmail.com"] = 1

	// bytes
	bytes := map[byte]string{}
	bytes[0] = "byByte"

	// arrays

	firstStudentScores := [5]int{10, 12, 11, 10, 9}
	secondStudentScores := [5]int{10, 12, 11, 10, 9}
	thirdStudentScores := [5]int{10, 12, 11, 10, 10}

	sameSequenceScore := map[[5]int]int{}

	sameSequenceScore[firstStudentScores]++
	sameSequenceScore[secondStudentScores]++
	sameSequenceScore[thirdStudentScores]++

	fmt.Println(sameSequenceScore)

	// some structs
	structs := map[testScore]string{}
	structs[testScore{studentName: "Greg", score: 12}] = "whaaat?!"

	positions := make(map[position]int, 1000)
	positions[position{1, 1}] = 10

	fmt.Println("occurrences with position", position{1, 1}, "are", positions[position{1, 1}])

	// Pointers
	// points := map[*testScore]string{&testScore{}: "pointer keys are allowed"}
}

func KeysMustBeDistinct() {
	scores := map[string]int{"Greg": 12, "Sasha": 10, "USERNAME_UNKNOWN": -1}

	for name, score := range scores {
		fmt.Println(name, " has a score of ", score)
	}
	fmt.Println("------------------")

	scores["Greg"] = 0

	for name, score := range scores {
		fmt.Println(name, " has a score of ", score)
	}
}

func Delete() {
	scores := map[string]int{"Greg": 12, "Sasha": 10, "USERNAME_UNKNOWN": -1}

	for name, score := range scores {
		fmt.Println(name, " has a score of ", score)
	}
	fmt.Println("------------------")

	delete(scores, "Greg")

	for name, score := range scores {
		fmt.Println(name, " has a score of ", score)
	}
	fmt.Println("------------------")

	delete(scores, "No-op")

	for name, score := range scores {
		fmt.Println(name, " has a score of ", score)
	}
	fmt.Println("------------------")

}

func MapsArePointers() {
	scores := map[string]int{"Greg": 12, "Sasha": 10, "USERNAME_UNKNOWN": -1}

	for name, score := range scores {
		fmt.Println(name, " has a score of ", score)
	}

	scores = nil

	for name, score := range scores {
		fmt.Println(name, " has a score of ", score)
	}

	fmt.Printf("len is %d, value is %p\n", len(scores), scores)
}

func TwoValuesAssignment() {

	// students := map[int]testScore{}

	// greg := students[1]

	// fmt.Printf("%+v, %d\n", greg, greg.score)

	scores := map[string]int{"Greg": 12, "Sasha": 10, "USERNAME_UNKNOWN": -1}

	gregScore := scores["Greg"]

	fmt.Println("For ", "Greg", "score ", "is ", gregScore)

	gregoryScore := scores["Gregory"]

	fmt.Println("For ", "Gregory", "score ", "is ", gregoryScore)

	// gregoryScore, ok := scores["Gregory"]

	// if ok {
	// 	// fmt
	// }

	if gregoryScore, ok := scores["Gregory"]; !ok {
		fmt.Println("Missing student")
	} else {
		fmt.Println("Gregory has a score of ", gregoryScore)
	}
}
