package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/AlexeyChudov/basics/context_pckg"
)

func main() {
	// f := closures.Outter(1)
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(f(1))
	// }
	// s := Student{
	// 	Rating: []int{4, 5, 3},
	// }
	// s.addToRating(5)
	// fmt.Println("non-pointer method result: ", s.Rating)
	// s.addToRatingP(4)
	// fmt.Println("non-pointer method result: ", s.Rating)
	// interfaces.InterfaceStruct()
	// interfaces.InterfaceAsType()
	// errors.MainFunc()

	context_pckg.MainFunc()
}

func fileTask() {
	json_task()
	file, err := os.Open("data.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var index int = 1
	reader := bufio.NewReader(file)
	delim := byte(';')
	for s, err := reader.ReadString(delim); err == nil; s, err = reader.ReadString(delim) {
		index++
		// fmt.Println(s)
		// if err != nil{fmt.Println(err)}
		if s == "0;" {
			fmt.Println(index)
			break
		}

	}
}

func map_example() {
	groupCity := map[int][]string{
		10:   []string{"Село", "Деревня", "ПГТ"},  // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Станица"},        // города с населением 100-999 тыс. человек
		1000: []string{"Мегаполис", "Миллионник"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"Город":     101,
		"Станица":   102,
		"Село":      103,
		"Мегаполис": 104,
	}
	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
	fmt.Printf("Type %T\n", x)
	for i := range cityPopulation {
		for _, el := range groupCity[10] {
			if i == el {
				delete(cityPopulation, i)
			}
		}
		for _, el := range groupCity[1000] {
			if i == el {
				delete(cityPopulation, i)
			}
		}

		fmt.Println(i, cityPopulation[i])
	}

}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Adress     string
	Phone      string
	Rating     []int
}
type Average struct{ average float32 }
type Group struct {
	Id       int
	Number   string
	Year     int
	Students []Student
}

func (s Student) addToRating(mark int) {
	s.Rating = append(s.Rating, mark)
}

func (s *Student) addToRatingP(mark int) {
	s.Rating = append(s.Rating, mark)
}

func json_task() {

	file, err := os.Open("data.txt")
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	var group Group = Group{}
	var sum int = 0

	if err := json.Unmarshal(data, &group); err != nil {
		fmt.Println(err)
		return
	}
	for _, student := range group.Students {
		sum += len(student.Rating)
	}
	var average Average
	average.average = float32(sum) / float32(len(group.Students))
	fmt.Println(sum, len(group.Students), average.average)
	s, _ := json.MarshalIndent(average, "", "    ")
	fmt.Printf("%s\n", string(s))

}
