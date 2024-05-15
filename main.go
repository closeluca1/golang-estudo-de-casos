package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    jsonData := `[
        {"name": "Alice", "age": 30},
        {"name": "Bob", "age": 25}
    ]`

    var people []Person
    err := json.Unmarshal([]byte(jsonData), &people)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    for _, person := range people {
        fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
    }
}
