package src

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GenCSV() {
	// Read the JSON file
	data := `[
		{
			"name": "A",
			"age": 1
		},
		{
			"name": "A2",
			"age": 2
		},
		{
			"name": "A3",
			"age": 1
		},
		{
			"name": "A4",
			"age": 2
		}
	]
	`
	// Unmarshal the JSON data
	var people []Person
	err := json.Unmarshal([]byte(data), &people)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create the CSV file
	file, err := os.Create("people.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Write the CSV data
	writer := csv.NewWriter(file)

	// Write title
	peopleTemp := Person{}
	t := reflect.TypeOf(peopleTemp)
	record := []string{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		record = append(record, field.Name)
	}

	err = writer.Write(record)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, person := range people {
		record := []string{person.Name, fmt.Sprintf("%d", person.Age)}
		err := writer.Write(record)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	writer.Flush()
}
